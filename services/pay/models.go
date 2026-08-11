package pay

// PromotionDetail
type PromotionDetail struct {
	// 券ID
	CouponId string `json:"coupon_id,omitempty"`
	// 优惠名称
	Name string `json:"name,omitempty"`
	// 优惠范围，GLOBAL：全场代金券；SINGLE：单品优惠
	Scope string `json:"scope,omitempty"`
	// 优惠类型，CASH：充值；NOCASH：预充值。
	Type string `json:"type,omitempty"`
	// 优惠券面额，单位为分
	Amount int64 `json:"amount,omitempty"`
	// 活动ID，批次ID
	StockId string `json:"stock_id,omitempty"`
	// 平台出资金额，单位为分
	DouyinpayContribute int64 `json:"douyinpay_contribute,omitempty"`
	// 商户出资金额，单位为分
	MerchantContribute int64 `json:"merchant_contribute,omitempty"`
	// 其他出资金额，单位为分
	OtherContribute int64 `json:"other_contribute,omitempty"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency string `json:"currency,omitempty"`
	// 字段含义：单品列表。
	GoodsDetail []PromotionGoodsDetail `json:"goods_detail,omitempty"`
}

// PromotionGoodsDetail
type PromotionGoodsDetail struct {
	// 商品编码
	GoodsId string `json:"goods_id"`
	// 商品数量
	Quantity int64 `json:"quantity"`
	// 商品单价，单位为分
	UnitPrice int64 `json:"unit_price"`
	// 商品优惠金额，单位为分
	DiscountAmount int64 `json:"discount_amount"`
	// 商品备注
	GoodsRemark string `json:"goods_remark,omitempty"`
}

// Transaction
type Transaction struct {
	// 字段含义：订单金额信息。
	Amount *TransactionAmount `json:"amount,omitempty"`
	// 字段含义：应用ID。
	Appid string `json:"appid,omitempty"`
	// 字段含义：附加数据。
	Attach string `json:"attach,omitempty"`
	// 字段含义：付款银行。
	BankType string `json:"bank_type,omitempty"`
	// 字段含义：直连商户号。
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：商户订单号。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 字段含义：支付者。
	Payer *TransactionPayer `json:"payer,omitempty"`
	// 字段含义：优惠功能。
	PromotionDetail []PromotionDetail `json:"promotion_detail,omitempty"`
	// 字段含义：支付完成时间。
	SuccessTime string `json:"success_time,omitempty"`
	// 字段含义：交易状态。
	TradeState string `json:"trade_state,omitempty"`
	// 字段含义：交易状态描述。
	TradeStateDesc string `json:"trade_state_desc,omitempty"`
	// 字段含义：交易类型。
	TradeType string `json:"trade_type,omitempty"`
	// 字段含义：抖音支付订单号。
	TransactionId string `json:"transaction_id,omitempty"`
	// 字段含义：场景信息。
	SceneInfo *TransactionSceneInfo `json:"scene_info,omitempty"`
	// 字段含义：委托代扣协议号。
	ContractId string `json:"contract_id,omitempty"`
	// 字段含义：错误代码。
	ErrCode string `json:"err_code,omitempty"`
	// 字段含义：错误代码描述。
	ErrCodeDes string `json:"err_code_des,omitempty"`
}

// TransactionAmount
type TransactionAmount struct {
	// 字段含义：货币类型。
	Currency string `json:"currency,omitempty"`
	// 字段含义：用户支付币种。
	PayerCurrency string `json:"payer_currency,omitempty"`
	// 字段含义：用户支付金额。
	PayerTotal int64 `json:"payer_total,omitempty"`
	// 字段含义：总金额。
	Total int64 `json:"total,omitempty"`
}

// TransactionPayer
type TransactionPayer struct {
	// 字段含义：用户标识。
	Openid string `json:"openid,omitempty"`
}

// TransactionSceneInfo
type TransactionSceneInfo struct {
	// 字段含义：商户端设备ID。
	DeviceId string `json:"device_id,omitempty"`
}
