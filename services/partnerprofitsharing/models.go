package partnerprofitsharing

import "fmt"

// ApiPartnerSplitFundRequest 服务商请求分账请求参数
type ApiPartnerSplitFundRequest struct {
	// 服务商户号
	// 字段含义：服务商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：服务商分账场景必填。
	// 示例：6020221212167701
	SpMchid string `json:"sp_mchid,omitempty"`
	// 服务商应用ID
	// 字段含义：服务商在抖音开放平台申请的应用ID。
	// 格式规则：长度为 1-32。
	// 业务规则：包含 PERSONAL_SP_OPENID 类型接收方时必填。
	// 示例：awofz9bncda6w2w4
	SpAppid string `json:"sp_appid,omitempty"`
	// 特约商户号
	// 字段含义：特约商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：服务商分账场景必填。
	// 示例：6020221212167702
	SubMchid string `json:"sub_mchid,omitempty"`
	// 特约商户应用ID
	// 字段含义：特约商户在抖音开放平台申请的应用ID。
	// 格式规则：长度为 1-32。
	// 业务规则：包含 PERSONAL_SUB_OPENID 类型接收方时必填。
	// 示例：awofz9bncda6w2w4
	SubAppid string `json:"sub_appid,omitempty"`
	// 抖音支付订单号
	// 字段含义：原支付交易对应的交易订单号。
	// 格式规则：长度为 1-32。
	// 业务规则：请求分账时必填。
	// 示例：TP2022101317144741443210681000
	TransactionID string `json:"transaction_id,omitempty"`
	// 商户分账单号
	// 字段含义：商户系统内部的分账单号。
	// 格式规则：长度为 6-32，只能是数字、大小写字母、`_`、`-`、`*`。
	// 业务规则：在商户系统内部唯一；同一单号重复请求视为同一次分账。
	// 示例：OUT_1666688488
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 分账接收方列表
	// 字段含义：本次分账的接收方明细列表。
	// 格式规则：单次请求最多 50 个接收方。
	// 业务规则：每个接收方的字段语义见 Receiver。
	// 示例：TODO（请参考 Receiver 示例）
	Receivers []ApiPartnerReceiver `json:"receivers,omitempty"`
	// 是否解冻剩余未分账资金
	// 字段含义：是否将剩余未分账金额解冻回商户。
	// 格式规则：布尔值。
	// 业务规则：true 表示剩余未分账金额解冻回商户；false 表示保留待后续分账。
	// 示例：true
	UnfreezeUnsplit bool `json:"unfreeze_unsplit,omitempty"`
	// 分账回调地址
	// 字段含义：用于接收分账结果通知的回调地址。
	// 格式规则：必须为 HTTPS 且不能携带查询串。
	// 业务规则：请求分账时必填。
	// 示例：https://www.mock.douyinpay.com
	NotifyURL string `json:"notify_url,omitempty"`
}

func (r ApiPartnerSplitFundRequest) Validate() error {
	if r.SpMchid == "" {
		return fmt.Errorf("field `SpMchid` is required and must be specified in ApiPartnerSplitFundRequest")
	}
	if r.SubMchid == "" {
		return fmt.Errorf("field `SubMchid` is required and must be specified in ApiPartnerSplitFundRequest")
	}
	if r.TransactionID == "" {
		return fmt.Errorf("field `TransactionID` is required and must be specified in ApiPartnerSplitFundRequest")
	}
	if r.OutOrderNo == "" {
		return fmt.Errorf("field `OutOrderNo` is required and must be specified in ApiPartnerSplitFundRequest")
	}
	if len(r.Receivers) == 0 {
		return fmt.Errorf("field `Receivers` is required and must be specified in ApiPartnerSplitFundRequest")
	}
	if r.NotifyURL == "" {
		return fmt.Errorf("field `NotifyURL` is required and must be specified in ApiPartnerSplitFundRequest")
	}
	return nil
}

// ApiPartnerReceiver 分账接收方
type ApiPartnerReceiver struct {
	// 分账接收方类型
	// 字段含义：分账接收方类型。
	// 格式规则：枚举值为 MERCHANT_ID、PERSONAL_SP_OPENID、PERSONAL_SUB_OPENID。
	// 业务规则：决定 account、name 与 appid 字段的语义。
	// 示例：MERCHANT_ID
	Type string `json:"type,omitempty"`
	// 分账接收方账号
	// 字段含义：分账接收方账号。
	// 格式规则：长度为 1-64。
	// 业务规则：MERCHANT_ID 传商户号；PERSONAL_SP_OPENID 与 PERSONAL_SUB_OPENID 传对应 AppID 下的 OpenID。
	// 示例：6020230307605084
	Account string `json:"account,omitempty"`
	// 分账接收方全称
	// 字段含义：分账接收方商户全称或个人姓名。
	// 格式规则：长度为 1-1024，敏感字段需用平台证书公钥加密。
	// 业务规则：MERCHANT_ID 类型必传；个人类型选传，传入时会校验实名匹配。
	// 示例：TODO（官方文档未提供明文示例）
	Name string `json:"name,omitempty"`
	// 分账金额
	// 字段含义：本次分给该接收方的金额。
	// 格式规则：单位为分，只能为整数。
	// 业务规则：不能超过原订单支付金额及最大分账比例金额。
	// 示例：100
	Amount int64 `json:"amount,omitempty"`
	// 分账描述
	// 字段含义：分账原因描述。
	// 格式规则：长度为 1-80。
	// 业务规则：会在分账账单中体现。
	// 示例：分给合作方
	Description string `json:"description,omitempty"`
}

// ApiPartnerSplitFundResponse 请求分账同步返回
type ApiPartnerSplitFundResponse struct {
	// 字段含义：服务商户号。
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：特约商户号。
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：抖音支付订单号。
	TransactionID string `json:"transaction_id,omitempty"`
	// 字段含义：商户分账单号。
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 字段含义：抖音支付分账单号。
	OrderID string `json:"order_id,omitempty"`
	// 字段含义：分账单状态。
	State string `json:"state,omitempty"`
	// 字段含义：分账接收方列表。
	Receivers []ApiPartnerReceiverResult `json:"receivers,omitempty"`
	// 字段含义：完结分账金额。
	FinishAmount int64 `json:"finish_amount,omitempty"`
	// 字段含义：完结分账描述。
	FinishDescription string `json:"finish_description,omitempty"`
	// 字段含义：完结分账时间。
	SplitFinishTime string `json:"split_finish_time,omitempty"`
}

// ApiPartnerQuerySplitFundRequest 查询分账结果请求参数
type ApiPartnerQuerySplitFundRequest struct {
	// 商户分账单号
	// 字段含义：商户系统内部的分账单号。
	// 格式规则：长度为 6-32，只能是数字、大小写字母、`_`、`-`、`*`。
	// 业务规则：当前 SDK 通过 path 传入该值。
	// 示例：OUT_3135780230025060619983034
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 服务商户号
	// 字段含义：服务商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：服务商分账查询场景必填。
	// 示例：6020221212167701
	SpMchid string `json:"sp_mchid,omitempty"`
	// 特约商户号
	// 字段含义：特约商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：服务商分账查询场景必填。
	// 示例：6020221212167702
	SubMchid string `json:"sub_mchid,omitempty"`
	// 抖音支付订单号
	// 字段含义：原支付交易对应的交易订单号。
	// 格式规则：长度为 1-32。
	// 业务规则：按官方文档作为 query 参数上送。
	// 示例：2100012501030500000618413371
	TransactionID string `json:"transaction_id,omitempty"`
	// 抖音支付分账单号
	// 字段含义：抖音支付生成的分账单号。
	// 格式规则：长度为 1-32。
	// 业务规则：当需要按抖音支付分账单号辅助查询时传入。
	// 示例：11777200250103110500000223502022
	OrderID string `json:"order_id,omitempty"`
}

func (r ApiPartnerQuerySplitFundRequest) Validate() error {
	if r.OutOrderNo == "" {
		return fmt.Errorf("field `OutOrderNo` is required and must be specified in ApiPartnerQuerySplitFundRequest")
	}
	if r.SpMchid == "" {
		return fmt.Errorf("field `SpMchid` is required and must be specified in ApiPartnerQuerySplitFundRequest")
	}
	if r.SubMchid == "" {
		return fmt.Errorf("field `SubMchid` is required and must be specified in ApiPartnerQuerySplitFundRequest")
	}
	return nil
}

// ApiPartnerQuerySplitFundResponse 查询分账结果响应
type ApiPartnerQuerySplitFundResponse ApiPartnerSplitFundResponse

// ApiPartnerReceiverResult 分账接收方执行结果
type ApiPartnerReceiverResult struct {
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

// ApiPartnerReturnSplitFundRequest 请求分账回退请求参数
type ApiPartnerReturnSplitFundRequest struct {
	// 服务商户号
	// 字段含义：服务商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：服务商分账回退场景必填。
	// 示例：6020230307605001
	SpMchid string `json:"sp_mchid,omitempty"`
	// 特约商户号
	// 字段含义：特约商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：服务商分账回退场景必填。
	// 示例：6020230307605002
	SubMchid string `json:"sub_mchid,omitempty"`
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
	// 业务规则：只能对原分账请求中成功分给商户接收方的金额发起回退。
	// 示例：6020231219024876
	ReturnMchid string `json:"return_mchid,omitempty"`
	// 回退金额
	// 字段含义：需要回退的金额。
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

func (r ApiPartnerReturnSplitFundRequest) Validate() error {
	if r.SpMchid == "" {
		return fmt.Errorf("field `SpMchid` is required and must be specified in ApiPartnerReturnSplitFundRequest")
	}
	if r.SubMchid == "" {
		return fmt.Errorf("field `SubMchid` is required and must be specified in ApiPartnerReturnSplitFundRequest")
	}
	if r.OrderID == "" && r.OutOrderNo == "" {
		return fmt.Errorf("one of `OrderID` or `OutOrderNo` is required in ApiPartnerReturnSplitFundRequest")
	}
	if r.OutReturnNo == "" {
		return fmt.Errorf("field `OutReturnNo` is required and must be specified in ApiPartnerReturnSplitFundRequest")
	}
	if r.ReturnMchid == "" {
		return fmt.Errorf("field `ReturnMchid` is required and must be specified in ApiPartnerReturnSplitFundRequest")
	}
	if r.Amount <= 0 {
		return fmt.Errorf("field `Amount` is required and must be greater than zero in ApiPartnerReturnSplitFundRequest")
	}
	if r.Description == "" {
		return fmt.Errorf("field `Description` is required and must be specified in ApiPartnerReturnSplitFundRequest")
	}
	return nil
}

// ApiPartnerReturnSplitFundResponse 分账回退响应
type ApiPartnerReturnSplitFundResponse struct {
	// 字段含义：服务商户号。
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：特约商户号。
	SubMchid string `json:"sub_mchid,omitempty"`
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

// ApiPartnerQueryReturnSplitFundResponse 查询分账回退结果响应
type ApiPartnerQueryReturnSplitFundResponse ApiPartnerReturnSplitFundResponse

// ApiPartnerQueryReturnSplitFundRequest 查询分账回退结果请求参数
type ApiPartnerQueryReturnSplitFundRequest struct {
	// 商户回退单号
	// 字段含义：商户在自己后台生成的新的回退单号。
	// 格式规则：长度为 1-32。
	// 业务规则：当前 SDK 通过 path 传入该值。
	// 示例：OUT_338004
	OutReturnNo string `json:"out_return_no,omitempty"`
	// 服务商户号
	// 字段含义：服务商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：服务商分账回退查询场景必填。
	// 示例：6020230307605001
	SpMchid string `json:"sp_mchid,omitempty"`
	// 特约商户号
	// 字段含义：特约商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：服务商分账回退查询场景必填。
	// 示例：6020230307605002
	SubMchid string `json:"sub_mchid,omitempty"`
	// 商户分账单号
	// 字段含义：商户系统内部的分账单号。
	// 格式规则：长度为 6-32，只能是数字、大小写字母、`_`、`-`、`*`。
	// 业务规则：按官方文档作为 query 参数上送。
	// 示例：OUT_3135780230025060619983
	OutOrderNo string `json:"out_order_no,omitempty"`
}

func (r ApiPartnerQueryReturnSplitFundRequest) Validate() error {
	if r.OutReturnNo == "" {
		return fmt.Errorf("field `OutReturnNo` is required and must be specified in ApiPartnerQueryReturnSplitFundRequest")
	}
	if r.SpMchid == "" {
		return fmt.Errorf("field `SpMchid` is required and must be specified in ApiPartnerQueryReturnSplitFundRequest")
	}
	if r.SubMchid == "" {
		return fmt.Errorf("field `SubMchid` is required and must be specified in ApiPartnerQueryReturnSplitFundRequest")
	}
	return nil
}

// ApiPartnerFinishSplitFundRequest 完结分账请求参数
type ApiPartnerFinishSplitFundRequest struct {
	// 服务商户号
	// 字段含义：服务商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：完结分账场景必填。
	// 示例：6020250310533405
	SpMchid string `json:"sp_mchid,omitempty"`
	// 特约商户号
	// 字段含义：特约商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：完结分账场景必填。
	// 示例：6020250314534907
	SubMchid string `json:"sub_mchid,omitempty"`
	// 交易订单号
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
	// 通知地址
	// 字段含义：用于接收完结分账结果通知的地址。
	// 格式规则：必须为 HTTPS 且不能携带查询串。
	// 业务规则：完结分账场景必填。
	// 示例：https://www.notify.com
	NotifyURL string `json:"notify_url,omitempty"`
}

func (r ApiPartnerFinishSplitFundRequest) Validate() error {
	if r.SpMchid == "" {
		return fmt.Errorf("field `SpMchid` is required and must be specified in ApiPartnerFinishSplitFundRequest")
	}
	if r.SubMchid == "" {
		return fmt.Errorf("field `SubMchid` is required and must be specified in ApiPartnerFinishSplitFundRequest")
	}
	if r.TransactionID == "" {
		return fmt.Errorf("field `TransactionID` is required and must be specified in ApiPartnerFinishSplitFundRequest")
	}
	if r.OutOrderNo == "" {
		return fmt.Errorf("field `OutOrderNo` is required and must be specified in ApiPartnerFinishSplitFundRequest")
	}
	if r.Description == "" {
		return fmt.Errorf("field `Description` is required and must be specified in ApiPartnerFinishSplitFundRequest")
	}
	if r.NotifyURL == "" {
		return fmt.Errorf("field `NotifyURL` is required and must be specified in ApiPartnerFinishSplitFundRequest")
	}
	return nil
}

// ApiPartnerFinishSplitFundResponse 完结分账同步返回
type ApiPartnerFinishSplitFundResponse ApiPartnerSplitFundResponse

// QueryUnsplitAmountRequest 查询剩余待分金额请求参数
type QueryUnsplitAmountRequest struct {
	// 抖音支付订单号
	// 字段含义：抖音支付订单号。
	// 格式规则：长度为 1-32。
	// 业务规则：当前 SDK 通过 path 传入该值。
	// 示例：TP2022101317144741443210681000
	TransactionID string `json:"transaction_id,omitempty"`
	// 服务商户号
	// 字段含义：服务商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：服务商查询剩余待分金额场景必填。
	// 示例：6000000000000001
	SpMchid string `json:"sp_mchid,omitempty"`
}

func (r QueryUnsplitAmountRequest) Validate() error {
	if r.TransactionID == "" {
		return fmt.Errorf("field `TransactionID` is required and must be specified in QueryUnsplitAmountRequest")
	}
	if r.SpMchid == "" {
		return fmt.Errorf("field `SpMchid` is required and must be specified in QueryUnsplitAmountRequest")
	}
	return nil
}

// QueryUnsplitAmountResponse 查询剩余待分金额响应
type QueryUnsplitAmountResponse struct {
	// 字段含义：服务商户号。
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：抖音支付订单号。
	TransactionID string `json:"transaction_id,omitempty"`
	// 字段含义：订单剩余待分金额。
	UnsplitAmount int64 `json:"unsplit_amount,omitempty"`
}

// QueryMerchantConfigRequest 查询特约商户分账配置请求参数
type QueryMerchantConfigRequest struct {
	// 特约商户号
	// 字段含义：特约商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：当前 SDK 通过 path 传入该值。
	// 示例：6020221212167702
	SubMchid string `json:"sub_mchid,omitempty"`
	// 服务商户号
	// 字段含义：服务商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：服务商查询配置场景必填。
	// 示例：6020221212167701
	SpMchid string `json:"sp_mchid,omitempty"`
}

func (r QueryMerchantConfigRequest) Validate() error {
	if r.SubMchid == "" {
		return fmt.Errorf("field `SubMchid` is required and must be specified in QueryMerchantConfigRequest")
	}
	if r.SpMchid == "" {
		return fmt.Errorf("field `SpMchid` is required and must be specified in QueryMerchantConfigRequest")
	}
	return nil
}

// QueryMerchantConfigResponse 查询特约商户分账配置响应
type QueryMerchantConfigResponse struct {
	// 字段含义：特约商户号。
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：最大分账比例，单位为万分比。
	MaxRatio int64 `json:"max_ratio,omitempty"`
}

// AddReceiverRequest 添加分账接收方请求参数
type AddReceiverRequest struct {
	// 服务商户号
	// 字段含义：服务商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：添加分账接收方场景必填。
	// 示例：6020221212167701
	SpMchid string `json:"sp_mchid,omitempty"`
	// 服务商应用ID
	// 字段含义：服务商在抖音开放平台申请的应用ID。
	// 格式规则：长度为 1-32。
	// 业务规则：包含 PERSONAL_SP_OPENID 类型接收方时必填。
	// 示例：awofz9bncda6w2w4
	SpAppid string `json:"sp_appid,omitempty"`
	// 特约商户号
	// 字段含义：特约商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：添加分账接收方场景必填。
	// 示例：6020221212167702
	SubMchid string `json:"sub_mchid,omitempty"`
	// 特约商户应用ID
	// 字段含义：特约商户在抖音开放平台申请的应用ID。
	// 格式规则：长度为 1-32。
	// 业务规则：包含 PERSONAL_SUB_OPENID 类型接收方时必填。
	// 示例：awofz9bncda6w2w4
	SubAppid string `json:"sub_appid,omitempty"`
	// 分账接收方账号类型
	// 字段含义：分账接收方类型。
	// 格式规则：枚举值为 MERCHANT_ID、PERSONAL_SP_OPENID、PERSONAL_SUB_OPENID。
	// 业务规则：决定 account、name 与 appid 字段的语义。
	// 示例：MERCHANT_ID
	Type string `json:"type,omitempty"`
	// 分账接收方账号
	// 字段含义：分账接收方账号。
	// 格式规则：长度为 1-64。
	// 业务规则：MERCHANT_ID 传商户号；个人类型传对应 AppID 下的 OpenID。
	// 示例：6020260126898210
	Account string `json:"account,omitempty"`
	// 分账接收方全称
	// 字段含义：分账接收方商户全称或个人姓名。
	// 格式规则：长度为 1-1024，敏感字段需用平台证书公钥加密。
	// 业务规则：MERCHANT_ID 类型必传；个人类型选传，传入时会校验实名匹配。
	// 示例：TODO（官方文档未提供明文示例）
	Name string `json:"name,omitempty"`
	// 与分账方的关系类型
	// 字段含义：分账发起方商户与分账接收方的关系。
	// 格式规则：枚举值包括 STORE、PARTNER、CUSTOM 等。
	// 业务规则：当为 CUSTOM 时需同时填写 CustomRelation。
	// 示例：STORE
	RelationType string `json:"relation_type,omitempty"`
	// 自定义的分账关系
	// 字段含义：特约商户与接收方的具体关系。
	// 格式规则：长度为 1-10。
	// 业务规则：仅当 RelationType 为 CUSTOM 时必填。
	// 示例：TODO（官方文档未提供示例）
	CustomRelation string `json:"custom_relation,omitempty"`
}

func (r AddReceiverRequest) Validate() error {
	if r.SpMchid == "" {
		return fmt.Errorf("field `SpMchid` is required and must be specified in AddReceiverRequest")
	}
	if r.SubMchid == "" {
		return fmt.Errorf("field `SubMchid` is required and must be specified in AddReceiverRequest")
	}
	if r.Type == "" {
		return fmt.Errorf("field `Type` is required and must be specified in AddReceiverRequest")
	}
	if r.Account == "" {
		return fmt.Errorf("field `Account` is required and must be specified in AddReceiverRequest")
	}
	if r.RelationType == "" {
		return fmt.Errorf("field `RelationType` is required and must be specified in AddReceiverRequest")
	}
	if r.Type == "MERCHANT_ID" && r.Name == "" {
		return fmt.Errorf("field `Name` is required when `Type` is MERCHANT_ID in AddReceiverRequest")
	}
	if r.Type == "PERSONAL_SP_OPENID" && r.SpAppid == "" {
		return fmt.Errorf("field `SpAppid` is required when `Type` is PERSONAL_SP_OPENID in AddReceiverRequest")
	}
	if r.Type == "PERSONAL_SUB_OPENID" && r.SubAppid == "" {
		return fmt.Errorf("field `SubAppid` is required when `Type` is PERSONAL_SUB_OPENID in AddReceiverRequest")
	}
	if r.RelationType == "CUSTOM" && r.CustomRelation == "" {
		return fmt.Errorf("field `CustomRelation` is required when `RelationType` is CUSTOM in AddReceiverRequest")
	}
	return nil
}

// ReceiverResponse 分账接收方响应
type ReceiverResponse struct {
	// 字段含义：特约商户号。
	SubMchid string `json:"sub_mchid,omitempty"`
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

// DeleteReceiverRequest 删除分账接收方请求参数
type DeleteReceiverRequest struct {
	// 服务商户号
	// 字段含义：服务商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：删除分账接收方场景必填。
	// 示例：6020221212167701
	SpMchid string `json:"sp_mchid,omitempty"`
	// 服务商应用ID
	// 字段含义：服务商在抖音开放平台申请的应用ID。
	// 格式规则：长度为 1-32。
	// 业务规则：官方文档提供该字段用于校验 appid 与 mchid 关系。
	// 示例：awofz9bncda6w2w4
	SpAppid string `json:"sp_appid,omitempty"`
	// 特约商户号
	// 字段含义：特约商户号，由抖音支付生成并下发。
	// 格式规则：长度为 1-32。
	// 业务规则：删除分账接收方场景必填。
	// 示例：6020221212167702
	SubMchid string `json:"sub_mchid,omitempty"`
	// 特约商户应用ID
	// 字段含义：特约商户在抖音开放平台申请的应用ID。
	// 格式规则：长度为 1-32。
	// 业务规则：官方文档提供该字段用于校验 appid 与 mchid 关系。
	// 示例：awofz9bncda6w2w4
	SubAppid string `json:"sub_appid,omitempty"`
	// 分账接收方类型
	// 字段含义：分账接收方类型。
	// 格式规则：枚举值为 MERCHANT_ID、PERSONAL_SP_OPENID、PERSONAL_SUB_OPENID。
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

func (r DeleteReceiverRequest) Validate() error {
	if r.SpMchid == "" {
		return fmt.Errorf("field `SpMchid` is required and must be specified in DeleteReceiverRequest")
	}
	if r.SubMchid == "" {
		return fmt.Errorf("field `SubMchid` is required and must be specified in DeleteReceiverRequest")
	}
	if r.Type == "" {
		return fmt.Errorf("field `Type` is required and must be specified in DeleteReceiverRequest")
	}
	if r.Account == "" {
		return fmt.Errorf("field `Account` is required and must be specified in DeleteReceiverRequest")
	}
	return nil
}

// DeleteReceiverResponse 删除分账接收方响应
type DeleteReceiverResponse struct {
	// 字段含义：特约商户号。
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：分账接收方类型。
	Type string `json:"type,omitempty"`
	// 字段含义：分账接收方账号。
	Account string `json:"account,omitempty"`
}

// ProfitSharingNotify 服务商分账结果通知明文
type ProfitSharingNotify struct {
	// 字段含义：服务商户号。
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：特约商户号。
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：抖音支付订单号。
	TransactionID string `json:"transaction_id,omitempty"`
	// 字段含义：商户分账单号。
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 字段含义：抖音支付分账单号。
	OrderID string `json:"order_id,omitempty"`
	// 字段含义：分账单状态。
	State string `json:"state,omitempty"`
	// 字段含义：分账接收方列表。
	Receivers []NotifyReceiver `json:"receivers,omitempty"`
	// 字段含义：完结分账金额。
	FinishAmount int64 `json:"finish_amount,omitempty"`
	// 字段含义：完结分账描述。
	FinishDescription string `json:"finish_description,omitempty"`
	// 字段含义：完结分账时间。
	SplitFinishTime string `json:"split_finish_time,omitempty"`
}

// NotifyReceiver 服务商分账结果通知中的接收方结果
type NotifyReceiver struct {
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

// ReceiverNotify 服务商分账接收方入账通知明文
type ReceiverNotify struct {
	// 字段含义：服务商户号。
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：特约商户号。
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：抖音支付订单号。
	TransactionID string `json:"transaction_id,omitempty"`
	// 字段含义：商户分账单号。
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 字段含义：抖音支付分账单号。
	OrderID string `json:"order_id,omitempty"`
	// 字段含义：分账接收方类型。
	Type string `json:"type,omitempty"`
	// 字段含义：分账接收方账号。
	Account string `json:"account,omitempty"`
	// 字段含义：通知类型。
	EventType string `json:"event_type,omitempty"`
}
