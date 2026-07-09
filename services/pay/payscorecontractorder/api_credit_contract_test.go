package payscorecontractorder

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/douyinpay/douyinpay-go/client"

	"github.com/douyinpay/douyinpay-go/services"
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

/*
*
测试用例
*/
func TestRequest(t *testing.T) {
	t.Run("预下单App", func(t *testing.T) {
		PrepayTestApp(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("预下单JsApi", func(t *testing.T) {
		PrepayTestJsApi(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("预下单H5", func(t *testing.T) {
		PrepayTestH5(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("关单", func(t *testing.T) {
		AppCloseOrderTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("根据交易订单ID查询详情", func(t *testing.T) {
		QueryOrderByIdTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("根据交易外部单号查询详情", func(t *testing.T) {
		QueryOrderByOutTradeNoTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})
}

func PrepayTestApp(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := CreditContractOrderApiService{Client: c}

	outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	ServiceId := "569"

	resp, result, err := svc.Prepay(ctx,
		PrepayRequest{
			Appid:       appId,
			Mchid:       mchID,
			Description: "抖音支付测试",
			OutTradeNo:  outTradeNo,
			TimeExpire:  time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			Attach:      "自定义数据",
			NotifyUrl:   "https://www.mock.douyinpay.com",
			GoodsTag:    "DouyingPay",
			TradeType:   consts.TradeTypeApp,
			Amount: &Amount{
				Total: 100,
			},
			Detail: &Detail{
				CostPrice: 608800,
				GoodsDetail: []GoodsDetail{GoodsDetail{
					GoodsName:        "测试商品",
					MerchantGoodsId:  "ABC",
					Quantity:         1,
					UnitPrice:        828800,
					DouyinpayGoodsId: "1001",
				}},
				InvoiceId: "dy123",
			},
			SceneInfo: &SceneInfo{
				DeviceId:      "013467007045764",
				PayerClientIp: "14.23.150.211",
			},
			SettleInfo: &SettleInfo{
				ProfitSharing: false,
			},
			ContractInfo: &ContractInfo{
				ContractMchId:          mchID,
				ContractAppId:          appId,
				OutContractCode:        fmt.Sprintf("%s%d", "contract", time.Now().Unix()),
				RequestSerial:          123,
				ContractNotifyUrl:      "https://www.mock.douyinpay.com",
				ContractDisplayAccount: "123",
				ServiceId:              &ServiceId,
			},
		},
	)

	log.Printf("outTradeNo:%s\nRequest-Id:%s", outTradeNo, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, len(resp.PrepayId) > 0)
}
func PrepayTestJsApi(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	ServiceId := "569"
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := CreditContractOrderApiService{Client: c}

	outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.Prepay(ctx,
		PrepayRequest{
			Appid:       appId,
			Mchid:       mchID,
			Description: "抖音支付测试",
			OutTradeNo:  outTradeNo,
			TimeExpire:  time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			Attach:      "自定义数据",
			NotifyUrl:   "https://www.mock.douyinpay.com",
			GoodsTag:    "DouyingPay",
			Openid:      "invalidopenid",
			TradeType:   consts.TradeTypeJsapi,

			Amount: &Amount{
				Total: 100,
			},
			Detail: &Detail{
				CostPrice: 608800,
				GoodsDetail: []GoodsDetail{GoodsDetail{
					GoodsName:        "测试商品",
					MerchantGoodsId:  "ABC",
					Quantity:         1,
					UnitPrice:        828800,
					DouyinpayGoodsId: "1001",
				}},
				InvoiceId: "dy123",
			},
			SceneInfo: &SceneInfo{
				DeviceId:      "013467007045764",
				PayerClientIp: "14.23.150.211",
			},
			SettleInfo: &SettleInfo{
				ProfitSharing: false,
			},
			ContractInfo: &ContractInfo{
				ContractMchId:          mchID,
				ContractAppId:          appId,
				OutContractCode:        fmt.Sprintf("%s%d", "contract", time.Now().Unix()),
				RequestSerial:          123,
				ContractNotifyUrl:      "https://www.mock.douyinpay.com",
				ContractDisplayAccount: "123",
				ServiceId:              &ServiceId,
			},
		},
	)

	log.Printf("outTradeNo:%s\nRequest-Id:%s", outTradeNo, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, len(resp.PrepayId) > 0)
}

func PrepayTestH5(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	ServiceId := "569"
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := CreditContractOrderApiService{Client: c}

	outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.Prepay(ctx,
		PrepayRequest{
			Appid:       appId,
			Mchid:       mchID,
			Description: "抖音支付测试",
			OutTradeNo:  outTradeNo,
			TimeExpire:  time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			Attach:      "自定义数据",
			NotifyUrl:   "https://www.mock.douyinpay.com",
			GoodsTag:    "DouyingPay",
			TradeType:   consts.TradeTypeMweb,
			Amount: &Amount{
				Total: 100,
			},
			Detail: &Detail{
				CostPrice: 608800,
				GoodsDetail: []GoodsDetail{GoodsDetail{
					GoodsName:        "测试商品",
					MerchantGoodsId:  "ABC",
					Quantity:         1,
					UnitPrice:        828800,
					DouyinpayGoodsId: "1001",
				}},
				InvoiceId: "dy123",
			},
			SceneInfo: &SceneInfo{
				DeviceId:      "013467007045764",
				PayerClientIp: "14.23.150.211",
				H5Info: &H5Info{
					Type:    "Wap",
					AppName: "测试应用",
					AppUrl:  "https://www.mock.douyinpay.com",
				},
			},
			SettleInfo: &SettleInfo{
				ProfitSharing: false,
			},
			ContractInfo: &ContractInfo{
				ContractMchId:          mchID,
				ContractAppId:          appId,
				OutContractCode:        fmt.Sprintf("%s%d", "contract", time.Now().Unix()),
				RequestSerial:          123,
				ContractNotifyUrl:      "https://www.mock.douyinpay.com",
				ContractDisplayAccount: "123",
				ServiceId:              &ServiceId,
			},
		},
	)

	log.Printf("outTradeNo:%s\nRequest-Id:%s", outTradeNo, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, len(resp.H5Url) > 0)
}
func AppCloseOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := CreditContractOrderApiService{Client: c}

	outTradeNo := "OUT_1729140216"
	result, err := svc.CloseOrder(ctx,
		CloseOrderRequest{
			OutTradeNo: outTradeNo,
			Mchid:      mchID,
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
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := CreditContractOrderApiService{Client: c}

	transId := "2105012410170500036421644194"
	resp, result, err := svc.QueryOrderById(ctx,
		QueryOrderByIdRequest{
			TransactionId: transId,
			Mchid:         mchID,
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
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := CreditContractOrderApiService{Client: c}

	outTradeNo := "OUT_1729140382"
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
		QueryOrderByOutTradeNoRequest{
			OutTradeNo: outTradeNo,
			Mchid:      mchID,
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
