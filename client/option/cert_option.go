package option

import (
	"crypto/rsa"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/tools/auth/signers"
	"github.com/douyinpay/douyinpay-go/tools/auth/verifiers"
	"github.com/douyinpay/douyinpay-go/tools/crypto/decryptors"
	"github.com/douyinpay/douyinpay-go/tools/crypto/encryptors"
	"github.com/douyinpay/douyinpay-go/utils"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

type clientOption struct{ settings client.DialSettings }

// Apply 设置 client.DialSettings 的 Signer、Validator 以及 Cipher
func (w clientOption) Apply(o *client.DialSettings) error {
	o.Signer = w.settings.Signer
	o.Verifier = w.settings.Verifier
	if w.settings.AgentName != "" {
		o.AgentName = w.settings.AgentName
	}
	return nil
}

func OptionSignAndVerifyWithRSA(mchID string, mchCertificateSerialNo string, privateKey *rsa.PrivateKey, platCertificate *x509.Certificate) client.ClientOption {
	return clientOption{
		settings: client.DialSettings{
			Signer: &signers.SHA256WithRSASigner{
				MchID:               mchID,
				CertificateSerialNo: mchCertificateSerialNo,
				PrivateKey:          privateKey,
			},
			Verifier: &verifiers.SHA256WithRSAVerifier{
				PublicKey:    platCertificate.PublicKey.(*rsa.PublicKey),
				SerialNumber: utils.ConvertSerailNo(platCertificate.SerialNumber),
			},
			Encryptor: &encryptors.RSAEncryptor{
				PlatformPublicKey: platCertificate.PublicKey.(*rsa.PublicKey),
				PlatformSerial:    utils.ConvertSerailNo(platCertificate.SerialNumber),
			},
			Decryptor: &decryptors.RSADecryptor{
				MerchantPrivateKey: privateKey,
			},
		},
	}
}

func OptionSignAndVerifyWithSM2(mchID string, mchCertificateSerialNo string, privateKey *sm2.PrivateKey, platCertificate *x509.Certificate) client.ClientOption {
	return clientOption{
		settings: client.DialSettings{
			Signer: &signers.Sm2Signer{
				MchID:               mchID,
				CertificateSerialNo: mchCertificateSerialNo,
				PrivateKey:          privateKey,
			},
			Verifier: &verifiers.Sm2Verifier{
				PublicKey:    utils.CovertToSm2PublicKey(platCertificate),
				SerialNumber: utils.ConvertSerailNo(platCertificate.SerialNumber),
			},
			Encryptor: &encryptors.SM2Encryptor{
				PlatformPublicKey: utils.CovertToSm2PublicKey(platCertificate),
				PlatformSerial:    utils.ConvertSerailNo(platCertificate.SerialNumber),
			},
			Decryptor: &decryptors.SM2Decryptor{
				MerchantPrivateKey: privateKey,
			},
		},
	}
}
