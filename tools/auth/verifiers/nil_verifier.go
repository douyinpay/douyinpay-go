package verifiers

import (
	"context"
	"fmt"
)

// NilVerifier 空验证器，不对报文进行验证，对任意报文均不会返回错误，
// 在不需要对报文签名进行验证的情况（如微信支付账单文件下载）下使用
type NilVerifier struct {
}

// Validate 跳过报文签名验证
func (v *NilVerifier) Verify(ctx context.Context, serialNumber, message, signature string) error {
	return nil
}

func (v *NilVerifier) GetSerial(ctx context.Context) (string, error) {
	return "", fmt.Errorf("NilVerifier has no serial")
}
