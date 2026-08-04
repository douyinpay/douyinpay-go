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

		// 处理支付结果通知
		switch notifyReq.EventType {
		case "TRANSACTION.SUCCESS":
			// 支付成功，读取 amount、payer、success_time 等字段。
		case "TRANSACTION.FAIL":
			// 支付失败。
		}
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
		// 支付结果通知使用 pay.Transaction 解析。
		// 可通过 notifyReq.EventType 获取通知结果类型（TRANSACTION.SUCCESS / TRANSACTION.FAIL）。
		content := new(pay.Transaction)
		notifyReq, err := handler.ParseCallback(context.Background(), request, content)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 处理支付结果通知
		switch notifyReq.EventType {
		case "TRANSACTION.SUCCESS":
			// 支付成功，读取 amount、payer、success_time 等字段。
		case "TRANSACTION.FAIL":
			// 支付失败。
		}
		log.Printf("content=%+s, notifyReq=%+s", utils.Json2Str(content), utils.Json2Str(notifyReq))

		assert.Equal(t, true, notifyReq != nil && notifyReq.Resource != nil && len(notifyReq.Resource.Plaintext) > 0)
	})
}

func TestNewDeductNotifyHandler(t *testing.T) {
	t.Run("代扣结果通知", func(t *testing.T) {
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

		// 代扣结果通知使用 pay.Transaction 解析。
		content := new(pay.Transaction)
		notifyReq, err := handler.ParseCallback(context.Background(), request, content)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 处理代扣结果通知
		switch notifyReq.EventType {
		case "TRANSACTION.SUCCESS":
			// 代扣成功，读取 amount、payer、success_time 等字段。
		case "TRANSACTION.FAIL":
			// 代扣失败。
			// 可通过 content.TradeType 判断交易类型：
			//   consts.TradeTypeSinglePay（SGP）：商户代扣
			//   consts.TradeTypeNoPwdPay（NPP）：免密支付
			// 读取 contract_id、err_code、err_code_des 获取失败原因。
		}
		log.Printf("content=%+s, notifyReq=%+s", utils.Json2Str(content), utils.Json2Str(notifyReq))

		assert.Equal(t, true, notifyReq != nil && notifyReq.Resource != nil && len(notifyReq.Resource.Plaintext) > 0)
	})
}

func TestNewDeductNotifyHandlerWithAutoClient(t *testing.T) {
	t.Run("代扣结果通知", func(t *testing.T) {
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

		// 代扣结果通知使用 pay.Transaction 解析。
		content := new(pay.Transaction)
		notifyReq, err := handler.ParseCallback(context.Background(), request, content)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 处理代扣结果通知
		switch notifyReq.EventType {
		case "TRANSACTION.SUCCESS":
			// 代扣成功，读取 amount、payer、success_time 等字段。
		case "TRANSACTION.FAIL":
			// 代扣失败。
			// 可通过 content.TradeType 判断交易类型：
			//   consts.TradeTypeSinglePay（SGP）：商户代扣
			//   consts.TradeTypeNoPwdPay（NPP）：免密支付
			// 读取 contract_id、err_code、err_code_des 获取失败原因。
		}
		log.Printf("content=%+s, notifyReq=%+s", utils.Json2Str(content), utils.Json2Str(notifyReq))

		assert.Equal(t, true, notifyReq != nil && notifyReq.Resource != nil && len(notifyReq.Resource.Plaintext) > 0)
	})
}
