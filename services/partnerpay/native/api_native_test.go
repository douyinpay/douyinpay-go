package native

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
	NativeSpMchID             = "" // 需要商户自行替换
	NativeSpAppID             = "" // 需要商户自行替换
	NativeSubMchID            = "" // 需要商户自行替换
	NativeSubAppID            = "" // 需要商户自行替换
	NativeOutTradeNo          = "" // 需要商户自行替换
	NativeTransactionID       = "" // 需要商户自行替换
	NativeMerchantSerialNo    = "" // 商户证书序列号，需要商户自行替换
	NativeMerchantPrivateKey  = "" // 需要商户自行替换
	NativePlatformCertificate = "" // 需要商户自行替换
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
		c = services.InitClientRSA(ctx, NativeSpMchID, NativeMerchantSerialNo, NativeMerchantPrivateKey, NativePlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, NativeSpMchID, NativeMerchantSerialNo, NativeMerchantPrivateKey, NativePlatformCertificate)
	}
	svc := NativeApiService{Client: c}

	outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.Prepay(ctx,
		PrepayRequest{
			SubMchid:    NativeSubMchID,
			SubAppid:    NativeSubAppID,
			SpMchid:     NativeSpMchID,
			SpAppid:     NativeSpAppID,
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

	assert.Equal(t, true, len(resp.CodeUrl) > 0)
}

func CloseOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, NativeSpMchID, NativeMerchantSerialNo, NativeMerchantPrivateKey, NativePlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, NativeSpMchID, NativeMerchantSerialNo, NativeMerchantPrivateKey, NativePlatformCertificate)
	}
	svc := NativeApiService{Client: c}

	outTradeNo := NativeOutTradeNo
	result, err := svc.CloseOrder(ctx,
		CloseOrderRequest{
			OutTradeNo: outTradeNo,
			SubMchid:   NativeSubMchID,
			SpMchid:    NativeSpMchID,
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
		c = services.InitClientRSA(ctx, NativeSpMchID, NativeMerchantSerialNo, NativeMerchantPrivateKey, NativePlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, NativeSpMchID, NativeMerchantSerialNo, NativeMerchantPrivateKey, NativePlatformCertificate)
	}
	svc := NativeApiService{Client: c}

	transId := NativeTransactionID
	resp, result, err := svc.QueryOrderById(ctx,
		QueryOrderByIdRequest{
			TransactionId: transId,
			SubMchid:      NativeSubMchID,
			SpMchid:       NativeSpMchID,
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
		c = services.InitClientRSA(ctx, NativeSpMchID, NativeMerchantSerialNo, NativeMerchantPrivateKey, NativePlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, NativeSpMchID, NativeMerchantSerialNo, NativeMerchantPrivateKey, NativePlatformCertificate)
	}
	svc := NativeApiService{Client: c}

	outTradeNo := NativeOutTradeNo
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
		QueryOrderByOutTradeNoRequest{
			OutTradeNo: outTradeNo,
			SubMchid:   NativeSubMchID,
			SpMchid:    NativeSpMchID,
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
