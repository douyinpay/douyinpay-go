package pay

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

// Transaction
type Transaction struct {
	Amount          *TransactionAmount    `json:"amount,omitempty"`
	Appid           string                `json:"appid,omitempty"`
	Attach          string                `json:"attach,omitempty"`
	BankType        string                `json:"bank_type,omitempty"`
	Mchid           string                `json:"mchid,omitempty"`
	OutTradeNo      string                `json:"out_trade_no,omitempty"`
	Payer           *TransactionPayer     `json:"payer,omitempty"`
	PromotionDetail []PromotionDetail     `json:"promotion_detail,omitempty"`
	SuccessTime     string                `json:"success_time,omitempty"`
	TradeState      string                `json:"trade_state,omitempty"`
	TradeStateDesc  string                `json:"trade_state_desc,omitempty"`
	TradeType       string                `json:"trade_type,omitempty"`
	TransactionId   string                `json:"transaction_id,omitempty"`
	SceneInfo       *TransactionSceneInfo `json:"scene_info,omitempty"`
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
	Openid string `json:"openid,omitempty"`
}

// TransactionSceneInfo
type TransactionSceneInfo struct {
	DeviceId string `json:"device_id,omitempty"`
}
