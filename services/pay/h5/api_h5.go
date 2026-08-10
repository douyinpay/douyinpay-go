package h5

// 官方文档：
// - H5下单：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/63f444f0b7d2f20202eaa928
// - 查询订单：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/63f445410b970c0209070167
// - 关闭订单：https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/63f4450f0b970c02090700df

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

type H5ApiService services.Service

// Prepay H5支付下单
func (a *H5ApiService) Prepay(ctx context.Context, req PrepayRequest) (resp *PrepayResponse, result *client.APIResult, err error) {
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.H5PrepayPath

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
func (a *H5ApiService) CloseOrder(ctx context.Context, req CloseOrderRequest) (result *client.APIResult, err error) {
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
func (a *H5ApiService) QueryOrderById(ctx context.Context, req QueryOrderByIdRequest) (resp *pay.Transaction, result *client.APIResult, err error) {
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
func (a *H5ApiService) QueryOrderByOutTradeNo(ctx context.Context, req QueryOrderByOutTradeNoRequest) (resp *pay.Transaction, result *client.APIResult, err error) {
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
