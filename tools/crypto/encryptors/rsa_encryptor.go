package encryptors

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"

	"github.com/douyinpay/douyinpay-go/tools/crypto"
)

type RSAEncryptor struct {
	PlatformPublicKey *rsa.PublicKey
	PlatformSerial    string
}

func (e *RSAEncryptor) Encrypt(ctx context.Context, plaintext string) (string, error) {
	if plaintext == "" {
		return plaintext, nil
	}
	publicKey, _, err := e.getPlatformPublicKey(ctx)
	if err != nil {
		return "", err
	}
	if isAlreadyEncryptedRSA(plaintext, publicKey) {
		return plaintext, nil
	}
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plaintext))
	if err != nil {
		return "", fmt.Errorf("rsa encrypt error:%s", err.Error())
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (e *RSAEncryptor) GetPlatformSerial(ctx context.Context) (string, error) {
	if e.PlatformPublicKey == nil || e.PlatformSerial == "" {
		return "", fmt.Errorf("platform public key is empty")
	}
	return e.PlatformSerial, nil
}

func (e *RSAEncryptor) getPlatformPublicKey(ctx context.Context) (*rsa.PublicKey, string, error) {
	if e.PlatformPublicKey == nil || e.PlatformSerial == "" {
		return nil, "", fmt.Errorf("platform public key is empty")
	}
	return e.PlatformPublicKey, e.PlatformSerial, nil
}

func isAlreadyEncryptedRSA(value string, publicKey *rsa.PublicKey) bool {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return false
	}
	keySize := (publicKey.N.BitLen() + 7) / 8
	return len(decoded) == keySize
}

var _ crypto.Encryptor = (*RSAEncryptor)(nil)
