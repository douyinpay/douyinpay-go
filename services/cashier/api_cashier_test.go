package cashier_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/services/cashier"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

var (
	mchID               = "" // 商户号
	appId               = "" // 商户AppId
	merchantSerialNo    = "" // 商户证书序列号
	merchantPrivateKey  = "" // 商户私钥
	platformCertificate = "" // 平台证书

	blindMobile = "" // SHA256 盲化后的手机号
)

func TestCashierRequest(t *testing.T) {
	t.Run("前置咨询-仅查询用户标签", func(t *testing.T) {
		PrePayConsultTagTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("前置咨询-查询用户标签和用户营销", func(t *testing.T) {
		PrePayConsultPromotionTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})
}

// PrePayConsultTagTest 仅查询用户标签（tag_code/tag_name）
func PrePayConsultTagTest(ctx context.Context, signType string, t *testing.T) {
	svc := cashier.ApiCashierService{Client: initClient(ctx, signType)}

	resp, result, err := svc.PrePayConsult(ctx,
		cashier.PrePayConsultRequest{
			Appid:           appId,
			Mchid:           mchID,
			GoodsTag:        `{"biz_scene":"pre_consult"}`,
			EncryptType:     "SHA256",
			BlindMobileList: []string{blindMobile},
		})

	if err != nil {
		// 处理错误
		log.Printf("call PrePayConsult err: %s", err)
	} else {
		// 处理返回结果
		jsonData, _ := json.MarshalIndent(resp, "", "  ")
		log.Printf("status=%d resp=%s", result.Response.StatusCode, string(jsonData))
	}

	assert.Equal(t, true, resp != nil)
}

// PrePayConsultPromotionTest 查询用户标签（tag_code/tag_name）和用户营销（operation_tip）
func PrePayConsultPromotionTest(ctx context.Context, signType string, t *testing.T) {
	svc := cashier.ApiCashierService{Client: initClient(ctx, signType)}

	resp, result, err := svc.PrePayConsult(ctx,
		cashier.PrePayConsultRequest{
			Appid:                 appId,
			Mchid:                 mchID,
			TotalAmount:           "2000",
			ProductCode:           []string{"NormalPay"},
			CommericalProductCode: "CO_PAY_APP",
			TradeType:             consts.TradeTypeApp,
			GoodsTag:              `{"biz_scene":"pre_consult"}`,
			EncryptType:           "SHA256",
			BlindMobileList:       []string{blindMobile},
		})

	if err != nil {
		// 处理错误
		log.Printf("call PrePayConsult err: %s", err)
	} else {
		// 处理返回结果
		jsonData, _ := json.MarshalIndent(resp, "", "  ")
		log.Printf("status=%d resp=%s", result.Response.StatusCode, string(jsonData))
	}

	assert.Equal(t, true, resp != nil)
}

func initClient(ctx context.Context, signType string) *client.Client {
	if signType == consts.CRYPTO_TYPE_RSA {
		return services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	return services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
}
