package services

import (
	"context"
	"log"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/client/option"
	"github.com/douyinpay/douyinpay-go/utils"
)

/**
 * 初始化RSA证书类型的client
 * 同时提供证书定时更新功能（因此需要提供 encryptKey 用于证书解密），不再需要本地提供平台证书
 */
func InitAutoClientRSA(ctx context.Context, mchID, merchantCertSerialNo, merchantPrivateKeyString, encryptKey string) *client.Client {
	// 加载商户私钥
	mchPrivateKey, err := utils.LoadPrivateKey(merchantPrivateKeyString)
	if err != nil {
		log.Printf("load merchant private key error:%+v", err)
	}
	// 使用商户私钥等初始化client
	c, err := client.NewClient(ctx,
		option.WithClientAgentName("AutoRSA", mchID),
		option.WithRSAClientAutoVisitor(mchID, merchantCertSerialNo, mchPrivateKey, encryptKey),
	)
	if err != nil {
		log.Printf("new client err:%+v", err)
	}
	return c
}

/**
 * 初始化SM2证书类型的client
 * 同时提供证书定时更新功能（因此需要提供 encryptKey 用于证书解密），不再需要本地提供平台证书
 */
func InitAutoClientSM2(ctx context.Context, mchID, merchantCertSerialNo, merchantPrivateKeyString, encryptKey string) *client.Client {
	// 加载商户私钥
	mchPrivateKey, err := utils.LoadSm2PrivateKey(merchantPrivateKeyString)
	if err != nil {
		log.Printf("load merchant private key error:%+v", err)
	}
	//初始化client
	c, err := client.NewClient(ctx,
		option.WithClientAgentName("AutoSM2", mchID),
		option.WithSm2ClientAutoVisitor(mchID, merchantCertSerialNo, mchPrivateKey, encryptKey),
	)
	if err != nil {
		log.Printf("new client err:%+v", err)
	}
	return c
}
