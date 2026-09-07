package profitsharing

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

// ProfitSharingApiService 直连商户分账服务。
//
// 官方文档：
//   - 请求分账：POST /v1/trade/profitsharing/orders
//     https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/694937c648fd720521ddf214
//   - 查询分账结果：GET /v1/trade/profitsharing/orders/{out_order_no}
//     https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/6949383d7f605b05358e7cc7
//   - 请求分账回退：POST /v1/trade/profitsharing/return-orders
//     https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/694938cbb9dd320544606cf7
//   - 查询分账回退结果：GET /v1/trade/profitsharing/return-orders/{out_return_no}
//     https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/69493917b9dd3205446070a0
//   - 完结分账：POST /v1/trade/profitsharing/finish-orders
//     https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/69c125ae5e28e105291b6c3d
//   - 查询剩余待分金额：GET /v1/trade/profitsharing/order/{transaction_id}/amounts
//     https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/69c125c573c1f605a3d496d6
//   - 添加分账接收方：POST /v1/trade/profitsharing/receivers/add
//     https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/6949394b7f605b05358e83b4
//   - 删除分账接收方：POST /v1/trade/profitsharing/receivers/delete
//     https://pay.douyinpay.com/wiki/639fd48f17c2f3021d237f61/694939886002be055131fa0a
type ProfitSharingApiService services.Service

// SplitFund 请求分账。
func (a *ProfitSharingApiService) SplitFund(ctx context.Context, req SplitFundRequest) (resp *SplitFundResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	resp = new(SplitFundResponse)
	result, err = a.doRequest(ctx, nethttp.MethodPost, consts.SplitFundPath, nil, req, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QuerySplitFund 查询分账结果。
func (a *ProfitSharingApiService) QuerySplitFund(ctx context.Context, req QuerySplitFundRequest) (resp *QuerySplitFundResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	queryParams := neturl.Values{}
	queryParams.Add("mchid", req.Mchid)
	if req.TransactionID != "" {
		queryParams.Add("transaction_id", req.TransactionID)
	}
	if req.OrderID != "" {
		queryParams.Add("order_id", req.OrderID)
	}
	resp = new(QuerySplitFundResponse)
	path := strings.Replace(consts.QuerySplitFundPath, "{out_order_no}", neturl.PathEscape(req.OutOrderNo), 1)
	result, err = a.doRequest(ctx, nethttp.MethodGet, path, queryParams, nil, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ReturnSplitFund 请求分账回退。
func (a *ProfitSharingApiService) ReturnSplitFund(ctx context.Context, req ReturnSplitFundRequest) (resp *ReturnSplitFundResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	resp = new(ReturnSplitFundResponse)
	result, err = a.doRequest(ctx, nethttp.MethodPost, consts.ReturnSplitFundPath, nil, req, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryReturnSplitFund 查询分账回退结果。
func (a *ProfitSharingApiService) QueryReturnSplitFund(ctx context.Context, req QueryReturnSplitFundRequest) (resp *ReturnSplitFundResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	queryParams := neturl.Values{}
	queryParams.Add("mchid", req.Mchid)
	if req.OutOrderNo != "" {
		queryParams.Add("out_order_no", req.OutOrderNo)
	}
	resp = new(ReturnSplitFundResponse)
	path := strings.Replace(consts.QueryReturnSplitFundPath, "{out_return_no}", neturl.PathEscape(req.OutReturnNo), 1)
	result, err = a.doRequest(ctx, nethttp.MethodGet, path, queryParams, nil, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// FinishSplitFund 完结分账。
func (a *ProfitSharingApiService) FinishSplitFund(ctx context.Context, req FinishSplitFundRequest) (resp *FinishSplitFundResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	resp = new(FinishSplitFundResponse)
	result, err = a.doRequest(ctx, nethttp.MethodPost, consts.FinishSplitFundPath, nil, req, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryUnsplitAmount 查询剩余待分金额。
func (a *ProfitSharingApiService) QueryUnsplitAmount(ctx context.Context, req QueryUnsplitAmountRequest) (resp *QueryUnsplitAmountResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	queryParams := neturl.Values{}
	queryParams.Add("mchid", req.Mchid)
	resp = new(QueryUnsplitAmountResponse)
	path := strings.Replace(consts.QueryUnsplitAmountPath, "{transaction_id}", neturl.PathEscape(req.TransactionID), 1)
	result, err = a.doRequest(ctx, nethttp.MethodGet, path, queryParams, nil, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// AddSplitReceiver 添加分账接收方。
func (a *ProfitSharingApiService) AddSplitReceiver(ctx context.Context, req AddSplitReceiverRequest) (resp *SplitReceiverResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	resp = new(SplitReceiverResponse)
	result, err = a.doRequest(ctx, nethttp.MethodPost, consts.AddSplitReceiverPath, nil, req, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// DeleteSplitReceiver 删除分账接收方。
func (a *ProfitSharingApiService) DeleteSplitReceiver(ctx context.Context, req DeleteSplitReceiverRequest) (resp *DeleteSplitReceiverResponse, result *client.APIResult, err error) {
	if err = req.Validate(); err != nil {
		return nil, nil, err
	}
	resp = new(DeleteSplitReceiverResponse)
	result, err = a.doRequest(ctx, nethttp.MethodPost, consts.DeleteSplitReceiverPath, nil, req, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ParseProfitSharingNotify 解析直连商户分账结果通知。
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

// ParseReceiverNotify 解析直连商户分账动态通知。
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

func (a *ProfitSharingApiService) doRequest(ctx context.Context, method, path string, queryParams neturl.Values, postBody interface{}, resp interface{}) (*client.APIResult, error) {
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
	if err = client.UnMarshalResponse(result.Response, resp); err != nil {
		return result, err
	}
	return result, nil
}
