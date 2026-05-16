package crypto

import "context"

// Decryptor 敏感字段解密
type Decryptor interface {
	Decrypt(ctx context.Context, ciphertext string) (string, error)
}
