package encryptors

import (
	"context"
	"fmt"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/tools/crypto"
	"github.com/douyinpay/douyinpay-go/utils"
	"github.com/tjfoc/gmsm/sm2"
)

type SM2EncryptorWithGetter struct {
	CertGetter client.CertificateGetter
}

func (e *SM2EncryptorWithGetter) Encrypt(ctx context.Context, plaintext string) (string, error) {
	if plaintext == "" {
		return plaintext, nil
	}
	publicKey, serial, err := e.getPlatformPublicKey(ctx)
	if err != nil {
		return "", err
	}
	base := &SM2Encryptor{
		PlatformPublicKey: publicKey,
		PlatformSerial:    serial,
	}
	return base.Encrypt(ctx, plaintext)
}

func (e *SM2EncryptorWithGetter) GetPlatformSerial(ctx context.Context) (string, error) {
	_, serial, err := e.getPlatformPublicKey(ctx)
	return serial, err
}

func (e *SM2EncryptorWithGetter) getPlatformPublicKey(ctx context.Context) (*sm2.PublicKey, string, error) {
	if e.CertGetter == nil {
		return nil, "", fmt.Errorf("cert getter is empty")
	}
	serial := e.CertGetter.GetNewestSerial(ctx)
	if serial == "" {
		return nil, "", fmt.Errorf("platform serial is empty")
	}
	cert, ok := e.CertGetter.Get(ctx, serial)
	if !ok || cert == nil {
		return nil, "", fmt.Errorf("platform certificate not found")
	}
	publicKey := utils.CovertToSm2PublicKey(cert)
	if publicKey == nil {
		return nil, "", fmt.Errorf("platform public key is empty")
	}
	return publicKey, serial, nil
}

var _ crypto.Encryptor = (*SM2EncryptorWithGetter)(nil)
