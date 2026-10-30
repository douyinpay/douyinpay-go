package verifiers

import (
	"context"
	"fmt"
)

// NilVerifier 空验证器，不对报文进行验证，对任意报文均不会返回错误。
//
// Deprecated: 该类型仅用于兼容历史版本。业务请求与回包验签应使用明确的证书或公钥
// verifier；SDK 内部需要跳过验签的场景应使用内部实现，避免将内部逻辑作为对外 API 暴露。
type NilVerifier struct{}

// Verify 跳过报文签名验证。
func (v *NilVerifier) Verify(ctx context.Context, serialNumber, message, signature string) error {
	return nil
}

// GetSerial 返回空验证器无法提供证书序列号的错误。
func (v *NilVerifier) GetSerial(ctx context.Context) (string, error) {
	return "", fmt.Errorf("NilVerifier has no serial")
}
