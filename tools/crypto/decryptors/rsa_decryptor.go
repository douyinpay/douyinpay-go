package decryptors

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"

	"github.com/douyinpay/douyinpay-go/tools/crypto"
)

type RSADecryptor struct {
	MerchantPrivateKey *rsa.PrivateKey
}

func (d *RSADecryptor) Decrypt(ctx context.Context, ciphertext string) (string, error) {
	if ciphertext == "" {
		return ciphertext, nil
	}
	if d.MerchantPrivateKey == nil {
		return "", fmt.Errorf("rsa privateKey is required")
	}
	raw, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("decode rsa ciphertext error:%s", err.Error())
	}
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, d.MerchantPrivateKey, raw)
	if err != nil {
		return "", fmt.Errorf("rsa decrypt error:%s", err.Error())
	}
	return string(plaintext), nil
}

var _ crypto.Decryptor = (*RSADecryptor)(nil)
