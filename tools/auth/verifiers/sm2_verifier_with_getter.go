package verifiers

import (
	"context"
	"fmt"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/utils"
)

type Sm2VerifierWithGetter struct {
	CertGetter client.CertificateGetter
}

// Verify 对数字签名信息进行验证
func (verifier *Sm2VerifierWithGetter) Verify(ctx context.Context, serialNumber, message, signature string) error {
	err := checkParameter(ctx, serialNumber, message, signature)
	if err != nil {
		return err
	}
	if verifier.CertGetter == nil {
		return fmt.Errorf("verifier has no validator")
	}
	cert, ok := verifier.CertGetter.Get(ctx, serialNumber)
	if !ok {
		return fmt.Errorf("certificate[%s] not found in verifier", serialNumber)
	}
	v := &Sm2Verifier{PublicKey: utils.CovertToSm2PublicKey(cert), SerialNumber: serialNumber}
	return v.Verify(ctx, serialNumber, message, signature)
}

// GetSerial 获取可验签的平台证书序列号
func (verifier *Sm2VerifierWithGetter) GetSerial(ctx context.Context) (string, error) {
	return verifier.CertGetter.GetNewestSerial(ctx), nil
}
