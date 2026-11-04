package cashier_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/douyinpay/douyinpay-go/services/cashier"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

func TestQueryInvokeInfoRequiredFields(t *testing.T) {
	ctx := context.Background()
	// 与现有收银台接入样例保持一致；必填字段校验发生在 Client 发起请求前。
	svc := cashier.ApiCashierService{Client: initClient(ctx, consts.CRYPTO_TYPE_RSA)}

	_, _, err := svc.QueryInvokeInfo(ctx, cashier.QueryInvokeInfoRequest{})
	assert.Equal(t, "QueryInvokeInfoRequest required field `Mchid` is empty", requiredFieldMessage(t, err))

	_, _, err = svc.QueryInvokeInfo(ctx, cashier.QueryInvokeInfoRequest{Mchid: "1230000109"})
	assert.Equal(t, "QueryInvokeInfoRequest required field `Prepayid` is empty", requiredFieldMessage(t, err))
}

func TestQueryInvokeInfoResponseUnmarshal(t *testing.T) {
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
