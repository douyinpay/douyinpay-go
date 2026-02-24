package refund_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/services/refund"
	"github.com/douyinpay/douyinpay-go/services/secret"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"
)

var (
	JsapiSpMchID                  = secret.JsapiSpMchID                  // 需要商户自行替换
	JsapiSpAppID                  = secret.JsapiSpAppID                  // 需要商户自行替换
	JsapiSubMchID                 = secret.JsapiSubMchID                 // 需要商户自行替换
	JsapiSubAppID                 = secret.JsapiSubAppID                 // 需要商户自行替换
	JsapiOpenID                   = secret.JsapiOpenID                   // 需要商户自行替换
	JsapiOutTradeNo               = secret.JsapiOutTradeNo               // 需要商户自行替换
	JsapiTransactionID            = secret.JsapiTransactionID            // 需要商户自行替换
	JsapiMerchantSerialNo         = secret.JsapiMerchantSerialNo         // 需要商户自行替换                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // 商户证书序列号
	JsapiMerchantPrivateKey       = secret.JsapiMerchantPrivateKey       // 需要商户自行替换
	JsapiPlatformCertificate      = secret.JsapiPlatformCertificate      // 需要商户自行替换
	RefundMchID                   = secret.RefundMchID                   // 需要商户自行替换
	RefundAppID                   = secret.RefundAppID                   // 需要商户自行替换
	RefundSubMchID                = secret.RefundSubMchID                // 需要商户自行替换
	RefundTransactionID           = secret.RefundTransactionID           // 需要商户自行替换
	RefundOutRefundNo             = secret.RefundOutRefundNo             // 需要商户自行替换
	RefundMerchantSerialNo        = secret.RefundMerchantSerialNo        // 需要商户自行替换                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   // 商户证书序列号
	RefundMerchantPrivateKey      = secret.RefundMerchantPrivateKey      // 需要商户自行替换
	RefundPlatformCertificate     = secret.RefundPlatformCertificate     // 需要商户自行替换
	DeductMchID                   = secret.DeductMchID                   // 需要商户自行替换
	DeductAppID                   = secret.DeductAppID                   // 需要商户自行替换
	DeductMerchantSerialNo        = secret.DeductMerchantSerialNo        // 需要商户自行替换
	DeductMerchantPrivateKey      = secret.DeductMerchantPrivateKey      // 需要商户自行替换
	DeductPlatformCertificate     = secret.DeductPlatformCertificate     // 需要商户自行替换
	DeductOutContractCode         = secret.DeductOutContractCode         // 需要商户自行替换
	DeductPlanID                  = secret.DeductPlanID                  // 需要商户自行替换
	DeductContractID              = secret.DeductContractID              // 需要商户自行替换
	DeductDeletePlanID            = secret.DeductDeletePlanID            // 需要商户自行替换
	DeductDeleteOutContractCode   = secret.DeductDeleteOutContractCode   // 需要商户自行替换
	DeductPayApplyOutContractCode = secret.DeductPayApplyOutContractCode // 需要商户自行替换
	DeductPayApplyContractID      = secret.DeductPayApplyContractID      // 需要商户自行替换
	DeductNotifyAppID             = secret.DeductNotifyAppID             // 需要商户自行替换
	DeductNotifyContractID        = secret.DeductNotifyContractID        // 需要商户自行替换
	MchID                         = secret.MchID                         // 需要商户自行替换
	AppID                         = secret.AppID                         // 需要商户自行替换
	MerchantSerialNo              = secret.MerchantSerialNo              // 需要商户自行替换
	MerchantPrivateKey            = secret.MerchantPrivateKey            // 需要商户自行替换
	PlatformCertificate           = secret.PlatformCertificate           // 需要商户自行替换
	SubMchID                      = secret.SubMchID                      // 需要商户自行替换
	LocalAppID                    = secret.LocalAppID                    // 需要商户自行替换
	LocalOutContractCode          = secret.LocalOutContractCode          // 需要商户自行替换
	LocalPlanID                   = secret.LocalPlanID                   // 需要商户自行替换
	LocalPlanIdStr                = secret.LocalPlanIdStr                // 需要商户自行替换
	LocalOpenID                   = secret.LocalOpenID                   // 需要商户自行替换
	LocalSubOpenID                = secret.LocalSubOpenID                // 需要商户自行替换
	LocalContractID               = secret.LocalContractID               // 需要商户自行替换
	LocalOutContractCode2         = secret.LocalOutContractCode2         // 需要商户自行替换
	LocalContractIDQuery          = secret.LocalContractIDQuery          // 需要商户自行替换
	PayMchID                      = secret.PayMchID                      // 需要商户自行替换
	PayAppID                      = secret.PayAppID                      // 需要商户自行替换
	PayMerchantSerialNo           = secret.PayMerchantSerialNo           // 需要商户自行替换
	PayMerchantPrivateKey         = secret.PayMerchantPrivateKey         // 需要商户自行替换
	PayPlatformCertificate        = secret.PayPlatformCertificate        // 需要商户自行替换
	PayOutTradeNo                 = secret.PayOutTradeNo                 // 需要商户自行替换
	PayTransactionId              = secret.PayTransactionId              // 需要商户自行替换
	PayQueryTradeNo               = secret.PayQueryTradeNo               // 需要商户自行替换
)

func TestRefundRequest(t *testing.T) {
	t.Run("退款", func(t *testing.T) {
		RefundCreateTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("退款查询", func(t *testing.T) {
		QueryByOutRefundNoTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

}

func RefundCreateTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, RefundMchID, RefundMerchantSerialNo, RefundMerchantPrivateKey, RefundPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, RefundMchID, RefundMerchantSerialNo, RefundMerchantPrivateKey, RefundPlatformCertificate)
	}
	svc := refund.RefundsApiService{Client: c}

	outRefundNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.Create(ctx,
		refund.CreateRequest{
			Amount: &refund.AmountReq{
				Currency: "CNY",
				// From: []AmountFrom{
				// 	{
				// 		Account: "AVAILABLE",
				// 		Amount:  1,
				// 	},
				// },
				Refund: 1,
				Total:  100,
			},
			//Mchid:    mchID,
			SpMchid:  RefundMchID,
			SubMchid: RefundSubMchID,
			//FundsAccount: "AVAILABLE",
			// Goo:[
			// 	{
			// 		"douyinpay_goods_id":"1001",
			// 		"goods_name":"iPhone6s 16G",
			// 		"merchant_goods_id":"1217752501201407033233368018",
			// 		"refund_amount":528800,
			// 		"refund_quantity":1,
			// 		"unit_price":528800
			// 	}
			// ],
			NotifyUrl:   "https://mock.douyinpay.com",
			OutRefundNo: outRefundNo,
			//OutTradeNo:  "lt_2025102919270231603160",
			TransactionId: RefundTransactionID,
			Reason:        "商品已售完",
		},
	)

	log.Printf("outTradeNo:%s\nRequest-Id:%s", outRefundNo, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call Create err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, utils.Json2Str(resp))
	}
}

func QueryByOutRefundNoTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, RefundMchID, RefundMerchantSerialNo, RefundMerchantPrivateKey, RefundPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, RefundMchID, RefundMerchantSerialNo, RefundMerchantPrivateKey, RefundPlatformCertificate)
	}
	svc := refund.RefundsApiService{Client: c}

	outRefundNo := RefundOutRefundNo
	resp, result, err := svc.QueryByOutRefundNo(ctx,
		refund.QueryByOutRefundNoRequest{
			OutRefundNo: outRefundNo,
			SpMchid:     RefundMchID,
			SubMchid:    RefundSubMchID,
		},
	)
	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call QueryByOutRefundNo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, len(resp.TransactionId) > 0)
}
