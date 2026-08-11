package contract

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/services/secret"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"
	"github.com/smartystreets/goconvey/convey"
)

func TestApiPartnerContractService_PartnerContractOrder(t *testing.T) {
	ctx := context.Background()

	t.Run("支付中签约下单-a", func(t *testing.T) {
		PartnerContractOrderTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("申请扣款-a", func(t *testing.T) {
		PartnerPayApplyTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

}

func PartnerContractOrderTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.MchID, secret.MerchantSerialNo, secret.MerchantPrivateKey, secret.PlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.MchID, secret.MerchantSerialNo, secret.MerchantPrivateKey, secret.PlatformCertificate)
	}
	outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())
	// info
	contractInfo := &ContractInfo{
		ContractMchId:          secret.MchID,
		ContractAppId:          secret.LocalAppID,
		PlanId:                 secret.LocalPlanIdStr,
		OutContractCode:        secret.LocalOutContractCode,
		RequestSerial:          1756379139,
		ContractDisplayAccount: "测试账号",
		ContractNotifyUrl:      "",
	}
	contractInfo2 := &ContractInfo{
		ContractMchId:          secret.MchID,
		ContractAppId:          secret.LocalAppID,
		PlanId:                 secret.LocalPlanIdStr,
		OutContractCode:        secret.LocalOutContractCode,
		RequestSerial:          1756379139,
		ContractDisplayAccount: "测试账号",
		ContractNotifyUrl:      "",
	}
	//jsonToObj
	req := PartnerContractOrderRequest{
		SpAppid:     secret.LocalAppID,
		SpMchid:     secret.MchID,
		SubAppid:    secret.LocalAppID,
		SubMchid:    secret.SubMchID,
		TimeExpire:  "2025-12-05T22:43:00+08:00",
		TradeType:   "APP",
		Description: "CS",
		NotifyUrl:   "https://www.mock.com",
		OutTradeNo:  outTradeNo,
		//TradeType:   "JSAPI",
		Amount: &Amount{
			Currency: "CNY",
			Total:    100,
		},
		ContractInfo: contractInfo2,
		Payer: &Payer{
			SpOpenid:  secret.LocalOpenID,
			SubOpenid: secret.LocalSubOpenID,
		},
	}
	svc := ApiPartnerContractService{Client: c}
	param := req
	param.SpAppid = ""
	_, _, err := svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.SpMchid = ""
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.SubMchid = ""
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.SubAppid = ""
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.TimeExpire = ""
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.NotifyUrl = ""
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.TradeType = ""
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.Description = ""
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.Amount.Total = -1
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.ContractInfo.ContractMchId = ""
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	contractInfo2.ContractMchId = contractInfo.ContractMchId
	param = req
	param.ContractInfo.PlanId = ""
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	contractInfo2.PlanId = contractInfo.PlanId
	param = req
	param.ContractInfo.OutContractCode = ""
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	contractInfo2.OutContractCode = contractInfo.OutContractCode
	param = req
	param.ContractInfo.RequestSerial = -1
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	contractInfo2.RequestSerial = contractInfo.RequestSerial
	param = req
	param.ContractInfo.ContractDisplayAccount = ""
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	contractInfo2.ContractDisplayAccount = contractInfo.ContractDisplayAccount
	param = req
	param.ContractInfo.ContractNotifyUrl = ""
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)
	contractInfo2.ContractNotifyUrl = contractInfo.ContractNotifyUrl

	param = req
	param.OutTradeNo = "1s"
	_, _, err = svc.PartnerContractOrder(ctx, param)
	convey.ShouldNotBeNil(t, err)

	resp, result, err := svc.PartnerContractOrder(ctx, req)

	log.Printf("OutTradeNo: %s\n resp : %v", req.OutTradeNo, resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	convey.ShouldNotBeNil(t, resp != nil && resp.PrepayId != "")
}

func PartnerPayApplyTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.MchID, secret.MerchantSerialNo, secret.MerchantPrivateKey, secret.PlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.MchID, secret.MerchantSerialNo, secret.MerchantPrivateKey, secret.PlatformCertificate)
	}
	//jsonToObj
	goodsArr := []GoodsDetail{
		{
			DouyinpayGoodsId: "",
			GoodsName:        "测试商品",
			MerchantGoodsId:  "ABC",
			Quantity:         1,
			UnitPrice:        828800,
		},
	}
	outTradeNo := fmt.Sprintf("%s_%d", "OUT", time.Now().Unix())

	req := PartnerPayApplyRequest{
		SpAppid:    secret.LocalAppID,
		SpMchid:    secret.MchID,
		SubAppid:   secret.LocalAppID,
		SubMchid:   secret.SubMchID,
		TimeExpire: "2025-09-29T19:06:08+08:00",
		OutTradeNo: outTradeNo,
		//ContractId:  "MSN2508281905379027263142846427",
		ContractId: secret.LocalContractID,
		//ContractId:  "MSN2502281043179010865055165676",
		TradeType:   "SGP",
		Description: "测试商品",
		NotifyUrl:   "",
		Attach:      "",
		Detail: &PayApplyDetail{
			CostPrice:   608800,
			GoodsDetail: goodsArr,
			InvoiceId:   "dy123",
		},
		Amount: &Amount{
			Currency: "CNY",
			Total:    10,
		},
		SceneInfo: &SceneInfo{
			DeviceId:      "013467007045764",
			PayerClientIp: "14.23.150.211",
			PayerDeviceId: "",
		},
		SettleInfo: &SettleInfo{
			ProfitSharing: false,
		},
	}
	svc := ApiPartnerContractService{Client: c}

	param := req
	param.SpAppid = ""
	_, _, err := svc.PartnerPayApply(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.SpMchid = ""
	_, _, err = svc.PartnerPayApply(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.SubAppid = ""
	_, _, err = svc.PartnerPayApply(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.SubMchid = ""
	_, _, err = svc.PartnerPayApply(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.TimeExpire = ""
	_, _, err = svc.PartnerPayApply(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.NotifyUrl = ""
	_, _, err = svc.PartnerPayApply(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.TradeType = ""
	_, _, err = svc.PartnerPayApply(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.Description = ""
	_, _, err = svc.PartnerPayApply(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.Amount.Total = -1
	_, _, err = svc.PartnerPayApply(ctx, param)
	convey.ShouldNotBeNil(t, err)
	param = req
	param.ContractId = ""
	_, _, err = svc.PartnerPayApply(ctx, param)
	convey.ShouldNotBeNil(t, err)

	param = req
	param.ContractId = "123"
	_, _, err = svc.PartnerPayApply(ctx, param)
	convey.ShouldNotBeNil(t, err)

	resp, result, err := svc.PartnerPayApply(ctx, req)

	log.Printf("OutTradeNo: %s\n resp : %v", req.OutTradeNo, resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	convey.ShouldBeTrue(t, resp != nil && resp.ResultCode == "SUCCESS")
}
