package contractorder

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
	serviceId           = "" // 服务id
	merchantSerialNo    = "" // 商户证书序列号
	merchantPrivateKey  = "" // 商户私钥
	platformCertificate = "" // 平台证书
)

func TestCreateAndCompleteServiceOrder(t *testing.T) {
	ctx := context.Background()

	t.Run("创建服务订单", func(t *testing.T) {
		CreateServiceOrderTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("完结服务订单", func(t *testing.T) {
		CompleteServiceOrderTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("修改订单金额", func(t *testing.T) {
		ModifyServiceAmountTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("取消服务订单", func(t *testing.T) {
		CancelServiceOrderTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("同步服务订单信息", func(t *testing.T) {
		SynchronizeServiceOrderInfoTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("商户发起催收扣款", func(t *testing.T) {
		ServiceOrderPayTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("查询服务订单", func(t *testing.T) {
		QueryServiceOrderTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("申请服务授权", func(t *testing.T) {
		CreditSrvSignApplyTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("查询用户授权记录", func(t *testing.T) {
		CreditSrvSignQueryTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("解除用户授权关系", func(t *testing.T) {
		CloseCreditServiceTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
}

func QueryServiceOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}
	outOrderNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.QueryServiceOrder(ctx, ApiQueryServiceOrderRequest{
		ServiceId:  serviceId,
		Mchid:      mchID,
		OutOrderNo: outOrderNo,
		Appid:      appId,
	})
	log.Printf("outOrderNo:%s\nRequest-Id:%s", outOrderNo, result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call QueryServiceOrder err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+s", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, resp != nil && resp.OrderId != "")
}

func CreditSrvSignApplyTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}
	authorizationCode := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.CreditSrvSignApply(ctx, ApiCreditSrvSignApplyRequest{
		Appid:             appId,
		Mchid:             mchID,
		ServiceId:         serviceId,
		AuthorizationCode: authorizationCode,
		NotifyUrl:         "https://www.mock.douyinpay.com",
		GoodsTag:          "{}",
		ExtInfo:           "{}"})
	log.Printf("authorizationCode:%s\nRequest-Id:%s", authorizationCode, result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call CreditSrvSignApply err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+s", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, resp != nil && resp.PayscoreApplyToken != "")

}

func CreateServiceOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}
	outOrderNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.CreateServiceOrder(ctx, ApiCreateServiceOrderRequest{
		Appid:               appId,
		Mchid:               mchID,
		OutOrderNo:          outOrderNo,
		ServiceId:           serviceId,
		ServiceIntroduction: "抖音先享后付测试",
		NotifyUrl:           "https://www.mock.douyinpay.com",
		Attach:              "自定义数据",
		GoodsTag:            "{}",
		OpenId:              "_000LKnU79LxPwYIYZ1tClr1duYm4",
		RiskFund: &RiskFund{
			Name:        "ESTIMATE_ORDER_COST",
			Amount:      1,
			Description: "预估订单费用",
		},
		PostPayments: []*PostItem{
			{
				Name:        "测试",
				Amount:      1,
				Description: "测试",
				Count:       1,
			},
		},
		PostDiscounts: nil,
		TimeRange: &TimeRange{
			StartTime:       "20220806091010",
			StartTimeRemark: "开始时间",
			EndTime:         "20220806092020",
			EndTimeRemark:   "结束时间",
		},
		Location: &Location{
			StartLocation: "北京市海淀区北三环西路43号",
			EndLocation:   "北京市海淀区北三环西路43号",
		},
		SceneInfo: &SceneInfo{
			ClientIp: "14.23.150.211",
			DeviceId: "013467007045764",
		},
	})
	log.Printf("outOrderNo:%s\nRequest-Id:%s", outOrderNo, result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call CreateServiceOrder err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+s", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, resp != nil && resp.OrderId != "")
}

func SynchronizeServiceOrderInfoTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}

	outOrderNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.SynchronizeServiceOrderInfo(ctx,
		ApiSynchronizeServiceOrderInfoRequest{
			Appid:      appId,
			Mchid:      mchID,
			OutOrderNo: outOrderNo,
			ServiceId:  serviceId,
			Type:       "ORDER_PAID",
			Detail: struct {
				PaidTime string `json:"paid_time"`
			}{
				PaidTime: "20220806091010",
			},
		},
	)

	log.Printf("outOrderNo: %s\nRequest-Id: %s", outOrderNo, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call CreateAndCompleteServiceOrder err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil && resp.OrderId != "")
}

func CompleteServiceOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}

	outOrderNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.CompleteServiceOrder(ctx,
		ApiCompleteServiceOrderRequest{
			Appid:       appId,
			Mchid:       mchID,
			OutOrderNo:  outOrderNo,
			ServiceId:   serviceId,
			TotalAmount: 1,
			Attach:      "自定义数据",
			GoodsTag:    "",
			PostPayments: []*PostItem{
				{
					Name:        "测试",
					Amount:      1,
					Description: "测试",
					Count:       1,
				},
			},
			PostDiscounts: nil,
			TimeRange: &TimeRange{
				StartTime:       "20220806091010",
				StartTimeRemark: "开始时间",
				EndTime:         "20220806092020",
				EndTimeRemark:   "结束时间",
			},
			Location: &Location{
				StartLocation: "北京市海淀区北三环西路43号",
				EndLocation:   "北京市海淀区北三环西路43号",
			},
			SceneInfo: &SceneInfo{
				ClientIp: "14.23.150.211",
				DeviceId: "013467007045764",
			},
		},
	)

	log.Printf("outOrderNo: %s\nRequest-Id: %s", outOrderNo, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call CompleteServiceOrder err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil)
}

func ModifyServiceAmountTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}

	outOrderNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.ModifyAmount(ctx,
		ApiModifyAmountRequest{
			Appid:       appId,
			Mchid:       mchID,
			OutOrderNo:  outOrderNo,
			ServiceId:   serviceId,
			TotalAmount: 1,
			PostPayments: []*PostItem{
				{
					Name:        "测试",
					Amount:      1,
					Description: "测试",
					Count:       1,
				},
			},
			Reason: "测试",
		},
	)

	log.Printf("outOrderNo: %s\nRequest-Id: %s", outOrderNo, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call ModifyServiceOrder err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil)
}

func CancelServiceOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}

	outOrderNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.CancelServiceOrder(ctx,
		ApiCancelServiceOrderRequest{
			Appid:      appId,
			Mchid:      mchID,
			OutOrderNo: outOrderNo,
			ServiceId:  serviceId,
			Reason:     "测试",
		},
	)

	log.Printf("outOrderNo: %s\nRequest-Id: %s", outOrderNo, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call CancelServiceOrder err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil)
}

func CreditSrvSignQueryTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}

	authorizationCode := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.CreditSrvSignQuery(ctx,
		ApiCreditSrvSignQueryRequest{
			Mchid:             mchID,
			AuthorizationCode: authorizationCode,
			ServiceId:         serviceId,
		},
	)

	log.Printf("authorizationCode: %s\nRequest-Id: %s", authorizationCode, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call CancelServiceOrder err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil)

}

func ServiceOrderPayTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}

	outOrderNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.ServiceOrderPay(ctx,
		ApiServiceOrderPayRequest{
			Appid:      appId,
			Mchid:      mchID,
			OutOrderNo: outOrderNo,
			ServiceId:  serviceId,
		},
	)

	log.Printf("outOrderNo: %s\nRequest-Id: %s", outOrderNo, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call CancelServiceOrder err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil)
}

func CloseCreditServiceTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}

	authorizationCode := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	resp, result, err := svc.CloseCreditService(ctx,
		ApiCloseCreditServiceRequest{
			Appid:             appId,
			Mchid:             mchID,
			ServiceId:         serviceId,
			AuthorizationCode: authorizationCode,
			Reason:            "测试",
		},
	)

	log.Printf("authorizationCode: %s\nRequest-Id: %s", authorizationCode, result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call CloseCreditService err: %v", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil)
}
