package partnercontract

import (
	"context"
	"log"
	"testing"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"

	"github.com/douyinpay/douyinpay-go/services/secret"
	"github.com/smartystreets/goconvey/convey"
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

func TestPartnerContract(t *testing.T) {
	ctx := context.Background()

	t.Run("查询签约关系-c", func(t *testing.T) {
		PartnerQueryContractRelaTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("协议解约-c", func(t *testing.T) {
		PartnerTerminateContractTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

}

func PartnerQueryContractRelaTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, MchID, MerchantSerialNo, MerchantPrivateKey, PlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, MchID, MerchantSerialNo, MerchantPrivateKey, PlatformCertificate)
	}

	req := PartnerQueryContractRequest{
		OutContractCode: LocalOutContractCode2,
		PlanId:          LocalPlanID,
		SpMchid:         MchID,
		SubMchid:        SubMchID,
	}

	svc := ApiPartnerContractOrderService{Client: c}

	param := req
	param.OutContractCode = ""
	_, _, err := svc.PartnerQueryContract(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.PlanId = -1
	_, _, err = svc.PartnerQueryContract(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.SpMchid = ""
	_, _, err = svc.PartnerQueryContract(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.SubMchid = ""
	_, _, err = svc.PartnerQueryContract(ctx, param)
	convey.ShouldNotBeNil(t, err)

	param = req
	param.OutContractCode = "12"
	_, _, err = svc.PartnerQueryContract(ctx, param)
	convey.ShouldNotBeNil(t, err)

	resp, result, err := svc.PartnerQueryContract(ctx, req)

	log.Printf("OutContractCode: %s\n resp : %v", req.OutContractCode, resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	convey.ShouldBeTrue(t, resp != nil && resp.ContractId != "")
}

func PartnerTerminateContractTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, MchID, MerchantSerialNo, MerchantPrivateKey, PlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, MchID, MerchantSerialNo, MerchantPrivateKey, PlatformCertificate)
	}

	//jsonToObj
	req := PartnerTerminateContractRequest{
		OutContractCode:           LocalOutContractCode2,
		PlanId:                    LocalPlanID,
		SpMchid:                   MchID,
		SubMchid:                  SubMchID,
		ContractTerminationRemark: "测试解约",
	}
	svc := ApiPartnerContractOrderService{Client: c}

	param := req
	param.OutContractCode = ""
	_, _, err := svc.PartnerTerminateContract(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.PlanId = -1
	_, _, err = svc.PartnerTerminateContract(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.SpMchid = ""
	_, _, err = svc.PartnerTerminateContract(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.SubMchid = ""
	_, _, err = svc.PartnerTerminateContract(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.ContractTerminationRemark = ""
	_, _, err = svc.PartnerTerminateContract(ctx, param)
	convey.ShouldNotBeNil(t, err)

	param = req
	param.PlanId = -1
	_, _, err = svc.PartnerTerminateContract(ctx, param)
	convey.ShouldNotBeNil(t, err)

	resp, result, err := svc.PartnerTerminateContract(ctx, req)

	log.Printf("requestBody: %v\n resp : %v \n logId: %s", req, resp, result.Response.Header.Get("Log-Id"))

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	convey.ShouldBeTrue(t, true, resp != nil && resp.ContractTerminateInfo != nil)
}
