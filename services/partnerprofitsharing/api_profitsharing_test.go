package partnerprofitsharing

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/services/callback"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"
)

var (
	spMchID             = "" // 服务商商户号
	spAppID             = "" // 服务商AppId
	subMchID            = "" // 特约商户号
	subAppID            = "" // 特约商户AppId
	merchantSerialNo    = "" // 商户证书序列号
	merchantPrivateKey  = "" // 商户私钥
	platformCertificate = "" // 平台证书
	transactionID       = "" // 抖音支付订单号
	outOrderNo          = "" // 商户分账单号
	orderID             = "" // 抖音支付分账单号
	outReturnNo         = "" // 商户回退单号
	returnMchID         = "" // 回退商户号
	receiverType        = "MERCHANT_ID"
	receiverAccount     = "" // 分账接收方账号
	receiverName        = "" // 分账接收方全称，需按文档要求使用平台证书公钥加密
	notifyURL           = "" // 分账通知地址
)

func TestPartnerProfitSharingRequest(t *testing.T) {
	ctx := context.Background()

	t.Run("请求分账", func(t *testing.T) {
		SplitFundTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("查询分账结果", func(t *testing.T) {
		QuerySplitFundTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("请求分账回退", func(t *testing.T) {
		ReturnSplitFundTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("查询分账回退结果", func(t *testing.T) {
		QueryReturnSplitFundTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("完结分账", func(t *testing.T) {
		FinishSplitFundTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("查询剩余待分金额", func(t *testing.T) {
		QueryUnsplitAmountTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("查询特约商户分账配置", func(t *testing.T) {
		QueryMerchantConfigTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("添加分账接收方", func(t *testing.T) {
		AddReceiverTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("删除分账接收方", func(t *testing.T) {
		DeleteReceiverTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
}

func SplitFundTest(ctx context.Context, signType string, t *testing.T) {
	svc := initService(ctx, signType)

	reqOutOrderNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.SplitFund(ctx,
		ApiPartnerSplitFundRequest{
			SpMchid:       spMchID,
			SpAppid:       spAppID,
			SubMchid:      subMchID,
			SubAppid:      subAppID,
			TransactionID: transactionID,
			OutOrderNo:    reqOutOrderNo,
			Receivers: []ApiPartnerReceiver{
				{
					Type:        receiverType,
					Account:     receiverAccount,
					Name:        receiverName,
					Amount:      1,
					Description: "分给合作方",
				},
			},
			UnfreezeUnsplit: true,
			NotifyURL:       notifyURL,
		},
	)

	logResult("SplitFund", result, resp, err)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, resp != nil)
}

func QuerySplitFundTest(ctx context.Context, signType string, t *testing.T) {
	svc := initService(ctx, signType)

	resp, result, err := svc.QuerySplitFund(ctx,
		ApiPartnerQuerySplitFundRequest{
			OutOrderNo:    outOrderNo,
			SpMchid:       spMchID,
			SubMchid:      subMchID,
			TransactionID: transactionID,
			OrderID:       orderID,
		},
	)

	logResult("QuerySplitFund", result, resp, err)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, resp != nil)
}

func ReturnSplitFundTest(ctx context.Context, signType string, t *testing.T) {
	svc := initService(ctx, signType)

	reqOutReturnNo := fmt.Sprintf("%s_%d", "RETURN", time.Now().Unix())
	resp, result, err := svc.ReturnSplitFund(ctx,
		ApiPartnerReturnSplitFundRequest{
			SpMchid:     spMchID,
			SubMchid:    subMchID,
			OrderID:     orderID,
			OutOrderNo:  outOrderNo,
			OutReturnNo: reqOutReturnNo,
			ReturnMchid: returnMchID,
			Amount:      1,
			Description: "退分账",
		},
	)

	logResult("ReturnSplitFund", result, resp, err)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, resp != nil)
}

func QueryReturnSplitFundTest(ctx context.Context, signType string, t *testing.T) {
	svc := initService(ctx, signType)

	resp, result, err := svc.QueryReturnSplitFund(ctx,
		ApiPartnerQueryReturnSplitFundRequest{
			OutReturnNo: outReturnNo,
			SpMchid:     spMchID,
			SubMchid:    subMchID,
			OutOrderNo:  outOrderNo,
		},
	)

	logResult("QueryReturnSplitFund", result, resp, err)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, resp != nil)
}

func FinishSplitFundTest(ctx context.Context, signType string, t *testing.T) {
	svc := initService(ctx, signType)

	reqOutOrderNo := fmt.Sprintf("%s_%d", "FINISH", time.Now().Unix())
	resp, result, err := svc.FinishSplitFund(ctx,
		ApiPartnerFinishSplitFundRequest{
			SpMchid:       spMchID,
			SubMchid:      subMchID,
			TransactionID: transactionID,
			OutOrderNo:    reqOutOrderNo,
			Description:   "测试商品分账",
			NotifyURL:     notifyURL,
		},
	)

	logResult("FinishSplitFund", result, resp, err)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, resp != nil)
}

func QueryUnsplitAmountTest(ctx context.Context, signType string, t *testing.T) {
	svc := initService(ctx, signType)

	resp, result, err := svc.QueryUnsplitAmount(ctx,
		QueryUnsplitAmountRequest{
			TransactionID: transactionID,
			SpMchid:       spMchID,
		},
	)

	logResult("QueryUnsplitAmount", result, resp, err)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, resp != nil)
}

func QueryMerchantConfigTest(ctx context.Context, signType string, t *testing.T) {
	svc := initService(ctx, signType)

	resp, result, err := svc.QueryMerchantConfig(ctx,
		QueryMerchantConfigRequest{
			SubMchid: subMchID,
			SpMchid:  spMchID,
		},
	)

	logResult("QueryMerchantConfig", result, resp, err)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, resp != nil)
}

func AddReceiverTest(ctx context.Context, signType string, t *testing.T) {
	svc := initService(ctx, signType)

	resp, result, err := svc.AddReceiver(ctx,
		AddReceiverRequest{
			SpMchid:      spMchID,
			SpAppid:      spAppID,
			SubMchid:     subMchID,
			SubAppid:     subAppID,
			Type:         receiverType,
			Account:      receiverAccount,
			Name:         receiverName,
			RelationType: "STORE",
		},
	)

	logResult("AddReceiver", result, resp, err)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, resp != nil)
}

func DeleteReceiverTest(ctx context.Context, signType string, t *testing.T) {
	svc := initService(ctx, signType)

	resp, result, err := svc.DeleteReceiver(ctx,
		DeleteReceiverRequest{
			SpMchid:  spMchID,
			SpAppid:  spAppID,
			SubMchid: subMchID,
			SubAppid: subAppID,
			Type:     receiverType,
			Account:  receiverAccount,
		},
	)

	logResult("DeleteReceiver", result, resp, err)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, resp != nil)
}

func ParseProfitSharingNotifyTest(ctx context.Context, handler *callback.Handler, request *http.Request, t *testing.T) {
	notifyReq, content, err := ParseProfitSharingNotify(ctx, handler, request)
	if err != nil {
		// 处理错误
		log.Printf("call ParseProfitSharingNotify err:%s", err)
		return
	}
	log.Printf("notify=%s content=%s", utils.Json2Str(notifyReq), utils.Json2Str(content))
}

func ParseReceiverNotifyTest(ctx context.Context, handler *callback.Handler, request *http.Request, t *testing.T) {
	notifyReq, content, err := ParseReceiverNotify(ctx, handler, request)
	if err != nil {
		// 处理错误
		log.Printf("call ParseReceiverNotify err:%s", err)
		return
	}
	log.Printf("notify=%s content=%s", utils.Json2Str(notifyReq), utils.Json2Str(content))
}

func initService(ctx context.Context, signType string) ApiService {
	if signType == consts.CRYPTO_TYPE_RSA {
		return ApiService{Client: services.InitClientRSA(ctx,
			spMchID, merchantSerialNo, merchantPrivateKey, platformCertificate)}
	}
	return ApiService{Client: services.InitClientSM2(ctx,
		spMchID, merchantSerialNo, merchantPrivateKey, platformCertificate)}
}

func logResult(api string, result *client.APIResult, resp interface{}, err error) {
	if result != nil && result.Response != nil {
		log.Printf("%s Request-Id:%s", api, result.Response.Header.Get(consts.RequestID))
	}
	if err != nil {
		// 处理错误
		log.Printf("call %s err:%s", api, err)
		return
	}
	if result != nil && result.Response != nil {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, utils.Json2Str(resp))
	}
}
