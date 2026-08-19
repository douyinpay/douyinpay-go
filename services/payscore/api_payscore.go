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

// ApiPayScoreService 先享后付服务 API。
type ApiPayScoreService services.Service

var (
	// reqUrl 先享后付服务 API 的请求地址。
	reqUrl = consts.DouyPayServer
)

// CreateServiceOrder 创建服务订单。
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

// CompleteServiceOrder 完结服务订单。
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

// QueryServiceOrder 查询服务订单信息。
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

// CancelServiceOrder 取消服务订单。
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

// ModifyAmount 修改订单金额。
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

// SynchronizeServiceOrderInfo 同步服务订单信息。
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

// ServiceOrderPay 商户发起催收扣款。
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

// CreditSrvSignApply 申请服务授权。
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

// CreditSrvSignQuery 查询用户授权记录。
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

// CloseCreditService 解除用户授权关系。
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

// CreateServiceOrderForSP 服务商创建服务订单。
func (a *ApiPayScoreService) CreateServiceOrderForSP(ctx context.Context, req ApiPartnerCreateServiceOrderRequest) (
	resp *ApiPartnerCreateServiceOrderResponse, result *client.APIResult, err error) {
	result, err = a.requestPartnerServiceOrder(ctx, consts.PartnerCreateServiceOrderPath, nethttp.MethodPost, req, nil)
	if err != nil {
		return nil, result, err
	}
	resp = new(ApiPartnerCreateServiceOrderResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, err
}

// CompleteServiceOrderForSP 服务商完结服务订单。
func (a *ApiPayScoreService) CompleteServiceOrderForSP(ctx context.Context, req ApiPartnerCompleteServiceOrderRequest) (
	resp *ApiPartnerCompleteServiceOrderResponse, result *client.APIResult, err error) {
	result, err = a.requestPartnerServiceOrder(ctx, consts.PartnerCompleteServiceOrderPath, nethttp.MethodPost, req, nil)
	if err != nil {
		return nil, result, err
	}
	resp = new(ApiPartnerCompleteServiceOrderResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, err
}

// QueryServiceOrderForSP 服务商查询服务订单。
func (a *ApiPayScoreService) QueryServiceOrderForSP(ctx context.Context, req ApiPartnerQueryServiceOrderRequest) (
	resp *ApiPartnerQueryServiceOrderResponse, result *client.APIResult, err error) {
	result, err = a.requestPartnerServiceOrder(ctx, consts.PartnerQueryServiceOrderPath, nethttp.MethodGet, nil,
		neturl.Values{
			"sp_mchid":     []string{req.SpMchid},
			"sub_mchid":    []string{req.SubMchid},
			"out_order_no": []string{req.OutOrderNo},
			"service_id":   []string{req.ServiceId},
		})
	if err != nil {
		return nil, result, err
	}
	resp = new(ApiPartnerQueryServiceOrderResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, err
}

// CancelServiceOrderForSP 服务商取消服务订单。
func (a *ApiPayScoreService) CancelServiceOrderForSP(ctx context.Context, req ApiPartnerCancelServiceOrderRequest) (
	resp *ApiPartnerCancelServiceOrderResponse, result *client.APIResult, err error) {
	result, err = a.requestPartnerServiceOrder(ctx, consts.PartnerCancelServiceOrderPath, nethttp.MethodPost, req, nil)
	if err != nil {
		return nil, result, err
	}
	resp = new(ApiPartnerCancelServiceOrderResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, err
}

// SynchronizeServiceOrderInfoForSP 服务商同步服务订单信息。
func (a *ApiPayScoreService) SynchronizeServiceOrderInfoForSP(ctx context.Context, req ApiPartnerSynchronizeServiceOrderInfoRequest) (
	resp *ApiPartnerSynchronizeServiceOrderInfoResponse, result *client.APIResult, err error) {
	path := strings.Replace(consts.PartnerSyncServiceOrderPath, "{out_order_no}", neturl.PathEscape(req.OutOrderNo), -1)
	result, err = a.requestPartnerServiceOrder(ctx, path, nethttp.MethodPost, req, nil)
	if err != nil {
		return nil, result, err
	}
	resp = new(ApiPartnerSynchronizeServiceOrderInfoResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, err
}

// ModifyAmountForSP 服务商修改订单金额。
func (a *ApiPayScoreService) ModifyAmountForSP(ctx context.Context, req ApiPartnerModifyAmountRequest) (
	resp *ApiPartnerModifyAmountResponse, result *client.APIResult, err error) {
	result, err = a.requestPartnerServiceOrder(ctx, consts.PartnerModifyServiceOrderPath, nethttp.MethodPost, req, nil)
	if err != nil {
		return nil, result, err
	}
	resp = new(ApiPartnerModifyAmountResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, err
}

// CreditSrvSignApplyForSP 服务商申请服务授权。
func (a *ApiPayScoreService) CreditSrvSignApplyForSP(ctx context.Context, req ApiPartnerCreditSrvSignApplyRequest) (
	resp *ApiPartnerCreditSrvSignApplyResponse, result *client.APIResult, err error) {
	result, err = a.requestPartnerServiceOrder(ctx, consts.PartnerCreditSrvSignApplyPath,
		nethttp.MethodPost, req, nil)
	if err != nil {
		return nil, result, err
	}
	resp = new(ApiPartnerCreditSrvSignApplyResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, err
}

// CreditSrvSignQueryForSP 服务商查询用户授权记录。
func (a *ApiPayScoreService) CreditSrvSignQueryForSP(ctx context.Context, req ApiPartnerCreditSrvSignQueryRequest) (
	resp *ApiPartnerCreditSrvSignQueryResponse, result *client.APIResult, err error) {
	path := strings.Replace(consts.PartnerCreditSrvSignQueryPath,
		"{authorization_code}", neturl.PathEscape(req.AuthorizationCode), -1)
	result, err = a.requestPartnerServiceOrder(ctx, path, nethttp.MethodGet, nil,
		neturl.Values{
			"sp_mchid":   []string{req.SpMchid},
			"sub_mchid":  []string{req.SubMchid},
			"service_id": []string{req.ServiceId},
		})
	if err != nil {
		return nil, result, err
	}
	resp = new(ApiPartnerCreditSrvSignQueryResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, err
}

// CloseCreditServiceForSP 服务商解除用户授权关系。
func (a *ApiPayScoreService) CloseCreditServiceForSP(ctx context.Context, req ApiPartnerCloseCreditServiceRequest) (
	resp *ApiPartnerCloseCreditServiceResponse, result *client.APIResult, err error) {
	path := reqUrl + strings.Replace(consts.PartnerCloseCreditServicePath,
		"{authorization_code}", neturl.PathEscape(req.AuthorizationCode), -1)
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
		RequestPath: path,
	}
	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	resp = new(ApiPartnerCloseCreditServiceResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, err
}

func (a *ApiPayScoreService) requestPartnerServiceOrder(
	ctx context.Context, path, method string, body interface{}, query neturl.Values,
) (result *client.APIResult, err error) {
	r := &client.RequestEntity{
		Method:      method,
		ContentType: consts.ApplicationJSON,
		PostBody:    body,
		QueryParams: query,
		Header:      nethttp.Header{},
		RequestPath: reqUrl + path,
	}
	result, err = a.Client.Request(ctx, r)
	return result, err
}
