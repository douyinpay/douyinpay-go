package partnercontract

// PartnerQueryContractRequest 查询签约关系请求
type PartnerQueryContractRequest struct {
	// 字段含义：委托代扣模板 ID。
	// 格式规则：int。
	// 业务规则：商户在抖音平台申请的委托代扣模版 id。
	PlanId int `json:"plan_id"`
	// 字段含义：商户签约协议号。
	// 格式规则：string[1,64]。
	// 业务规则：商户发起签约时传入的签约协议号，商户侧需保证唯一。
	OutContractCode string `json:"out_contract_code,omitempty"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	SubMchid string `json:"sub_mchid"`
}

// PartnerQueryContractResponse 查询签约关系响应
type PartnerQueryContractResponse struct {
	// 字段含义：服务商应用 ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用 ID，全局唯一。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用 ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用 ID，全局唯一。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：委托代扣协议号。
	// 格式规则：string[1,64]。
	// 业务规则：委托代扣签约成功后，抖音返回的委托代扣协议号。
	ContractId string `json:"contract_id,omitempty"`
	// 字段含义：委托代扣模板 ID。
	// 格式规则：int。
	// 业务规则：商户在抖音平台申请的委托代扣模版 id。
	PlanId int `json:"plan_id"`
	// 字段含义：商户签约协议号。
	// 格式规则：string[1,64]。
	// 业务规则：商户发起签约时传入的签约协议号，商户侧需保证唯一。
	OutContractCode string `json:"out_contract_code,omitempty"`
	// 字段含义：用户账户展示名称。
	// 格式规则：string[1,128]。
	// 业务规则：签约用户的名称，用于页面展示，在签约时由商户传入。
	ContractDisplayAccount string `json:"contract_display_account,omitempty"`
	// 字段含义：委托代扣协议状态。
	// 格式规则：string[1,16]。
	// 业务规则：协议状态枚举值。
	// 枚举值：SIGNED：签约协议生效中, TERMINATED：生效的签约协议已被解约（此时协议已经到达终态，该协议无法再次进行签约；可更换协议号再发起签约）, FROZEN：协议被冻结。
	ContractState string `json:"contract_state,omitempty"`
	// 字段含义：协议到期时间。
	// 格式规则：string[1,32]，遵循 RFC3339 标准格式，格式为 yyyy-MM-DDTHH:mm:ss+TIMEZONE。
	// 业务规则：协议到期时间（目前协议时间为长期有效，可以忽略该字段）。例如：2015-05-20T13:29:35+08:00 表示北京时间 2015 年 5 月 20 日 13 点 29 分 35 秒。
	ContractExpiredTime string `json:"contract_expired_time,omitempty"`
	// 字段含义：协议解约信息。
	// 格式规则：object。
	// 业务规则：仅当 contract_state=TERMINATED 时，该值有效。
	ContractTerminateInfo ContractTerminateInfo `json:"contract_terminate_info,omitempty"`
	// 字段含义：子商户 AppID 对应的用户 OpenID。
	// 格式规则：string[1,128]。
	// 业务规则：商户 AppID 下的用户唯一标识。
	SubOpenid string `json:"sub_openid"`
	// 字段含义：用户 OpenID。
	// 格式规则：string[1,128]。
	// 业务规则：商户 AppID 下的用户唯一标识。
	SpOpenid string `json:"sp_openid"`
}

// ContractTerminateInfo 协议解约信息
type ContractTerminateInfo struct {
	// 字段含义：协议解约时间。
	// 格式规则：string[1,32]，遵循 RFC3339 标准格式，格式为 yyyy-MM-DDTHH:mm:ss+TIMEZONE。
	// 业务规则：仅当 contract_state=TERMINATED 时该值有效。例如：2015-05-20T13:29:35+08:00 表示北京时间 2015 年 5 月 20 日 13 点 29 分 35 秒。
	ContractTerminatedTime string `json:"contract_terminated_time,omitempty"`
	// 字段含义：协议解约方式。
	// 格式规则：string[1,64]。
	// 业务规则：解约发起方枚举值。
	// 枚举值：USER_TERMINATE：用户发起的解约, MCH_API_TERMINATE：商户通过 API 发起的解约, DOUPAY_WEB_TERMINATE：商户在商户平台发起的解约, CUSTOMER_SERVICE_TERMINATE：用户联系客服发起的解约, SYSTEM_TERMINATE：抖音支付系统主动发起的解约。
	ContractTerminationMode string `json:"contract_termination_mode,omitempty"`
	// 字段含义：解约备注。
	// 格式规则：string[1,128]。
	// 业务规则：解约原因的备注说明，如：签约信息有误，须重新签约。
	ContractTerminationRemark string `json:"contract_termination_remark,omitempty"`
}

// PartnerTerminateContractRequest 协议解约请求
type PartnerTerminateContractRequest struct {
	// 字段含义：委托代扣模板 ID。
	// 格式规则：int。
	// 业务规则：商户在抖音平台申请的委托代扣模版 id。
	PlanId int `json:"plan_id"`
	// 字段含义：商户签约协议号。
	// 格式规则：string[1,64]。
	// 业务规则：商户发起签约时传入的签约协议号，商户侧需保证唯一。
	OutContractCode string `json:"out_contract_code,omitempty"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：解约备注。
	// 格式规则：string[1,128]。
	// 业务规则：解约备注，传入解约原因。
	ContractTerminationRemark string `json:"contract_termination_remark,omitempty"`
}

// PartnerTerminateContractResponse 协议解约响应
type PartnerTerminateContractResponse struct {
	// 字段含义：服务商应用 ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用 ID，全局唯一。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用 ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用 ID，全局唯一。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：委托代扣协议号。
	// 格式规则：string[1,64]。
	// 业务规则：委托代扣签约成功后，抖音返回的委托代扣协议号。
	ContractId string `json:"contract_id,omitempty"`
	// 字段含义：委托代扣模板 ID。
	// 格式规则：int。
	// 业务规则：商户在抖音平台申请的委托代扣模版 id。
	PlanId int `json:"plan_id"`
	// 字段含义：商户签约协议号。
	// 格式规则：string[1,64]。
	// 业务规则：商户发起签约时传入的签约协议号，商户侧需保证唯一。
	OutContractCode string `json:"out_contract_code,omitempty"`
	// 字段含义：用户账户展示名称。
	// 格式规则：string[1,64]。
	// 业务规则：签约用户的名称，用于页面展示，在签约时由商户传入。
	ContractDisplayAccount string `json:"contract_display_account,omitempty"`
	// 字段含义：委托代扣协议状态。
	// 格式规则：string[1,16]。
	// 业务规则：协议状态枚举值。
	// 枚举值：SIGNED：签约协议生效中, TERMINATED：生效的签约协议已被解约（此时协议已经到达终态，该协议无法再次进行签约；可更换协议号再发起签约）, FROZEN：协议被冻结。
	ContractState string `json:"contract_state,omitempty"`
	// 字段含义：协议签署时间。
	// 格式规则：string[1,32]，遵循 RFC3339 标准格式，格式为 yyyy-MM-DDTHH:mm:ss+TIMEZONE。
	// 业务规则：协议签署时间。例如：2015-05-20T13:29:35+08:00 表示北京时间 2015 年 5 月 20 日 13 点 29 分 35 秒。
	ContractSignedTime string `json:"contract_signed_time,omitempty"`
	// 字段含义：协议到期时间。
	// 格式规则：string[1,32]，遵循 RFC3339 标准格式，格式为 yyyy-MM-DDTHH:mm:ss+TIMEZONE。
	// 业务规则：协议到期时间（目前协议时间为长期有效，可以忽略该字段）。例如：2015-05-20T13:29:35+08:00 表示北京时间 2015 年 5 月 20 日 13 点 29 分 35 秒。
	ContractExpiredTime string `json:"contract_expired_time,omitempty"`
	// 字段含义：子商户 AppID 对应的用户 OpenID。
	// 格式规则：string[1,128]。
	// 业务规则：商户 AppID 下的用户唯一标识。
	SubOpenid string `json:"sub_openid,omitempty"`
	// 字段含义：用户 OpenID。
	// 格式规则：string[1,128]。
	// 业务规则：商户 AppID 下的用户唯一标识。
	SpOpenid string `json:"sp_openid,omitempty"`
	// 字段含义：协议解约信息。
	// 格式规则：object。
	// 业务规则：仅当 contract_state=TERMINATED 时，该值有效。
	ContractTerminateInfo *ContractTerminateInfo `json:"contract_terminate_info,omitempty"`
}

// SignContractNotify 签解约通知消息体
type SignContractNotify struct {
	// 字段含义：服务商应用 ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用 ID，全局唯一。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用 ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用 ID，全局唯一。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：委托代扣模板 ID。
	// 格式规则：int。
	// 业务规则：商户在抖音平台申请的委托代扣模版 id。
	PlanId int `json:"plan_id"`
	// 字段含义：商户签约协议号。
	// 格式规则：string[1,64]。
	// 业务规则：商户发起签约时传入的签约协议号，商户侧需保证唯一。
	OutContractCode string `json:"out_contract_code,omitempty"`
	// 字段含义：子商户 AppID 对应的用户 OpenID。
	// 格式规则：string[1,128]。
	// 业务规则：商户 AppID 下的用户唯一标识。
	SubOpenid string `json:"sub_openid,omitempty"`
	// 字段含义：用户 OpenID。
	// 格式规则：string[1,128]。
	// 业务规则：商户 AppID 下的用户唯一标识。
	SpOpenid string `json:"sp_openid,omitempty"`
	// 字段含义：委托代扣协议号。
	// 格式规则：string[1,64]。
	// 业务规则：委托代扣签约成功后，抖音返回的委托代扣协议号。
	ContractId string `json:"contract_id,omitempty"`
	// 字段含义：协议到期时间。
	// 格式规则：string[1,32]，遵循 RFC3339 标准格式，格式为 yyyy-MM-DDTHH:mm:ss+TIMEZONE。
	// 业务规则：协议到期时间（目前协议时间为长期有效，可以忽略该字段）。例如：2015-05-20T13:29:35+08:00 表示北京时间 2015 年 5 月 20 日 13 点 29 分 35 秒。
	ContractExpiredTime string `json:"contract_expired_time,omitempty"`
	// 字段含义：用户账户展示名称。
	// 格式规则：string[1,128]。
	// 业务规则：签约用户的名称，用于页面展示，在签约时由商户传入。
	ContractDisplayAccount string `json:"contract_display_account,omitempty"`
	// 字段含义：委托代扣协议状态。
	// 格式规则：string[1,16]。
	// 业务规则：协议状态枚举值。
	// 枚举值：SIGNED：签约协议生效中, TERMINATED：生效的签约协议已被解约（此时协议已经到达终态，该协议无法再次进行签约；可更换协议号再发起签约）, FROZEN：协议被冻结。
	ContractState string `json:"contract_state,omitempty"`
	// 字段含义：协议签署时间。
	// 格式规则：string[1,32]，遵循 RFC3339 标准格式，格式为 yyyy-MM-DDTHH:mm:ss+TIMEZONE。
	// 业务规则：协议签署时间。例如：2015-05-20T13:29:35+08:00 表示北京时间 2015 年 5 月 20 日 13 点 29 分 35 秒。
	ContractSignedTime string `json:"contract_signed_time,omitempty"`
	// 字段含义：协议解约信息。
	// 格式规则：object。
	// 业务规则：仅当 contract_state=TERMINATED 时，该值有效。
	ContractTerminateInfo *ContractTerminateInfo `json:"contract_terminate_info,omitempty"`
}
