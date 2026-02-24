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
 */
func InitClientRSA(ctx context.Context, mchID, merchantCertSerialNo, merchantPrivateKeyString, plantCertString string) *client.Client {
	// 加载商户私钥
	mchPrivateKey, err := utils.LoadPrivateKey(merchantPrivateKeyString)
	if err != nil {
		log.Printf("load merchant private key error:%+v", err)
	}
	// 加载平台证书
	platCertificate, err := utils.LoadCertificate(plantCertString)
	if err != nil {
		log.Printf("load plantform certificate error:%+v", err)
	}
	// 使用商户私钥等初始化client
	c, err := client.NewClient(ctx,
		option.WithClientAgentName("RSA", mchID),
		option.OptionSignAndVerifyWithRSA(mchID, merchantCertSerialNo, mchPrivateKey, platCertificate),
	)
	if err != nil {
		log.Printf("new client err:%+v", err)
	}
	return c
}

/**
 * 初始化SM2证书类型的client
 */
func InitClientSM2(ctx context.Context, mchID, merchantCertSerialNo, merchantPrivateKeyString, plantCertString string) *client.Client {
	// 加载商户私钥
	mchPrivateKey, err := utils.LoadSm2PrivateKey(merchantPrivateKeyString)
	if err != nil {
		log.Printf("load merchant private key error:%+v", err)
	}
	// 加载平台证书
	platCertificate, err := utils.LoadSm2Certificate(plantCertString)
	if err != nil {
		log.Printf("load plantform certificate error:%+v", err)
	}
	//初始化client
	c, err := client.NewClient(ctx,
		option.WithClientAgentName("SM2", mchID),
		option.OptionSignAndVerifyWithSM2(mchID, merchantCertSerialNo, mchPrivateKey, platCertificate),
	)
	if err != nil {
		log.Printf("new client err:%+v", err)
	}
	return c
}
