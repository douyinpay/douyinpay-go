package partnerapp

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
	AppSpMchID             = "" // 需要商户自行替换
	AppSpAppID             = "" // 需要商户自行替换
	AppSubMchID            = "" // 需要商户自行替换
	AppSubAppID            = "" // 需要商户自行替换
	AppOutTradeNo          = "" // 需要商户自行替换
	AppTransactionID       = "" // 需要商户自行替换
	AppMerchantSerialNo    = "" // 商户证书序列号，需要商户自行替换
	AppMerchantPrivateKey  = "" // 需要商户自行替换
	AppPlatformCertificate = "" // 需要商户自行替换
)

func TestAppRequest(t *testing.T) {
	t.Run("预下单", func(t *testing.T) {
		AppPrepayTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("关单", func(t *testing.T) {
		AppCloseOrderTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("根据交易订单ID查询详情", func(t *testing.T) {
		AppQueryOrderByIdTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("根据交易外部单号查询详情", func(t *testing.T) {
		AppQueryOrderByOutTradeNoTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})
}

func AppPrepayTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, AppSpMchID, AppMerchantSerialNo, AppMerchantPrivateKey, AppPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, AppSpMchID, AppMerchantSerialNo, AppMerchantPrivateKey, AppPlatformCertificate)
	}
	svc := AppApiService{Client: c}

	outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.Prepay(ctx,
		PrePayRequest{
			SubMchid:      AppSubMchID,
			SubAppid:      AppSubAppID,
			SpMchid:       AppSpMchID,
			SpAppid:       AppSpAppID,
			Description:   "抖音支付测试",
			OutTradeNo:    outTradeNo,
			TimeExpire:    time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			Attach:        "自定义数据",
			NotifyUrl:     "https://www.mock.douyinpay.com",
			GoodsTag:      "DouyingPay",
			SupportFapiao: false,
			Amount: &Amount{
				Currency: "CNY",
				Total:    100,
			},
			Detail: &Detail{
				CostPrice: 608800,
				GoodsDetail: []*GoodsDetail{&GoodsDetail{
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

func AppCloseOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, AppSpMchID, AppMerchantSerialNo, AppMerchantPrivateKey, AppPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, AppSpMchID, AppMerchantSerialNo, AppMerchantPrivateKey, AppPlatformCertificate)
	}
	svc := AppApiService{Client: c}

	outTradeNo := AppOutTradeNo
	result, err := svc.CloseOrder(ctx,
		CloseOrderRequest{
			OutTradeNo: outTradeNo,
			SubMchid:   AppSubMchID,
			SpMchid:    AppSpMchID,
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

func AppQueryOrderByIdTest(ctx context.Context, signType string, t *testing.T) {

	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, AppSpMchID, AppMerchantSerialNo, AppMerchantPrivateKey, AppPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, AppSpMchID, AppMerchantSerialNo, AppMerchantPrivateKey, AppPlatformCertificate)
	}
	svc := AppApiService{Client: c}

	transId := AppTransactionID
	resp, result, err := svc.QueryOrderById(ctx,
		QueryOrderByIdRequest{
			TransactionId: transId,
			SubMchid:      AppSubMchID,
			SpMchid:       AppSpMchID,
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

func AppQueryOrderByOutTradeNoTest(ctx context.Context, signType string, t *testing.T) {

	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, AppSpMchID, AppMerchantSerialNo, AppMerchantPrivateKey, AppPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, AppSpMchID, AppMerchantSerialNo, AppMerchantPrivateKey, AppPlatformCertificate)
	}
	svc := AppApiService{Client: c}

	outTradeNo := AppOutTradeNo
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
		QueryOrderByOutTradeNoRequest{
			OutTradeNo: outTradeNo,
			SubMchid:   AppSubMchID,
			SpMchid:    AppSpMchID,
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
