package partnerh5

// 官方文档：
// - H5下单：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/6911f7c7f09d4f04f443b5e4
// - 查询订单：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/6852bb22fe022d05166966c5
// - 关闭订单：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/6852bb25d479e6051ac20fb4

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/services/partnerpay"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

type H5ApiService services.Service

// Prepay 服务商H5下单
func (a *H5ApiService) Prepay(ctx context.Context, req PrepayRequest) (resp *PrepayResponse, result *client.APIResult, err error) {
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.PartnerH5PrepayPath
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
		SpMchid:  req.SpMchid,
		SubMchid: req.SubMchid,
	}

	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    localVarPostBody,
		Header:      nethttp.Header{},
	}

	rawUrl := consts.DouyinPayServer + consts.PartnerClosePath
	r.RequestPath = strings.Replace(rawUrl, "{out_trade_no}", neturl.PathEscape(req.OutTradeNo), -1)

	// Perform Http Request
	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return result, err
	}

	return result, nil
}

// QueryOrderById 抖音支付订单号查询订单
func (a *H5ApiService) QueryOrderById(ctx context.Context, req QueryOrderByIdRequest) (resp *partnerpay.Transaction, result *client.APIResult, err error) {
	if req.TransactionId == "" {
		return nil, nil, fmt.Errorf("QueryOrderByIdRequest required field `TransactionId` is empty")
	}
	if req.SubMchid == "" {
		return nil, nil, fmt.Errorf("QueryOrderByIdRequest required field `SubMchid` is empty")
	}
	if req.SpMchid == "" {
		return nil, nil, fmt.Errorf("QueryOrderByIdRequest required field `SpMchid` is empty")
	}

	// Setup Query Params
	localVarQueryParams := neturl.Values{}
	localVarQueryParams.Add("sp_mchid", req.SpMchid)
	localVarQueryParams.Add("sub_mchid", req.SubMchid)

	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
		QueryParams: localVarQueryParams,
	}

	rawUrl := consts.DouyinPayServer + consts.PartnerQueryByIdPath
	r.RequestPath = strings.Replace(rawUrl, "{transaction_id}", neturl.PathEscape(req.TransactionId), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	// Extract payments.Transaction from Http Response
	resp = new(partnerpay.Transaction)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryOrderByOutTradeNo 商户订单号查询订单
func (a *H5ApiService) QueryOrderByOutTradeNo(ctx context.Context, req QueryOrderByOutTradeNoRequest) (resp *partnerpay.Transaction, result *client.APIResult, err error) {
	// Make sure Path Params are properly set
	if req.OutTradeNo == "" {
		return nil, nil, fmt.Errorf("QueryOrderByOutTradeNoRequest required field `OutTradeNo` is empty")
	}
	if req.SubMchid == "" {
		return nil, nil, fmt.Errorf("QueryOrderByIdRequest required field `SubMchid` is empty")
	}
	if req.SpMchid == "" {
		return nil, nil, fmt.Errorf("QueryOrderByIdRequest required field `SpMchid` is empty")
	}

	localVarQueryParams := neturl.Values{}
	localVarQueryParams.Add("sp_mchid", req.SpMchid)
	localVarQueryParams.Add("sub_mchid", req.SubMchid)

	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
		QueryParams: localVarQueryParams,
	}

	rawUrl := consts.DouyinPayServer + consts.PartnerQueryByOutTradeNoPath
	r.RequestPath = strings.Replace(rawUrl, "{out_trade_no}", neturl.PathEscape(req.OutTradeNo), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	// Extract payments.Transaction from Http Response
	resp = new(partnerpay.Transaction)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
