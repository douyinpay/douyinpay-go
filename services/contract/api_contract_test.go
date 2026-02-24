package contract

import (
	"context"
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

/*
*
测试用例
*/
func TestContract(t *testing.T) {
	ctx := context.Background()

	t.Run("contract_id 查询代扣签约协议 - success", func(t *testing.T) {
		QueryContractTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("contract_code+plan_id 查询协议", func(t *testing.T) {
		QueryContractByPlantIdAndContractCodeTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("contract_id 解除代扣签约协议", func(t *testing.T) {
		DeleteContractTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("contract_code+plan_id 删除协议", func(t *testing.T) {
		DeleteContractByPlantIdAndContractCodeTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("App预签约 -c", func(t *testing.T) {
		PreentrustWebTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("App预签约 - param,rpc error", func(t *testing.T) {
		PreentrustWebParamErrorTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("H5预签约- c", func(t *testing.T) {
		H5entrustWebTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("H5预签约- param,rpc error", func(t *testing.T) {
		H5entrustWebParamErrorTest(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
}
func QueryContractTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	}
	svc := ApiContractService{Client: c}
	req := QueryContractRequest{
		Appid:      secret.DeductAppID,
		Mchid:      secret.DeductMchID,
		ContractId: secret.DeductContractID,
	}
	resp, result, err := svc.QueryContract(ctx, req)

	log.Printf("Contractid: %s\n resp : %v", req.ContractId, resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil && resp.ContractId != "")
}

func QueryContractByPlantIdAndContractCodeTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	}
	svc := ApiContractService{Client: c}
	req := QueryContractRequest{
		Appid:        secret.DeductAppID,
		Mchid:        secret.DeductMchID,
		PlanId:       "48",
		ContractCode: "OutSignNo",
	}
	resp, result, err := svc.QueryContract(ctx, req)

	log.Printf("ContractCode: %s\n resp : %v", req.ContractCode, resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil && resp.ContractId != "")
}
func DeleteContractTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	}
	svc := ApiContractService{Client: c}
	req := DeleteContractRequest{
		Appid: secret.DeductAppID,
		Mchid: secret.DeductMchID,
		//	Contractid:                "MSN2508271037592330733690893420",
		ContractTerminationRemark: "test",
		ContractCode:              secret.DeductDeleteOutContractCode,
		PlanId:                    secret.DeductDeletePlanID,
	}
	resp, result, err := svc.DeleteContract(ctx, req)

	log.Printf("ContractCode: %s\n resp : %v", req.ContractCode, resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil && resp.ContractId != "")
}
func DeleteContractByPlantIdAndContractCodeTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	}
	svc := ApiContractService{Client: c}
	req := DeleteContractRequest{
		Appid:                     secret.DeductAppID,
		Mchid:                     secret.DeductMchID,
		PlanId:                    "48",
		ContractCode:              "OutSignNo",
		ContractTerminationRemark: "test",
	}
	resp, result, err := svc.DeleteContract(ctx, req)

	log.Printf("ContractCode: %s\n resp : %v", req.ContractCode, resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil && resp.ContractId != "")
}

func PreentrustWebTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	}
	svc := ApiContractService{Client: c}
	req := PreEntrustWebRequest{
		Appid:                  secret.DeductAppID,
		Mchid:                  secret.DeductMchID,
		OutContractCode:        secret.DeductOutContractCode,
		PlanId:                 secret.DeductPlanID,
		RequestSerial:          1000,
		ContractDisplayAccount: "张三",
		NotifyUrl:              "https://abc.com",
		ContractExt:            "{\"goods_tag\":\"{\\\"product_tag\\\":\\\"商品信息\\\",\\\"third_part_right_info\\\":\\\"{\\\\\\\"right_id\\\\\\\":\\\\\\\"123\\\\\\\",\\\\\\\"prize_type\\\\\\\":\\\\\\\"XXX\\\\\\\"}\\\"}\"}",
	}

	resp, result, err := svc.PreEntrustWeb(ctx, req)

	log.Printf("OutContractCode: %s\n resp : %v", req.OutContractCode, resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil && resp.PreEntrustWebId != "")
}
func PreentrustWebAppIdIsNullTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	}
	svc := ApiContractService{Client: c}
	req := PreEntrustWebRequest{
		Appid:                  "",
		Mchid:                  secret.DeductMchID,
		OutContractCode:        secret.DeductOutContractCode,
		PlanId:                 secret.DeductPlanID,
		RequestSerial:          1000,
		ContractDisplayAccount: "张三",
		NotifyUrl:              "https://abc.com",
		ContractExt:            "{\"goods_tag\":\"{\\\"product_tag\\\":\\\"商品信息\\\",\\\"third_part_right_info\\\":\\\"{\\\\\\\"right_id\\\\\\\":\\\\\\\"123\\\\\\\",\\\\\\\"prize_type\\\\\\\":\\\\\\\"XXX\\\\\\\"}\\\"}\"}",
	}

	resp, result, err := svc.PreEntrustWeb(ctx, req)

	log.Printf("OutContractCode: %s\n resp : %v", req.OutContractCode, resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil && resp.PreEntrustWebId != "")
}

func PreentrustWebParamErrorTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	}
	svc := ApiContractService{Client: c}
	req := PreEntrustWebRequest{
		Appid:                  secret.DeductAppID,
		Mchid:                  secret.DeductMchID,
		OutContractCode:        secret.DeductOutContractCode,
		PlanId:                 secret.DeductPlanID,
		RequestSerial:          1000,
		ContractDisplayAccount: "张三",
		NotifyUrl:              "https://abc.com",
		ContractExt:            "{\"goods_tag\":\"{\\\"product_tag\\\":\\\"商品信息\\\",\\\"third_part_right_info\\\":\\\"{\\\\\\\"right_id\\\\\\\":\\\\\\\"123\\\\\\\",\\\\\\\"prize_type\\\\\\\":\\\\\\\"XXX\\\\\\\"}\\\"}\"}",
	}

	param := req

	param.Appid = ""
	_, _, err := svc.PreEntrustWeb(ctx, param)
	assert.Equal(t, true, err != nil)

	param = req

	param.Mchid = ""
	_, _, err = svc.PreEntrustWeb(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req

	param.OutContractCode = ""
	_, _, err = svc.PreEntrustWeb(ctx, param)
	assert.Equal(t, true, err != nil)

	param = req
	param.PlanId = ""
	_, _, err = svc.PreEntrustWeb(ctx, param)
	assert.Equal(t, true, err != nil)

	param = req
	param.RequestSerial = 0
	_, _, err = svc.PreEntrustWeb(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.ContractDisplayAccount = ""
	_, _, err = svc.PreEntrustWeb(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.NotifyUrl = ""
	_, _, err = svc.PreEntrustWeb(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req

	param.OutContractCode = "cs"
	param.Appid = "123"
	_, _, err = svc.PreEntrustWeb(ctx, param)
	assert.Equal(t, true, err != nil)

	// log.Printf("OutContractCode: %s\n resp : %v", req.OutContractCode, resp)

	// if err != nil {
	// 	// 处理错误
	// 	log.Printf("call Prepay err: %s", err)
	// } else {
	// 	// 处理返回结果
	// 	log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	// }

}

func H5entrustWebTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	}
	svc := ApiContractService{Client: c}
	req := H5EntrustwebRequest{
		Appid:                  secret.DeductAppID,
		Mchid:                  secret.DeductMchID,
		OutContractCode:        secret.DeductOutContractCode,
		PlanId:                 secret.DeductPlanID,
		RequestSerial:          1000,
		ContractDisplayAccount: "张三",
		NotifyUrl:              "https://abc.com",
		ContractExt:            "{\"goods_tag\":\"{\\\"product_tag\\\":\\\"商品信息\\\",\\\"third_part_right_info\\\":\\\"{\\\\\\\"right_id\\\\\\\":\\\\\\\"123\\\\\\\",\\\\\\\"prize_type\\\\\\\":\\\\\\\"XXX\\\\\\\"}\\\"}\"}",
		Timestamp:              time.Now().String(),
		ClientIp:               "127.0.0.1",
	}
	resp, result, err := svc.H5Entrustweb(ctx, req)

	log.Printf("ContractCode: %s\n resp : %v", req.OutContractCode, resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+v", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp != nil && resp.RedirectUrl != "")
}

func H5entrustWebParamErrorTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, secret.DeductMchID, secret.DeductMerchantSerialNo, secret.DeductMerchantPrivateKey, secret.DeductPlatformCertificate)
	}
	svc := ApiContractService{Client: c}
	req := H5EntrustwebRequest{
		Appid:                  secret.DeductAppID,
		Mchid:                  secret.DeductMchID,
		OutContractCode:        secret.DeductOutContractCode,
		PlanId:                 secret.DeductPlanID,
		RequestSerial:          1000,
		ContractDisplayAccount: "张三",
		NotifyUrl:              "https://abc.com",
		ContractExt:            "{\"goods_tag\":\"{\\\"product_tag\\\":\\\"商品信息\\\",\\\"third_part_right_info\\\":\\\"{\\\\\\\"right_id\\\\\\\":\\\\\\\"123\\\\\\\",\\\\\\\"prize_type\\\\\\\":\\\\\\\"XXX\\\\\\\"}\\\"}\"}",
		Timestamp:              time.Now().String(),
		ClientIp:               "127.0.0.1",
	}

	param := req
	param.Appid = ""
	_, _, err := svc.H5Entrustweb(ctx, param)
	assert.Equal(t, true, err != nil)

	param = req
	param.Mchid = ""
	_, _, err = svc.H5Entrustweb(ctx, param)
	assert.Equal(t, true, err != nil)

	param = req
	param.OutContractCode = ""
	_, _, err = svc.H5Entrustweb(ctx, param)
	assert.Equal(t, true, err != nil)

	param = req
	param.PlanId = ""
	_, _, err = svc.H5Entrustweb(ctx, param)
	assert.Equal(t, true, err != nil)

	param = req
	param.RequestSerial = 0
	_, _, err = svc.H5Entrustweb(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.ContractDisplayAccount = ""
	_, _, err = svc.H5Entrustweb(ctx, param)
	assert.Equal(t, true, err != nil)
	param = req
	param.NotifyUrl = ""
	_, _, err = svc.H5Entrustweb(ctx, param)
	assert.Equal(t, true, err != nil)

	param = req
	param.ClientIp = ""
	_, _, err = svc.H5Entrustweb(ctx, param)
	assert.Equal(t, true, err != nil)

	param = req
	param.Timestamp = ""
	_, _, err = svc.H5Entrustweb(ctx, param)
	assert.Equal(t, true, err != nil)

	req.Mchid = "1234"
	_, _, err = svc.H5Entrustweb(ctx, req)

	assert.Equal(t, true, err != nil)
}
