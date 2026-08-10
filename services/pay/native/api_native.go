package native

// 官方文档：
// - Native下单：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/65bf88e4fcc2570315fd1457
// - 查询订单：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/65bf8a15ea861802f27235b3
// - 关闭订单：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/65bf8b844d4516030eec0e12

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/services/pay"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

type NativeApiService services.Service

// Prepay Native支付下单
func (a *NativeApiService) Prepay(ctx context.Context, req PrepayRequest) (resp *PrepayResponse, result *client.APIResult, err error) {
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}

	r.RequestPath = consts.DouyinPayServer + consts.NativePrepayPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(PrepayResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// CloseOrder 关闭订单
func (a *NativeApiService) CloseOrder(ctx context.Context, req CloseOrderRequest) (result *client.APIResult, err error) {
	if req.OutTradeNo == "" {
		return nil, fmt.Errorf("CloseOrderRequest required field `OutTradeNo` is empty")
	}

	localVarPostBody := &CloseRequest{
		Mchid: req.Mchid,
	}

	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    localVarPostBody,
		Header:      nethttp.Header{},
	}

	rawUrl := consts.DouyinPayServer + consts.ClosePath
	r.RequestPath = strings.Replace(rawUrl, "{out_trade_no}", neturl.PathEscape(req.OutTradeNo), -1)

	// Perform Http Request
	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return result, err
	}

	return result, nil
}

// QueryOrderById 抖音支付订单号查询订单
func (a *NativeApiService) QueryOrderById(ctx context.Context, req QueryOrderByIdRequest) (resp *pay.Transaction, result *client.APIResult, err error) {
	if req.TransactionId == "" {
		return nil, nil, fmt.Errorf("QueryOrderByIdRequest required field `TransactionId` is empty")
	}
	if req.Mchid == "" {
		return nil, nil, fmt.Errorf("QueryOrderByIdRequest required field `Mchid` is empty")
	}

	// Setup Query Params
	localVarQueryParams := neturl.Values{}
	localVarQueryParams.Add("mchid", req.Mchid)

	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
		QueryParams: localVarQueryParams,
	}

	rawUrl := consts.DouyinPayServer + consts.QueryByIdPath
	r.RequestPath = strings.Replace(rawUrl, "{transaction_id}", neturl.PathEscape(req.TransactionId), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	// Extract payments.Transaction from Http Response
	resp = new(pay.Transaction)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryOrderByOutTradeNo 商户订单号查询订单
func (a *NativeApiService) QueryOrderByOutTradeNo(ctx context.Context, req QueryOrderByOutTradeNoRequest) (resp *pay.Transaction, result *client.APIResult, err error) {
	// Make sure Path Params are properly set
	if req.OutTradeNo == "" {
		return nil, nil, fmt.Errorf("QueryOrderByOutTradeNoRequest required field `OutTradeNo` is empty")
	}
	if req.Mchid == "" {
		return nil, nil, fmt.Errorf("QueryOrderByOutTradeNoRequest required field `Mchid` is empty")
	}

	localVarQueryParams := neturl.Values{}
	localVarQueryParams.Add("mchid", req.Mchid)

	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
		QueryParams: localVarQueryParams,
	}

	rawUrl := consts.DouyinPayServer + consts.QueryByOutTradeNoPath
	r.RequestPath = strings.Replace(rawUrl, "{out_trade_no}", neturl.PathEscape(req.OutTradeNo), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	// Extract payments.Transaction from Http Response
	resp = new(pay.Transaction)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
