package cashier

import (
	"context"
	"fmt"
	nethttp "net/http"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

type ApiCashierService services.Service

// PrePayConsult 前置咨询
//
// 商户在支付前咨询当前用户的支付渠道可用性与营销内容，用于在商户侧提前展示营销文案。
// 当前支持「普通支付」（APP支付、H5支付、JSAPI支付、Native支付）、「签约并支付」、「免密支付」、
// 「商户代扣」、「先享后付」，具体的 product_code、commerical_product_code、trade_type 组合见官方文档映射表。
func (a *ApiCashierService) PrePayConsult(ctx context.Context, req PrePayConsultRequest) (resp *PrePayConsultResponse, result *client.APIResult, err error) {
	if req.Appid == "" {
		return nil, nil, fmt.Errorf("PrePayConsultRequest required field `Appid` is empty")
	}
	if req.Mchid == "" {
		return nil, nil, fmt.Errorf("PrePayConsultRequest required field `Mchid` is empty")
	}

	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.PrePayConsultPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(PrePayConsultResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
