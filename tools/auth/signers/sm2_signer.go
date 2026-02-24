package signers

import (
	"context"
	"fmt"
	"strings"

	"github.com/douyinpay/douyinpay-go/tools/auth"
	"github.com/douyinpay/douyinpay-go/utils"

	"github.com/tjfoc/gmsm/sm2"
)

type Sm2Signer struct {
	MchID               string          // 商户号
	CertificateSerialNo string          // 商户证书序列号
	PrivateKey          *sm2.PrivateKey // 商户私钥
}

func (s *Sm2Signer) Sign(_ context.Context, message string) (*auth.SignatureResult, error) {
	if strings.TrimSpace(s.CertificateSerialNo) == "" {
		return nil, fmt.Errorf("you must set mch certificate serial no to use Sm2Signer")
	}
	signature, err := utils.SignSm2(message, s.PrivateKey)
	if err != nil {
		return nil, err
	}
	return &auth.SignatureResult{MchID: s.MchID, CertificateSerialNo: s.CertificateSerialNo, Signature: signature}, nil
}

// 返回使用的签名算法：SM2
func (s *Sm2Signer) Algorithm() string {
	return "SM2"
}
