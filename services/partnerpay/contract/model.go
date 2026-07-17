package contract

// PromotionDetail
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
	// 抖音出资，单位为分
	DouyinpayContribute int64 `json:"douyinpay_contribute,omitempty"`
	// 商户出资，单位为分
	MerchantContribute int64 `json:"merchant_contribute,omitempty"`
	// 其他出资金额，单位为分
	OtherContribute int64 `json:"other_contribute,omitempty"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency    string                 `json:"currency,omitempty"`
	GoodsDetail []PromotionGoodsDetail `json:"goods_detail,omitempty"`
}

// PromotionGoodsDetail
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

type PartnerContractOrderRequest struct {
	// 服务商应用Id
	SpAppid string `json:"sp_appid"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户应用Id
	SubAppid string `json:"sub_appid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
	// 商品描述
	Description string `json:"description,omitempty"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 交易结束时间
	TimeExpire string `json:"time_expire,omitempty"`
	//交易类型
	TradeType string `json:"trade_type,omitempty"`
	// 附加数据
	Attach string `json:"attach,omitempty"`
	// 通知地址
	NotifyUrl string `json:"notify_url,omitempty"`
	// 优惠标记
	GoodsTag string `json:"goods_tag,omitempty"`
	// 电子发票入口开放标识
	SupportFapiao bool `json:"support_fapiao,omitempty"`
	// 订单金额
	Amount *Amount `json:"amount"`
	// 支付者
	Payer *Payer `json:"payer"`
	//优惠信息
	Detail *Detail `json:"detail,omitempty"`
	// 商品列表
	GoodsDetail []GoodsDetail `json:"goods_detail"`
	// 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 商户门店信息（预留字段）
	//StoreInfo *StoreInfo `json:"store_info"`
	// 结算信息
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
	//
	ContractInfo *ContractInfo `json:"contract_info,omitempty"`
}

// ContractPayer 支付者
type Payer struct {
	// 用户服务标识
	SpOpenid string `json:"sp_openid"`
	// 用户子标识
	SubOpenid string `json:"sub_openid"`
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
	CostPrice int64  `json:"cost_price"`
	InvoiceId string `json:"invoice_id,omitempty"`
}

type PayApplyDetail struct {
	CostPrice   int64         `json:"cost_price"`
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
	Quantity int64 `json:"quantity,omitempty"`
	// 商品单价，单位为分。
	UnitPrice int64 `json:"unit_price,omitempty"`
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	// 用户终端IP
	PayerClientIp string `json:"payer_client_ip,omitempty"`
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

type ContractInfo struct {
	//签约商户号
	ContractMchId string `json:"contract_mchid,omitempty"`
	//签约appid
	ContractAppId string `json:"contract_appid,omitempty"`
	//模版id
	PlanId string `json:"plan_id,omitempty"`
	//签约协议号
	OutContractCode string `json:"out_contract_code,omitempty"`
	//请求序号
	RequestSerial int64 `json:"request_serial,omitempty"`
	//用户账户展示名称
	ContractDisplayAccount string `json:"contract_display_account,omitempty"`
	//签约信息通知url
	ContractNotifyUrl string `json:"contract_notify_url,omitempty"`
	// 签约拓展业务参数
	ContractExt string `json:"contract_ext,omitempty"`
}

// ContractOrderResponse
type PartnerContractOrderResponse struct {
	// 预支付交易会话标识
	PrepayId string `json:"prepay_id"`
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

// 申请扣款
type PartnerPayApplyRequest struct {
	// 服务商应用Id
	SpAppid string `json:"sp_appid"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户应用Id
	SubAppid string `json:"sub_appid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 交易结束时间
	TimeExpire string `json:"time_expire,omitempty"`
	// 代扣协议id
	ContractId string `json:"contract_id,omitempty"`
	//交易类型
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
	Detail *PayApplyDetail `json:"detail,omitempty"`
	// 商品列表
	// GoodsDetail []PayApplyDetail `json:"goods_detail,omitempty"`
	// 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 结算信息
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
	// 订单金额
	Amount *Amount `json:"amount"`
}

type PartnerPayApplyResponse struct {
	// 业务结果
	ResultCode string `json:"result_code"`
}
