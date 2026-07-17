package contractorder

// ApiDeductRequest 申请扣款请求参数
type ApiDeductRequest struct {
	// 公众号ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no"`
	// 订单失效时间，格式为rfc3339格式
	TimeExpire string `json:"time_expire,omitempty"`
	// 签约成功后抖音支付返回的委托代扣协议id
	ContractId string `json:"contract_id"`
	// 交易类型
	TradeType string `json:"trade_type"`
	// 商品描述
	Description string `json:"description"`
	// 通知地址 有效性：1. HTTPS；2. 不允许携带查询串。
	NotifyUrl string `json:"notify_url"`
	// 附加数据
	Attach string `json:"attach,omitempty"`
	// 优惠标记
	GoodsTag string `json:"goods_tag"`

	Detail     *Detail     `json:"detail,omitempty"`
	Amount     *Amount     `json:"amount"`
	SceneInfo  *SceneInfo  `json:"scene_info,omitempty"`
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
}

// ApiDeductResponse 申请扣款响应参数
type ApiDeductResponse struct {
	// 业务结果
	ResultCode string `json:"result_code"`

	//预支付交易会话标识
	PrepayId string `json:"prepay_id"`
}

// CloseOrderRequest 关单请求参数
type CloseOrderRequest struct {
	OutTradeNo string `json:"out_trade_no"`
	Mchid      string `json:"mchid"`
}

// CloseRequest 关单请求参数
type CloseRequest struct {
	Mchid string `json:"mchid"`
}

// QueryOrderByIdRequest 查单请求参数
type QueryOrderByIdRequest struct {
	TransactionId string `json:"transaction_id"`
	Mchid         string `json:"mchid"`
}

// QueryOrderByOutTradeNoRequest 查单请求参数
type QueryOrderByOutTradeNoRequest struct {
	OutTradeNo string `json:"out_trade_no"`
	Mchid      string `json:"mchid"`
}

// Amount 金额信息
type Amount struct {
	// 订单总金额，单位为分
	Total int64 `json:"total"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency string `json:"currency,omitempty"`
}

// Detail 优惠功能
type Detail struct {
	CostPrice   int64         `json:"cost_price,omitempty"`
	InvoiceId   string        `json:"invoice_id,omitempty"`
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// GoodsDetail 商品明细
type GoodsDetail struct {
	// 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成。
	MerchantGoodsId string `json:"merchant_goods_id"`
	// 抖音支付定义的统一商品编号（没有可不传）。
	DouyinpayGoodsId string `json:"douyinpay_goods_id,omitempty"`
	// 商品的实际名称。
	GoodsName string `json:"goods_name,omitempty"`
	// 用户购买的数量。
	Quantity int64 `json:"quantity"`
	// 商品单价，单位为分。
	UnitPrice int64 `json:"unit_price"`
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	// 用户终端IP
	PayerClientIp string `json:"payer_client_ip"`
	// 商户端设备号（预留字段）
	DeviceId string `json:"device_id,omitempty"`
	// 用户终端设备号
	PayerDeviceId string `json:"payer_device_id,omitempty"`
}

// SettleInfo 结算信息
type SettleInfo struct {
	// 是否指定分账
	ProfitSharing bool `json:"profit_sharing,omitempty"`
}

// Transaction 交易信息
type Transaction struct {
	Amount          *TransactionAmount `json:"amount,omitempty"`
	Appid           string             `json:"appid,omitempty"`
	Attach          string             `json:"attach,omitempty"`
	BankType        string             `json:"bank_type,omitempty"`
	Mchid           string             `json:"mchid,omitempty"`
	OutTradeNo      string             `json:"out_trade_no,omitempty"`
	Payer           *TransactionPayer  `json:"payer,omitempty"`
	PromotionDetail []PromotionDetail  `json:"promotion_detail,omitempty"`
	SuccessTime     string             `json:"success_time,omitempty"`
	TradeState      string             `json:"trade_state,omitempty"`
	TradeStateDesc  string             `json:"trade_state_desc,omitempty"`
	TradeType       string             `json:"trade_type,omitempty"`
	TransactionId   string             `json:"transaction_id,omitempty"`
}

// TransactionAmount 交易金额信息
type TransactionAmount struct {
	Currency      string `json:"currency,omitempty"`
	PayerCurrency string `json:"payer_currency,omitempty"`
	PayerTotal    int64  `json:"payer_total,omitempty"`
	Total         int64  `json:"total,omitempty"`
}

// TransactionPayer 付款方信息
type TransactionPayer struct {
	Openid string `json:"openid,omitempty"`
}

// PromotionDetail 营销信息
type PromotionDetail struct {
	// 券ID
	CouponId string `json:"coupon_id,omitempty"`
	// 优惠名称
	Name string `json:"name,omitempty"`
	// GLOBAL：全场代金券；SINGLE：单品优惠
	Scope string `json:"scope,omitempty"`
	// CASH：充值；NOCASH：预充值。
	Type string `json:"type,omitempty"`
	// 优惠券面额
	Amount int64 `json:"amount,omitempty"`
	// 活动ID，批次ID
	StockId string `json:"stock_id,omitempty"`
	// 单位为分
	DouyinpayContribute int64 `json:"douyinpay_contribute,omitempty"`
	// 单位为分
	MerchantContribute int64 `json:"merchant_contribute,omitempty"`
	// 单位为分
	OtherContribute int64 `json:"other_contribute,omitempty"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency    string                 `json:"currency,omitempty"`
	GoodsDetail []PromotionGoodsDetail `json:"goods_detail,omitempty"`
}

// PromotionGoodsDetail 营销明细信息
type PromotionGoodsDetail struct {
	// 商品编码
	GoodsId string `json:"goods_id"`
	// 商品数量
	Quantity int64 `json:"quantity"`
	// 商品价格
	UnitPrice int64 `json:"unit_price"`
	// 商品优惠金额
	DiscountAmount int64 `json:"discount_amount"`
	// 商品备注
	GoodsRemark string `json:"goods_remark,omitempty"`
}

// DeductNotifyRequest 预约扣费通知请求参数
type DeductNotifyRequest struct {
	// 委托代扣协议id
	ContractId string `json:"contract_id,omitempty"`
	// 直连商户号
	Mchid string `json:"mchid,omitempty"`
	// 应用id
	Appid string `json:"appid,omitempty"`
	// 预约扣费金额信息
	EstimatedAmount DeductAmount `json:"estimated_amount,omitempty"`
}

type DeductAmount struct {
	// 预约扣费金额
	Amount int64 `json:"amount,omitempty"`
	// 预约扣费币种
	Currency string `json:"currency,omitempty"`
}

// 申请扣款
type PayApplyRequest struct {
	//应用Id
	Appid string `json:"appid"`
	//商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no,omitempty"`
	//交易结束时间
	TimeExpire string `json:"time_expire,omitempty"`
	//委托代扣协议id
	ContractId string `json:"contract_id,omitempty"`
	// 交易类型
	TradeType string `json:"trade_type,omitempty"`
	// 商品描述
	Description string `json:"description,omitempty"`
	// 通知地址
	NotifyUrl string `json:"notify_url,omitempty"`
	// 附加数据
	Attach string `json:"attach,omitempty"`
	// 优惠标记
	GoodsTag string `json:"goods_tag,omitempty"`
	// 优惠信息
	Detail Detail `json:"detail,omitempty"`
	// 订单金额
	Amount *Amount `json:"amount,omitempty"`
	// 场景信息
	SceneInfo SceneInfo `json:"scene_info,omitempty"`
	// 结算信息
	SettleInfo SettleInfo `json:"settle_info,omitempty"`
}

// 申请扣款 响应体
type PayApplyResponse struct {
	// 业务结果
	ResultCode string `json:"result_code,omitempty"`
	// 预支付交易会话标识
	PrepayId string `json:"prepay_id,omitempty"`
}

// 预约扣费
type PartnerContractScheduleRequest struct {
	// 代扣协议id
	ContractId string `json:"contract_id,omitempty"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
	// 预约扣费金额
	ScheduleAmount Amount `json:"schedule_amount,omitempty"`
}

// 预约扣费
type PartnerContractScheduleResponse struct {
	// 可扣费开始日期
	DeductStartDate string `json:"deduct_start_date,omitempty"`
	// 可扣费结束日期
	DeductEndDate string `json:"deduct_end_date,omitempty"`
	// 预约扣费金额
	ScheduleAmount Amount `json:"schedule_amount,omitempty"`
}

// 预约扣费结果查询
type PartnerContractScheduleQueryRequest struct {
	// 代扣协议id
	ContractId string `json:"contract_id,omitempty"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
}

type PartnerContractScheduleQueryResponse struct {
	// 可扣费开始日期
	DeductStartDate string `json:"deduct_start_date"`
	// 可扣费结束日期
	DeductEndDate string `json:"deduct_end_date"`
	// 预约扣费金额
	ScheduleAmount Amount `json:"schedule_amount"`
	// 预扣费通知状态
	ScheduleState string `json:"schedule_state"`
	// 实际扣费金额
	DeductAmount Amount `json:"deduct_amount"`
	// 实际扣费日期
	DeductDate string `json:"deduct_date"`
}
