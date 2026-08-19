package partnercontract

// 官方文档：
// - 查询签约关系：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/685b4b8777f5de0546f6af5d
// - 协议解约：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/685b4b8af6235c04f84bdea6
import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strconv"
	"strings"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

// ApiPartnerContractOrderService 服务商签约关系服务。
//
// 接口路径：
// - 查询签约关系：GET /v1/agreementauth/partner/contracts/plan-id/{plan_id}/out-contract-code/{out_contract_code}
// - 协议解约：POST /v1/agreementauth/partner/contracts/plan-id/{plan_id}/out-contract-code/{out_contract_code}/terminate
type ApiPartnerContractOrderService services.Service

// PartnerQueryContract 查询签约关系。
// 接口路径：GET /v1/agreementauth/partner/contracts/plan-id/{plan_id}/out-contract-code/{out_contract_code}。
func (a *ApiPartnerContractOrderService) PartnerQueryContract(ctx context.Context, req PartnerQueryContractRequest) (resp *PartnerQueryContractResponse, result *client.APIResult, err error) {
	// 参数校验
	if req.PlanId < 1 {
		return nil, nil, fmt.Errorf("field `PlanId` is required and must be specified in PartnerQueryContractRequest")
	}
	if req.OutContractCode == "" {
		return nil, nil, fmt.Errorf("field `OutContractCode` is required and must be specified in PartnerQueryContractRequest")
	}
	if req.SpMchid == "" {
		return nil, nil, fmt.Errorf("field `SpMchid` is required and must be specified in PartnerQueryContractRequest")
	}
	if req.SubMchid == "" {
		return nil, nil, fmt.Errorf("field `SubMchid` is required and must be specified in PartnerQueryContractRequest")
	}

	localVarQueryParams := neturl.Values{}
	localVarQueryParams.Add("sp_mchid", req.SpMchid)
	localVarQueryParams.Add("sub_mchid", req.SubMchid)

	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		QueryParams: localVarQueryParams,
		Header:      nethttp.Header{},
	}
	//获取请求路径
	rawURL := consts.DouyinPayServer + consts.PartnerQueryContractPath
	currURL := strings.Replace(rawURL, "{plan_id}", neturl.PathEscape(strconv.Itoa(req.PlanId)), -1)
	r.RequestPath = strings.Replace(currURL, "{out_contract_code}", neturl.PathEscape(req.OutContractCode), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(PartnerQueryContractResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// PartnerTerminateContract 协议解约。
// 接口路径：POST /v1/agreementauth/partner/contracts/plan-id/{plan_id}/out-contract-code/{out_contract_code}/terminate。
func (a *ApiPartnerContractOrderService) PartnerTerminateContract(ctx context.Context, req PartnerTerminateContractRequest) (resp *PartnerTerminateContractResponse, result *client.APIResult, err error) {
	// 参数校验
	if req.PlanId < 1 {
		return nil, nil, fmt.Errorf("field `PlanId` is required and must be specified in PartnerTerminateContractRequest")
	}
	if req.OutContractCode == "" {
		return nil, nil, fmt.Errorf("field `OutContractCode` is required and must be specified in PartnerTerminateContractRequest")
	}
	if req.SpMchid == "" {
		return nil, nil, fmt.Errorf("field `SpMchid` is required and must be specified in PartnerTerminateContractRequest")
	}
	if req.SubMchid == "" {
		return nil, nil, fmt.Errorf("field `SubMchid` is required and must be specified in PartnerTerminateContractRequest")
	}
	if req.ContractTerminationRemark == "" {
		return nil, nil, fmt.Errorf("field `ContractTerminationRemark` is required and must be specified in PartnerTerminateContractRequest")
	}
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	//获取请求路径
	rawURL := consts.DouyinPayServer + consts.PartnerTerminateContractPath
	currURL := strings.Replace(rawURL, "{plan_id}", neturl.PathEscape(strconv.Itoa(req.PlanId)), -1)
	r.RequestPath = strings.Replace(currURL, "{out_contract_code}", neturl.PathEscape(req.OutContractCode), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}

	resp = new(PartnerTerminateContractResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
