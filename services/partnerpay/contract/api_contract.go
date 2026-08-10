package contract

// 官方文档：
// - 签约并支付下单：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/685b4b7975a00b04c8aa2fb4
// - 申请扣款：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/685b4b7e75a00b04c8aa2fd1

import (
	"context"
	"fmt"
	nethttp "net/http"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

type ApiPartnerContractService services.Service

// PartnerContractOrder 支付中签约下单
func (a *ApiPartnerContractService) PartnerContractOrder(ctx context.Context, req PartnerContractOrderRequest) (resp *PartnerContractOrderResponse, result *client.APIResult, err error) {
	err = ValidParam(req)
	if err != nil {
		return nil, nil, err
	}
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.PartnerContractOrderPath
	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	resp = new(PartnerContractOrderResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// PartnerPayApply 申请扣款
func (a *ApiPartnerContractService) PartnerPayApply(ctx context.Context, req PartnerPayApplyRequest) (resp *PartnerPayApplyResponse, result *client.APIResult, err error) {
	if req.SpAppid == "" {
		return nil, nil, fmt.Errorf("field `SpAppid` is required and must be specified in PartnerPayApplyRequest")
	}
	if req.SpMchid == "" {
		return nil, nil, fmt.Errorf("field `SpMchid` is required and must be specified in PartnerPayApplyRequest")
	}
	if req.SubMchid == "" {
		return nil, nil, fmt.Errorf("field `SubMchid` is required and must be specified in PartnerPayApplyRequest")
	}
	if req.OutTradeNo == "" {
		return nil, nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in PartnerPayApplyRequest")
	}
	if req.ContractId == "" {
		return nil, nil, fmt.Errorf("field `ContractId` is required and must be specified in PartnerPayApplyRequest")
	}
	if req.TradeType == "" {
		return nil, nil, fmt.Errorf("field `TradeType` is required and must be specified in PartnerPayApplyRequest")
	}
	if req.Description == "" {
		return nil, nil, fmt.Errorf("field `Description` is required and must be specified in PartnerPayApplyRequest")
	}
	if req.NotifyUrl == "" {
		return nil, nil, fmt.Errorf("field `NotifyUrl` is required and must be specified in PartnerPayApplyRequest")
	}
	if req.Amount == nil || req.Amount.Total < 0 {
		return nil, nil, fmt.Errorf("field `Amount or Amount#Total` is required and must be specified in PartnerPayApplyRequest")
	}
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.PartnerPayApplyPath
	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	resp = new(PartnerPayApplyResponse)
	err = client.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func ValidParam(req PartnerContractOrderRequest) error {
	if req.SpMchid == "" {
		return fmt.Errorf("field `SpMchid` is required and must be specified in PartnerContractOrderRequest")
	}
	if req.SpAppid == "" {
		return fmt.Errorf("field `SpAppid` is required and must be specified in PartnerContractOrderRequest")
	}
	if (req.Payer != nil && req.Payer.SubOpenid != "") && req.SubAppid == "" {
		return fmt.Errorf("[`SubOpenid`&'SubAppid'] is required and must be specified in PartnerContractOrderRequest")
	}
	if req.SubMchid == "" {
		return fmt.Errorf("field `SubMchid` is required and must be specified in PartnerContractOrderRequest")
	}
	if req.Description == "" {
		return fmt.Errorf("field `Description` is required and must be specified in PartnerContractOrderRequest")
	}
	if req.OutTradeNo == "" {
		return fmt.Errorf("field `OutTradeNo` is required and must be specified in PartnerContractOrderRequest")
	}
	if req.TradeType == "" {
		return fmt.Errorf("field `TradeType` is required and must be specified in PartnerContractOrderRequest")
	}
	if req.NotifyUrl == "" {
		return fmt.Errorf("field `NotifyUrl` is required and must be specified in PartnerContractOrderRequest")
	}
	if req.Amount == nil || req.Amount.Total < 0 {
		return fmt.Errorf("field `Amount or Amount#Total` is required and must be specified in PartnerContractOrderRequest")
	}
	if req.ContractInfo.ContractMchId == "" {
		return fmt.Errorf("field `ContractInfo#ContractMchId` is required and must be specified in PartnerContractOrderRequest")
	}
	if req.ContractInfo.PlanId == "" {
		return fmt.Errorf("field `ContractInfo#PlanId` is required and must be specified in PartnerContractOrderRequest")
	}
	if req.ContractInfo.OutContractCode == "" {
		return fmt.Errorf("field `ContractInfo#OutContractCode` is required and must be specified in PartnerContractOrderRequest")
	}
	if req.ContractInfo.RequestSerial < 1 {
		return fmt.Errorf("field `ContractInfo#RequestSerial` is invalid and must be specified in PartnerContractOrderRequest")
	}
	if req.ContractInfo.ContractDisplayAccount == "" {
		return fmt.Errorf("field `ContractInfo#ContractDisplayAccount` is invalid and must be specified in PartnerContractOrderRequest")
	}
	if req.ContractInfo.ContractNotifyUrl == "" {
		return fmt.Errorf("field `ContractInfo#ContractNotifyUrl` is invalid and must be specified in PartnerContractOrderRequest")
	}
	return nil
}
