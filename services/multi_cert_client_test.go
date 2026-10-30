package services

import (
	"context"
	"time"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/client/option"
	"github.com/douyinpay/douyinpay-go/tools/downloader"
	"github.com/douyinpay/douyinpay-go/utils"
)

func ExampleNewAutoClientRSA_default() {
	// 示例参数，实际使用时请自行初始化
	var (
		mchID                  string
		mchCertificateSerialNo string
		mchPrivateKeyString    string
		encryptKeyString       string

		signType string
	)

	ctx := context.Background()

	var cli *client.Client
	var err error
	if signType == "RSA" {
		cli, err = NewAutoClientRSA(ctx, mchID, mchCertificateSerialNo, mchPrivateKeyString, encryptKeyString)
	} else {
		cli, err = NewAutoClientSM2(ctx, mchID, mchCertificateSerialNo, mchPrivateKeyString, encryptKeyString)
	}
	if err != nil {
		// err handle
	}

	// 接下来使用 cli 注册ApiService 进行请求发送
	//svc := app.AppApiService{Client: cli}
	//svc.Prepay()

	_ = cli
}

func ExampleNewAutoClientRSA_customMgr() {

	// 示例参数，实际使用时请自行初始化
	var (
		mchID                  string
		mchCertificateSerialNo string
		mchPrivateKeyString    string
		encryptKeyString       string

		signType string
	)

	ctx := context.Background()

	// 先设置定时器，设置定时刷新证书间隔
	// 更新间隔最大不建议超过 2 天，以免错过平台证书平滑切换窗口；
	// 同时亦不建议小于 1 小时，以避免过多请求导致浪费
	mgr := downloader.NewCertificateDownloaderMgrWithInterval(ctx, time.Second)
	defer mgr.Stop()

	var op client.ClientOption
	if signType == "RSA" {
		privateKey, err := utils.LoadPrivateKey(mchPrivateKeyString)
		if err != nil {
			// err handle
		}
		op = option.WithRSAClientAutoVisitor(mchID, mchCertificateSerialNo, privateKey, encryptKeyString)
	} else {
		privateKey, err := utils.LoadSm2PrivateKey(mchPrivateKeyString)
		if err != nil {
			// err handle
		}
		op = option.WithSm2ClientAutoVisitor(mchID, mchCertificateSerialNo, privateKey, encryptKeyString)
	}

	cli, err := client.NewClient(ctx,
		option.WithClientAgentName("custom", mchID), op)
	if err != nil {
		// err handle
	}

	// 接下来使用 cli 注册ApiService 进行请求发送
	//svc := app.AppApiService{Client: cli}
	//svc.Prepay()

	_ = cli

}
