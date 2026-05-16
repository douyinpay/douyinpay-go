package decryptors

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/douyinpay/douyinpay-go/tools/crypto"
	"github.com/tjfoc/gmsm/sm2"
)

type SM2Decryptor struct {
	MerchantPrivateKey *sm2.PrivateKey
}

func (d *SM2Decryptor) Decrypt(ctx context.Context, ciphertext string) (string, error) {
	if ciphertext == "" {
		return ciphertext, nil
	}
	if d.MerchantPrivateKey == nil {
		return "", fmt.Errorf("sm2 privateKey is required")
	}
	raw, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("decode sm2 ciphertext error:%s", err.Error())
	}
	plaintext, err := sm2.DecryptAsn1(d.MerchantPrivateKey, raw)
	if err != nil {
		return "", fmt.Errorf("sm2 decrypt error:%s", err.Error())
	}
	return string(plaintext), nil
}

var _ crypto.Decryptor = (*SM2Decryptor)(nil)
