package bill

import (
	"compress/gzip"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/douyinpay/douyinpay-go/client"

	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

var (
	mchID               = "" // 商户号
	appId               = "" // 商户AppId
	merchantSerialNo    = "" // 商户证书序列号
	merchantPrivateKey  = "" // 商户私钥
	platformCertificate = "" // 平台证书
)

func TestBillRequest(t *testing.T) {
	ctx := context.Background()
	billDate := "2025-06-30"
	tarType := "GZIP"
	t.Run("交易账单下载", func(t *testing.T) {
		testCommon(ctx, consts.CRYPTO_TYPE_RSA, func(ctx context.Context, ser BillApiService) (
			resp *Bill, result *client.APIResult, err error) {
			return ser.BillApply(ctx,
				BillApplyRequest{
					Mchid:    mchID,
					BillDate: billDate,
					BillType: "TRADE",
					TarType:  tarType,
				},
			)
		}, t)
	})
	t.Run("结算账单下载", func(t *testing.T) {
		testCommon(ctx, consts.CRYPTO_TYPE_RSA, func(ctx context.Context, ser BillApiService) (
			resp *Bill, result *client.APIResult, err error) {
			return ser.BillApply(ctx,
				BillApplyRequest{
					Mchid:    mchID,
					BillDate: billDate,
					BillType: "SETTLEMENT",
					TarType:  tarType,
				},
			)
		}, t)
	})
	t.Run("资金账单下载", func(t *testing.T) {
		testCommon(ctx, consts.CRYPTO_TYPE_RSA, func(ctx context.Context, ser BillApiService) (
			resp *Bill, result *client.APIResult, err error) {
			return ser.ApplyFundFlowBill(ctx,
				ApplyFundFlowBillRequest{
					Mchid:       mchID,
					BillDate:    billDate,
					AccountType: "BaseAccount",
					TarType:     tarType,
				},
			)
		}, t)
	})
	t.Run("分账账单下载", func(t *testing.T) {
		testCommon(ctx, consts.CRYPTO_TYPE_RSA, func(ctx context.Context, ser BillApiService) (
			resp *Bill, result *client.APIResult, err error) {
			return ser.ApplySplitBill(ctx,
				ApplySplitBillRequest{
					Mchid:    mchID,
					BillDate: billDate,
					TarType:  tarType,
				},
			)
		}, t)
	})
}

type ApplyFunc func(ctx context.Context, ser BillApiService) (resp *Bill, result *client.APIResult, err error)

func testCommon(ctx context.Context, signType string, applyFunc ApplyFunc, t *testing.T) {
	resp, result, err := applyFunc(ctx, initService(ctx, signType))
	assert.Equal(t, nil, err)
	assert.Equal(t, true, result != nil)
	assert.Equal(t, true, resp != nil)
	if result == nil || resp == nil {
		return
	}
	log.Printf("Request-Id:%s", result.Response.Header.Get(consts.RequestID))
	log.Printf("status=%d resp=%+v", result.Response.StatusCode, resp)
	checkSHA(resp.DownloadUrl, resp.HashValue, t)
}

func initService(ctx context.Context, signType string) BillApiService {
	if signType == consts.CRYPTO_TYPE_RSA {
		return BillApiService{Client: services.InitClientRSA(ctx,
			mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)}
	}
	return BillApiService{Client: services.InitClientSM2(ctx,
		mchID, merchantSerialNo, merchantPrivateKey, platformCertificate)}
}

func checkSHA(downloadUrl, hashVal string, t *testing.T) {
	calVal, err := getDownloadFileSHA(downloadUrl)
	assert.Equal(t, nil, err)
	assert.Equal(t, hashVal, calVal)
}

func getDownloadFileSHA(downloadUrl string) (string, error) {
	resp, err := http.Get(downloadUrl)
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", errors.New("http resp empty")
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get download url failed, status code:%d", resp.StatusCode)
	}
	defer resp.Body.Close()
	gzReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return "", err
	}
	h := sha1.New()
	_, err = io.Copy(h, gzReader)
	if err != nil {
		return "", err
	}
	hash := h.Sum(nil)
	hashStr := hex.EncodeToString(hash[:])
	return hashStr, nil
}
