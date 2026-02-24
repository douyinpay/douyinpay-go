package partnercontract

type PartnerQueryContractRequest struct {
	// 模版id
	PlanId int `json:"plan_id"`
	// 签约协议号
	OutContractCode string `json:"out_contract_code,omitempty"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
}

type PartnerQueryContractResponse struct {
	// 服务商应用Id
	SpAppid string `json:"sp_appid"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户应用Id
	SubAppid string `json:"sub_appid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
	// 代扣协议id
	ContractId string `json:"contract_id,omitempty"`
	// 模版id
	PlanId int `json:"plan_id"`
	// 委托签约协议号
	OutContractCode string `json:"out_contract_code,omitempty"`
	// 用户账户显示名称
	ContractDisplayAccount string `json:"contract_display_account,omitempty"`
	// 委托代扣协议状态
	ContractState string `json:"contract_state,omitempty"`
	// 协议到期时间
	ContractExpiredTime string `json:"contract_expired_time,omitempty"`
	// 解约信息
	ContractTerminateInfo ContractTerminateInfo `json:"contract_terminate_info,omitempty"`
	// 子商户AppID对应的用户OpenID
	SubOpenid string `json:"sub_openid"`
	// 用户OpenID
	SpOpenid string `json:"sp_openid"`
}

type ContractTerminateInfo struct {
	// 协议解约时间
	ContractTerminatedTime string `json:"contract_terminated_time,omitempty"`
	// 协议解约方式
	ContractTerminationMode string `json:"contract_termination_mode,omitempty"`
	// 协议解约备注
	ContractTerminationRemark string `json:"contract_termination_remark,omitempty"`
}

type PartnerTerminateContractRequest struct {
	//委托代扣模板ID
	PlanId int `json:"plan_id"`
	//商户签约协议号
	OutContractCode string `json:"out_contract_code,omitempty"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
	// 解约备注
	ContractTerminationRemark string `json:"contract_termination_remark,omitempty"`
}

type PartnerTerminateContractResponse struct {
	// 服务商应用Id
	SpAppid string `json:"sp_appid"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户应用Id
	SubAppid string `json:"sub_appid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
	// 代扣协议id
	ContractId string `json:"contract_id,omitempty"`
	// 模版id
	PlanId int `json:"plan_id"`
	// 委托签约协议号
	OutContractCode string `json:"out_contract_code,omitempty"`
	// 用户账户显示名称
	ContractDisplayAccount string `json:"contract_display_account,omitempty"`
	// 委托代扣协议状态
	ContractState string `json:"contract_state,omitempty"`
	// 协议签署时间
	ContractSignedTime string `json:"contract_signed_time,omitempty"`
	// 协议到期时间
	ContractExpiredTime string `json:"contract_expired_time,omitempty"`
	// 子商户AppID对应的用户OpenID
	SubOpenid string `json:"sub_openid,omitempty"`
	// 用户OpenID
	SpOpenid string `json:"sp_openid,omitempty"`
	//协议解约信息
	ContractTerminateInfo *ContractTerminateInfo `json:"contract_terminate_info,omitempty"`
}

// 签解约通知消息体
type SignContractNotify struct {
	// 服务商应用Id
	SpAppid string `json:"sp_appid"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户应用Id
	SubAppid string `json:"sub_appid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
	// 模版id
	PlanId int `json:"plan_id"`
	// 委托签约协议号
	OutContractCode string `json:"out_contract_code,omitempty"`
	// 子商户AppID对应的用户OpenID
	SubOpenid string `json:"sub_openid,omitempty"`
	// 用户OpenID
	SpOpenid string `json:"sp_openid,omitempty"`
	// 代扣协议id
	ContractId string `json:"contract_id,omitempty"`
	// 	协议到期时间
	ContractExpiredTime string `json:"contract_expired_time,omitempty"`
	//	用户账户展示名称
	ContractDisplayAccount string `json:"contract_display_account,omitempty"`
	// 委托代扣协议状态
	ContractState string `json:"contract_state,omitempty"`
	// 协议签署时间
	ContractSignedTime string `json:"contract_signed_time,omitempty"`
	//协议解约信息
	ContractTerminateInfo *ContractTerminateInfo `json:"contract_terminate_info,omitempty"`
}
