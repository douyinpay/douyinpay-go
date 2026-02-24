package certificates

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

var (
	MchID               = ""
	MerchantSerialNo    = ""
	MerchantPrivateKey  = ""
	PlatformCertificate = ""

	encryptKey = ""
)

/*
*
测试用例
*/
func TestDownloadCertificates(t *testing.T) {
	ctx := context.Background()

	t.Run("下载RSA平台证书列-RSA", func(t *testing.T) {
		DownloadCertificates(ctx, consts.CRYPTO_TYPE_RSA, t)
	})
	// t.Run("下载SM2平台证书列表 -SM2", func(t *testing.T) {
	// 	DownloadCertificates(ctx, consts.CRYPTO_TYPE_SM2, t)
	// })
}

func DownloadCertificates(ctx context.Context, signType string, t *testing.T) {
	var c *client.Client
	if signType == "RSA" {
		c = services.InitClientRSA(ctx, MchID, MerchantSerialNo, MerchantPrivateKey, PlatformCertificate)
	} else {
		c = services.InitClientSM2(ctx, MchID, MerchantSerialNo, MerchantPrivateKey, PlatformCertificate)
	}
	svc := ApiCertificatesService{Client: c}
	resp, result, err := svc.DownloadCertificates(ctx)

	log.Printf("resp : %v", resp)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err: %s", err)
	} else {
		// 处理返回结果
		jsonData, _ := json.MarshalIndent(resp.Certificates, "", "  ")
		log.Printf("status=%d resp=%s", result.Response.StatusCode, string(jsonData))
	}

	assert.Equal(t, true, resp != nil && resp.Certificates != nil)
}
