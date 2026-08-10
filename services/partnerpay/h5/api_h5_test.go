package partnerh5

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
	H5SpMchID             = "" // 需要商户自行替换
	H5SpAppID             = "" // 需要商户自行替换
	H5SubMchID            = "" // 需要商户自行替换
	H5SubAppID            = "" // 需要商户自行替换
	H5OutTradeNo          = "" // 需要商户自行替换
	H5TransactionID       = "" // 需要商户自行替换
	H5MerchantSerialNo    = "" // 商户证书序列号，需要商户自行替换
	H5MerchantPrivateKey  = "" // 需要商户自行替换
	H5PlatformCertificate = "" // 需要商户自行替换
)

func TestH5Request(t *testing.T) {
	t.Run("预下单", func(t *testing.T) {
		H5PrepayTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("关单", func(t *testing.T) {
		H5CloseOrderTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("根据交易订单ID查询详情", func(t *testing.T) {
		H5QueryOrderByIdTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("根据交易外部单号查询详情", func(t *testing.T) {
		H5QueryOrderByOutTradeNoTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})
}

func H5PrepayTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, H5SpMchID, H5MerchantSerialNo, H5MerchantPrivateKey, H5PlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, H5SpMchID, H5MerchantSerialNo, H5MerchantPrivateKey, H5PlatformCertificate)
	}
	svc := H5ApiService{Client: c}

	outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.Prepay(ctx,
		PrepayRequest{
			SubMchid:    H5SubMchID,
			SubAppid:    H5SubAppID,
			SpMchid:     H5SpMchID,
			SpAppid:     H5SpAppID,
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
				H5Info: &H5Info{
					Type:    "Wap",
					AppName: "测试应用",
					AppUrl:  "https://www.mock.douyinpay.com",
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

	assert.Equal(t, true, len(resp.H5_url) > 0)
}

func H5CloseOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, H5SpMchID, H5MerchantSerialNo, H5MerchantPrivateKey, H5PlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, H5SpMchID, H5MerchantSerialNo, H5MerchantPrivateKey, H5PlatformCertificate)
	}
	svc := H5ApiService{Client: c}

	outTradeNo := H5OutTradeNo
	result, err := svc.CloseOrder(ctx,
		CloseOrderRequest{
			OutTradeNo: outTradeNo,
			SubMchid:   H5SubMchID,
			SpMchid:    H5SpMchID,
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

func H5QueryOrderByIdTest(ctx context.Context, signType string, t *testing.T) {

	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, H5SpMchID, H5MerchantSerialNo, H5MerchantPrivateKey, H5PlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, H5SpMchID, H5MerchantSerialNo, H5MerchantPrivateKey, H5PlatformCertificate)
	}
	svc := H5ApiService{Client: c}

	transId := H5TransactionID
	resp, result, err := svc.QueryOrderById(ctx,
		QueryOrderByIdRequest{
			TransactionId: transId,
			SubMchid:      H5SubMchID,
			SpMchid:       H5SpMchID,
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

func H5QueryOrderByOutTradeNoTest(ctx context.Context, signType string, t *testing.T) {

	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, H5SpMchID, H5MerchantSerialNo, H5MerchantPrivateKey, H5PlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, H5SpMchID, H5MerchantSerialNo, H5MerchantPrivateKey, H5PlatformCertificate)
	}
	svc := H5ApiService{Client: c}

	outTradeNo := H5OutTradeNo
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
		QueryOrderByOutTradeNoRequest{
			OutTradeNo: outTradeNo,
			SubMchid:   H5SubMchID,
			SpMchid:    H5SpMchID,
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
