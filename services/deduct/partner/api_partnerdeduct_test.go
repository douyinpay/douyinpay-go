package partner

import (
	"context"
	"log"
	"testing"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	deduct "github.com/douyinpay/douyinpay-go/services/deduct"
	"github.com/douyinpay/douyinpay-go/services/secret"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"

	"github.com/smartystreets/goconvey/convey"
)

/*
*
测试用例
*/
func TestPartnerdeduct(t *testing.T) {
	ctx := context.Background()

	t.Run("预约扣费-a", func(t *testing.T) {
		PartnerContractScheduleTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("预约扣费结果查询-a", func(t *testing.T) {
		PartnerContractScheduleQueryTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

}
func PartnerContractScheduleQueryTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.MchID, secret.MerchantSerialNo, secret.MerchantPrivateKey, secret.PlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.MchID, secret.MerchantSerialNo, secret.MerchantPrivateKey, secret.PlatformCertificate)
	}

	svc := ApiPartnerDeductService{Client: c}
	req := deduct.PartnerContractScheduleQueryRequest{
		SpMchid:    secret.MchID,
		SubMchid:   secret.SubMchID,
		ContractId: secret.LocalContractIDQuery,
	}

	param := req
	param.ContractId = ""
	_, _, err := svc.PartnerContractScheduleQuery(ctx, req)
	convey.ShouldNotBeNil(t, err)

	param = req
	param.SpMchid = ""
	_, _, err = svc.PartnerContractScheduleQuery(ctx, param)
	convey.ShouldNotBeNil(t, err)

	param = req
	param.SubMchid = ""
	_, _, err = svc.PartnerContractScheduleQuery(ctx, param)
	convey.ShouldNotBeNil(t, err)

	param = req
	param.ContractId = ""
	_, _, err = svc.PartnerContractScheduleQuery(ctx, param)
	convey.ShouldNotBeNil(t, err)

	param = req
	param.ContractId = "sfsfs"
	_, _, err = svc.PartnerContractScheduleQuery(ctx, param)
	convey.ShouldNotBeNil(t, err)

	resp, result, err := svc.PartnerContractScheduleQuery(ctx, req)

	log.Printf("Contractid: %s\n resp : %v", req.ContractId, resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	convey.ShouldBeNil(t, true, resp != nil && resp.ScheduleState != "")
}

func PartnerContractScheduleTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	// 修改为测试用的值

	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.MchID, secret.MerchantSerialNo, secret.MerchantPrivateKey, secret.PlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.MchID, secret.MerchantSerialNo, secret.MerchantPrivateKey, secret.PlatformCertificate)
	}

	//jsonToObj
	req := deduct.PartnerContractScheduleRequest{
		//Contractid: "MSN2509012115599031543301564688",
		//
		ContractId: secret.LocalContractIDQuery,
		SpMchid:    secret.MchID,
		SubMchid:   secret.SubMchID,
		ScheduleAmount: deduct.Amount{
			Currency: "CNY",
			Total:    2,
		},
	}

	svc := ApiPartnerDeductService{Client: c}
	param := req
	param.ContractId = ""
	_, _, err := svc.PartnerContractSchedule(ctx, param)
	convey.ShouldNotBeNil(t, err)

	param = req
	param.SpMchid = ""
	_, _, err = svc.PartnerContractSchedule(ctx, param)
	convey.ShouldNotBeNil(t, err)

	param = req
	param.SubMchid = ""
	_, _, err = svc.PartnerContractSchedule(ctx, param)
	convey.ShouldNotBeNil(t, err)

	param = req
	param.ScheduleAmount.Currency = ""
	_, _, err = svc.PartnerContractSchedule(ctx, param)
	convey.ShouldNotBeNil(t, err)

	param = req
	param.ScheduleAmount.Total = -1
	_, _, err = svc.PartnerContractSchedule(ctx, param)
	convey.ShouldNotBeNil(t, err)

	param = req
	param.ContractId = "s123"
	_, _, err = svc.PartnerContractSchedule(ctx, param)
	convey.ShouldNotBeNil(t, err)

	resp, result, err := svc.PartnerContractSchedule(ctx, req)

	log.Printf("Contractid: %s\n resp : %v", req.ContractId, resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	convey.ShouldBeNil(t, true, resp != nil && resp.DeductEndDate != "")
}
