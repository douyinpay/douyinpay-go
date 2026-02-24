package contractorder

import (
	"context"
	"encoding/json"
	"log"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

type ApiPayScoreService services.Service

var (
	reqUrl = consts.DouyPayServer
)

func (a *ApiPayScoreService) CreateServiceOrder(ctx context.Context, req ApiCreateServiceOrderRequest) (resp *ApiCreateServiceOrderResponse, result *client.APIResult, err error) {
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = reqUrl + consts.CreateServiceOrderPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(ApiCreateServiceOrderResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ApiPayScoreService) CompleteServiceOrder(
	ctx context.Context, req ApiCompleteServiceOrderRequest) (
	resp *ApiCompleteServiceOrderResponse, result *client.APIResult, err error) {
	reqBytes, _ := json.Marshal(req)
	log.Printf("request: %s\n", string(reqBytes))
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = reqUrl + consts.CompleteServiceOrderPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(ApiCompleteServiceOrderResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ApiPayScoreService) QueryServiceOrder(
	ctx context.Context, req ApiQueryServiceOrderRequest) (
	resp *ApiQueryServiceOrderResponse, result *client.APIResult, err error) {
	reqBytes, _ := json.Marshal(req)
	log.Printf("request: %s\n", string(reqBytes))
	localVarQueryParams := neturl.Values{}
	localVarQueryParams.Add("mchid", req.Mchid)
	localVarQueryParams.Add("appid", req.Appid)
	localVarQueryParams.Add("service_id", req.ServiceId)
	localVarQueryParams.Add("out_order_no", req.OutOrderNo)

	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		QueryParams: localVarQueryParams,
		Header:      nethttp.Header{},
	}
	r.RequestPath = reqUrl + consts.QueryServiceOrderPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(ApiQueryServiceOrderResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ApiPayScoreService) CancelServiceOrder(
	ctx context.Context, req ApiCancelServiceOrderRequest) (
	resp *ApiCancelServiceOrderResponse, result *client.APIResult, err error) {
	reqBytes, _ := json.Marshal(req)
	log.Printf("request: %s\n", string(reqBytes))
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = reqUrl + consts.CancelServiceOrderPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(ApiCancelServiceOrderResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ModifyAmount 修改订单金额
func (a *ApiPayScoreService) ModifyAmount(
	ctx context.Context, req ApiModifyAmountRequest) (
	resp *ApiModifyAmountResponse, result *client.APIResult, err error) {
	reqBytes, _ := json.Marshal(req)
	log.Printf("request: %s\n", string(reqBytes))
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = reqUrl + consts.ModifyServiceOrderPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(ApiModifyAmountResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// SynchronizeServiceOrderInfo 同步服务订单信息
func (a *ApiPayScoreService) SynchronizeServiceOrderInfo(
	ctx context.Context, req ApiSynchronizeServiceOrderInfoRequest) (
	resp *ApiSynchronizeServiceOrderInfoResponse, result *client.APIResult, err error) {
	reqBytes, _ := json.Marshal(req)
	log.Printf("request: %s\n", string(reqBytes))
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = reqUrl + strings.Replace(consts.SyncServiceOrderPath, "{out_order_no}", req.OutOrderNo, 1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(ApiSynchronizeServiceOrderInfoResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ApiPayScoreService) ServiceOrderPay(ctx context.Context, req ApiServiceOrderPayRequest) (
	resp *ApiServiceOrderPayResponse, result *client.APIResult, err error) {
	reqBytes, _ := json.Marshal(req)
	log.Printf("request: %s\n", string(reqBytes))
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = reqUrl + strings.Replace(consts.ServiceOrderPayPath, "{out_order_no}", req.OutOrderNo, 1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(ApiServiceOrderPayResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ApiPayScoreService) CreditSrvSignApply(ctx context.Context, req ApiCreditSrvSignApplyRequest) (
	resp *ApiCreditSrvSignApplyResponse, result *client.APIResult, err error) {
	reqBytes, _ := json.Marshal(req)
	log.Printf("request: %s\n", string(reqBytes))
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = reqUrl + consts.CreditSrvSignApplyPath
	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(ApiCreditSrvSignApplyResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ApiPayScoreService) CreditSrvSignQuery(ctx context.Context, req ApiCreditSrvSignQueryRequest) (
	resp *ApiCreditSrvSignQueryResponse, result *client.APIResult, err error) {
	reqBytes, _ := json.Marshal(req)
	log.Printf("request: %s\n", string(reqBytes))
	localVarQueryParams := neturl.Values{}
	localVarQueryParams.Add("mchid", req.Mchid)
	localVarQueryParams.Add("service_id", req.ServiceId)
	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		QueryParams: localVarQueryParams,
		Header:      nethttp.Header{},
	}
	r.RequestPath = reqUrl + strings.Replace(consts.CreditSrvSignQueryPath, "{authorization_code}", req.AuthorizationCode, 1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(ApiCreditSrvSignQueryResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ApiPayScoreService) CloseCreditService(
	ctx context.Context, req ApiCloseCreditServiceRequest) (
	resp *ApiCloseCreditServiceResponse, result *client.APIResult, err error) {
	reqBytes, _ := json.Marshal(req)
	log.Printf("request: %s\n", string(reqBytes))
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = reqUrl + strings.Replace(consts.CloseCreditServicePath,
		"{authorization_code}", req.AuthorizationCode, 1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(ApiCloseCreditServiceResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
