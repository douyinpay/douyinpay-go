package cashier_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/services/cashier"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

var (
	mchID               = "" // 商户号
	appId               = "" // 商户AppId
	merchantSerialNo    = "" // 商户证书序列号
	merchantPrivateKey  = "" // 商户私钥
	platformCertificate = "" // 平台证书

	blindMobile = "" // SHA256 盲化后的手机号
)

func TestCashierRequest(t *testing.T) {
	t.Run("前置咨询-仅查询用户标签", func(t *testing.T) {
		PrePayConsultTagTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("前置咨询-查询用户标签和用户营销", func(t *testing.T) {
		PrePayConsultPromotionTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("唤起收银台信息查询-必填字段校验", func(t *testing.T) {
		QueryInvokeInfoRequiredFieldsTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("唤起收银台信息查询-响应解析", func(t *testing.T) {
		QueryInvokeInfoResponseUnmarshalTest(t)
	})
}

// PrePayConsultTagTest 仅查询用户标签（tag_code/tag_name）
func PrePayConsultTagTest(ctx context.Context, signType string, t *testing.T) {
	svc := cashier.ApiCashierService{Client: initClient(ctx, signType)}

	resp, result, err := svc.PrePayConsult(ctx,
		cashier.PrePayConsultRequest{
			Appid:           appId,
			Mchid:           mchID,
			GoodsTag:        `{"biz_scene":"pre_consult"}`,
			EncryptType:     "SHA256",
			BlindMobileList: []string{blindMobile},
		})

	if err != nil {
		// 处理错误
		log.Printf("call PrePayConsult err: %s", err)
	} else {
		// 处理返回结果
		jsonData, _ := json.MarshalIndent(resp, "", "  ")
		log.Printf("status=%d resp=%s", result.Response.StatusCode, string(jsonData))
	}

	assert.Equal(t, true, resp != nil)
}

// PrePayConsultPromotionTest 查询用户标签（tag_code/tag_name）和用户营销（operation_tip）
func PrePayConsultPromotionTest(ctx context.Context, signType string, t *testing.T) {
	svc := cashier.ApiCashierService{Client: initClient(ctx, signType)}

	resp, result, err := svc.PrePayConsult(ctx,
		cashier.PrePayConsultRequest{
			Appid:                 appId,
			Mchid:                 mchID,
			TotalAmount:           "2000",
			ProductCode:           []string{"NormalPay"},
			CommericalProductCode: "CO_PAY_APP",
			TradeType:             consts.TradeTypeApp,
			GoodsTag:              `{"biz_scene":"pre_consult"}`,
			EncryptType:           "SHA256",
			BlindMobileList:       []string{blindMobile},
		})

	if err != nil {
		// 处理错误
		log.Printf("call PrePayConsult err: %s", err)
	} else {
		// 处理返回结果
		jsonData, _ := json.MarshalIndent(resp, "", "  ")
		log.Printf("status=%d resp=%s", result.Response.StatusCode, string(jsonData))
	}

	assert.Equal(t, true, resp != nil)
}

// QueryInvokeInfoRequiredFieldsTest 校验唤起收银台信息查询的必填字段。
func QueryInvokeInfoRequiredFieldsTest(ctx context.Context, signType string, t *testing.T) {
	// 与现有收银台接入样例保持一致；必填字段校验发生在 Client 发起请求前。
	svc := cashier.ApiCashierService{Client: initClient(ctx, signType)}

	_, _, err := svc.QueryInvokeInfo(ctx, cashier.QueryInvokeInfoRequest{})
	assert.Equal(t, "QueryInvokeInfoRequest required field `Mchid` is empty", requiredFieldMessage(t, err))

	_, _, err = svc.QueryInvokeInfo(ctx, cashier.QueryInvokeInfoRequest{Mchid: "1230000109"})
	assert.Equal(t, "QueryInvokeInfoRequest required field `Prepayid` is empty", requiredFieldMessage(t, err))
}

// QueryInvokeInfoResponseUnmarshalTest 校验唤起收银台信息查询的响应解析。
func QueryInvokeInfoResponseUnmarshalTest(t *testing.T) {
	var resp cashier.QueryInvokeInfoResponse
	body := []byte(`{"invoke_url":"https://cashier.douyinpay.com/bytepay-cashdesk/bytepay-invoke"}`)

	err := json.Unmarshal(body, &resp)

	assert.Equal(t, nil, err)
	assert.Equal(t, "https://cashier.douyinpay.com/bytepay-cashdesk/bytepay-invoke", resp.InvokeURL)
}

func requiredFieldMessage(t *testing.T, err error) string {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
	return err.Error()
}

func initClient(ctx context.Context, signType string) *client.Client {
	if signType == consts.CRYPTO_TYPE_RSA {
		return services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	return services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
}
