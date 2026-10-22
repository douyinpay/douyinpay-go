package indirectmerchant

import (
	"context"
	"log"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"
)

var (
	mchID               = "" // 商户号
	appId               = "" // 商户AppId
	merchantSerialNo    = "" // 商户证书序列号
	merchantPrivateKey  = "" // 商户私钥
	platformCertificate = "" // 平台证书
)

func TestIndirectMerchantRequest(t *testing.T) {
	t.Run("提交入驻申请", func(t *testing.T) {
		CreateApplyTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("查询入驻申请单", func(t *testing.T) {
		QueryApplyTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("撤销入驻申请单", func(t *testing.T) {
		CancelApplyTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("查询间连商户", func(t *testing.T) {
		QueryMerchantTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("修改商户资料", func(t *testing.T) {
		ModifyTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("绑定应用APPID", func(t *testing.T) {
		BindAppIDTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("查询应用APPID", func(t *testing.T) {
		QueryAppIDTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})

	t.Run("上传商户资质图片", func(t *testing.T) {
		UploadTest(context.Background(), consts.CRYPTO_TYPE_RSA, t)
	})
}

func CreateApplyTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiService{Client: c}

	resp, result, err := svc.CreateApply(ctx,
		ApplyRequest{
			OutIndirectMerchantID: "20434920032",
			OutPaymentInstID:      "4043424032",
			BusinessCode:          "biz_6120241204339424",
			SubjectInfo: &SubjectInfo{
				SubjectType: "SUBJECT_TYPE_ENTERPRISE",
				CertificateInfo: &CertificateInfo{
					CertNo:          "91110000000000000X",
					MerchantName:    "示例科技有限公司",
					CertCopy:        "media-cert-copy",
					LegalPersonName: "张三",
				},
			},
			IdentificationInfo: &IdentificationInfo{
				IdentificationType:   "IDENTIFICATION_TYPE_IDCARD",
				IdentificationName:   "张三",
				IdentificationNumber: "110101199001011234",
			},
			BusinessInfo: &BusinessInfo{
				MerchantShortName: "示例商户",
				ServicePhone:      "4001234567",
				MCC:               "1234",
			},
		},
	)

	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call CreateApply err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+s", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, len(resp.ApplymentID) > 0)
}

func QueryApplyTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiService{Client: c}

	applymentID := "NIM7696048001306091521"
	resp, result, err := svc.QueryApply(ctx,
		QueryApplyRequest{
			ApplymentID: applymentID,
		},
	)

	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call QueryApply err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+s", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, len(resp.ApplymentState) > 0)
}

func CancelApplyTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiService{Client: c}

	applymentID := "NIM7696048001306091521"
	result, err := svc.CancelApply(ctx,
		CancelRequest{
			ApplymentID: applymentID,
		},
	)
	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call CancelApply err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d,header=%+v", result.Response.StatusCode, result.Response.Header)
	}

	assert.Equal(t, true, 200 == result.Response.StatusCode || 400 == result.Response.StatusCode)
}

func QueryMerchantTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiService{Client: c}

	indirectMerchantID := "6120241204339424"
	resp, result, err := svc.QueryMerchant(ctx,
		Locator{
			IndirectMerchantID: indirectMerchantID,
		},
	)

	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call QueryMerchant err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+s", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, len(resp.IndirectMerchantID) > 0)
}

func ModifyTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiService{Client: c}

	indirectMerchantID := "6120241204339424"
	result, err := svc.Modify(ctx,
		ModifyRequest{
			Locator:           Locator{IndirectMerchantID: indirectMerchantID},
			MerchantShortName: "示例商户",
			ServicePhone:      "4001234567",
		},
	)
	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))

	if err != nil {
		// 处理错误
		log.Printf("call Modify err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d,header=%+v", result.Response.StatusCode, result.Response.Header)
	}

	assert.Equal(t, true, 200 == result.Response.StatusCode || 400 == result.Response.StatusCode)
}

func BindAppIDTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiService{Client: c}

	indirectMerchantID := "6120241204339424"
	resp, result, err := svc.BindAppID(ctx,
		BindAppRequest{
			Locator: Locator{IndirectMerchantID: indirectMerchantID},
			AppID:   appId,
		},
	)

	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call BindAppID err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+s", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, resp.AppInfo != nil)
}

func QueryAppIDTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiService{Client: c}

	indirectMerchantID := "6120241204339424"
	resp, result, err := svc.QueryAppID(ctx,
		QueryAppRequest{
			Locator: Locator{IndirectMerchantID: indirectMerchantID},
			AppID:   appId,
		},
	)

	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call QueryAppID err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+s", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, len(resp.AppInfoList) > 0)
}

func UploadTest(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	} else {
		c = services.InitClientSM2(ctx, mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)
	}
	svc := ApiService{Client: c}

	resp, result, err := svc.Upload(ctx, "license.png", "image/png", []byte("fake-image-bytes"))

	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))
	if err != nil {
		// 处理错误
		log.Printf("call Upload err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%+s", result.Response.StatusCode, utils.Json2Str(resp))
	}

	assert.Equal(t, true, len(resp.MediaID) > 0)
}
