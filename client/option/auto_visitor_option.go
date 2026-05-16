package option

import (
	"context"
	"crypto/rsa"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/tools/auth/signers"
	"github.com/douyinpay/douyinpay-go/tools/auth/verifiers"
	"github.com/douyinpay/douyinpay-go/tools/crypto/decryptors"
	"github.com/douyinpay/douyinpay-go/tools/crypto/encryptors"
	"github.com/douyinpay/douyinpay-go/tools/downloader"

	"github.com/tjfoc/gmsm/sm2"
)

// WithRSAClientAutoVisitor 一键初始化 RSAClient，使其具备「签名/验签/敏感字段加解密」能力。
// 同时提供证书定时更新功能（因此需要提供 encryptKey 用于证书解密），不再需要本地提供平台证书
func WithRSAClientAutoVisitor(
	mchID string, mchCertificateSerialNo string, privateKey *rsa.PrivateKey, encryptKey string,
) client.ClientOption {
	mgr := downloader.MgrInstance()

	if !mgr.HasDownloader(context.Background(), mchID) {
		err := mgr.RegisterRSADownloaderWithPrivateKey(context.Background(), mchID,
			mchCertificateSerialNo, privateKey, encryptKey)
		if err != nil {
			return client.ErrorOption{Error: err}
		}
	}
	return NewRSAClientOptionAuthVisitorUsingDownloaderMgr(mchID, mchCertificateSerialNo, privateKey, mgr)
}

// NewRSAClientOptionAuthVisitorUsingDownloaderMgr 一键初始化 SM2Client，使其具备「签名/验签/敏感字段加解密」能力。
// 需要使用者自行提供 CertificateDownloaderMgr 已实现平台证书的自动更新
//
// 【注意】本函数的能力与 WithRSAClientAutoVisitor 完全一致，除非有自行管理 CertificateDownloaderMgr 的需求，
// 否则推荐直接使用 WithRSAClientAutoVisitor
func NewRSAClientOptionAuthVisitorUsingDownloaderMgr(
	mchID string, mchCertificateSerialNo string, privateKey *rsa.PrivateKey, mgr *downloader.CertificateDownloaderMgr,
) client.ClientOption {
	return clientOption{
		settings: client.DialSettings{
			Signer: &signers.SHA256WithRSASigner{
				MchID:               mchID,
				CertificateSerialNo: mchCertificateSerialNo,
				PrivateKey:          privateKey,
			},
			Verifier: &verifiers.SHA256WithRSAVerifierWithGetter{
				CertGetter: mgr.GetCertificateVisitor(mchID),
			},
			Encryptor: &encryptors.RSAEncryptorWithGetter{
				CertGetter: mgr.GetCertificateVisitor(mchID),
			},
			Decryptor: &decryptors.RSADecryptor{
				MerchantPrivateKey: privateKey,
			},
		},
	}
}

// WithSm2ClientAutoVisitor 一键初始化 SM2Client，使其具备「签名/验签/敏感字段加解密」能力。
// 同时提供证书定时更新功能（因此需要提供 encryptKey 用于证书解密），不再需要本地提供平台证书
func WithSm2ClientAutoVisitor(
	mchID string, mchCertificateSerialNo string, privateKey *sm2.PrivateKey, encryptKey string,
) client.ClientOption {
	mgr := downloader.MgrInstance()

	if !mgr.HasDownloader(context.Background(), mchID) {
		err := mgr.RegisterSM2DownloaderWithPrivateKey(context.Background(), mchID,
			mchCertificateSerialNo, privateKey, encryptKey)
		if err != nil {
			return client.ErrorOption{Error: err}
		}
	}
	return NewSm2ClientOptionAuthVisitorUsingDownloaderMgr(mchID, mchCertificateSerialNo, privateKey, mgr)
}

// NewSm2ClientOptionAuthVisitorUsingDownloaderMgr 一键初始化 SM2Client，使其具备「签名/验签/敏感字段加解密」能力。
// 需要使用者自行提供 CertificateDownloaderMgr 已实现平台证书的自动更新
//
// 【注意】本函数的能力与 WithSm2ClientAutoVisitor 完全一致，除非有自行管理 CertificateDownloaderMgr 的需求，
// 否则推荐直接使用 WithSm2ClientAutoVisitor
func NewSm2ClientOptionAuthVisitorUsingDownloaderMgr(
	mchID string, mchCertificateSerialNo string, privateKey *sm2.PrivateKey, mgr *downloader.CertificateDownloaderMgr,
) client.ClientOption {
	return clientOption{
		settings: client.DialSettings{
			Signer: &signers.Sm2Signer{
				MchID:               mchID,
				CertificateSerialNo: mchCertificateSerialNo,
				PrivateKey:          privateKey,
			},
			Verifier: &verifiers.Sm2VerifierWithGetter{
				CertGetter: mgr.GetCertificateVisitor(mchID),
			},
			Encryptor: &encryptors.SM2EncryptorWithGetter{
				CertGetter: mgr.GetCertificateVisitor(mchID),
			},
			Decryptor: &decryptors.SM2Decryptor{
				MerchantPrivateKey: privateKey,
			},
		},
	}
}
