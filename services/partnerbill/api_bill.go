package partnerbill

import (
	"context"
	nethttp "net/http"
	neturl "net/url"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

// BillApiService 提供服务商账单相关接口，包括交易账单、资金账单和分账账单下载申请。
//
// 适用于普通服务商和平台商户。申请成功后会返回 download_url、hash_type 和 hash_value，
// 其中 download_url 有效期为 5 分钟，建议下载完成后对比 hash_value 校验账单完整性。
type BillApiService services.Service

// ApplyTradeBill 申请服务商交易账单的下载地址。
//
// 交易账单按天生成，包含交易相关的金额、时间、营销等信息，供商户核对订单交易完成、退款、撤销等情况。
// 二级商户不单独提供对账单下载；如需下载某个子商户下的交易或退款数据，可传入 sub_mchid，平台商户不支持该字段。
func (a *BillApiService) ApplyTradeBill(ctx context.Context, req ApplyTradeBillRequest) (
	resp *Bill, result *client.APIResult, err error) {
	return a.applyBill(ctx, req)
}

// ApplyFundFlowBill 申请服务商资金账单的下载地址。
//
// 资金账单按天生成，反映商户账户的资金变动情况，包含业务单号、收支金额和记账时间等信息。
// account_type 选填，可选值包括 BaseAccount（基本账户）、OperationAccount（运营账户）和 FeeAccount（手续费账户），默认值为 BaseAccount。
func (a *BillApiService) ApplyFundFlowBill(ctx context.Context, req ApplyFundFlowBillRequest) (
	resp *Bill, result *client.APIResult, err error) {
	return a.applyBill(ctx, req)
}

// ApplySplitBill 申请服务商分账账单的下载地址。
//
// 分账账单按天生成，包含分账相关的金额、时间等信息，供商户核对到账等情况。
// 抖音侧未成功的分账单不会出现在对账单中；如需下载某个子商户下的分账账单，可传入 sub_mchid。
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
