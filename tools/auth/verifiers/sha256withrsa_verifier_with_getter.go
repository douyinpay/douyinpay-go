package verifiers

import (
	"context"
	"crypto/rsa"
	"fmt"

	"github.com/douyinpay/douyinpay-go/client"
)

type SHA256WithRSAVerifierWithGetter struct {
	CertGetter client.CertificateGetter
}

// Verify 对数字签名信息进行验证
func (verifier *SHA256WithRSAVerifierWithGetter) Verify(ctx context.Context, serialNumber, message, signature string) error {
	err := checkParameter(ctx, serialNumber, message, signature)
	if err != nil {
		return err
	}
	if verifier.CertGetter == nil {
		return fmt.Errorf("verifier has no validator")
	}
	certificate, ok := verifier.CertGetter.Get(ctx, serialNumber)
	if !ok {
		return fmt.Errorf("certificate[%s] not found in verifier", serialNumber)
	}
	v := &SHA256WithRSAVerifier{PublicKey: certificate.PublicKey.(*rsa.PublicKey), SerialNumber: serialNumber}
	return v.Verify(ctx, serialNumber, message, signature)
}

// GetSerial 获取可验签的平台证书序列号
func (verifier *SHA256WithRSAVerifierWithGetter) GetSerial(ctx context.Context) (string, error) {
	return verifier.CertGetter.GetNewestSerial(ctx), nil
}
