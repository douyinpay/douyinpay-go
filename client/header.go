package client

import (
	"context"
	"fmt"
	"time"

	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"
)

func (c *Client) GenerateAuthorizationHeader(ctx context.Context, method, canonicalURL, signBody string) (string, error) {
	if c.signer == nil {
		return "", fmt.Errorf("please init client with signer")
	}
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Unix()
	message := fmt.Sprintf(consts.SignatureMessageFormat, method, canonicalURL, timestamp, nonce, signBody)
	signatureResult, err := c.signer.Sign(ctx, message)
	if err != nil {
		return "", err
	}
	authorization := fmt.Sprintf(
		consts.HeaderAuthorizationFormat, c.getAuthorizationType(),
		signatureResult.MchID, nonce, timestamp, signatureResult.CertificateSerialNo, signatureResult.Signature,
	)
	return authorization, nil
}

func (c *Client) getAuthorizationType() string {
	return "DouyinPay-" + c.signer.Algorithm()
}
