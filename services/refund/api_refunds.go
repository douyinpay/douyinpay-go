package refund

// 官方文档：
// - 申请退款（App支付）：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/639fd61817c2f3021d238235
// - 查询单笔退款（App支付）：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/639fd62c17c2f3021d23826f
// - 申请退款（JSAPI支付）：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/64413ff8ef6db1021d438dcf
// - 申请退款（Native支付）：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/65bf8d013ba7e102fb44d69e
// - 查询单笔退款（H5支付）：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/63f4451e88f4740227f6e36c

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

type RefundsApiService services.Service

// Create 退款申请
func (a *RefundsApiService) Create(ctx context.Context, req CreateRequest) (resp *Refund, result *client.APIResult, err error) {
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}

	r.RequestPath = consts.DouyinPayServer + consts.RefundPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	// Extract Refund from Http Response
	resp = new(Refund)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryByOutRefundNo 查询单笔退款（通过商户退款单号）
func (a *RefundsApiService) QueryByOutRefundNo(ctx context.Context, req QueryByOutRefundNoRequest) (resp *Refund, result *client.APIResult, err error) {
	if req.OutRefundNo == "" {
		return nil, nil, fmt.Errorf("field `OutRefundNo` is required and must be specified in QueryByOutRefundNoRequest")
	}

	localVarQueryParams := neturl.Values{}
	// 直连商户
	if req.Mchid != "" {
		localVarQueryParams.Add("mchid", req.Mchid)
	}
	if req.Appid != "" {
		localVarQueryParams.Add("appid", req.Appid)
	}
	// 服务商
	if req.SpMchid != "" {
		localVarQueryParams.Add("sp_mchid", req.SpMchid)
	}
	if req.SubMchid != "" {
		localVarQueryParams.Add("sub_mchid", req.SubMchid)
	}

	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		QueryParams: localVarQueryParams,
		Header:      nethttp.Header{},
	}
	rawUrl := consts.DouyinPayServer + consts.RefundQueryByOutTradeNoPath
	r.RequestPath = strings.Replace(rawUrl, "{out_refund_no}", neturl.PathEscape(req.OutRefundNo), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	// Extract Refund from Http Response
	resp = new(Refund)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
