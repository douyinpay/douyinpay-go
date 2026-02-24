package partnerbill

import (
	"context"
	nethttp "net/http"
	neturl "net/url"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

type BillApiService services.Service

func (a *BillApiService) ApplyTradeBill(ctx context.Context, req ApplyTradeBillRequest) (
	resp *Bill, result *client.APIResult, err error) {
	return a.applyBill(ctx, req)
}

func (a *BillApiService) ApplyFundFlowBill(ctx context.Context, req ApplyFundFlowBillRequest) (
	resp *Bill, result *client.APIResult, err error) {
	return a.applyBill(ctx, req)
}

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
