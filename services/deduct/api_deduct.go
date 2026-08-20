package contractorder

import (
	"context"
	"fmt"
	nethttp "net/http"
	netUrl "net/url"
	"strings"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

// ApiDeductService 直连商户代扣服务。
//
// 接口路径：
// - 申请扣款：POST /v1/deduct/payapply
// - 关闭订单：POST /v1/trade/transactions/out-trade-no/{out_trade_no}/close
// - 按抖音支付订单号查询订单：GET /v1/trade/transactions/id/{transaction_id}
// - 按商户订单号查询订单：GET /v1/trade/transactions/out-trade-no/{out_trade_no}
// - 预约扣费通知：POST /v1/agreementauth/deductNotify
type ApiDeductService services.Service

// Deduct 已废弃，请改用 PayApply 申请扣款。
//
// Deprecated: 请改用 PayApply 申请扣款。
func (a *ApiDeductService) Deduct(ctx context.Context, req ApiDeductRequest) (
	resp *ApiDeductResponse, result *client.APIResult, err error) {
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.DeductPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(ApiDeductResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// CloseOrder 关闭订单
func (a *ApiDeductService) CloseOrder(ctx context.Context, req CloseOrderRequest) (result *client.APIResult, err error) {
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
	r.RequestPath = strings.Replace(rawUrl, "{out_trade_no}", netUrl.PathEscape(req.OutTradeNo), -1)

	// Perform Http Request
	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return result, err
	}

	return result, nil
}

// QueryOrderById 抖音支付订单号查询订单
func (a *ApiDeductService) QueryOrderById(ctx context.Context, req QueryOrderByIdRequest) (resp *Transaction, result *client.APIResult, err error) {
	if req.TransactionId == "" {
		return nil, nil, fmt.Errorf("QueryOrderByIdRequest required field `TransactionId` is empty")
	}
	if req.Mchid == "" {
		return nil, nil, fmt.Errorf("QueryOrderByIdRequest required field `Mchid` is empty")
	}

	// Setup Query Params
	localVarQueryParams := netUrl.Values{}
	localVarQueryParams.Add("mchid", req.Mchid)

	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
		QueryParams: localVarQueryParams,
	}
	rawUrl := consts.DouyinPayServer + consts.QueryByIdPath
	r.RequestPath = strings.Replace(rawUrl, "{transaction_id}", netUrl.PathEscape(req.TransactionId), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	// Extract payments.Transaction from Http Response
	resp = new(Transaction)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryOrderByOutTradeNo 商户订单号查询订单
func (a *ApiDeductService) QueryOrderByOutTradeNo(ctx context.Context, req QueryOrderByOutTradeNoRequest) (resp *Transaction, result *client.APIResult, err error) {
	// Make sure Path Params are properly set
	if req.OutTradeNo == "" {
		return nil, nil, fmt.Errorf("QueryOrderByOutTradeNoRequest required field `OutTradeNo` is empty")
	}
	if req.Mchid == "" {
		return nil, nil, fmt.Errorf("QueryOrderByOutTradeNoRequest required field `Mchid` is empty")
	}

	localVarQueryParams := netUrl.Values{}
	localVarQueryParams.Add("mchid", req.Mchid)

	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
		QueryParams: localVarQueryParams,
	}
	rawUrl := consts.DouyinPayServer + consts.QueryByOutTradeNoPath
	r.RequestPath = strings.Replace(rawUrl, "{out_trade_no}", netUrl.PathEscape(req.OutTradeNo), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	// Extract payments.Transaction from Http Response
	resp = new(Transaction)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// DeductNotify 预约扣费通知。
// 接口路径：POST /v1/agreementauth/deductNotify
func (a *ApiDeductService) DeductNotify(ctx context.Context, req DeductNotifyRequest) (result *client.APIResult, err error) {
	if req.ContractId == "" {
		return nil, fmt.Errorf("DeductNotifyRequest required field `ContractId` is empty")
	}
	if req.Appid == "" {
		return nil, fmt.Errorf("DeductNotifyRequest required field `Appid` is empty")
	}
	if req.Mchid == "" {
		return nil, fmt.Errorf("DeductNotifyRequest required field `Mchid` is empty")
	}
	if req.EstimatedAmount.Amount < 0 {
		return nil, fmt.Errorf("DeductNotifyRequest required field `EstimatedAmount#Amount` is invalid")
	}
	if req.EstimatedAmount.Currency == "" {
		return nil, fmt.Errorf("DeductNotifyRequest required field `EstimatedAmount#Currency` is empty")
	}

	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.DeductNotifyPath

	// Perform Http Request
	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return result, err
	}

	return result, nil
}

// PayApply 申请扣款。
// 接口路径：POST /v1/deduct/payapply
func (a *ApiDeductService) PayApply(ctx context.Context, req PayApplyRequest) (resp *PayApplyResponse, result *client.APIResult, err error) {
	// Make sure Path Params are properly set
	if req.ContractId == "" {
		return nil, nil, fmt.Errorf("PayApplyRequest required field `ContractId` is empty")
	}
	if req.Appid == "" {
		return nil, nil, fmt.Errorf("PayApplyRequest required field `Appid` is empty")
	}
	if req.Mchid == "" {
		return nil, nil, fmt.Errorf("PayApplyRequest required field `Mchid` is empty")
	}
	if req.OutTradeNo == "" {
		return nil, nil, fmt.Errorf("PayApplyRequest required field `OutTradeNo` is empty")
	}
	if req.TradeType == "" {
		return nil, nil, fmt.Errorf("PayApplyRequest required field `TradeType` is empty")
	}
	if req.Description == "" {
		return nil, nil, fmt.Errorf("PayApplyRequest required field `Description` is empty")
	}
	if req.NotifyUrl == "" {
		return nil, nil, fmt.Errorf("PayApplyRequest required field `NotifyUrl` is empty")
	}
	if req.Amount == nil || req.Amount.Total < 0 {
		return nil, nil, fmt.Errorf("PayApplyRequest required field `Amount or Amount#Total` is empty")
	}

	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
		PostBody:    req,
	}

	r.RequestPath = consts.DouyinPayServer + consts.PayApplyPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	// Extract PayApplyResponse from Http Response
	resp = new(PayApplyResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// GetServerAddress 返回抖音支付服务地址。
func GetServerAddress() string {
	return consts.DouyinPayServer
}
