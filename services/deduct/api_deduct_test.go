package contractorder

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
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

/*
*
测试用例
*/
func TestDeduct(t *testing.T) {
	ctx := context.Background()

	t.Run("申请扣款", func(t *testing.T) {
		DeductTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("关单", func(t *testing.T) {
		CloseOrderTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("根据交易订单ID查询详情", func(t *testing.T) {
		QueryOrderByIdTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("根据交易外部单号查询详情", func(t *testing.T) {
		QueryOrderByOutTradeNoTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	//
	t.Run("预扣费通知 - a", func(t *testing.T) {
		DeductNotifyTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("预扣费通知 - param,rpc error", func(t *testing.T) {
		DeductNotifyParamErrorTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("申请扣款-PayApplyTest-a", func(t *testing.T) {
		PayApplyTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("申请扣款-PayApplyTest-param、rpc error", func(t *testing.T) {
		PayApplyParamErrorTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

}

func DeductTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	}
	svc := ApiDeductService{Client: c}

	outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.Deduct(ctx,
		ApiDeductRequest{
			Appid:       DeductAppID,
			Mchid:       DeductMchID,
			OutTradeNo:  outTradeNo,
			TimeExpire:  time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			ContractId:  "",
			TradeType:   consts.TradeTypeCyclePay,
			Description: "抖音支付申请扣款测试",
			NotifyUrl:   "https://www.mock.douyinpay.com",
			Attach:      "自定义数据",
			GoodsTag:    "test_tag",
			Detail: &Detail{
				CostPrice: 608800,
				GoodsDetail: []GoodsDetail{{
					GoodsName:        "测试商品",
					MerchantGoodsId:  "ABC",
					Quantity:         1,
					UnitPrice:        828800,
					DouyinpayGoodsId: "1001",
				}},
				InvoiceId: "dy123",
			},
			Amount: &Amount{
				Currency: "CNY",
				Total:    1,
			},
			SceneInfo: &SceneInfo{
				DeviceId:      "",
				PayerClientIp: "14.23.150.211",
			},
			SettleInfo: &SettleInfo{
				ProfitSharing: false,
			},
		},
	)

	log.Printf("outTradeNo: %s\nRequest-Id: %s", outTradeNo, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call Deduct err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil && resp.ResultCode != "")
}

func CloseOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	}
	svc := ApiDeductService{Client: c}

	outTradeNo := "OUT_1671077833"
	result, err := svc.CloseOrder(ctx,
		CloseOrderRequest{
			OutTradeNo: outTradeNo,
			Mchid:      DeductMchID,
		},
	)
	log.Printf("Request-Id: %s", result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call CloseOrder err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d,header=%+v", result.Response.StatusCode, result.Response.Header)
	}

	assert.Equal(t, true, 200 == result.Response.StatusCode || 400 == result.Response.StatusCode)
}

func QueryOrderByIdTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	}
	svc := ApiDeductService{Client: c}

	transId := "TP2022121512161718679398479660"
	resp, result, err := svc.QueryOrderById(ctx,
		QueryOrderByIdRequest{
			TransactionId: transId,
			Mchid:         DeductMchID,
		},
	)

	log.Printf("Request-Id: %s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call QueryOrderById err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil && len(resp.TransactionId) > 0)
}

func DeductNotifyTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	}
	svc := ApiDeductService{Client: c}
	req := DeductNotifyRequest{
		//ContractId: "MSN2508271037592330733690893420",

		ContractId: DeductNotifyContractID,
		//Appid:      appId,
		Mchid: DeductMchID,
		Appid: DeductNotifyAppID,
		EstimatedAmount: DeductAmount{
			Amount:   1,
			Currency: "CNY",
		},
	}
	result, err := svc.DeductNotify(ctx, req)

	log.Printf("ContractId: %s\n resp : %v", req.ContractId, result.Response)
	if err != nil {
		// 处理错误
		log.Printf("call QueryOrderById err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(result.Response))
	}
	assert.Equal(t, true, err == nil || strings.Contains(err.Error(), "RESOURCE_ALREADY_EXISTS"))
}

func DeductNotifyParamErrorTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	}
	svc := ApiDeductService{Client: c}
	req := DeductNotifyRequest{
		//Contractid: "MSN2508271037592330733690893420",

		ContractId: DeductNotifyContractID,
		//Appid:      appId,
		Mchid: DeductMchID,
		Appid: DeductNotifyAppID,
		EstimatedAmount: DeductAmount{
			Amount:   1,
			Currency: "CNY",
		},
	}

	param := req
	param.ContractId = ""
	_, err := svc.DeductNotify(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.Mchid = ""
	_, err = svc.DeductNotify(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.Appid = ""
	_, err = svc.DeductNotify(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.EstimatedAmount.Amount = -1
	_, err = svc.DeductNotify(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.EstimatedAmount.Currency = ""
	_, err = svc.DeductNotify(ctx, param)
	assert.Equal(t, true, err != nil)
}

func QueryOrderByOutTradeNoTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	}
	svc := ApiDeductService{Client: c}

	outTradeNo := "OUT_1671077833"
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
		QueryOrderByOutTradeNoRequest{
			OutTradeNo: outTradeNo,
			Mchid:      DeductMchID,
		},
	)

	log.Printf("Request-Id: %s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call QueryOrderByOutTradeNo err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil && len(resp.TransactionId) > 0)
}

func PayApplyTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	}
	svc := ApiDeductService{Client: c}
	goodsArr := []GoodsDetail{
		{
			GoodsName:       "",
			MerchantGoodsId: "ABC",
			Quantity:        1,
			UnitPrice:       1,
		},
	}
	//outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	outTradeNo := DeductPayApplyOutContractCode
	resp, result, err := svc.PayApply(ctx,
		PayApplyRequest{
			Appid:       DeductAppID,
			Mchid:       DeductMchID,
			OutTradeNo:  outTradeNo,
			TimeExpire:  "2022-12-01T20:59:16+08:00",
			ContractId:  DeductPayApplyContractID,
			Attach:      "自定义数据",
			Description: "抖音支付测试",
			Detail: Detail{
				CostPrice:   1,
				GoodsDetail: goodsArr,
			},
			TradeType: "SGP",
			Amount: &Amount{
				Currency: "CNY",
				Total:    1,
			},
			NotifyUrl: "https://www.mock.com",
			GoodsTag:  "",
			// SceneInfo: SceneInfo{
			// 	MerchantDeviceId: "013467007045764",
			// 	PayerClientIp: "14.23.150.211",
			// 	PayerDeviceId: "a0e4b456-c9e5-3783-a422",
			// },
			// SettleInfo: SettleInfo{
			// 	ProfitSharing: false,
			// },
		},
	)

	log.Printf("Request-Id: %s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call PayApplyTest err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil && resp.ResultCode == "SUCCESS")
}

func PayApplyParamErrorTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, DeductMchID, DeductMerchantSerialNo, DeductMerchantPrivateKey, DeductPlatformCertificate)
	}
	svc := ApiDeductService{Client: c}
	goodsArr := []GoodsDetail{
		{
			GoodsName:       "",
			MerchantGoodsId: "ABC",
			Quantity:        1,
			UnitPrice:       1,
		},
	}
	outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())

	req := PayApplyRequest{
		Appid:       DeductAppID,
		Mchid:       DeductMchID,
		OutTradeNo:  outTradeNo,
		TimeExpire:  "2022-12-01T20:59:16+08:00",
		ContractId:  DeductPayApplyContractID,
		Attach:      "自定义数据",
		Description: "抖音支付测试",
		Detail: Detail{
			CostPrice:   1,
			GoodsDetail: goodsArr,
		},
		TradeType: "",
		Amount: &Amount{
			Currency: "CNY",
			Total:    1,
		},
		NotifyUrl: "https://www.mock.com",
		GoodsTag:  "",
		// SceneInfo: SceneInfo{
		// 	MerchantDeviceId: "013467007045764",
		// 	PayerClientIp: "14.23.150.211",
		// 	PayerDeviceId: "a0e4b456-c9e5-3783-a422",
		// },
		// SettleInfo: SettleInfo{
		// 	ProfitSharing: false,
		// },
	}

	param := req
	param.Appid = ""
	_, _, err := svc.PayApply(ctx, param)
	assert.Equal(t, true, err != nil)

	param = req
	param.Mchid = ""
	_, _, err = svc.PayApply(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.OutTradeNo = ""
	_, _, err = svc.PayApply(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.ContractId = ""
	_, _, err = svc.PayApply(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.TradeType = ""
	_, _, err = svc.PayApply(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.Description = ""
	_, _, err = svc.PayApply(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.Amount.Total = -1
	_, _, err = svc.PayApply(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.NotifyUrl = ""
	_, _, err = svc.PayApply(ctx, param)
	assert.Equal(t, true, err != nil)

	param = req
	param.ContractId = "csss"
	_, _, err = svc.PayApply(ctx, param)
	assert.Equal(t, true, err != nil)

}
