package partnerpay

// Transaction
type Transaction struct {
	SpMchid         string             `json:"sp_mchid,omitempty"`
	SpAppid         string             `json:"sp_appid,omitempty"`
	SubMchid        string             `json:"sub_mchid,omitempty"`
	SubAppid        string             `json:"sub_appid,omitempty"`
	Amount          *TransactionAmount `json:"amount,omitempty"`
	Attach          string             `json:"attach,omitempty"`
	BankType        string             `json:"bank_type,omitempty"`
	OutTradeNo      string             `json:"out_trade_no,omitempty"`
	Payer           *TransactionPayer  `json:"payer,omitempty"`
	PromotionDetail []PromotionDetail  `json:"promotion_detail,omitempty"`
	SuccessTime     string             `json:"success_time,omitempty"`
	TradeState      string             `json:"trade_state,omitempty"`
	TradeStateDesc  string             `json:"trade_state_desc,omitempty"`
	TradeType       string             `json:"trade_type,omitempty"`
	TransactionId   string             `json:"transaction_id,omitempty"`
	SceneInfo       *SceneInfo         `json:"scene_info,omitempty"`
	// 委托代扣协议号，扣款失败通知中返回。
	ContractId string `json:"contract_id,omitempty"`
	// 错误代码，申请扣款失败且交易状态为 CLOSED 时返回。
	ErrCode string `json:"err_code,omitempty"`
	// 错误代码描述，扣款失败原因描述。
	ErrCodeDes string `json:"err_code_des,omitempty"`
}

// TransactionAmount
type TransactionAmount struct {
	Currency      string `json:"currency,omitempty"`
	PayerCurrency string `json:"payer_currency,omitempty"`
	PayerTotal    int64  `json:"payer_total,omitempty"`
	Total         int64  `json:"total,omitempty"`
}

// TransactionPayer
type TransactionPayer struct {
	SpOpenId  string `json:"sp_openid,omitempty"`
	SubOpenId string `json:"sub_openid,omitempty"`
}

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
