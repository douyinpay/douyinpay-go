package partnerbill

// 官方文档：
// - 申请交易账单：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/69e2ee7acad2c105c439a809
// - 申请资金账单：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/684a53064037d5050b11863d
// - 申请分账账单：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/684a53090efadf054e0489f0

import (
	"context"
	nethttp "net/http"
	neturl "net/url"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

// BillApiService 提供服务商账单相关接口，包括交易账单、资金账单和分账账单下载申请。
type BillApiService services.Service

// ApplyTradeBill 申请服务商交易账单的下载地址。
//
// 服务商需传入服务商商户号和账单日期；账单日期格式为 yyyy-MM-dd，
// 仅支持近三个月内且为昨日及以前的账单。sub_mchid 为选填，传入后用于过滤指定子商户账单；
// tar_type 常用值为 GZIP。
func (a *BillApiService) ApplyTradeBill(ctx context.Context, req ApplyTradeBillRequest) (
	resp *Bill, result *client.APIResult, err error) {
	return a.applyBill(ctx, req)
}

// ApplyFundFlowBill 申请服务商资金账单的下载地址。
//
// 服务商需传入服务商商户号和账单日期；账单日期格式为 yyyy-MM-dd，
// 仅支持近三个月内且为昨日及以前的账单。sub_mchid 为选填，account_type 常见取值包括
// BaseAccount（基本户）和 OperationAccount（运营户），tar_type 常用值为 GZIP。
func (a *BillApiService) ApplyFundFlowBill(ctx context.Context, req ApplyFundFlowBillRequest) (
	resp *Bill, result *client.APIResult, err error) {
	return a.applyBill(ctx, req)
}

// ApplySplitBill 申请服务商分账账单的下载地址。
//
// 服务商需传入服务商商户号和账单日期；账单日期格式为 yyyy-MM-dd，
// 仅支持近三个月内且为昨日及以前的账单。sub_mchid 为选填，tar_type 常用值为 GZIP。
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
