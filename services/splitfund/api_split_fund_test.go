package splitfund_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/douyinpay/douyinpay-go/client"

	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/services/splitfund"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"
)

var (
	mchID               = "" // 商户号
	appId               = "" // 商户AppId
	merchantSerialNo    = "" // 商户证书序列号
	merchantPrivateKey  = "" // 商户私钥
	platformCertificate = "" // 平台证书
)

func TestSplitFund(t *testing.T) {
	t.Run("分账", func(t *testing.T) {
		DoSplitFund(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("分账查询", func(t *testing.T) {
		QuerySplitFundTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

}

func DoSplitFund(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := splitfund.SplitFundService{Client: c}
	outOrderNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.SplitFund(ctx,
		splitfund.SplitFundReq{
			MerchantId:      mchID,
			AppId:           appId,
			OutOrderNo:      outOrderNo,
			TradeNo:         "",
			UnfreezeUnsplit: true,
			Receivers: []*splitfund.ReceiverInfoDTO{&splitfund.ReceiverInfoDTO{
				Type:        "MERCHANT_ID",
				Account:     "",
				Amount:      1,
				Description: "",
			}},
		},
	)
	log.Printf("Request-Id:%s\nRequest-Id:%s", outOrderNo, result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call SplitFund err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, utils.Json2Str(resp))
	}
}

func QuerySplitFundTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := splitfund.SplitFundService{Client: c}

	resp, result, err := svc.QuerySplitFund(ctx,
		splitfund.QuerySplitFundReq{
			MerchantId: mchID,
			TradeNo:    "",
			OutOrderNo: "",
		},
	)
	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call QuerySplitFund err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, len(resp.TradeNo) > 0)
}
