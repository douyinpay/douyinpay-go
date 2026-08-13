package contract

import (
	"context"
	"fmt"
	nethttp "net/http"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

// ApiContractService 直连商户签约关系服务。
//
// 接口路径：
// - 查询签约协议：POST /v1/member/querycontract
// - 删除签约协议：POST /v1/member/deletecontract
// - APP 预签约下单：POST /v1/agreementauth/preentrustweb
// - H5 预签约下单：POST /v1/agreementauth/h5entrustweb
type ApiContractService services.Service

// QueryContract 查询代扣签约协议。
// 接口路径：POST /v1/member/querycontract。
func (a *ApiContractService) QueryContract(ctx context.Context, req QueryContractRequest) (resp *QueryContractResponse, result *client.APIResult, err error) {
	// 参数校验
	if req.Mchid == "" {
		return nil, nil, fmt.Errorf("field `Mchid` is required and must be specified in QueryContractRequest")
	}
	if req.Appid == "" {
		return nil, nil, fmt.Errorf("field `Appid` is required and must be specified in QueryContractRequest")
	}
	if req.ContractId == "" && (req.ContractCode == "" || req.PlanId == "") {
		return nil, nil, fmt.Errorf("field `ContractId` or [`ContractCode`&'PlanId'] is required and must be specified in QueryContractRequest")
	}
	// 发送请求
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.QueryContractPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	// 解析结果
	resp = new(QueryContractResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// DeleteContract 解除代扣签约协议。
// 接口路径：POST /v1/member/deletecontract。
func (a *ApiContractService) DeleteContract(ctx context.Context, req DeleteContractRequest) (resp *DeleteContractResponse, result *client.APIResult, err error) {
	// 参数校验
	if req.Mchid == "" {
		return nil, nil, fmt.Errorf("field `Mchid` is required and must be specified in DeleteContractRequest")
	}
	if req.Appid == "" {
		return nil, nil, fmt.Errorf("field `Appid` is required and must be specified in DeleteContractRequest")
	}
	if req.ContractId == "" && (req.ContractCode == "" || req.PlanId == "") {
		return nil, nil, fmt.Errorf("field `ContractId` or [`ContractCode`&'PlanId'] is required and must be specified in DeleteContractRequest")
	}
	if req.ContractTerminationRemark == "" {
		return nil, nil, fmt.Errorf("field `ContractTerminationRemark` is required and must be specified in DeleteContractRequest")
	}
	// 发送请求
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.DeleteContractPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	// 解析结果
	resp = new(DeleteContractResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// PreEntrustWeb 独立签约下单。
// 接口路径：POST /v1/agreementauth/preentrustweb。
func (a *ApiContractService) PreEntrustWeb(ctx context.Context, req PreEntrustWebRequest) (resp *PreEntrustWebResponse, result *client.APIResult, err error) {
	// 参数校验
	if req.Mchid == "" {
		return nil, nil, fmt.Errorf("field `Mchid` is required and must be specified in PreEntrustWebRequest")
	}
	if req.Appid == "" {
		return nil, nil, fmt.Errorf("field `Appid` is required and must be specified in PreEntrustWebRequest")
	}
	if req.OutContractCode == "" {
		return nil, nil, fmt.Errorf("field `OutContractCode` is required and must be specified in PreEntrustWebRequest")
	}
	if req.PlanId == "" {
		return nil, nil, fmt.Errorf("field `PlanId` is required and must be specified in PreEntrustWebRequest")
	}
	if req.RequestSerial < 1 {
		return nil, nil, fmt.Errorf("field `RequestSerial` is required and must be specified in PreEntrustWebRequest")
	}
	if req.ContractDisplayAccount == "" {
		return nil, nil, fmt.Errorf("field `ContractDisplayAccount` is required and must be specified in PreEntrustWebRequest")
	}
	if req.NotifyUrl == "" {
		return nil, nil, fmt.Errorf("field `NotifyUrl` is required and must be specified in PreEntrustWebRequest")
	}
	// 发送请求
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.PreentrustWebPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	// 解析结果
	resp = new(PreEntrustWebResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// H5Entrustweb H5 预签约。
// 接口路径：POST /v1/agreementauth/h5entrustweb。
func (a *ApiContractService) H5Entrustweb(ctx context.Context, req H5EntrustwebRequest) (resp *H5EntrustwebResponse, result *client.APIResult, err error) {
	// 参数校验
	if req.Mchid == "" {
		return nil, nil, fmt.Errorf("field `Mchid` is required and must be specified in H5EntrustwebRequest")
	}
	if req.Appid == "" {
		return nil, nil, fmt.Errorf("field `Appid` is required and must be specified in H5EntrustwebRequest")
	}
	if req.OutContractCode == "" {
		return nil, nil, fmt.Errorf("field `OutContractCode` is required and must be specified in H5EntrustwebRequest")
	}
	if req.PlanId == "" {
		return nil, nil, fmt.Errorf("field `PlanId` is required and must be specified in H5EntrustwebRequest")
	}
	if req.RequestSerial < 1 {
		return nil, nil, fmt.Errorf("field `RequestSerial` is required and must be specified in H5EntrustwebRequest")
	}
	if req.ContractDisplayAccount == "" {
		return nil, nil, fmt.Errorf("field `ContractDisplayAccount` is required and must be specified in H5EntrustwebRequest")
	}
	if req.NotifyUrl == "" {
		return nil, nil, fmt.Errorf("field `NotifyUrl` is required and must be specified in H5EntrustwebRequest")
	}
	if req.ClientIp == "" {
		return nil, nil, fmt.Errorf("field `ClientIp` is required and must be specified in H5EntrustwebRequest")
	}
	if req.Timestamp == "" {
		return nil, nil, fmt.Errorf("field `Timestamp` is required and must be specified in H5EntrustwebRequest")
	}
	// 发送请求
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.H5EntrustwebPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	// 解析结果
	resp = new(H5EntrustwebResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
