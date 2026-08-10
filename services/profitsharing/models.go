package profitsharing

import "fmt"

// SplitFundRequest 请求分账请求参数
type SplitFundRequest struct {
	// 应用ID
	// 字段含义：商户在抖音开放平台申请的应用ID，全局唯一。
	// 格式规则：长度为 1-32。
	// 业务规则：商户需确保该 appid 和 mchid 有绑定关系，且和下单接口的 appid 保持一致。
	// 示例：awofz9bncda6w2w4
	Appid string `json:"appid,omitempty"`
	// 直连商户号
	// 字段含义：直连商户的商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：请求分账时必填。
	// 示例：6020221212167701
	Mchid string `json:"mchid,omitempty"`
	// 抖音支付订单号
	// 字段含义：原支付交易对应的交易订单号。
	// 格式规则：长度为 1-32。
	// 业务规则：请求分账时必填。
	// 示例：TP2022101317144741443210681000
	TransactionID string `json:"transaction_id,omitempty"`
	// 商户分账单号
	// 字段含义：商户系统内部的分账单号。
	// 格式规则：长度为 6-32，只能是数字、大小写字母、`_`、`-`、`*`。
	// 业务规则：在商户系统内部唯一；同一分账单号多次请求等同一次。
	// 示例：OUT_1666688488
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 分账接收方列表
	// 字段含义：分账接收方列表。
	// 格式规则：单次请求最多 50 个分账接收方。
	// 业务规则：可以设置出资商户作为分账接收方。
	// 示例：TODO（请参考 SplitReceiver 示例）
	Receivers []SplitReceiver `json:"receivers,omitempty"`
	// 是否解冻剩余未分账资金
	// 字段含义：是否将该笔订单剩余未分账金额解冻给商户。
	// 格式规则：布尔值。
	// 业务规则：true 表示剩余未分账金额会结算给商户；false 表示剩余未分账金额不解冻，可再次分账。
	// 示例：true
	UnfreezeUnsplit bool `json:"unfreeze_unsplit,omitempty"`
	// 分账回调地址
	// 字段含义：分账结果通知地址。
	// 格式规则：必须为 HTTPS 且不能携带查询串。
	// 业务规则：交易成功后，通过该地址通知分账结果。
	// 示例：https://www.mock.douyinpay.com
	NotifyURL string `json:"notify_url,omitempty"`
}

func (r SplitFundRequest) Validate() error {
	if r.Appid == "" {
		return fmt.Errorf("field `Appid` is required and must be specified in SplitFundRequest")
	}
	if r.Mchid == "" {
		return fmt.Errorf("field `Mchid` is required and must be specified in SplitFundRequest")
	}
	if r.TransactionID == "" {
		return fmt.Errorf("field `TransactionID` is required and must be specified in SplitFundRequest")
	}
	if r.OutOrderNo == "" {
		return fmt.Errorf("field `OutOrderNo` is required and must be specified in SplitFundRequest")
	}
	if len(r.Receivers) == 0 {
		return fmt.Errorf("field `Receivers` is required and must be specified in SplitFundRequest")
	}
	for idx, receiver := range r.Receivers {
		if err := receiver.Validate(r.Appid); err != nil {
			return fmt.Errorf("field `Receivers[%d]` is invalid: %w", idx, err)
		}
	}
	if r.NotifyURL == "" {
		return fmt.Errorf("field `NotifyURL` is required and must be specified in SplitFundRequest")
	}
	return nil
}

// SplitReceiver 分账接收方
type SplitReceiver struct {
	// 分账接收方类型
	// 字段含义：分账接收方类型。
	// 格式规则：枚举字符串，取值 MERCHANT_ID、PERSONAL_OPENID。
	// 业务规则：MERCHANT_ID 表示商户号；PERSONAL_OPENID 表示用户在商户 appid 下的唯一标识。
	// 示例：MERCHANT_ID
	Type string `json:"type,omitempty"`
	// 分账接收方账号
	// 字段含义：分账接收方账号。
	// 格式规则：长度为 1-64。
	// 业务规则：类型为 MERCHANT_ID 时传商户号；类型为 PERSONAL_OPENID 时传个人 openid。
	// 示例：6020230307605084
	Account string `json:"account,omitempty"`
	// 分账接收方全称
	// 字段含义：分账接收方商户全称或个人姓名。
	// 格式规则：长度为 1-1024，敏感字段需使用抖音支付平台证书公钥加密。
	// 业务规则：MERCHANT_ID 类型必传；PERSONAL_OPENID 类型选传，传入时会检查实名匹配。
	// 示例：
	// TODO: 官方文档未提供标准示例值，暂留空。
	Name string `json:"name,omitempty"`
	// 分账金额
	// 字段含义：分账金额。
	// 格式规则：单位为分，只能为整数。
	// 业务规则：不能超过原订单支付金额和最大分账比例金额。
	// 示例：100
	Amount int64 `json:"amount,omitempty"`
	// 分账描述
	// 字段含义：分账原因描述。
	// 格式规则：长度为 1-80。
	// 业务规则：会在分账账单中体现。
	// 示例：分给合作方
	Description string `json:"description,omitempty"`
}

func (r SplitReceiver) Validate(appid string) error {
	if r.Type == "" {
		return fmt.Errorf("field `Type` is required")
	}
	if r.Account == "" {
		return fmt.Errorf("field `Account` is required")
	}
	if r.Type == "MERCHANT_ID" && r.Name == "" {
		return fmt.Errorf("field `Name` is required when `Type` is MERCHANT_ID")
	}
	if r.Type == "PERSONAL_OPENID" && appid == "" {
		return fmt.Errorf("field `Appid` is required when `Type` is PERSONAL_OPENID")
	}
	if r.Amount <= 0 {
		return fmt.Errorf("field `Amount` is required and must be greater than zero")
	}
	if r.Description == "" {
		return fmt.Errorf("field `Description` is required")
	}
	return nil
}

// SplitFundResponse 请求分账同步返回
type SplitFundResponse struct {
	// 字段含义：直连商户号。
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：抖音支付订单号。
	TransactionID string `json:"transaction_id,omitempty"`
	// 字段含义：商户分账单号。
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 字段含义：抖音支付分账单号。
	OrderID string `json:"order_id,omitempty"`
	// 字段含义：分账单状态。
	State string `json:"state,omitempty"`
	// 字段含义：分账接收方列表。
	Receivers []SplitReceiverResult `json:"receivers,omitempty"`
	// 字段含义：完结分账金额。
	FinishAmount int64 `json:"finish_amount,omitempty"`
	// 字段含义：完结分账描述。
	FinishDescription string `json:"finish_description,omitempty"`
	// 字段含义：完结分账时间。
	SplitFinishTime string `json:"split_finish_time,omitempty"`
}

// QuerySplitFundRequest 查询分账结果请求参数
type QuerySplitFundRequest struct {
	// 商户分账单号
	// 字段含义：商户系统内部的分账单号。
	// 格式规则：长度为 6-32，只能是数字、大小写字母、`_`、`-`、`*`。
	// 业务规则：当前 SDK 通过 path 传入该值。
	// 示例：OUT_3135780230025060619983034
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 直连商户号
	// 字段含义：直连商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：查询分账结果时作为 query 参数上送。
	// 示例：6020221212167701
	Mchid string `json:"mchid,omitempty"`
	// 抖音支付订单号
	// 字段含义：原支付交易对应的交易订单号。
	// 格式规则：长度为 1-32。
	// 业务规则：按官方文档作为 query 参数上送。
	// 示例：2100012501030500000618413371
	TransactionID string `json:"transaction_id,omitempty"`
	// 抖音支付分账单号
	// 字段含义：抖音支付生成的分账单号。
	// 格式规则：长度为 1-32。
	// 业务规则：商户分账单号和抖音支付分账单号二选一时，可传入该字段辅助查询。
	// 示例：11777200250103110500000223502022
	OrderID string `json:"order_id,omitempty"`
}

func (r QuerySplitFundRequest) Validate() error {
	if r.OutOrderNo == "" {
		return fmt.Errorf("field `OutOrderNo` is required and must be specified in QuerySplitFundRequest")
	}
	if r.Mchid == "" {
		return fmt.Errorf("field `Mchid` is required and must be specified in QuerySplitFundRequest")
	}
	return nil
}

// QuerySplitFundResponse 查询分账结果响应
type QuerySplitFundResponse SplitFundResponse

// SplitReceiverResult 分账接收方执行结果
type SplitReceiverResult struct {
	// 字段含义：分账金额。
	Amount int64 `json:"amount,omitempty"`
	// 字段含义：分账描述。
	Description string `json:"description,omitempty"`
	// 字段含义：分账接收方类型。
	Type string `json:"type,omitempty"`
	// 字段含义：分账接收方账号。
	Account string `json:"account,omitempty"`
	// 字段含义：分账结果。
	Result string `json:"result,omitempty"`
	// 字段含义：分账失败原因。
	FailReason string `json:"fail_reason,omitempty"`
	// 字段含义：分账创建时间。
	CreateTime string `json:"create_time,omitempty"`
	// 字段含义：分账完成时间。
	FinishTime string `json:"finish_time,omitempty"`
	// 字段含义：分账明细单号。
	DetailID string `json:"detail_id,omitempty"`
}

// ReturnSplitFundRequest 请求分账回退请求参数
type ReturnSplitFundRequest struct {
	// 直连商户号
	// 字段含义：直连商户的商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：请求分账回退时必填。
	// 示例：6020230307605084
	Mchid string `json:"mchid,omitempty"`
	// 抖音支付分账单号
	// 字段含义：抖音支付分账单号。
	// 格式规则：长度为 1-32。
	// 业务规则：与 OutOrderNo 二选一填写。
	// 示例：11777200250103110500000223512022
	OrderID string `json:"order_id,omitempty"`
	// 商户分账单号
	// 字段含义：商户系统内部的分账单号。
	// 格式规则：长度为 6-32，只能是数字、大小写字母、`_`、`-`、`*`。
	// 业务规则：与 OrderID 二选一填写。
	// 示例：OUT_31357802300250606199830
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 商户回退单号
	// 字段含义：商户在自己后台生成的新的回退单号。
	// 格式规则：长度为 1-32。
	// 业务规则：在商户后台唯一。
	// 示例：OUT_338004
	OutReturnNo string `json:"out_return_no,omitempty"`
	// 回退商户号
	// 字段含义：分账回退的出资商户。
	// 格式规则：长度为 1-32。
	// 业务规则：只能对原分账请求中成功分给商户接收方进行回退。
	// 示例：6020231219024876
	ReturnMchid string `json:"return_mchid,omitempty"`
	// 回退金额
	// 字段含义：需要从分账接收方回退的金额。
	// 格式规则：单位为分，只能为整数。
	// 业务规则：不能超过原始分账单分给该接收方的金额。
	// 示例：10
	Amount int64 `json:"amount,omitempty"`
	// 回退描述
	// 字段含义：分账回退的原因描述。
	// 格式规则：长度为 1-80。
	// 业务规则：会在分账账单中体现。
	// 示例：退分账
	Description string `json:"description,omitempty"`
}

func (r ReturnSplitFundRequest) Validate() error {
	if r.Mchid == "" {
		return fmt.Errorf("field `Mchid` is required and must be specified in ReturnSplitFundRequest")
	}
	if r.OrderID == "" && r.OutOrderNo == "" {
		return fmt.Errorf("one of `OrderID` or `OutOrderNo` is required in ReturnSplitFundRequest")
	}
	if r.OutReturnNo == "" {
		return fmt.Errorf("field `OutReturnNo` is required and must be specified in ReturnSplitFundRequest")
	}
	if r.ReturnMchid == "" {
		return fmt.Errorf("field `ReturnMchid` is required and must be specified in ReturnSplitFundRequest")
	}
	if r.Amount <= 0 {
		return fmt.Errorf("field `Amount` is required and must be greater than zero in ReturnSplitFundRequest")
	}
	if r.Description == "" {
		return fmt.Errorf("field `Description` is required and must be specified in ReturnSplitFundRequest")
	}
	return nil
}

// ReturnSplitFundResponse 分账回退响应
type ReturnSplitFundResponse struct {
	// 字段含义：抖音支付分账单号。
	OrderID string `json:"order_id,omitempty"`
	// 字段含义：商户分账单号。
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 字段含义：商户回退单号。
	OutReturnNo string `json:"out_return_no,omitempty"`
	// 字段含义：抖音支付回退单号。
	ReturnID string `json:"return_id,omitempty"`
	// 字段含义：回退商户号。
	ReturnMchid string `json:"return_mchid,omitempty"`
	// 字段含义：回退金额。
	Amount int64 `json:"amount,omitempty"`
	// 字段含义：回退描述。
	Description string `json:"description,omitempty"`
	// 字段含义：回退结果。
	Result string `json:"result,omitempty"`
	// 字段含义：失败原因。
	FailReason string `json:"fail_reason,omitempty"`
	// 字段含义：创建时间。
	CreateTime string `json:"create_time,omitempty"`
	// 字段含义：完成时间。
	FinishTime string `json:"finish_time,omitempty"`
}

// QueryReturnSplitFundRequest 查询分账回退结果请求参数
type QueryReturnSplitFundRequest struct {
	// 商户回退单号
	// 字段含义：商户在自己后台生成的新的回退单号。
	// 格式规则：长度为 1-32。
	// 业务规则：当前 SDK 通过 path 传入该值。
	// 示例：OUT_338004
	OutReturnNo string `json:"out_return_no,omitempty"`
	// 直连商户号
	// 字段含义：直连商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：查询分账回退结果时作为 query 参数上送。
	// 示例：6020230307605084
	Mchid string `json:"mchid,omitempty"`
	// 商户分账单号
	// 字段含义：商户系统内部的分账单号。
	// 格式规则：长度为 6-32，只能是数字、大小写字母、`_`、`-`、`*`。
	// 业务规则：按官方文档作为 query 参数上送。
	// 示例：OUT_3135780230025060619983
	OutOrderNo string `json:"out_order_no,omitempty"`
}

func (r QueryReturnSplitFundRequest) Validate() error {
	if r.OutReturnNo == "" {
		return fmt.Errorf("field `OutReturnNo` is required and must be specified in QueryReturnSplitFundRequest")
	}
	if r.Mchid == "" {
		return fmt.Errorf("field `Mchid` is required and must be specified in QueryReturnSplitFundRequest")
	}
	return nil
}

// FinishSplitFundRequest 完结分账请求参数
type FinishSplitFundRequest struct {
	// 抖音支付订单号
	// 字段含义：抖音支付订单号。
	// 格式规则：长度为 1-32。
	// 业务规则：完结分账场景必填。
	// 示例：TP2022101317144741443210681000
	TransactionID string `json:"transaction_id,omitempty"`
	// 商户分账单号
	// 字段含义：商户系统内部的分账单号。
	// 格式规则：长度为 6-32，只能是数字、大小写字母、`_`、`-`、`*`。
	// 业务规则：完结分账场景必填。
	// 示例：OUT_1666688488
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 完结分账描述
	// 字段含义：完结分账的原因描述。
	// 格式规则：长度为 1-64。
	// 业务规则：会在分账账单中体现。
	// 示例：测试商品分账
	Description string `json:"description,omitempty"`
	// 直连商户号
	// 字段含义：直连商户号。
	// 格式规则：长度为 1-32。
	// 业务规则：完结分账场景必填。
	// 示例：6000000000000001
	Mchid string `json:"mchid,omitempty"`
	// 通知地址
	// 字段含义：完结分账结果通知地址。
	// 格式规则：必须为 HTTPS 且不能携带查询串。
	// 业务规则：交易成功后，通过该地址通知完结分账结果。
	// 示例：https://www.notify.com
	NotifyURL string `json:"notify_url,omitempty"`
}

func (r FinishSplitFundRequest) Validate() error {
	if r.Mchid == "" {
		return fmt.Errorf("field `Mchid` is required and must be specified in FinishSplitFundRequest")
	}
	if r.TransactionID == "" {
		return fmt.Errorf("field `TransactionID` is required and must be specified in FinishSplitFundRequest")
	}
	if r.OutOrderNo == "" {
		return fmt.Errorf("field `OutOrderNo` is required and must be specified in FinishSplitFundRequest")
	}
	if r.Description == "" {
		return fmt.Errorf("field `Description` is required and must be specified in FinishSplitFundRequest")
	}
	if r.NotifyURL == "" {
		return fmt.Errorf("field `NotifyURL` is required and must be specified in FinishSplitFundRequest")
	}
	return nil
}

// FinishSplitFundResponse 完结分账同步返回
type FinishSplitFundResponse SplitFundResponse

// QueryUnsplitAmountRequest 查询剩余待分金额请求参数
type QueryUnsplitAmountRequest struct {
	// 抖音支付订单号
	// 字段含义：抖音支付订单号。
	// 格式规则：长度为 1-32。
	// 业务规则：当前 SDK 通过 path 传入该值。
	// 示例：TP2022101317144741443210681000
	TransactionID string `json:"transaction_id,omitempty"`
	// 直连商户号
	// 字段含义：直连商户号。
	// 格式规则：长度为 1-32。
	// 业务规则：查询剩余待分金额时作为 query 参数上送。
	// 示例：6000000000000001
	Mchid string `json:"mchid,omitempty"`
}

func (r QueryUnsplitAmountRequest) Validate() error {
	if r.TransactionID == "" {
		return fmt.Errorf("field `TransactionID` is required and must be specified in QueryUnsplitAmountRequest")
	}
	if r.Mchid == "" {
		return fmt.Errorf("field `Mchid` is required and must be specified in QueryUnsplitAmountRequest")
	}
	return nil
}

// QueryUnsplitAmountResponse 查询剩余待分金额响应
type QueryUnsplitAmountResponse struct {
	// 字段含义：直连商户号。
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：抖音支付订单号。
	TransactionID string `json:"transaction_id,omitempty"`
	// 字段含义：订单剩余待分金额。
	UnsplitAmount int64 `json:"unsplit_amount,omitempty"`
}

// AddSplitReceiverRequest 添加分账接收方请求参数
type AddSplitReceiverRequest struct {
	// 直连商户号
	// 字段含义：直连商户号。
	// 格式规则：长度为 1-32。
	// 业务规则：添加分账接收方场景必填。
	// 示例：6020240223833009
	Mchid string `json:"mchid,omitempty"`
	// 商户应用号
	// 字段含义：商户应用号。
	// 格式规则：长度为 1-32。
	// 业务规则：添加个人 OpenID 类型接收方时用于标识 OpenID 所属应用。
	// 示例：byOOJzkcOJWYmSPBuPWLbDjSSqf
	Appid string `json:"appid,omitempty"`
	// 分账接收方账号类型
	// 字段含义：分账接收方账号类型。
	// 格式规则：枚举字符串，取值 MERCHANT_ID、PERSONAL_OPENID。
	// 业务规则：决定 account 和 name 字段的语义。
	// 示例：MERCHANT_ID
	Type string `json:"type,omitempty"`
	// 分账接收方账号
	// 字段含义：分账接收方账号。
	// 格式规则：长度为 1-64。
	// 业务规则：MERCHANT_ID 类型传商户号；PERSONAL_OPENID 类型传个人 OpenID。
	// 示例：6020260126898210
	Account string `json:"account,omitempty"`
	// 分账接收方全称
	// 字段含义：分账接收方商户全称或个人姓名。
	// 格式规则：长度为 1-1024，敏感字段需使用抖音支付平台证书公钥加密。
	// 业务规则：MERCHANT_ID 类型必传；PERSONAL_OPENID 类型选传，传入时会检查实名匹配。
	// 示例：CDEgKhcAkOQVESRENiMsdtfoRDOsLPOfCmJPR
	Name string `json:"name,omitempty"`
	// 与分账方的关系类型
	// 字段含义：分账发起方商户与分账接收方的关系。
	// 格式规则：枚举字符串，包括 SERVICE_PROVIDER、STORE、STAFF、STORE_OWNER、PARTNER、HEADQUARTER、BRAND、DISTRIBUTOR、USER、SUPPLIER、CUSTOM。
	// 业务规则：当取值为 CUSTOM 时需同时填写 CustomRelation。
	// 示例：STORE
	RelationType string `json:"relation_type,omitempty"`
	// 自定义的分账关系
	// 字段含义：商户与接收方具体的关系。
	// 格式规则：最多 10 个字。
	// 业务规则：当 RelationType 为 CUSTOM 时必填；其他关系类型无需填写。
	// 示例：
	// TODO: 官方文档未提供标准示例值，暂留空。
	CustomRelation string `json:"custom_relation,omitempty"`
}

func (r AddSplitReceiverRequest) Validate() error {
	if r.Mchid == "" {
		return fmt.Errorf("field `Mchid` is required and must be specified in AddSplitReceiverRequest")
	}
	if r.Type == "" {
		return fmt.Errorf("field `Type` is required and must be specified in AddSplitReceiverRequest")
	}
	if r.Account == "" {
		return fmt.Errorf("field `Account` is required and must be specified in AddSplitReceiverRequest")
	}
	if r.RelationType == "" {
		return fmt.Errorf("field `RelationType` is required and must be specified in AddSplitReceiverRequest")
	}
	if r.Type == "MERCHANT_ID" && r.Name == "" {
		return fmt.Errorf("field `Name` is required when `Type` is MERCHANT_ID in AddSplitReceiverRequest")
	}
	if r.Type == "PERSONAL_OPENID" && r.Appid == "" {
		return fmt.Errorf("field `Appid` is required when `Type` is PERSONAL_OPENID in AddSplitReceiverRequest")
	}
	if r.RelationType == "CUSTOM" && r.CustomRelation == "" {
		return fmt.Errorf("field `CustomRelation` is required when `RelationType` is CUSTOM in AddSplitReceiverRequest")
	}
	return nil
}

// SplitReceiverResponse 分账接收方响应
type SplitReceiverResponse struct {
	// 字段含义：分账接收方类型。
	Type string `json:"type,omitempty"`
	// 字段含义：分账接收方账号。
	Account string `json:"account,omitempty"`
	// 字段含义：分账接收方全称。
	Name string `json:"name,omitempty"`
	// 字段含义：与分账方的关系类型。
	RelationType string `json:"relation_type,omitempty"`
	// 字段含义：自定义的分账关系。
	CustomRelation string `json:"custom_relation,omitempty"`
}

// DeleteSplitReceiverRequest 删除分账接收方请求参数
type DeleteSplitReceiverRequest struct {
	// 直连商户号
	// 字段含义：直连商户的商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：删除分账接收方场景必填。
	// 示例：6020221212167701
	Mchid string `json:"mchid,omitempty"`
	// 应用ID
	// 字段含义：商户在抖音开放平台申请的应用ID，全局唯一。
	// 格式规则：长度为 1-32。
	// 业务规则：用于校验 appid 与 mchid 关系。
	// 示例：awofz9bncda6w2w4
	Appid string `json:"appid,omitempty"`
	// 分账接收方类型
	// 字段含义：分账接收方类型。
	// 格式规则：枚举字符串，取值 MERCHANT_ID、PERSONAL_OPENID。
	// 业务规则：决定 account 的语义。
	// 示例：MERCHANT_ID
	Type string `json:"type,omitempty"`
	// 分账接收方账号
	// 字段含义：分账接收方账号。
	// 格式规则：长度为 1-64。
	// 业务规则：需与已添加的接收方账号一致。
	// 示例：6020230307605084
	Account string `json:"account,omitempty"`
}

func (r DeleteSplitReceiverRequest) Validate() error {
	if r.Mchid == "" {
		return fmt.Errorf("field `Mchid` is required and must be specified in DeleteSplitReceiverRequest")
	}
	if r.Type == "" {
		return fmt.Errorf("field `Type` is required and must be specified in DeleteSplitReceiverRequest")
	}
	if r.Account == "" {
		return fmt.Errorf("field `Account` is required and must be specified in DeleteSplitReceiverRequest")
	}
	return nil
}

// DeleteSplitReceiverResponse 删除分账接收方响应
type DeleteSplitReceiverResponse struct {
	// 字段含义：直连商户号。
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：分账接收方类型。
	Type string `json:"type,omitempty"`
	// 字段含义：分账接收方账号。
	Account string `json:"account,omitempty"`
}

// ProfitSharingNotify 直连商户分账结果通知明文
type ProfitSharingNotify struct {
	// 字段含义：直连商户号。
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：抖音支付订单号。
	TransactionID string `json:"transaction_id,omitempty"`
	// 字段含义：商户分账单号。
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 字段含义：抖音支付分账单号。
	OrderID string `json:"order_id,omitempty"`
	// 字段含义：分账单状态。
	State string `json:"state,omitempty"`
	// 字段含义：分账接收方列表。
	Receivers []SplitReceiverResult `json:"receivers,omitempty"`
	// 字段含义：完结分账金额。
	FinishAmount int64 `json:"finish_amount,omitempty"`
	// 字段含义：完结分账描述。
	FinishDescription string `json:"finish_description,omitempty"`
	// 字段含义：完结分账时间。
	SplitFinishTime string `json:"split_finish_time,omitempty"`
}

// ReceiverNotify 直连商户分账动态通知明文
type ReceiverNotify struct {
	// 字段含义：直连商户号。
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：抖音支付订单号。
	TransactionID string `json:"transaction_id,omitempty"`
	// 字段含义：商户分账单号。
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 字段含义：抖音支付分账单号。
	OrderID string `json:"order_id,omitempty"`
	// 字段含义：分账接收方入账结果。
	Receiver NotifyReceiver `json:"receiver,omitempty"`
	// 字段含义：分账接收方入账成功时间。
	SuccessTime string `json:"success_time,omitempty"`
}

// NotifyReceiver 直连商户分账动态通知中的接收方信息
type NotifyReceiver struct {
	// 字段含义：分账金额。
	Amount int64 `json:"amount,omitempty"`
	// 字段含义：分账描述。
	Description string `json:"description,omitempty"`
	// 字段含义：分账接收方类型。
	Type string `json:"type,omitempty"`
	// 字段含义：分账接收方账号。
	Account string `json:"account,omitempty"`
}
