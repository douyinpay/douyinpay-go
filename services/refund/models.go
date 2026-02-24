package refund

// Account
/*
`AVAILABLE` - 可用余额, 多账户资金准备退款可用余额出资账户类型
`UNAVAILABLE` - 不可用余额, 多账户资金准备退款不可用余额出资账户类型
*/
type Account string

func (e Account) Ptr() *Account {
	return &e
}

// Enums of Account
const (
	ACCOUNT_AVAILABLE   Account = "AVAILABLE"
	ACCOUNT_UNAVAILABLE Account = "UNAVAILABLE"
)

// Amount
type Amount struct {
	// 订单总金额，单位为分
	Total int64 `json:"total"`
	// 退款标价金额，单位为分，可以做部分退款
	Refund int64 `json:"refund"`
	// 退款出资的账户类型及金额信息
	From []FundsFromItem `json:"from,omitempty"`
	// 现金支付金额，单位为分，只能为整数
	PayerTotal int64 `json:"payer_total"`
	// 退款给用户的金额，不包含所有优惠券金额
	PayerRefund int64 `json:"payer_refund"`
	// 去掉非充值代金券退款金额后的退款金额，单位为分，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	SettlementRefund int64 `json:"settlement_refund"`
	// 应结订单金额=订单金额-免充值代金券金额，应结订单金额<=订单金额，单位为分
	SettlementTotal int64 `json:"settlement_total"`
	// 优惠退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金，说明详见代金券或立减优惠，单位为分
	DiscountRefund int64 `json:"discount_refund"`
	// 符合ISO 4217标准的三位字母代码，目前只支持人民币：CNY。
	Currency string `json:"currency"`
	//手续费退款金额
	RefundFee int64 `json:"refund_fee"`
}

// AmountReq
type AmountReq struct {
	// 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额。
	Refund int64 `json:"refund"`
	// 退款需要从指定账户出资时，传递此参数指定出资金额（币种的最小单位，只能为整数）。 同时指定多个账户出资退款的使用场景需要满足以下条件：1、未开通退款支出分离产品功能；2、订单属于分账订单，且分账处于待分账或分账中状态。 参数传递需要满足条件：1、基本账户可用余额出资金额与基本账户不可用余额出资金额之和等于退款金额；2、账户类型不能重复。 上述任一条件不满足将返回错误
	From []FundsFromItem `json:"from,omitempty"`
	// 原支付交易的订单总金额，币种的最小单位，只能为整数。
	Total int64 `json:"total"`
	// 符合ISO 4217标准的三位字母代码，目前只支持人民币：CNY。
	Currency string `json:"currency"`
}

// Channel
type Channel string

func (e Channel) Ptr() *Channel {
	return &e
}

// Enums of Channel
const (
	CHANNEL_ORIGINAL       Channel = "ORIGINAL"       //原路退款
	CHANNEL_BALANCE        Channel = "BALANCE"        //退回到余额
	CHANNEL_OTHER_BALANCE  Channel = "OTHER_BALANCE"  //原账户异常退到其他余额账户
	CHANNEL_OTHER_BANKCARD Channel = "OTHER_BANKCARD" //原银行卡异常退到其他银行卡
)

// CreateRequest
type CreateRequest struct {
	// 商户号
	Mchid string `json:"mchid,omitempty"`
	// 应用id
	Appid string `json:"appid,omitempty"`
	// 服务商户号
	SpMchid string `json:"sp_mchid,omitempty"`
	// 子商户的商户号，由抖音支付生成并下发。服务商模式下必须传递此参数
	SubMchid string `json:"sub_mchid,omitempty"`
	// 原支付交易对应的抖音订单号
	TransactionId string `json:"transaction_id,omitempty"`
	// 原支付交易对应的商户订单号
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo string `json:"out_refund_no"`
	// 若商户传入，会在下发给用户的退款消息中体现退款原因
	Reason string `json:"reason,omitempty"`
	// 异步接收抖音支付退款结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效，优先回调当前传的这个地址。
	NotifyUrl string `json:"notify_url,omitempty"`
	// 若传递此参数则使用对应的资金账户退款，否则默认使用未结算资金退款（仅对老资金流商户适用）  枚举值： - AVAILABLE：可用余额账户    * `AVAILABLE` - 可用余额
	FundsAccount *ReqFundsAccount `json:"funds_account,omitempty"`
	// 订单金额信息
	Amount *AmountReq `json:"amount"`
	// 指定商品退款需要传此参数，其他场景无需传递
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// FundsAccount
type FundsAccount string

func (e FundsAccount) Ptr() *FundsAccount {
	return &e
}

// Enums of FundsAccount
const (
	FUNDSACCOUNT_UNSETTLED   FundsAccount = "UNSETTLED"   //未结算资金
	FUNDSACCOUNT_AVAILABLE   FundsAccount = "AVAILABLE"   //可用余额
	FUNDSACCOUNT_UNAVAILABLE FundsAccount = "UNAVAILABLE" //不可用余额
	FUNDSACCOUNT_OPERATION   FundsAccount = "OPERATION"   //运营户
	FUNDSACCOUNT_BASIC       FundsAccount = "BASIC"       //基本账户（含可用余额和不可用余额）
)

// FundsFromItem
type FundsFromItem struct {
	// 下面枚举值多选一。 枚举值： AVAILABLE : 可用余额 UNAVAILABLE : 不可用余额 * `AVAILABLE` - 可用余额 * `UNAVAILABLE` - 不可用余额
	Account *Account `json:"account"`
	// 对应账户出资金额
	Amount int64 `json:"amount"`
}

// GoodsDetail
type GoodsDetail struct {
	// 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成
	MerchantGoodsId string `json:"merchant_goods_id"`
	// 抖音支付定义的统一商品编号（没有可不传）
	DouyinpayGoodsId string `json:"douyinpay_goods_id,omitempty"`
	// 商品的实际名称
	GoodsName string `json:"goods_name,omitempty"`
	// 商品单价金额，单位为分
	UnitPrice int64 `json:"unit_price"`
	// 商品退款金额，单位为分
	RefundAmount int64 `json:"refund_amount"`
	// 对应商品的退货数量
	RefundQuantity int64 `json:"refund_quantity"`
}

// Promotion
type Promotion struct {
	// 券或者立减优惠id
	PromotionId string `json:"promotion_id"`
	// 枚举值： - GLOBAL- 全场代金券 - SINGLE- 单品优惠 * `GLOBAL` - 全场代金券 * `SINGLE` - 单品优惠
	Scope *Scope `json:"scope"`
	// 枚举值： - COUPON- 代金券，需要走结算资金的充值型代金券 - DISCOUNT- 优惠券，不走结算资金的免充值型优惠券 * `COUPON` - 代金券 * `DISCOUNT` - 优惠券
	Type *Type `json:"type"`
	// 用户享受优惠的金额（优惠券面额=抖音出资金额+商家出资金额+其他出资方金额 ），单位为分
	Amount int64 `json:"amount"`
	// 优惠退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为用户支付的现金，说明详见代金券或立减优惠，单位为分
	RefundAmount int64 `json:"refund_amount"`
	// 优惠商品发生退款时返回商品信息
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// QueryByOutRefundNoRequest
type QueryByOutRefundNoRequest struct {
	// 商户号
	Mchid string `json:"mchid,omitempty"`
	// 应用id
	Appid string `json:"appid,omitempty"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo string `json:"out_refund_no"`
	//服务商商户号
	SpMchid string `json:"sp_mchid,omitempty"`
	// 子商户的商户号，由抖音支付生成并下发。服务商模式下必须传递此参数
	SubMchid string `json:"sub_mchid,omitempty"`
}

// Refund
type Refund struct {
	// 抖音支付退款单号
	RefundId string `json:"refund_id"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo string `json:"out_refund_no"`
	// 抖音支付交易订单号
	TransactionId string `json:"transaction_id"`
	// 原支付交易对应的商户订单号
	OutTradeNo string `json:"out_trade_no"`
	// 枚举值： - ORIGINAL—原路退款 - BALANCE—退回到余额 - OTHER_BALANCE—原账户异常退到其他余额账户 - OTHER_BANKCARD—原银行卡异常退到其他银行卡 * `ORIGINAL` - 原路退款 * `BALANCE` - 退回到余额 * `OTHER_BALANCE` - 原账户异常退到其他余额账户 * `OTHER_BANKCARD` - 原银行卡异常退到其他银行卡
	Channel *Channel `json:"channel"`
	// 取当前退款单的退款入账方，有以下几种情况： 1）退回银行卡：{银行名称}{卡类型}{卡尾号} 2）退回支付用户零钱:支付用户零钱 3）退还商户:商户基本账户商户结算银行账户 4）退回支付用户零钱通:支付用户零钱通
	UserReceivedAccount string `json:"user_received_account"`
	// 退款成功时间，退款状态status为SUCCESS（退款成功）时，返回该字段。遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。
	SuccessTime string `json:"success_time,omitempty"`
	// 退款受理时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。
	CreateTime string `json:"create_time"`
	// 退款到银行发现用户的卡作废或者冻结了，导致原路退款银行卡失败，可前往商户平台（douyinpay.com），手动处理此笔退款。 枚举值： - SUCCESS—退款成功 - CLOSED—退款关闭 - PROCESSING—退款处理中 - ABNORMAL—退款异常 * `SUCCESS` - 退款成功 * `CLOSED` - 退款关闭 * `PROCESSING` - 退款处理中 * `ABNORMAL` - 退款异常
	Status *Status `json:"status"`
	// 退款所使用资金对应的资金账户类型 枚举值： - UNSETTLED : 未结算资金 - AVAILABLE : 可用余额 - UNAVAILABLE : 不可用余额 - OPERATION : 运营户 - BASIC : 基本账户（含可用余额和不可用余额） * `UNSETTLED` - 未结算资金 * `AVAILABLE` - 可用余额 * `UNAVAILABLE` - 不可用余额 * `OPERATION` - 运营户 * `BASIC` - 基本账户（含可用余额和不可用余额）
	FundsAccount *FundsAccount `json:"funds_account,omitempty"`
	// 金额详细信息
	Amount *Amount `json:"amount"`
	// 优惠退款信息
	PromotionDetail []Promotion `json:"promotion_detail,omitempty"`
}

// ReqFundsAccount
type ReqFundsAccount string

func (e ReqFundsAccount) Ptr() *ReqFundsAccount {
	return &e
}

// Enums of ReqFundsAccount
const (
	REQFUNDSACCOUNT_AVAILABLE ReqFundsAccount = "AVAILABLE" //可用余额, 仅对老资金流商户适用，指定从可用余额账户出资
)

// Scope
type Scope string

func (e Scope) Ptr() *Scope {
	return &e
}

// Enums of Scope
const (
	SCOPE_GLOBAL Scope = "GLOBAL" //全场代金券, 全场优惠类型
	SCOPE_SINGLE Scope = "SINGLE" //单品优惠, 单品优惠类型
)

// Status
type Status string

func (e Status) Ptr() *Status {
	return &e
}

// Enums of Status
const (
	STATUS_SUCCESS    Status = "SUCCESS"    //退款成功
	STATUS_CLOSED     Status = "CLOSED"     //退款关闭
	STATUS_PROCESSING Status = "PROCESSING" //退款处理中
	STATUS_ABNORMAL   Status = "ABNORMAL"   //退款异常
)

// Type
type Type string

func (e Type) Ptr() *Type {
	return &e
}

// Enums of Type
const (
	TYPE_COUPON   Type = "COUPON"   //代金券, 代金券类型，需要走结算资金的充值型代金券
	TYPE_DISCOUNT Type = "DISCOUNT" //优惠券, 优惠券类型，不走结算资金的免充值型优惠券
)
