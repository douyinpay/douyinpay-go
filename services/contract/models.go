package contract

type QueryContractRequest struct {
	Mchid        string `json:"mchid"`                       // 商户号
	Appid        string `json:"appid"`                       // 商户appid
	ContractId   string `json:"contract_id,omitempty"`       // 抖音支付代扣协议号,ContractId 和 PlanId+ContractCode 选填其一
	PlanId       string `json:"plan_id,omitempty"`           // 代扣模板id
	ContractCode string `json:"out_contract_code,omitempty"` // 商户生成的代扣协议号
}

type QueryContractResponse struct {
	Mchid                     string `json:"mchid,omitempty"`                       // 三方商户号
	Appid                     string `json:"appid,omitempty"`                       // 商户appid
	ContractId                string `json:"contract_id,omitempty"`                 // 签约成功后，抖音支付返回的委托代扣协议id
	PlanId                    string `json:"plan_id,omitempty"`                     // 协议模板id
	RequestSerial             int64  `json:"request_serial,omitempty"`              // 商户请求签约时的序列号，商户侧须唯一。序列号主要用于排序，不作为查询条件，纯数字,范围不能超过int64的范围（9223372036854775807）。
	ContractCode              string `json:"out_contract_code,omitempty"`           // 商户生成的代扣协议号
	ContractDisplayAccount    string `json:"contract_display_account,omitempty"`    // 用户账户展示名称，签约用户的名称，用于页面展示
	ContractStatus            int64  `json:"contract_status,omitempty"`             // 协议状态，枚举值： 0：已签约 1：未签约 9：签约进行中
	ContractSignedTime        string `json:"contract_signed_time,omitempty"`        // 协议签署时间
	ContractExpiredTime       string `json:"contract_expired_time,omitempty"`       // 协议到期时间（目前协议时间为长期有效，可以忽略该字段）
	ContractTerminatedTime    string `json:"contract_terminated_time,omitempty"`    // 协议解约时间,当contract_state=1时，该值有效
	ContractTerminationMode   int64  `json:"contract_termination_mode,omitempty"`   // 当contract_state=1时，该值有效 1：有效期过自动解约（预留功能） 2：用户主动解约 3：商户API解约 4：商户平台解约 5：注销（用户微信账户注销） 7：用户联系客服发起的解约
	ContractTerminationRemark string `json:"contract_termination_remark,omitempty"` // 解约备注 当contract_state=1时，该值有效
	OpenId                    string `json:"openid,omitempty"`                      // 用户标识
}

// 申请解约
type DeleteContractRequest struct {
	Mchid                     string `json:"mchid"`                       // 三方商户号
	Appid                     string `json:"appid"`                       // 三方商户appid
	PlanId                    string `json:"plan_id,omitempty"`           // 签约模版ID
	ContractCode              string `json:"out_contract_code,omitempty"` // 外部签约协议号
	ContractId                string `json:"contract_id,omitempty"`       // 抖音支付代扣协议号
	ContractTerminationRemark string `json:"contract_termination_remark"` // 委托代扣签约成功后由微信返回的委托代扣协议id，选择contract_id解约
}

type DeleteContractResponse struct {
	Appid        string `json:"appid"`
	ContractId   string `json:"contract_id,omitempty"`       // 抖音支付代扣协议号
	PlanId       string `json:"plan_id,omitempty"`           // 协议模板id
	ContractCode string `json:"out_contract_code,omitempty"` // 商户生成的代扣协议号
	ResultCode   string `json:"result_code,omitempty"`       // SUCCESS/FAIL，标识业务处理结果  解约结果以查询或回调为准

}

// 签解约通知消息体
type SignContractNotifyBody struct {
	// 服务商户号
	Mchid string `json:"mchid"`
	// 委托签约协议号
	OutContractCode string `json:"out_contract_code,omitempty"`
	// 模版id
	PlanId string `json:"plan_id,omitempty"`
	// 子商户AppID对应的用户OpenID
	Openid string `json:"openid,omitempty"`
	// 变更类型
	ChangeType string `json:"change_type,omitempty"`
	// 操作时间
	OperateTime string `json:"operate_time,omitempty"`
	// 代扣协议id
	ContractId string `json:"contract_id,omitempty"`
	// 	协议到期时间
	ContractExpiredTime string `json:"contract_expired_time,omitempty"`
	// 请求序列号
	RequestSerial int64 `json:"request_serial,omitempty"`
}

type PreEntrustWebRequest struct {
	Mchid                  string `json:"mchid"`                    // 商户号
	Appid                  string `json:"appid"`                    // 商户appid
	OutContractCode        string `json:"out_contract_code"`        // 签约协议号
	PlanId                 string `json:"plan_id"`                  // 协议模板id
	RequestSerial          int64  `json:"request_serial"`           // 商户请求签约时的序列号，要求唯一性。禁止使用0开头，序列号主要用于排序，不作为查询条件，纯数字，范围不能超过int64的范围（9223372036854775807）
	ContractDisplayAccount string `json:"contract_display_account"` // 签约用户的名称，用于页面展示
	NotifyUrl              string `json:"notify_url"`               // 用于接收签约成功消息的回调通知地址，以http或https开头，通知url必须为外网可访问的url，不能携带参数
	ContractExt            string `json:"contract_ext"`             // 签约拓展参数，json格式，注：仅与抖音支付线下约定后使用
}

type PreEntrustWebResponse struct {
	PreEntrustWebId string `json:"pre_entrustweb_id"` // 预签约id，两个小时内有效
}

type H5EntrustwebRequest struct {
	Mchid                  string `json:"mchid"`                    // 商户号
	Appid                  string `json:"appid"`                    // 商户appid
	OutContractCode        string `json:"out_contract_code"`        // 签约协议号
	PlanId                 string `json:"plan_id"`                  // 协议模板id
	RequestSerial          int64  `json:"request_serial"`           // 商户请求签约时的序列号，要求唯一性。禁止使用0开头，序列号主要用于排序，不作为查询条件，纯数字，范围不能超过int64的范围（9223372036854775807）
	ContractDisplayAccount string `json:"contract_display_account"` // 签约用户的名称，用于页面展示
	ContractExt            string `json:"contract_ext"`             // 签约拓展参数，json格式，注：仅与抖音支付线下约定后使用
	NotifyUrl              string `json:"notify_url"`               // 用于接收签约成功消息的回调通知地址，以http或https开头，通知url必须为外网可访问的url，不能携带参数
	Timestamp              string `json:"timestamp"`                // 系统当前时间
	ClientIp               string `json:"client_ip"`                // 用户客户端ip地址
}

type H5EntrustwebResponse struct {
	RedirectUrl string `json:"redirect_url"` // 有效期十分钟，用户通过此路径跳转抖音签约支付页面
}
