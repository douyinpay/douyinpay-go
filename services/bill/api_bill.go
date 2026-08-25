package bill

// 官方文档：
// - 申请交易账单：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/667e6bc444a74902ead102ad
// - 申请资金账单：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/667e69daa998e00305dcec48
// - 申请分账账单：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/68355a92994b190515a2af86

import (
	"context"
	nethttp "net/http"
	neturl "net/url"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

// BillApiService 提供直连商户账单相关接口，包括交易账单、资金账单和分账账单下载申请。
//
// 申请成功后会返回 download_url、hash_type 和 hash_value。download_url 有效期为 5 分钟，
// 建议商户下载文件后对比 hash_value 校验账单完整性。
type BillApiService services.Service

// BillApply 申请交易账单的下载地址。
//
// 交易账单按天生成，包含交易相关的金额、时间、营销等信息，供商户核对订单交易完成、退款、撤销等情况。
// 抖音侧未成功下单的交易不会出现在对账单中，支付成功后撤销的交易会出现在对账单中且沿用原支付单订单号。
// 账单涉及金额字段的单位为元，账单文件通常建议在 T+1 日 10 点后获取。
func (a *BillApiService) BillApply(ctx context.Context, req BillApplyRequest) (
	resp *Bill, result *client.APIResult, err error) {
	return a.applyBill(ctx, req)
}

// ApplyFundFlowBill 申请资金账单的下载地址。
//
// 资金账单按天生成，用于反映抖音支付账户的资金变动情况，包含业务单号、收支金额和记账时间等信息。
// 账单涉及金额字段的单位为元；account_type 可选值包括 BaseAccount（基本账户）和
// OperationAccount（运营账户），OpenAPI 文档中默认值为 BaseAccount。
func (a *BillApiService) ApplyFundFlowBill(ctx context.Context, req ApplyFundFlowBillRequest) (
	resp *Bill, result *client.APIResult, err error) {
	return a.applyBill(ctx, req)
}

// ApplySplitBill 申请分账账单的下载地址。
//
// 分账账单按天生成，包含分账相关的金额、时间和状态等信息，供商户核对到账和分账结果。
// 账单涉及金额字段的单位为元，tar_type 常用值为 GZIP。
func (a *BillApiService) ApplySplitBill(ctx context.Context, req ApplySplitBillRequest) (
	resp *Bill, result *client.APIResult, err error) {
	return a.applyBill(ctx, req)
}

type ApplyBillReq interface {
	Validate() error
	GetQueryParams() neturl.Values
	GetPath() string
}

func (a *BillApiService) applyBill(ctx context.Context, req ApplyBillReq) (
	resp *Bill, result *client.APIResult, err error) {
	// Validate
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	// Set Params
	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		QueryParams: req.GetQueryParams(),
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + req.GetPath()
	// Do Http Request
	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	// Extract Http Response
	resp = new(Bill)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
