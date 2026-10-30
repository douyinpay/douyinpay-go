package services

import (
	"context"
	"fmt"
	"log"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/client/option"
	"github.com/douyinpay/douyinpay-go/utils"
)

/**
 * 初始化RSA证书类型的client
 */
// Deprecated: use NewClientRSA to receive initialization errors.
func InitClientRSA(ctx context.Context, mchID, merchantCertSerialNo, merchantPrivateKeyString, plantCertString string) *client.Client {
	c, err := NewClientRSA(ctx, mchID, merchantCertSerialNo, merchantPrivateKeyString, plantCertString)
	if err != nil {
		log.Printf("init rsa client error:%+v", err)
	}
	return c
}

// NewClientRSA 初始化 RSA 证书类型的 client，初始化失败时返回错误。
func NewClientRSA(ctx context.Context, mchID, merchantCertSerialNo, merchantPrivateKeyString, plantCertString string) (*client.Client, error) {
	// 加载商户私钥
	mchPrivateKey, err := utils.LoadPrivateKey(merchantPrivateKeyString)
	if err != nil {
		return nil, fmt.Errorf("init rsa client: load merchant private key: %w", err)
	}
	// 加载平台证书
	platCertificate, err := utils.LoadCertificate(plantCertString)
	if err != nil {
		return nil, fmt.Errorf("init rsa client: load platform certificate: %w", err)
	}
	// 使用商户私钥等初始化client
	c, err := client.NewClient(ctx,
		option.WithClientAgentName("RSA", mchID),
		option.OptionSignAndVerifyWithRSA(mchID, merchantCertSerialNo, mchPrivateKey, platCertificate),
	)
	if err != nil {
		return nil, fmt.Errorf("init rsa client: new client: %w", err)
	}
	return c, nil
}

/**
 * 初始化SM2证书类型的client
 */
// Deprecated: use NewClientSM2 to receive initialization errors.
func InitClientSM2(ctx context.Context, mchID, merchantCertSerialNo, merchantPrivateKeyString, plantCertString string) *client.Client {
	c, err := NewClientSM2(ctx, mchID, merchantCertSerialNo, merchantPrivateKeyString, plantCertString)
	if err != nil {
		log.Printf("init sm2 client error:%+v", err)
	}
	return c
}

// NewClientSM2 初始化 SM2 证书类型的 client，初始化失败时返回错误。
func NewClientSM2(ctx context.Context, mchID, merchantCertSerialNo, merchantPrivateKeyString, plantCertString string) (*client.Client, error) {
	// 加载商户私钥
	mchPrivateKey, err := utils.LoadSm2PrivateKey(merchantPrivateKeyString)
	if err != nil {
		return nil, fmt.Errorf("init sm2 client: load merchant private key: %w", err)
	}
	// 加载平台证书
	platCertificate, err := utils.LoadSm2Certificate(plantCertString)
	if err != nil {
		return nil, fmt.Errorf("init sm2 client: load platform certificate: %w", err)
	}
	//初始化client
	c, err := client.NewClient(ctx,
		option.WithClientAgentName("SM2", mchID),
		option.OptionSignAndVerifyWithSM2(mchID, merchantCertSerialNo, mchPrivateKey, platCertificate),
	)
	if err != nil {
		return nil, fmt.Errorf("init sm2 client: new client: %w", err)
	}
	return c, nil
}
