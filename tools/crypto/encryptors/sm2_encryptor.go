package encryptors

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/douyinpay/douyinpay-go/tools/crypto"
	"github.com/tjfoc/gmsm/sm2"
)

type SM2Encryptor struct {
	PlatformPublicKey *sm2.PublicKey
	PlatformSerial    string
}

func (e *SM2Encryptor) Encrypt(ctx context.Context, plaintext string) (string, error) {
	if plaintext == "" {
		return plaintext, nil
	}
	publicKey, _, err := e.getPlatformPublicKey(ctx)
	if err != nil {
		return "", err
	}
	if publicKey == nil {
		return "", fmt.Errorf("sm2 publicKey is required")
	}
	ciphertext, err := sm2.EncryptAsn1(publicKey, []byte(plaintext), rand.Reader)
	if err != nil {
		return "", fmt.Errorf("sm2 encrypt error:%s", err.Error())
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (e *SM2Encryptor) GetPlatformSerial(ctx context.Context) (string, error) {
	if e.PlatformPublicKey == nil || e.PlatformSerial == "" {
		return "", fmt.Errorf("platform public key is empty")
	}
	return e.PlatformSerial, nil
}

func (e *SM2Encryptor) getPlatformPublicKey(ctx context.Context) (*sm2.PublicKey, string, error) {
	if e.PlatformPublicKey == nil || e.PlatformSerial == "" {
		return nil, "", fmt.Errorf("platform public key is empty")
	}
	return e.PlatformPublicKey, e.PlatformSerial, nil
}

var _ crypto.Encryptor = (*SM2Encryptor)(nil)
