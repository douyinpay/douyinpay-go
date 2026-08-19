package partner

// 官方文档：
// - 预约扣费：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/685b4b9077f5de0546f6af87
// - 预约扣费结果查询：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/685b4b97a2b24a063d54b442

import (
	"context"
	"fmt"
	nethttp "net/http"
	netUrl "net/url"
	"strings"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	contractorder "github.com/douyinpay/douyinpay-go/services/deduct"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

// ApiPartnerDeductService 服务商预约扣费服务。
//
// 接口路径：
// - 预约扣费：POST /v1/agreementauth/partner/schedules/contract-id/{contract_id}/schedule
// - 预约扣费结果查询：GET /v1/agreementauth/partner/schedules/contract-id/{contract_id}
type ApiPartnerDeductService services.Service

// PartnerContractSchedule 预约扣费。
// 接口路径：POST /v1/agreementauth/partner/schedules/contract-id/{contract_id}/schedule
func (a *ApiPartnerDeductService) PartnerContractSchedule(ctx context.Context, req contractorder.PartnerContractScheduleRequest) (resp *contractorder.PartnerContractScheduleResponse, result *client.APIResult, err error) {
	// 参数校验
	if req.ContractId == "" {
		return nil, nil, fmt.Errorf("field `ContractId` is required and must be specified in PartnerContractScheduleRequest")
	}
	if req.SpMchid == "" {
		return nil, nil, fmt.Errorf("field `SpMchid` is required and must be specified in PartnerContractScheduleRequest")
	}
	if req.SubMchid == "" {
		return nil, nil, fmt.Errorf("field `SubMchid` is required and must be specified in PartnerContractScheduleRequest")
	}
	if req.ScheduleAmount.Total < 0 {
		return nil, nil, fmt.Errorf("field `ScheduleAmount#Total` is required and must be specified in PartnerContractScheduleRequest")
	}
	if req.ScheduleAmount.Currency == "" {
		return nil, nil, fmt.Errorf("field `ScheduleAmount#Currency` is required and must be specified in PartnerContractScheduleRequest")
	}
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	//获取请求路径
	rawURL := consts.DouyinPayServer + consts.PartnerContractSchedulePath
	r.RequestPath = strings.Replace(rawURL, "{contract_id}", netUrl.PathEscape(req.ContractId), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(contractorder.PartnerContractScheduleResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// PartnerContractScheduleQuery 预约扣费结果查询。
// 接口路径：GET /v1/agreementauth/partner/schedules/contract-id/{contract_id}。
func (a *ApiPartnerDeductService) PartnerContractScheduleQuery(ctx context.Context, req contractorder.PartnerContractScheduleQueryRequest) (resp *contractorder.PartnerContractScheduleQueryResponse, result *client.APIResult, err error) {
	// 参数校验
	if req.ContractId == "" {
		return nil, nil, fmt.Errorf("field `ContractId` is required and must be specified in PartnerContractScheduleQueryRequest")
	}
	if req.SpMchid == "" {
		return nil, nil, fmt.Errorf("field `SpMchid` is required and must be specified in PartnerContractScheduleQueryRequest")
	}
	if req.SubMchid == "" {
		return nil, nil, fmt.Errorf("field `SubMchid` is required and must be specified in PartnerContractScheduleQueryRequest")
	}

	localVarQueryParams := netUrl.Values{}
	localVarQueryParams.Add("sp_mchid", req.SpMchid)
	localVarQueryParams.Add("sub_mchid", req.SubMchid)

	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		QueryParams: localVarQueryParams,
		Header:      nethttp.Header{},
	}
	//获取请求路径
	rawURL := consts.DouyinPayServer + consts.PartnerContractScheduleQueryPath
	r.RequestPath = strings.Replace(rawURL, "{contract_id}", netUrl.PathEscape(req.ContractId), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(contractorder.PartnerContractScheduleQueryResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
