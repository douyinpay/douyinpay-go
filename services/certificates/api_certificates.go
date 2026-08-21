package certificates

import (
	"context"
	"fmt"
	nethttp "net/http"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"

	"github.com/tjfoc/gmsm/x509"
)

type ApiCertificatesService services.Service

// DownloadCertificates 下载平台证书列表
// 接口路径：GET /v1/merchant/certificates/getPlatformCertificates
func (a *ApiCertificatesService) DownloadCertificates(ctx context.Context) (resp *DownloadCertificatesResponse, result *client.APIResult, err error) {
	// 发送请求
	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.DownloadCertificatesPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	// 解析结果
	resp = new(DownloadCertificatesResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// Download 下载平台证书并解密、解析为 x509 证书对象列表。
// signType 用于过滤证书类型（如 RSA/SM2），仅保留同类型的证书；
// encryptKey 为商户在“产品中心-密钥管理”设置的接口加密密钥，用于解密证书密文。
func (a *ApiCertificatesService) Download(ctx context.Context, signType, encryptKey string) ([]*x509.Certificate, error) {

	resp, _, err := a.DownloadCertificates(ctx)
	if err != nil {
		return nil, err
	}

	certificateList := make([]*x509.Certificate, 0)
	for _, rawCertificate := range resp.Certificates {
		err = rawCertificate.Validate()
		if err != nil {
			return nil, err
		}
		// 判断证书类型 过滤不符合的证书
		if signType != *rawCertificate.CertType {
			continue
		}
		encryptCertificate := rawCertificate.EncryptCertificate
		err = encryptCertificate.Validate()
		if err != nil {
			return nil, err
		}
		// 按加密算法解密证书密文，证书密文不携带关联数据（AAD 为空）
		var certContent string
		if *encryptCertificate.Algorithm == consts.EncryptTypeAES256GCM {
			certContent, err = utils.DecryptAES256GCM(encryptKey, "",
				*encryptCertificate.Nonce, *encryptCertificate.Ciphertext)
		} else {
			certContent, err = utils.DecryptSM4(encryptKey, "",
				*encryptCertificate.Nonce, *encryptCertificate.Ciphertext)
		}
		if err != nil {
			return nil, fmt.Errorf("decrypt downloaded certificate failed: %v", err)
		}

		// 按证书类型选用解析器：RSA 用标准 x509 解析，SM2 用国密解析
		var certificate *x509.Certificate
		if *rawCertificate.CertType == consts.CRYPTO_TYPE_RSA {
			certificate, err = utils.LoadCertificate(certContent)
		} else {
			certificate, err = utils.LoadSm2Certificate(certContent)
		}
		if err != nil {
			return nil, fmt.Errorf("parse downlaoded certificate failed: %v, certcontent:%v", err, certContent)
		}

		certificateList = append(certificateList, certificate)
	}

	return certificateList, nil
}
