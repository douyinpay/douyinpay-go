package jsapi

// PrepayRequest
type PrepayRequest struct {
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 服务商应用ID
	SpAppid string `json:"sp_appid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
	// 子商户应用ID
	SubAppid string `json:"sub_appid"`
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
	// 商品标记，代金券或立减优惠功能的参数。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 传入true时，支付成功消息和支付详情页将出现开票入口。需要在抖音支付商户平台或抖音公众平台开通电子发票功能，传此字段才可生效。
	SupportFapiao bool        `json:"support_fapiao,omitempty"`
	Amount        *Amount     `json:"amount"`
	Detail        *Detail     `json:"detail,omitempty"`
	SceneInfo     *SceneInfo  `json:"scene_info,omitempty"`
	SettleInfo    *SettleInfo `json:"settle_info,omitempty"`
	PayerInfo     *PayerInfo  `json:"payer"`
}

// PrepayResponse
type PrepayResponse struct {
	// 预支付交易会话标识
	PrepayId string `json:"prepay_id"`
}

// CloseOrderRequest
type CloseOrderRequest struct {
	OutTradeNo string `json:"out_trade_no"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
}

// CloseRequest
type CloseRequest struct {
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
}

// QueryOrderByIdRequest
type QueryOrderByIdRequest struct {
	//支付订单号
	TransactionId string `json:"transaction_id"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
}

// QueryOrderByOutTradeNoRequest
type QueryOrderByOutTradeNoRequest struct {
	//支付订单号
	OutTradeNo string `json:"out_trade_no"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
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
	DeviceId  string     `json:"device_id,omitempty"`
	StoreInfo *StoreInfo `json:"store_info,omitempty"`
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

type PayerInfo struct {
	SpOpenId  string `json:"sp_openid"`
	SubOpenId string `json:"sub_openid"`
}
