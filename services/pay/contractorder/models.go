package contractorder

// PrepayRequest
type PrepayRequest struct {
	// 公众号ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 商品描述
	Description string `json:"description"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no"`
	// 订单失效时间，格式为rfc3339格式
	TimeExpire string `json:"time_expire,omitempty"`
	// 附加数据
	Attach string `json:"attach,omitempty"`
	// 有效性：1. HTTPS；2. 不允许携带查询串。
	NotifyUrl string `json:"notify_url"`
	//交易类型
	TradeType string `json:"trade_type"`
	//用户标识
	Openid string `json:"openid,omitempty"`
	// 商品标记，代金券或立减优惠功能的参数。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 传入true时，支付成功消息和支付详情页将出现开票入口。需要在抖音支付商户平台或抖音公众平台开通电子发票功能，传此字段才可生效。
	SupportFapiao bool          `json:"support_fapiao,omitempty"`
	Amount        *Amount       `json:"amount"`
	Detail        *Detail       `json:"detail,omitempty"`
	SceneInfo     *SceneInfo    `json:"scene_info,omitempty"`
	SettleInfo    *SettleInfo   `json:"settle_info,omitempty"`
	ContractInfo  *ContractInfo `json:"contract_info,omitempty"`
}

// PrepayResponse
type PrepayResponse struct {
	// 预支付交易会话标识
	PrepayId string `json:"prepay_id"`
	H5Url    string `json:"h5_url"`
}

// CloseOrderRequest
type CloseOrderRequest struct {
	OutTradeNo string `json:"out_trade_no"`
	Mchid      string `json:"mchid"`
}

// CloseRequest
type CloseRequest struct {
	Mchid string `json:"mchid"`
}

// QueryOrderByIdRequest
type QueryOrderByIdRequest struct {
	TransactionId string `json:"transaction_id"`
	Mchid         string `json:"mchid"`
}

// QueryOrderByOutTradeNoRequest
type QueryOrderByOutTradeNoRequest struct {
	OutTradeNo string `json:"out_trade_no"`
	Mchid      string `json:"mchid"`
}

// Amount
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

// GoodsDetail
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
	PayerDeviceId string     `json:"payer_device_id,omitempty"`
	StoreInfo     *StoreInfo `json:"store_info,omitempty"`
}

// SettleInfo
type SettleInfo struct {
	// 是否指定分账
	ProfitSharing bool `json:"profit_sharing,omitempty"`
}

// StoreInfo 商户门店信息（预留字段）
type StoreInfo struct {
	// 商户侧门店编号
	Id string `json:"id"`
	// 商户侧门店名称
	Name string `json:"name,omitempty"`
	// 地区编码，详细请见抖音支付提供的文档
	AreaCode string `json:"area_code,omitempty"`
	// 详细的商户门店地址
	Address string `json:"address,omitempty"`
}

type ContractInfo struct {
	ContractMchId          string `json:"contract_mchid,omitempty"`
	ContractAppId          string `json:"contract_appid,omitempty"`
	PlanId                 string `json:"plan_id,omitempty"`
	OutContractCode        string `json:"out_contract_code,omitempty"`
	RequestSerial          int64  `json:"request_serial,omitempty"`
	ContractDisplayAccount string `json:"contract_display_account,omitempty"`
	ContractNotifyUrl      string `json:"contract_notify_url,omitempty"`
}

type PayResultNotifyBody struct {
	// 应用Id
	Appid string `json:"appid"`
	// 商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no"`
	// 抖音支付订单号
	TransactionId string `json:"transaction_id"`
	// 委托代扣协议id
	ContractId string `json:"contract_id"`
	// 交易类型
	TradeType string `json:"trade_type"`
	// 交易状态
	TradeState string `json:"trade_state"`
	// 交易状态描述
	TradeStateDesc string `json:"trade_state_desc"`
	// 付款银行
	BankType string `json:"bank_type"`
	// 附加数据
	Attach string `json:"attach"`
	// 支付完成时间
	SuccessTime string `json:"success_time"`
	// 支付者
	Payer Payer `json:"payer"`
	// 订单金额
	Amount BaseAmountInfo `json:"amount"`
	// 场景信息
	SceneInfo SceneNotifyInfo `json:"scene_info"`
	// 优惠功能
	PromotionDetail []PromotionDetail `json:"promotion_detail"`
}

type Payer struct {
	// 用户标识
	Openid string `json:"openid"`
}

type BaseAmountInfo struct {
	// 订单总金额，单位为分
	Total int64 `json:"total"`
	// 用户支付金额，单位为分
	PayerTotal int64 `json:"payer_total"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency string `json:"currency"`
	// 用户支付币种
	PayerCurrency string `json:"payer_currency"`
}

type SceneNotifyInfo struct {
	// 商户端设备号（预留字段）
	DeviceId string `json:"device_id"`
}

type PromotionDetail struct {
	// 券Id
	CouponId string `json:"coupon_id"`
	// 优惠名称
	Name string `json:"name"`
	// 优惠范围
	Scope string `json:"scope"`
	// 优惠类型
	Type string `json:"type"`
	// 优惠券面额
	Amount int64 `json:"amount"`
	// 活动id
	StockId string `json:"stock_id"`
	// 平台出资
	DouyinpayContribute int64 `json:"douyinpay_contribute"`
	// 商户出资
	MerchantContribute int64 `json:"merchant_contribute"`
	// 其他出资
	OtherContribute int64 `json:"other_contribute"`
	// 优惠币种
	Currency string `json:"currency"`
	//
	GoodsDetail []GoodsDetailNotify `json:"goods_detail"`
	//

}

// 商品详情
type GoodsDetailNotify struct {
	// 商品编码
	GoodsId string `json:"goods_id"`
	// 商品数量
	Quantity int64 `json:"quantity"`
	// 商品单价
	UnitPrice int64 `json:"unit_price"`
	// 商品优惠金额
	DiscountAmount int64 `json:"discount_amount"`
	// 商品备注
	GoodsRemark string `json:"goods_remark"`
}
