package services

import (
	"context"

	"github.com/douyinpay/douyinpay-go/client"
)

func ExampleNewClient_default() {
	// 示例参数，实际使用时请自行初始化
	var (
		mchID                  string
		mchCertificateSerialNo string
		mchPrivateKeyString    string
		plantformCertString    string

		signType string
	)

	ctx := context.Background()

	var cli *client.Client
	var err error
	if signType == "RSA" {
		cli, err = NewClientRSA(ctx, mchID, mchCertificateSerialNo, mchPrivateKeyString, plantformCertString)
	} else {
		cli, err = NewClientSM2(ctx, mchID, mchCertificateSerialNo, mchPrivateKeyString, plantformCertString)
	}
	if err != nil {
		// err handle
	}

	// 接下来使用 cli 注册ApiService 进行请求发送
	//svc := app.AppApiService{Client: cli}
	//svc.Prepay()

	_ = cli
}
