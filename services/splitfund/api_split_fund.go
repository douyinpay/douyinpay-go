package splitfund

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

type SplitFundService services.Service

// 分账
func (a *SplitFundService) SplitFund(ctx context.Context, req SplitFundReq) (resp *SplitFundResp, result *client.APIResult, err error) {
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.SplitFundPath
	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	resp = new(SplitFundResp)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 分账查询
func (a *SplitFundService) QuerySplitFund(ctx context.Context, req QuerySplitFundReq) (resp *QuerySplitFundResp, result *client.APIResult, err error) {
	if req.OutOrderNo == "" {
		return nil, nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in QuerySplitFundReq")
	}
	localVarQueryParams := neturl.Values{}
	if req.MerchantId != "" {
		localVarQueryParams.Add("mchid", req.MerchantId)
	}
	if req.TradeNo != "" {
		localVarQueryParams.Add("transaction_id", req.TradeNo)
	}
	if req.OrderId != "" {
		localVarQueryParams.Add("order_id", req.OrderId)
	}
	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		QueryParams: localVarQueryParams,
		Header:      nethttp.Header{},
	}
	rawUrl := consts.DouyinPayServer + consts.QuerySplitFundPath
	r.RequestPath = strings.Replace(rawUrl, "{out_order_no}", neturl.PathEscape(req.OutOrderNo), -1)
	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	// Extract QuerySplitFundResp from Http Response
	resp = new(QuerySplitFundResp)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
