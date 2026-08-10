package jsapi

import (
	"context"
	"fmt"
	"log"
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

func TestRequest(t *testing.T) {
	t.Run("预下单", func(t *testing.T) {
		PrepayTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("关单", func(t *testing.T) {
		CloseOrderTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("根据交易订单ID查询详情", func(t *testing.T) {
		QueryOrderByIdTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("根据交易外部单号查询详情", func(t *testing.T) {
		QueryOrderByOutTradeNoTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})
}

func PrepayTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, JsapiSpMchID, JsapiMerchantSerialNo, JsapiMerchantPrivateKey, JsapiPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, JsapiSpMchID, JsapiMerchantSerialNo, JsapiMerchantPrivateKey, JsapiPlatformCertificate)
	}
	svc := JsapiApiService{Client: c}

	outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.Prepay(ctx,
		PrepayRequest{
			SubMchid:    JsapiSubMchID,
			SubAppid:    JsapiSubAppID,
			SpMchid:     JsapiSpMchID,
			SpAppid:     JsapiSpAppID,
			Description: "抖音支付测试",
			OutTradeNo:  outTradeNo,
			TimeExpire:  time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			Attach:      "自定义数据",
			NotifyUrl:   "https://www.mock.douyinpay.com",
			GoodsTag:    "DouyingPay",
			Amount: &Amount{
				Currency: "CNY",
				Total:    100,
			},
			SceneInfo: &SceneInfo{
				DeviceId:      "013467007045764",
				PayerClientIp: "14.23.150.211",
				StoreInfo: &StoreInfo{
					Address:  "北京市海淀区中关村大街",
					AreaCode: "100191",
					Id:       "0001",
					Name:     "测试店铺",
				},
			},
			SettleInfo: &SettleInfo{
				ProfitSharing: false,
			},
			PayerInfo: &PayerInfo{
				SpOpenId:  JsapiOpenID,
				SubOpenId: "",
			},
		},
	)

	log.Printf("outTradeNo:%s\nRequest-Id:%s", outTradeNo, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+s", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, len(resp.PrepayId) > 0)
}

func CloseOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, JsapiSpMchID, JsapiMerchantSerialNo, JsapiMerchantPrivateKey, JsapiPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, JsapiSpMchID, JsapiMerchantSerialNo, JsapiMerchantPrivateKey, JsapiPlatformCertificate)
	}
	svc := JsapiApiService{Client: c}

	outTradeNo := JsapiOutTradeNo
	result, err := svc.CloseOrder(ctx,
		CloseOrderRequest{
			OutTradeNo: outTradeNo,
			SubMchid:   JsapiSubMchID,
			SpMchid:    JsapiSpMchID,
		},
	)
	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call CloseOrder err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d,header=%+v", result.Response.StatusCode, result.Response.Header)
	}

	assert.Equal(t, true, 200 == result.Response.StatusCode || 400 == result.Response.StatusCode)
}

func QueryOrderByIdTest(ctx context.Context, signType string, t *testing.T) {

	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, JsapiSpMchID, JsapiMerchantSerialNo, JsapiMerchantPrivateKey, JsapiPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, JsapiSpMchID, JsapiMerchantSerialNo, JsapiMerchantPrivateKey, JsapiPlatformCertificate)
	}
	svc := JsapiApiService{Client: c}

	transId := JsapiTransactionID
	resp, result, err := svc.QueryOrderById(ctx,
		QueryOrderByIdRequest{
			TransactionId: transId,
			SubMchid:      JsapiSubMchID,
			SpMchid:       JsapiSpMchID,
		},
	)

	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call QueryOrderById err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, len(resp.TransactionId) > 0)
}

func QueryOrderByOutTradeNoTest(ctx context.Context, signType string, t *testing.T) {

	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, JsapiSpMchID, JsapiMerchantSerialNo, JsapiMerchantPrivateKey, JsapiPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, JsapiSpMchID, JsapiMerchantSerialNo, JsapiMerchantPrivateKey, JsapiPlatformCertificate)
	}
	svc := JsapiApiService{Client: c}

	outTradeNo := JsapiOutTradeNo
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
		QueryOrderByOutTradeNoRequest{
			OutTradeNo: outTradeNo,
			SubMchid:   JsapiSubMchID,
			SpMchid:    JsapiSpMchID,
		},
	)

	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call QueryOrderByOutTradeNo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, len(resp.TransactionId) > 0)
}
