package encryptors

import (
	"context"
	"crypto/rsa"
	"fmt"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/tools/crypto"
)

type RSAEncryptorWithGetter struct {
	CertGetter client.CertificateGetter
}

func (e *RSAEncryptorWithGetter) Encrypt(ctx context.Context, plaintext string) (string, error) {
	if plaintext == "" {
		return plaintext, nil
	}
	publicKey, serial, err := e.getPlatformPublicKey(ctx)
	if err != nil {
		return "", err
	}
	base := &RSAEncryptor{
		PlatformPublicKey: publicKey,
		PlatformSerial:    serial,
	}
	return base.Encrypt(ctx, plaintext)
}

func (e *RSAEncryptorWithGetter) GetPlatformSerial(ctx context.Context) (string, error) {
	_, serial, err := e.getPlatformPublicKey(ctx)
	return serial, err
}

func (e *RSAEncryptorWithGetter) getPlatformPublicKey(ctx context.Context) (*rsa.PublicKey, string, error) {
	if e.CertGetter == nil {
		return nil, "", fmt.Errorf("cert getter is empty")
	}
	serial := e.CertGetter.GetNewestSerial(ctx)
	if serial == "" {
		return nil, "", fmt.Errorf("platform serial is empty")
	}
	cert, ok := e.CertGetter.Get(ctx, serial)
	if !ok || cert == nil || cert.PublicKey == nil {
		return nil, "", fmt.Errorf("platform certificate not found")
	}
	return cert.PublicKey.(*rsa.PublicKey), serial, nil
}

var _ crypto.Encryptor = (*RSAEncryptorWithGetter)(nil)
