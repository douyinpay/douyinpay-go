package crypto

import "context"

// Encryptor 敏感字段加密
type Encryptor interface {
	Encrypt(ctx context.Context, plaintext string) (string, error)
	GetPlatformSerial(ctx context.Context) (string, error)
}
