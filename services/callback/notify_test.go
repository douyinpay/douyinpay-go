package callback_test

import (
	"bytes"
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/douyinpay/douyinpay-go/services/callback"
	"github.com/douyinpay/douyinpay-go/services/pay"
	"github.com/douyinpay/douyinpay-go/tools/auth/verifiers"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/tools/downloader"
	"github.com/douyinpay/douyinpay-go/utils"
)

const (
	platformCertificate = "" // 平台证书
)

func TestNewNotifyHandler(t *testing.T) {
	t.Run("通知", func(t *testing.T) {
		certificate, err := utils.LoadCertificate(platformCertificate)
		if err != nil {
			return
		}

		var handler = callback.NewNotifyHandler("", &verifiers.SHA256WithRSAVerifier{
			PublicKey:    certificate.PublicKey.(*rsa.PublicKey),
			SerialNumber: utils.ConvertSerailNo(certificate.SerialNumber),
		})
		var timestamp string
		var nonce string
		var signature string
		var body string
		var request = &http.Request{
			Header: map[string][]string{
				consts.DouyinPaySerial:    []string{utils.ConvertSerailNo(certificate.SerialNumber)},
				consts.DouyinPayTimestamp: []string{timestamp},
				consts.DouyinPayNonce:     []string{nonce},
				consts.DouyinPaySignature: []string{signature},
			},
			Body: ioutil.NopCloser(bytes.NewBufferString(body)),
		}

		content := new(pay.Transaction)
		notifyReq, err := handler.ParseCallback(context.Background(), request, content)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 处理通知内容
		log.Printf("content=%+s, notifyReq=%+s", utils.Json2Str(content), utils.Json2Str(notifyReq))

		assert.Equal(t, true, notifyReq != nil && notifyReq.Resource != nil && len(notifyReq.Resource.Plaintext) > 0)
	})
}

func TestNewNotifyHandlerWithAutoClient(t *testing.T) {
	t.Run("通知", func(t *testing.T) {
		certificate, err := utils.LoadCertificate(platformCertificate)
		if err != nil {
			return
		}
		mgr := downloader.MgrInstance()
		mchID := ""
		//err := mgr.RegisterSM2DownloaderWithPrivateKey(context.Background(), mchID,
		//	mchCertificateSerialNo, privateKey, encryptKey)
		// 从已经注册好下载器的Mgr中取出 verifier
		var handler = callback.NewNotifyHandler("", &verifiers.SHA256WithRSAVerifierWithGetter{
			CertGetter: mgr.GetCertificateVisitor(mchID),
		})
		var timestamp string
		var nonce string
		var signature string
		var body string
		var request = &http.Request{
			Header: map[string][]string{
				consts.DouyinPaySerial:    []string{utils.ConvertSerailNo(certificate.SerialNumber)},
				consts.DouyinPayTimestamp: []string{timestamp},
				consts.DouyinPayNonce:     []string{nonce},
				consts.DouyinPaySignature: []string{signature},
			},
			Body: ioutil.NopCloser(bytes.NewBufferString(body)),
		}

		content := new(pay.Transaction)
		notifyReq, err := handler.ParseCallback(context.Background(), request, content)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 处理通知内容
		log.Printf("content=%+s, notifyReq=%+s", utils.Json2Str(content), utils.Json2Str(notifyReq))

		assert.Equal(t, true, notifyReq != nil && notifyReq.Resource != nil && len(notifyReq.Resource.Plaintext) > 0)
	})
}
