package refund

// Account 出资账户类型
//
// 枚举值：
//   - AVAILABLE：可用余额，多账户资金准备退款可用余额出资账户类型
//   - UNAVAILABLE：不可用余额，多账户资金准备退款不可用余额出资账户类型
type Account string

func (e Account) Ptr() *Account {
	return &e
}

// Enums of Account
const (
	ACCOUNT_AVAILABLE   Account = "AVAILABLE"   // 可用余额
	ACCOUNT_UNAVAILABLE Account = "UNAVAILABLE" // 不可用余额
)

// Amount 退款响应中的金额详细信息
type Amount struct {
	// 字段含义：订单金额。
	// 格式规则：int。
	// 业务规则：订单总金额，单位为分。
	Total int64 `json:"total"`
	// 字段含义：退款金额。
	// 格式规则：int。
	// 业务规则：退款标价金额，单位为分，可以做部分退款。
	Refund int64 `json:"refund"`
	// 字段含义：退款出资账户及金额。
	// 格式规则：array。
	// 业务规则：退款出资的账户类型及金额信息。
	From []FundsFromItem `json:"from,omitempty"`
	// 字段含义：用户支付金额。
	// 格式规则：int。
	// 业务规则：用户现金支付金额，整型，单位为分。例如 10 元订单用户使用了 2 元全场代金券，则该金额为用户实际支付的 8 元。
	PayerTotal int64 `json:"payer_total"`
	// 字段含义：用户退款金额。
	// 格式规则：int。
	// 业务规则：指用户实际收到的现金退款金额，不包含所有优惠券金额，数据类型为整型，单位为分。例如在一个 10 元的订单中，用户使用了 2 元的全场代金券，若商户申请退款 5 元，则用户将收到 4 元的现金退款（即该字段所示金额）和 1 元的代金券退款。
	PayerRefund int64 `json:"payer_refund"`
	// 字段含义：应结退款金额。
	// 格式规则：int。
	// 业务规则：预留字段，商户不需要关注。去掉免充值代金券退款金额后的退款金额，整型，单位为分，应结退款金额 = 申请退款金额 - 非充值代金券退款金额，应结退款金额 <= 申请退款金额。例如 10 元订单用户使用了 2 元全场代金券（一张免充值 1 元 + 一张预充值 1 元），商户申请退款 5 元，则该金额为 退款金额 5 元 - 0.5 元免充值代金券退款金额 = 4.5 元。
	SettlementRefund int64 `json:"settlement_refund"`
	// 字段含义：应结订单金额。
	// 格式规则：int。
	// 业务规则：预留字段，商户不需要关注。去除免充值代金券金额后的订单金额，整型，单位为分。应结订单金额 = 订单金额 - 免充值代金券金额，应结订单金额 <= 订单金额。例如 10 元订单用户使用了 2 元全场代金券（一张免充值 1 元 + 一张预充值 1 元），则该金额为 订单金额 10 元 - 免充值代金券金额 1 元 = 9 元。
	SettlementTotal int64 `json:"settlement_total"`
	// 字段含义：优惠退款金额。
	// 格式规则：int。
	// 业务规则：申请退款后用户收到的优惠退款金额，整型，单位为分。例如 10 元订单用户使用了 2 元全场代金券，商户申请退款 5 元，用户收到的是 4 元现金 + 1 元代金券退款金额（该字段）。
	DiscountRefund int64 `json:"discount_refund"`
	// 字段含义：退款币种。
	// 格式规则：string[1,16]。
	// 业务规则：符合 ISO 4217 标准的三位字母代码，目前只支持人民币：CNY。
	Currency string `json:"currency"`
	// 字段含义：手续费退款金额。
	// 格式规则：int。
	// 业务规则：订单退款时退还的手续费金额，整型，单位为分。例如一笔 100 元的订单收了 0.6 元手续费，商户申请退款 50 元，该金额为等比退还的 0.3 元手续费。
	RefundFee int64 `json:"refund_fee"`
}

// AmountReq 申请退款请求中的金额信息
type AmountReq struct {
	// 字段含义：退款金额。
	// 格式规则：int。
	// 业务规则：退款金额，币种的最小单位，单位为分，只能为整数，不能超过原订单支付金额。
	Refund int64 `json:"refund"`
	// 字段含义：退款出资账户及金额。
	// 格式规则：array。
	// 业务规则：预留字段，商户不需要关注。退款需要从指定账户出资时，传递此参数指定出资金额（币种的最小单位，只能为整数）。同时指定多个账户出资退款的使用场景需要满足以下条件：1、未开通退款支出分离产品功能；2、订单属于分账订单，且分账处于待分账或分账中状态。参数传递需要满足条件：1、基本账户可用余额出资金额与基本账户不可用余额出资金额之和等于退款金额；2、账户类型不能重复。上述任一条件不满足将返回错误。
	From []FundsFromItem `json:"from,omitempty"`
	// 字段含义：原订单金额。
	// 格式规则：int。
	// 业务规则：原支付交易的订单总金额，币种的最小单位，单位为分，只能为整数。
	Total int64 `json:"total"`
	// 字段含义：退款币种。
	// 格式规则：string[1,16]。
	// 业务规则：符合 ISO 4217 标准的三位字母代码，目前只支持人民币：CNY。
	Currency string `json:"currency"`
}

// Channel 退款渠道
//
// 枚举值：
//   - ORIGINAL：原路退款
//   - BALANCE：退回到余额
//   - OTHER_BALANCE：原账户异常退到其他余额账户
//   - OTHER_BANKCARD：原银行卡异常退到其他银行卡
type Channel string

func (e Channel) Ptr() *Channel {
	return &e
}

// Enums of Channel
const (
	CHANNEL_ORIGINAL       Channel = "ORIGINAL"       // 原路退款
	CHANNEL_BALANCE        Channel = "BALANCE"        // 退回到余额
	CHANNEL_OTHER_BALANCE  Channel = "OTHER_BALANCE"  // 原账户异常退到其他余额账户
	CHANNEL_OTHER_BANKCARD Channel = "OTHER_BANKCARD" // 原银行卡异常退到其他银行卡
)

// CreateRequest 申请退款请求
type CreateRequest struct {
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：应用 ID。
	// 格式规则：string[1,32]。
	// 业务规则：预留字段，商户不需要关注。由抖音支付生成的应用 ID，全局唯一。
	Appid string `json:"appid,omitempty"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成并下发。服务商模式下必须传递此参数。
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：子商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。服务商模式下必须传递此参数。
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：抖音支付订单号。
	// 格式规则：string[1,32]。
	// 业务规则：原支付交易对应的支付订单号。transaction_id 和 out_trade_no 必须二选一进行传参。
	TransactionId string `json:"transaction_id,omitempty"`
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]。
	// 业务规则：原支付交易对应的商户订单号。transaction_id 和 out_trade_no 必须二选一进行传参。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 字段含义：商户退款单号。
	// 格式规则：string[1,64]，只能是数字、大小写字母_-|*@。
	// 业务规则：商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔。
	OutRefundNo string `json:"out_refund_no"`
	// 字段含义：退款原因。
	// 格式规则：string[1,80]。
	// 业务规则：若商户传入，会在下发给用户的退款消息中体现退款原因。
	Reason string `json:"reason,omitempty"`
	// 字段含义：退款结果回调 url。
	// 格式规则：string[8,256]，必须为外网可访问的 url，不能携带参数。
	// 业务规则：异步接收抖音支付退款结果通知的回调地址。如果参数中传了 notify_url，则商户平台上配置的回调地址将不会生效，优先回调当前传的这个地址。
	NotifyUrl string `json:"notify_url,omitempty"`
	// 字段含义：退款资金来源。
	// 格式规则：string[1,32]。
	// 业务规则：预留字段，商户不需要关注。若传递此参数则使用对应的资金账户退款，否则默认使用未结算资金退款（仅对老资金流商户适用）。
	// 枚举值：AVAILABLE：可用余额账户。
	FundsAccount *ReqFundsAccount `json:"funds_account,omitempty"`
	// 字段含义：金额信息。
	// 格式规则：object。
	// 业务规则：订单金额信息。
	Amount *AmountReq `json:"amount"`
	// 字段含义：退款商品。
	// 格式规则：array。
	// 业务规则：预留字段，商户不需要关注。填写需要指定退款的商品信息，所指定的商品信息需要与下单时传入的单品列表 goods_detail 中的对应商品信息一致，如无需按照指定商品退款，本字段不填。
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// FundsAccount 退款资金账户类型
//
// 枚举值：
//   - UNSETTLED：未结算资金
//   - AVAILABLE：可用余额
//   - UNAVAILABLE：不可用余额
//   - OPERATION：运营户
//   - BASIC：基本账户（含可用余额和不可用余额）
type FundsAccount string

func (e FundsAccount) Ptr() *FundsAccount {
	return &e
}

// Enums of FundsAccount
const (
	FUNDSACCOUNT_UNSETTLED   FundsAccount = "UNSETTLED"   // 未结算资金
	FUNDSACCOUNT_AVAILABLE   FundsAccount = "AVAILABLE"   // 可用余额
	FUNDSACCOUNT_UNAVAILABLE FundsAccount = "UNAVAILABLE" // 不可用余额
	FUNDSACCOUNT_OPERATION   FundsAccount = "OPERATION"   // 运营户
	FUNDSACCOUNT_BASIC       FundsAccount = "BASIC"       // 基本账户（含可用余额和不可用余额）
)

// FundsFromItem 退款出资账户及金额
type FundsFromItem struct {
	// 字段含义：出资账户类型。
	// 格式规则：string[1,32]。
	// 业务规则：下面枚举值多选一。
	// 枚举值：AVAILABLE：可用余额, UNAVAILABLE：不可用余额。
	Account *Account `json:"account"`
	// 字段含义：出资金额。
	// 格式规则：int。
	// 业务规则：对应账户出资金额。只能为整数。
	Amount int64 `json:"amount"`
}

// GoodsDetail 退款商品详情
type GoodsDetail struct {
	// 字段含义：商户侧商品编码。
	// 格式规则：string[1,32]，由半角的大小写字母、数字、中划线、下划线中的一种或几种组成。
	// 业务规则：预留字段。订单下单时传入的商户侧商品编码。
	MerchantGoodsId string `json:"merchant_goods_id"`
	// 字段含义：抖音支付商品编码。
	// 格式规则：string[1,32]。
	// 业务规则：预留字段。订单下单时传入的抖音支付侧商品编码（没有可不传）。
	DouyinpayGoodsId string `json:"douyinpay_goods_id,omitempty"`
	// 字段含义：商品名称。
	// 格式规则：string[1,256]。
	// 业务规则：预留字段。订单下单时传入的商品名称。
	GoodsName string `json:"goods_name,omitempty"`
	// 字段含义：商品单价。
	// 格式规则：int。
	// 业务规则：预留字段。订单下单时传入的商品单价，只能为整数。单位为分。
	UnitPrice int64 `json:"unit_price"`
	// 字段含义：商品退款金额。
	// 格式规则：int。
	// 业务规则：预留字段。商品退款金额，单位为分，只能为整数。
	RefundAmount int64 `json:"refund_amount"`
	// 字段含义：商品退货数量。
	// 格式规则：int。
	// 业务规则：预留字段。对应商品的退货数量，只能为整数。
	RefundQuantity int64 `json:"refund_quantity"`
}

// Promotion 优惠退款详情
type Promotion struct {
	// 字段含义：券 ID。
	// 格式规则：string[1,32]。
	// 业务规则：券或者立减优惠 id。
	PromotionId string `json:"promotion_id"`
	// 字段含义：优惠范围。
	// 格式规则：string[1,32]。
	// 业务规则：优惠活动的适用范围，分为两种类型。
	// 枚举值：GLOBAL：全场代金券（以订单整体可优惠的金额为优惠门槛）, SINGLE：单品优惠（以订单中具体某个单品的总金额为优惠门槛）。
	Scope *Scope `json:"scope"`
	// 字段含义：优惠类型。
	// 格式规则：string[1,32]。
	// 业务规则：分为两种类型。
	// 枚举值：COUPON：代金券（需要走结算资金的充值型代金券，会随订单结算给订单收款商户）, DISCOUNT：优惠券（不走结算资金的免充值型优惠券，无资金结算给订单收款商户）。
	Type *Type `json:"type"`
	// 字段含义：优惠券面额。
	// 格式规则：int。
	// 业务规则：用户享受优惠的金额（优惠券面额 = 抖音出资金额 + 商家出资金额 + 其他出资方金额），单位为分。
	Amount int64 `json:"amount"`
	// 字段含义：优惠退款金额。
	// 格式规则：int。
	// 业务规则：优惠退款金额 <= 退款金额，退款金额 - 代金券或立减优惠退款金额为用户支付的现金，单位为分。
	RefundAmount int64 `json:"refund_amount"`
	// 字段含义：商品列表。
	// 格式规则：array。
	// 业务规则：优惠商品发生退款时返回商品信息。
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// QueryByOutRefundNoRequest 查询单笔退款请求（通过商户退款单号）
type QueryByOutRefundNoRequest struct {
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：应用 ID。
	// 格式规则：string[1,32]。
	// 业务规则：预留字段，商户不需要关注。由抖音支付生成的应用 ID，全局唯一。
	Appid string `json:"appid,omitempty"`
	// 字段含义：商户退款单号。
	// 格式规则：string[1,64]，只能是数字、大小写字母_-|*@。
	// 业务规则：商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔。
	OutRefundNo string `json:"out_refund_no"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成并下发。服务商模式下必须传递此参数。
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：子商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。服务商模式下必须传递此参数。
	SubMchid string `json:"sub_mchid,omitempty"`
}

// Refund 退款响应
type Refund struct {
	// 字段含义：抖音支付退款单号。
	// 格式规则：string[1,32]。
	// 业务规则：抖音支付退款单号。
	RefundId string `json:"refund_id"`
	// 字段含义：商户退款单号。
	// 格式规则：string[1,64]，只能是数字、大小写字母_-|*@。
	// 业务规则：商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔。
	OutRefundNo string `json:"out_refund_no"`
	// 字段含义：抖音支付订单号。
	// 格式规则：string[1,32]。
	// 业务规则：抖音支付交易订单号。
	TransactionId string `json:"transaction_id"`
	// 字段含义：商户订单号。
	// 格式规则：string[1,32]。
	// 业务规则：原支付交易对应的商户订单号。
	OutTradeNo string `json:"out_trade_no"`
	// 字段含义：退款渠道。
	// 格式规则：string[1,16]。
	// 业务规则：退款成功时返回。
	// 枚举值：ORIGINAL：原路退款, BALANCE：退回到余额, OTHER_BALANCE：原账户异常退到其他余额账户, OTHER_BANKCARD：原银行卡异常退到其他银行卡。
	Channel *Channel `json:"channel"`
	// 字段含义：退款入账账户。
	// 格式规则：string[1,64]。
	// 业务规则：默认不返回，若需获取具体信息请联系抖音支付运营。取当前退款单的退款入账方，有以下几种情况：1）退回银行卡：{银行名称}{卡类型}{卡尾号}；2）退回支付用户零钱：支付用户零钱；3）退还商户：商户基本账户商户结算银行账户（暂未开放）。
	UserReceivedAccount string `json:"user_received_account"`
	// 字段含义：退款成功时间。
	// 格式规则：string[1,64]，遵循 RFC3339 标准格式，格式为 yyyy-MM-DDTHH:mm:ss+TIMEZONE。
	// 业务规则：退款成功时间，当退款状态为退款成功时有返回。例如：2015-05-20T13:29:35+08:00 表示北京时间 2015 年 5 月 20 日 13 点 29 分 35 秒。
	SuccessTime string `json:"success_time,omitempty"`
	// 字段含义：退款创建时间。
	// 格式规则：string[1,64]，遵循 RFC3339 标准格式，格式为 yyyy-MM-DDTHH:mm:ss+TIMEZONE。
	// 业务规则：退款受理时间。例如：2015-05-20T13:29:35+08:00 表示北京时间 2015 年 5 月 20 日 13 点 29 分 35 秒。
	CreateTime string `json:"create_time"`
	// 字段含义：退款状态。
	// 格式规则：string[1,32]。
	// 业务规则：枚举值如下。
	// 枚举值：SUCCESS：退款成功, CLOSED：退款关闭, PROCESSING：退款处理中, ABNORMAL：退款异常。
	Status *Status `json:"status"`
	// 字段含义：资金账户。
	// 格式规则：string[1,32]。
	// 业务规则：退款所使用资金对应的资金账户类型。
	// 枚举值：UNSETTLED：未结算资金, AVAILABLE：可用余额, UNAVAILABLE：不可用余额, OPERATION：运营户, BASIC：基本账户（含可用余额和不可用余额）。
	FundsAccount *FundsAccount `json:"funds_account,omitempty"`
	// 字段含义：金额信息。
	// 格式规则：object。
	// 业务规则：金额详细信息。
	Amount *Amount `json:"amount"`
	// 字段含义：优惠退款信息。
	// 格式规则：array。
	// 业务规则：订单优惠的退款详情，订单使用了优惠且优惠发生退款时返回（仅退款状态为退款成功时返回）。
	PromotionDetail []Promotion `json:"promotion_detail,omitempty"`
}

// ReqFundsAccount 退款资金来源（请求参数用）
//
// 枚举值：
//   - AVAILABLE：可用余额（仅对老资金流商户适用，指定从可用余额账户出资）
type ReqFundsAccount string

func (e ReqFundsAccount) Ptr() *ReqFundsAccount {
	return &e
}

// Enums of ReqFundsAccount
const (
	REQFUNDSACCOUNT_AVAILABLE ReqFundsAccount = "AVAILABLE" // 可用余额
)

// Scope 优惠范围
//
// 枚举值：
//   - GLOBAL：全场代金券（全场优惠类型，以订单整体可优惠的金额为优惠门槛）
//   - SINGLE：单品优惠（单品优惠类型，以订单中具体某个单品的总金额为优惠门槛）
type Scope string

func (e Scope) Ptr() *Scope {
	return &e
}

// Enums of Scope
const (
	SCOPE_GLOBAL Scope = "GLOBAL" // 全场代金券
	SCOPE_SINGLE Scope = "SINGLE" // 单品优惠
)

// Status 退款状态
//
// 枚举值：
//   - SUCCESS：退款成功
//   - CLOSED：退款关闭
//   - PROCESSING：退款处理中
//   - ABNORMAL：退款异常
type Status string

func (e Status) Ptr() *Status {
	return &e
}

// Enums of Status
const (
	STATUS_SUCCESS    Status = "SUCCESS"    // 退款成功
	STATUS_CLOSED     Status = "CLOSED"     // 退款关闭
	STATUS_PROCESSING Status = "PROCESSING" // 退款处理中
	STATUS_ABNORMAL   Status = "ABNORMAL"   // 退款异常
)

// Type 优惠类型
//
// 枚举值：
//   - COUPON：代金券（需要走结算资金的充值型代金券，会随订单结算给订单收款商户）
//   - DISCOUNT：优惠券（不走结算资金的免充值型优惠券，无资金结算给订单收款商户）
type Type string

func (e Type) Ptr() *Type {
	return &e
}

// Enums of Type
const (
	TYPE_COUPON   Type = "COUPON"   // 代金券
	TYPE_DISCOUNT Type = "DISCOUNT" // 优惠券
)
