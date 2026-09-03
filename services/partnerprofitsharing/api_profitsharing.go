package partnerprofitsharing

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/services/callback"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

type ApiService services.Service

// SplitFund 请求分账
func (a *ApiService) SplitFund(ctx context.Context, req ApiPartnerSplitFundRequest) (resp *ApiPartnerSplitFundResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	resp = new(ApiPartnerSplitFundResponse)
	result, err = a.doRequest(ctx, nethttp.MethodPost, consts.SplitFundPath, nil, req, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QuerySplitFund 查询分账结果
func (a *ApiService) QuerySplitFund(ctx context.Context, req ApiPartnerQuerySplitFundRequest) (resp *ApiPartnerQuerySplitFundResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	queryParams := neturl.Values{}
	queryParams.Add("sp_mchid", req.SpMchid)
	queryParams.Add("sub_mchid", req.SubMchid)
	if req.TransactionID != "" {
		queryParams.Add("transaction_id", req.TransactionID)
	}
	if req.OrderID != "" {
		queryParams.Add("order_id", req.OrderID)
	}
	resp = new(ApiPartnerQuerySplitFundResponse)
	path := strings.Replace(consts.QuerySplitFundPath, "{out_order_no}", neturl.PathEscape(req.OutOrderNo), 1)
	result, err = a.doRequest(ctx, nethttp.MethodGet, path, queryParams, nil, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ReturnSplitFund 请求分账回退
func (a *ApiService) ReturnSplitFund(ctx context.Context, req ApiPartnerReturnSplitFundRequest) (resp *ApiPartnerReturnSplitFundResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	resp = new(ApiPartnerReturnSplitFundResponse)
	result, err = a.doRequest(ctx, nethttp.MethodPost, consts.ReturnSplitFundPath, nil, req, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryReturnSplitFund 查询分账回退结果
func (a *ApiService) QueryReturnSplitFund(ctx context.Context, req ApiPartnerQueryReturnSplitFundRequest) (resp *ApiPartnerQueryReturnSplitFundResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	queryParams := neturl.Values{}
	queryParams.Add("sp_mchid", req.SpMchid)
	queryParams.Add("sub_mchid", req.SubMchid)
	if req.OutOrderNo != "" {
		queryParams.Add("out_order_no", req.OutOrderNo)
	}
	resp = new(ApiPartnerQueryReturnSplitFundResponse)
	path := strings.Replace(consts.QueryReturnSplitFundPath, "{out_return_no}", neturl.PathEscape(req.OutReturnNo), 1)
	result, err = a.doRequest(ctx, nethttp.MethodGet, path, queryParams, nil, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// FinishSplitFund 完结分账
func (a *ApiService) FinishSplitFund(ctx context.Context, req ApiPartnerFinishSplitFundRequest) (resp *ApiPartnerFinishSplitFundResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	resp = new(ApiPartnerFinishSplitFundResponse)
	result, err = a.doRequest(ctx, nethttp.MethodPost, consts.FinishSplitFundPath, nil, req, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryUnsplitAmount 查询剩余待分金额
func (a *ApiService) QueryUnsplitAmount(ctx context.Context, req QueryUnsplitAmountRequest) (resp *QueryUnsplitAmountResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	queryParams := neturl.Values{}
	queryParams.Add("sp_mchid", req.SpMchid)
	resp = new(QueryUnsplitAmountResponse)
	path := strings.Replace(consts.QueryUnsplitAmountPath, "{transaction_id}", neturl.PathEscape(req.TransactionID), 1)
	result, err = a.doRequest(ctx, nethttp.MethodGet, path, queryParams, nil, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryMerchantConfig 查询特约商户分账配置
func (a *ApiService) QueryMerchantConfig(ctx context.Context, req QueryMerchantConfigRequest) (resp *QueryMerchantConfigResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	queryParams := neturl.Values{}
	queryParams.Add("sp_mchid", req.SpMchid)
	resp = new(QueryMerchantConfigResponse)
	path := strings.Replace(consts.QueryMerchantConfigPath, "{sub_mchid}", neturl.PathEscape(req.SubMchid), 1)
	result, err = a.doRequest(ctx, nethttp.MethodGet, path, queryParams, nil, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// AddReceiver 添加分账接收方
func (a *ApiService) AddReceiver(ctx context.Context, req AddReceiverRequest) (resp *ReceiverResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	resp = new(ReceiverResponse)
	result, err = a.doRequest(ctx, nethttp.MethodPost, consts.AddSplitReceiverPath, nil, req, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// DeleteReceiver 删除分账接收方
func (a *ApiService) DeleteReceiver(ctx context.Context, req DeleteReceiverRequest) (resp *DeleteReceiverResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	resp = new(DeleteReceiverResponse)
	result, err = a.doRequest(ctx, nethttp.MethodPost, consts.DeleteSplitReceiverPath, nil, req, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ParseProfitSharingNotify 解析服务商分账结果通知
func ParseProfitSharingNotify(ctx context.Context, handler *callback.Handler, request *nethttp.Request) (notifyReq *callback.Request, content *ProfitSharingNotify, err error) {
	if handler == nil {
		return nil, nil, fmt.Errorf("handler is required")
	}
	content = new(ProfitSharingNotify)
	notifyReq, err = handler.ParseCallback(ctx, request, content)
	if err != nil {
		return nil, nil, err
	}
	return notifyReq, content, nil
}

// ParseReceiverNotify 解析服务商分账接收方入账通知
func ParseReceiverNotify(ctx context.Context, handler *callback.Handler, request *nethttp.Request) (notifyReq *callback.Request, content *ReceiverNotify, err error) {
	if handler == nil {
		return nil, nil, fmt.Errorf("handler is required")
	}
	content = new(ReceiverNotify)
	notifyReq, err = handler.ParseCallback(ctx, request, content)
	if err != nil {
		return nil, nil, err
	}
	return notifyReq, content, nil
}

func (a *ApiService) doRequest(ctx context.Context, method, path string, queryParams neturl.Values, postBody interface{}, resp interface{}) (*client.APIResult, error) {
	r := &client.RequestEntity{
		Method:      method,
		QueryParams: queryParams,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
		PostBody:    postBody,
	}
	r.RequestPath = consts.DouyinPayServer + path

	result, err := a.Client.Request(ctx, r)
	if err != nil {
		return result, err
	}
	if resp == nil {
		return result, nil
	}
	if err = client.UnMarshalResponse(result.Response, resp); err != nil {
		return result, err
	}
	return result, nil
}
