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
	spAppId             = "" // 服务商应用ID
	spMchId             = "" // 服务商商户号
	subAppId            = "" // 子商户应用ID
	subMchId            = "" // 子商户商户号
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
	t.Run("服务商创建服务订单", func(t *testing.T) {
		PartnerCreateServiceOrderTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("服务商完结服务订单", func(t *testing.T) {
		PartnerCompleteServiceOrderTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("服务商查询服务订单", func(t *testing.T) {
		PartnerQueryServiceOrderTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("服务商取消服务订单", func(t *testing.T) {
		PartnerCancelServiceOrderTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("服务商同步服务订单信息", func(t *testing.T) {
		PartnerSynchronizeServiceOrderInfoTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("服务商修改订单金额", func(t *testing.T) {
		PartnerModifyAmountTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("服务商申请服务授权", func(t *testing.T) {
		PartnerCreditSrvSignApplyTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("服务商查询用户授权记录", func(t *testing.T) {
		PartnerCreditSrvSignQueryTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	t.Run("服务商解除用户授权关系", func(t *testing.T) {
		PartnerCloseCreditServiceTest(ctx, consts.CRYPTO_TYPE_RSA, t)
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

func PartnerCreateServiceOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}
	resp, result, err := svc.CreateServiceOrderForSP(ctx, ApiPartnerCreateServiceOrderRequest{
		SpAppid:             spAppId,
		SpMchid:             spMchId,
		SubAppid:            subAppId,
		SubMchid:            subMchId,
		OutOrderNo:          "OUT_1666688488",
		ServiceId:           serviceId,
		ServiceIntroduction: "某某酒店",
		AuthorizationCode:   "AUTH_1666688488",
		NotifyUrl:           "https://callback.example.com/payscore",
		RiskFund:            &RiskFund{Name: "ESTIMATE_ORDER_COST", Amount: 10000},
	})
	log.Printf("outOrderNo: OUT_1666688488\nRequest-Id: %s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		log.Printf("call CreateServiceOrderForSP err: %v", err)
	} else {
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, resp != nil)
}

func PartnerCompleteServiceOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}
	resp, result, err := svc.CompleteServiceOrderForSP(ctx, ApiPartnerCompleteServiceOrderRequest{
		SpAppid:     spAppId,
		SpMchid:     spMchId,
		SubMchid:    subMchId,
		OutOrderNo:  "OUT_1666688488",
		ServiceId:   serviceId,
		TotalAmount: 10000,
		ChannelInfo: &ChannelInfo{
			PresetChannel: []*PresetChannel{
				{
					ChannelCode:   "OUTSIDE_MC",
					ChannelId:     "M2025042914432001250054700",
					ChannelAmount: 100,
				},
			},
		},
	})
	log.Printf("outOrderNo: OUT_1666688488\nRequest-Id: %s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		log.Printf("call CompleteServiceOrderForSP err: %v", err)
	} else {
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, resp != nil)
}

func PartnerQueryServiceOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}
	resp, result, err := svc.QueryServiceOrderForSP(ctx, ApiPartnerQueryServiceOrderRequest{
		SpAppid:    spAppId,
		SpMchid:    spMchId,
		SubAppid:   subAppId,
		SubMchid:   subMchId,
		OutOrderNo: "OUT_1",
		ServiceId:  serviceId,
	})
	log.Printf("outOrderNo: OUT_1\nRequest-Id: %s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		log.Printf("call QueryServiceOrderForSP err: %v", err)
	} else {
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, resp != nil)
}

func PartnerCancelServiceOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}
	resp, result, err := svc.CancelServiceOrderForSP(ctx, ApiPartnerCancelServiceOrderRequest{
		SpAppid:    spAppId,
		SpMchid:    spMchId,
		SubMchid:   subMchId,
		OutOrderNo: "OUT_1666688488",
		ServiceId:  serviceId,
		Reason:     "用户取消",
	})
	log.Printf("outOrderNo: OUT_1666688488\nRequest-Id: %s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		log.Printf("call CancelServiceOrderForSP err: %v", err)
	} else {
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, resp != nil)
}

func PartnerSynchronizeServiceOrderInfoTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}
	resp, result, err := svc.SynchronizeServiceOrderInfoForSP(ctx, ApiPartnerSynchronizeServiceOrderInfoRequest{
		SpAppid:    spAppId,
		SpMchid:    spMchId,
		SubMchid:   subMchId,
		OutOrderNo: "OUT_1666688488",
		ServiceId:  serviceId,
		Type:       "ORDER_PAID",
		Detail: struct {
			PaidTime string `json:"paid_time"`
		}{
			PaidTime: "20220806091010",
		},
	})
	log.Printf("outOrderNo: OUT_1666688488\nRequest-Id: %s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		log.Printf("call SynchronizeServiceOrderInfoForSP err: %v", err)
	} else {
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, resp != nil)
}

func PartnerModifyAmountTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}
	resp, result, err := svc.ModifyAmountForSP(ctx, ApiPartnerModifyAmountRequest{
		SpAppid:     spAppId,
		SpMchid:     spMchId,
		SubMchid:    subMchId,
		OutOrderNo:  "OUT_1666688488",
		ServiceId:   serviceId,
		TotalAmount: 9000,
		Reason:      "修改原因",
	})
	log.Printf("outOrderNo: OUT_1666688488\nRequest-Id: %s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		log.Printf("call ModifyAmountForSP err: %v", err)
	} else {
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, resp != nil)
}

func PartnerCreditSrvSignApplyTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}
	resp, result, err := svc.CreditSrvSignApplyForSP(ctx, ApiPartnerCreditSrvSignApplyRequest{
		SpAppid:           spAppId,
		SpMchid:           spMchId,
		SubAppid:          subAppId,
		SubMchid:          subMchId,
		ServiceId:         serviceId,
		AuthorizationCode: "AUTH_1",
		NotifyUrl:         "https://callback.example.com/payscore",
	})
	log.Printf("authorizationCode: AUTH_1\nRequest-Id: %s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		log.Printf("call CreditSrvSignApplyForSP err: %v", err)
	} else {
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, resp != nil)
}

func PartnerCreditSrvSignQueryTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}
	resp, result, err := svc.CreditSrvSignQueryForSP(ctx, ApiPartnerCreditSrvSignQueryRequest{
		SpMchid:           spMchId,
		SubMchid:          subMchId,
		ServiceId:         serviceId,
		AuthorizationCode: "AUTH_1",
	})
	log.Printf("authorizationCode: AUTH_1\nRequest-Id: %s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		log.Printf("call CreditSrvSignQueryForSP err: %v", err)
	} else {
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, resp != nil)
}

func PartnerCloseCreditServiceTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, spMchId, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiPayScoreService{Client: c}
	resp, result, err := svc.CloseCreditServiceForSP(ctx, ApiPartnerCloseCreditServiceRequest{
		SpAppid:           spAppId,
		SpMchid:           spMchId,
		SubAppid:          subAppId,
		SubMchid:          subMchId,
		ServiceId:         serviceId,
		AuthorizationCode: "AUTH_1",
		Reason:            "用户取消授权",
	})
	log.Printf("authorizationCode: AUTH_1\nRequest-Id: %s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		log.Printf("call CloseCreditServiceForSP err: %v", err)
	} else {
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}
	assert.Equal(t, true, resp != nil)
}
