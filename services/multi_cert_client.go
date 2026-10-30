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
 * 同时提供证书定时更新功能（因此需要提供 encryptKey 用于证书解密），不再需要本地提供平台证书
 */
// Deprecated: use NewAutoClientRSA to receive initialization errors.
func InitAutoClientRSA(ctx context.Context, mchID, merchantCertSerialNo, merchantPrivateKeyString, encryptKey string) *client.Client {
	c, err := NewAutoClientRSA(ctx, mchID, merchantCertSerialNo, merchantPrivateKeyString, encryptKey)
	if err != nil {
		log.Printf("init auto rsa client error:%+v", err)
	}
	return c
}

// NewAutoClientRSA 初始化 RSA 自动证书 client，初始化失败时返回错误。
func NewAutoClientRSA(ctx context.Context, mchID, merchantCertSerialNo, merchantPrivateKeyString, encryptKey string) (*client.Client, error) {
	// 加载商户私钥
	mchPrivateKey, err := utils.LoadPrivateKey(merchantPrivateKeyString)
	if err != nil {
		return nil, fmt.Errorf("init auto rsa client: load merchant private key: %w", err)
	}
	// 使用商户私钥等初始化client
	c, err := client.NewClient(ctx,
		option.WithClientAgentName("AutoRSA", mchID),
		option.WithRSAClientAutoVisitor(mchID, merchantCertSerialNo, mchPrivateKey, encryptKey),
	)
	if err != nil {
		return nil, fmt.Errorf("init auto rsa client: new client: %w", err)
	}
	return c, nil
}

/**
 * 初始化SM2证书类型的client
 * 同时提供证书定时更新功能（因此需要提供 encryptKey 用于证书解密），不再需要本地提供平台证书
 */
// Deprecated: use NewAutoClientSM2 to receive initialization errors.
func InitAutoClientSM2(ctx context.Context, mchID, merchantCertSerialNo, merchantPrivateKeyString, encryptKey string) *client.Client {
	c, err := NewAutoClientSM2(ctx, mchID, merchantCertSerialNo, merchantPrivateKeyString, encryptKey)
	if err != nil {
		log.Printf("init auto sm2 client error:%+v", err)
	}
	return c
}

// NewAutoClientSM2 初始化 SM2 自动证书 client，初始化失败时返回错误。
func NewAutoClientSM2(ctx context.Context, mchID, merchantCertSerialNo, merchantPrivateKeyString, encryptKey string) (*client.Client, error) {
	// 加载商户私钥
	mchPrivateKey, err := utils.LoadSm2PrivateKey(merchantPrivateKeyString)
	if err != nil {
		return nil, fmt.Errorf("init auto sm2 client: load merchant private key: %w", err)
	}
	//初始化client
	c, err := client.NewClient(ctx,
		option.WithClientAgentName("AutoSM2", mchID),
		option.WithSm2ClientAutoVisitor(mchID, merchantCertSerialNo, mchPrivateKey, encryptKey),
	)
	if err != nil {
		return nil, fmt.Errorf("init auto sm2 client: new client: %w", err)
	}
	return c, nil
}
