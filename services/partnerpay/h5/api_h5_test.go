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
	"github.com/douyinpay/douyinpay-go/services/secret"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"
)

var (
	mchID               = secret.AppPayMchID
	subMchID            = secret.AppPaySubMchID
	subAppID            = secret.AppPaySubAppID
	merchantSerialNo    = secret.AppPayMerchantSerialNo
	merchantPrivateKey  = secret.AppPayMerchantPrivateKey
	platformCertificate = secret.AppPayPlatformCertificate
	appID               = secret.AppPayAppID
	outTradeNo          = secret.PartnerH5OutTradeNo
	transId             = secret.PartnerH5TransactionId
)

/*
*
测试用例
*/
func TestParterH5(t *testing.T) {
	ctx := context.Background()

	t.Run("服务商- H5下单", func(t *testing.T) {
		PartnerH5CreateOrderTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("服务商- H5关单", func(t *testing.T) {
		PartnerH5CloseOrderTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("服务商- H5根据交易订单ID查询详情", func(t *testing.T) {
		PartnerH5QueryOrderByIdTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("服务商- H5根据交易外部单号查询详情", func(t *testing.T) {
		PartnerH5QueryOrderByOutTradeNoTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

}

func PartnerH5CreateOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	//DouyinPayServer := "https://bytepay-boe.byted.org"
	//mockey.Mock(contractorder.GetServerAddress).Return(DouyinPayServer).Build()
	svc := H5ApiService{Client: c}

	outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.Prepay(ctx,
		PrepayRequest{
			SpAppid:    appID,
			SpMchid:    mchID,
			SubMchid:   subMchID,
			SubAppid:   subAppID,
			OutTradeNo: outTradeNo,
			//TimeExpire:  time.Now().Add(24 * time.Hour).Unix(),
			Description: "抖音支付申请扣款测试",
			NotifyUrl:   "https://www.mock.douyinpay.com",
			Attach:      "自定义数据",
			GoodsTag:    "test_tag",
			Detail: &Detail{
				CostPrice: 100,
				GoodsDetail: []*GoodsDetail{{
					GoodsName:        "测试商品",
					MerchantGoodsId:  "ABC",
					Quantity:         1,
					UnitPrice:        100,
					DouyinpayGoodsId: "1001",
				}},
				InvoiceId: "dy123",
			},
			Amount: &Amount{
				Currency: "CNY",
				Total:    100,
			},
			SceneInfo: &SceneInfo{
				DeviceId:      "",
				PayerClientIp: "14.23.150.211",
				H5Info: &H5Info{
					Type: "ios",
				},
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

	assert.Equal(t, true, resp != nil && resp.H5_url != "")
}

func PartnerH5CloseOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	// DouyinPayServer := "https://bytepay-boe.byted.org"
	// mockey.Mock(contractorder.GetServerAddress).Return(DouyinPayServer).Build()
	svc := H5ApiService{Client: c}

	outTradeNo := outTradeNo
	result, err := svc.CloseOrder(ctx,
		CloseOrderRequest{
			OutTradeNo: outTradeNo,
			SubMchid:   subMchID,
			SpMchid:    mchID,
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

func PartnerH5QueryOrderByIdTest(ctx context.Context, signType string, t *testing.T) {

	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := H5ApiService{Client: c}

	transId := transId
	resp, result, err := svc.QueryOrderById(ctx,
		QueryOrderByIdRequest{
			TransactionId: transId,
			SubMchid:      subMchID,
			SpMchid:       mchID,
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

func PartnerH5QueryOrderByOutTradeNoTest(ctx context.Context, signType string, t *testing.T) {

	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := H5ApiService{Client: c}
	outTradeNo := outTradeNo
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
		QueryOrderByOutTradeNoRequest{
			OutTradeNo: outTradeNo,
			SubMchid:   subMchID,
			SpMchid:    mchID,
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

	assert.Equal(t, true, resp != nil)
}
