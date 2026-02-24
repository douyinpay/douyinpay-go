package auth

import "context"

// Verifier 数字签名验证器
type Verifier interface {
	Verify(ctx context.Context, serial, message, signature string) error // 对签名信息进行验证
	GetSerial(ctx context.Context) (string, error)                       // 客户端可以处理的证书或者公钥序列号
}
