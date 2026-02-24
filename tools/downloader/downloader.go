package downloader

import (
	"context"
	"fmt"
	nethttp "net/http"
	"sync"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/tools/auth"
	"github.com/douyinpay/douyinpay-go/tools/auth/verifiers"

	"github.com/tjfoc/gmsm/x509"

	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"
)

// isSameCertificateMap Check if two CertificateMaps stores same certificates.
// Normally, checking serial number set is enough.
func isSameCertificateMap(l, r map[string]*x509.Certificate) bool {
	if l == nil && r == nil {
		return true
	}

	if len(l) != len(r) {
		return false
	}

	for serialNumber := range l {
		if _, ok := r[serialNumber]; !ok {
			return false
		}
	}

	return true
}

// CertificateDownloader 平台证书下载器，下载完成后可直接获取 x509.Certificate 对象或导出证书内容
type CertificateDownloader struct {
	certContents map[string]string // 证书文本内容，用于导出
	certificates *CertificateMap   // 证书实例
	client       *client.Client    // HTTPClient
	encryptKey   string            // 商户接口密钥
	signType     string            // signType 签名类型
	lock         sync.RWMutex
}

// Get 获取证书序列号对应的平台证书
func (d *CertificateDownloader) Get(ctx context.Context, serialNumber string) (*x509.Certificate, bool) {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return d.certificates.Get(ctx, serialNumber)
}

// GetAll 获取平台证书Map
func (d *CertificateDownloader) GetAll(ctx context.Context) map[string]*x509.Certificate {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return d.certificates.GetAll(ctx)
}

// GetNewestSerial 获取最新的平台证书的证书序列号
func (d *CertificateDownloader) GetNewestSerial(ctx context.Context) string {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return d.certificates.GetNewestSerial(ctx)
}

// Export 获取证书序列号对应的平台证书内容
func (d *CertificateDownloader) Export(_ context.Context, serialNumber string) (string, bool) {
	d.lock.RLock()
	defer d.lock.RUnlock()

	content, ok := d.certContents[serialNumber]
	return content, ok
}

// ExportAll 获取平台证书内容Map
func (d *CertificateDownloader) ExportAll(_ context.Context) map[string]string {
	d.lock.RLock()
	defer d.lock.RUnlock()

	ret := make(map[string]string)
	for serialNumber, content := range d.certContents {
		ret[serialNumber] = content
	}

	return ret
}

// 这里需要看下 map 值
func (d *CertificateDownloader) updateCertificates(
	ctx context.Context, certContents map[string]string, certificates map[string]*x509.Certificate,
) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if isSameCertificateMap(d.certificates.GetAll(ctx), certificates) {
		return
	}

	d.certContents = certContents
	d.certificates.Reset(certificates)

	var v auth.Verifier
	if d.signType == consts.CRYPTO_TYPE_RSA {
		v = &verifiers.SHA256WithRSAVerifierWithGetter{
			CertGetter: d.certificates,
		}
	} else {
		v = &verifiers.Sm2VerifierWithGetter{
			CertGetter: d.certificates,
		}
	}

	d.client = client.NewClientWithVerifier(d.client, v)
}

func (d *CertificateDownloader) performDownloading(ctx context.Context) (*DownloadCertificatesResponse, error) {
	// 发送请求
	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.DownloadCertificatesPath

	result, err := d.client.Request(ctx, r)
	if err != nil {
		return nil, err
	}
	// 解析结果
	resp := new(DownloadCertificatesResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DownloadCertificates 立即下载平台证书列表
func (d *CertificateDownloader) DownloadCertificates(ctx context.Context) error {
	resp, err := d.performDownloading(ctx)
	if err != nil {
		return err
	}
	// 判断是否有证书
	if len(resp.Certificates) == 0 {
		return fmt.Errorf("no certificate downloaded")
	}
	rawCertContentMap := make(map[string]string)
	certificateMap := make(map[string]*x509.Certificate)
	for _, rawCertificate := range resp.Certificates {
		err = rawCertificate.Validate()
		if err != nil {
			return err
		}
		// 判断证书类型 过滤不符合的证书
		if d.signType != *rawCertificate.CertType {
			continue
		}
		encryptCertificate := rawCertificate.EncryptCertificate
		err = encryptCertificate.Validate()
		if err != nil {
			return err
		}
		// 解密
		var certContent string
		if *encryptCertificate.Algorithm == consts.EncryptTypeAES256GCM {
			certContent, err = utils.DecryptAES256GCM(d.encryptKey, "",
				*encryptCertificate.Nonce, *encryptCertificate.Ciphertext)
		} else {
			certContent, err = utils.DecryptSM4(d.encryptKey, "",
				*encryptCertificate.Nonce, *encryptCertificate.Ciphertext)
		}
		if err != nil {
			return fmt.Errorf("decrypt downloaded certificate failed: %v", err)
		}

		var certificate *x509.Certificate
		if *rawCertificate.CertType == consts.CRYPTO_TYPE_RSA {
			certificate, err = utils.LoadCertificate(certContent)
		} else {
			certificate, err = utils.LoadSm2Certificate(certContent)
		}
		if err != nil {
			return fmt.Errorf("parse downlaoded certificate failed: %v, certcontent:%v", err, certContent)
		}

		serialNumber := *rawCertificate.CertNo

		rawCertContentMap[serialNumber] = certContent
		certificateMap[serialNumber] = certificate
	}

	if len(certificateMap) == 0 {
		return fmt.Errorf("no certificate downloaded")
	}

	// 更新平台证书

	d.lock.Lock()
	defer d.lock.Unlock()
	if isSameCertificateMap(d.certificates.GetAll(ctx), certificateMap) {
		return nil
	}
	d.certContents = rawCertContentMap
	d.certificates.Reset(certificateMap)
	var v auth.Verifier
	if d.signType == consts.CRYPTO_TYPE_RSA {
		v = &verifiers.SHA256WithRSAVerifierWithGetter{CertGetter: d}
	} else {
		v = &verifiers.Sm2VerifierWithGetter{CertGetter: d}
	}
	d.client = client.NewClientWithVerifier(d.client, v)
	return nil
}
