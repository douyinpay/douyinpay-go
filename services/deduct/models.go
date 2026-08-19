package contractorder

// ApiDeductRequest 申请扣款请求参数
type ApiDeductRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用 ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]，只能是数字、大小写字母、下划线、中划线、星号。
	// 业务规则：商户系统内部订单号，同一商户号下唯一。
	// 示例：OUT_1666688488。
	OutTradeNo string `json:"out_trade_no"`
	// 字段含义：交易结束时间。
	// 格式规则：string[1,64]，遵循 RFC3339 标准格式。
	// 业务规则：该笔订单允许的最晚用户付款时间，并非订单关闭的时间；超过后用户无法支付，若需继续扣款需使用新的商户订单号发起扣款；传递时间需在请求时间 15 天以内，超过 15 天时抖音支付会自动调整为请求时间后的第 15 天。
	// 示例：2018-06-08T10:34:56+08:00。
	TimeExpire string `json:"time_expire,omitempty"`
	// 字段含义：委托代扣协议ID。
	// 格式规则：string[1,64]。
	// 业务规则：签约成功后抖音支付返回的委托代扣协议 ID。
	// 示例：MSN20230314112037389849955326013。
	ContractId string `json:"contract_id"`
	// 字段含义：交易类型。
	// 格式规则：string[1,3]。
	// 业务规则：SGP 表示商户代扣；NPP 表示免密支付。
	TradeType string `json:"trade_type"`
	// 字段含义：商品描述。
	// 格式规则：string[1,128]。
	// 业务规则：商品信息描述，会展示在用户抖音钱包账单的"商品说明"内，需传递能真实代表商品信息的描述。
	// 示例：测试商品。
	Description string `json:"description"`
	// 字段含义：通知地址。
	// 格式规则：URL，string[1,256]，必须为 HTTPS 地址，不允许携带查询串。
	// 业务规则：抖音支付通过该地址通知扣款结果。
	// 示例：https://www.mock.douyinpay.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：附加数据。
	// 格式规则：string[1,1024]。
	// 业务规则：在查询 API 和支付通知中原样返回，可作为自定义参数使用，实际情况下只有支付完成状态才会返回该字段。
	// 示例：自定义数据。
	Attach string `json:"attach,omitempty"`
	// 字段含义：优惠标记。
	// 格式规则：string[1,512]，JSON 字符串。
	// 业务规则：和抖音支付协商后可用。
	// 示例：""。
	GoodsTag string `json:"goods_tag"`

	// 字段含义：优惠信息。
	// 格式规则：object。
	// 业务规则：为预留字段，商户不需要传。
	Detail *Detail `json:"detail,omitempty"`
	// 字段含义：订单金额。
	// 格式规则：object。
	// 业务规则：订单金额信息。
	// 示例：{"currency":"CNY","total":100}。
	Amount *Amount `json:"amount"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 业务规则：支付场景描述，免密支付产品需要传。
	// 示例：{"payer_client_ip":"14.23.150.211","device_id":"13467007045764"}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 字段含义：结算信息。
	// 格式规则：object。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：{"profit_sharing":false}。
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
}

// ApiDeductResponse 申请扣款响应参数
type ApiDeductResponse struct {
	// 字段含义：业务结果，以回调或查单返回的扣款结果为准。
	ResultCode string `json:"result_code"`
	// 字段含义：预支付交易会话标识。
	PrepayId string `json:"prepay_id"`
}

// CloseOrderRequest 关单请求参数
type CloseOrderRequest struct {
	// 字段含义：商户订单号。
	// 格式规则：string。
	// 业务规则：商户系统内部订单号，同一商户号下唯一。
	OutTradeNo string `json:"out_trade_no"`
	// 字段含义：直连商户号。
	// 格式规则：string。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	Mchid string `json:"mchid"`
}

// CloseRequest 关单请求参数
type CloseRequest struct {
	// 字段含义：直连商户号。
	// 格式规则：string。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	Mchid string `json:"mchid"`
}

// QueryOrderByIdRequest 查单请求参数
type QueryOrderByIdRequest struct {
	// 字段含义：抖音支付订单号。
	// 格式规则：string。
	// 业务规则：抖音支付系统生成的订单号。
	TransactionId string `json:"transaction_id"`
	// 字段含义：直连商户号。
	// 格式规则：string。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	Mchid string `json:"mchid"`
}

// QueryOrderByOutTradeNoRequest 查单请求参数
type QueryOrderByOutTradeNoRequest struct {
	// 字段含义：商户订单号。
	// 格式规则：string。
	// 业务规则：商户系统内部订单号，同一商户号下唯一。
	OutTradeNo string `json:"out_trade_no"`
	// 字段含义：直连商户号。
	// 格式规则：string。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	Mchid string `json:"mchid"`
}

// Amount 金额信息
type Amount struct {
	// 字段含义：订单总金额，单位为分。
	// 格式规则：int64。
	// 业务规则：订单金额信息中的总金额。
	// 示例：100。
	Total int64 `json:"total"`
	// 字段含义：货币类型。
	// 格式规则：string[1,16]。
	// 业务规则：CNY 表示人民币，境内商户号仅支持人民币。
	// 示例：CNY。
	Currency string `json:"currency,omitempty"`
}

// Detail 优惠功能
type Detail struct {
	// 字段含义：订单原价。
	// 格式规则：int64。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：608800。
	CostPrice int64 `json:"cost_price,omitempty"`
	// 字段含义：商品小票ID。
	// 格式规则：string[1,32]。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：dy123。
	InvoiceId string `json:"invoice_id,omitempty"`
	// 字段含义：单品列表。
	// 格式规则：array。
	// 业务规则：为预留字段，商户不需要传。
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// GoodsDetail 商品明细
type GoodsDetail struct {
	// 字段含义：商户侧商品编码。
	// 格式规则：string[1,32]。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：ABC。
	MerchantGoodsId string `json:"merchant_goods_id"`
	// 字段含义：抖音支付商品编码。
	// 格式规则：string[1,32]。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：1001。
	DouyinpayGoodsId string `json:"douyinpay_goods_id,omitempty"`
	// 字段含义：商品名称。
	// 格式规则：string[1,256]。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：芒果TV-会员费。
	GoodsName string `json:"goods_name,omitempty"`
	// 字段含义：商品数量。
	// 格式规则：int64。
	// 业务规则：为预留字段，商户不需要传。
	Quantity int64 `json:"quantity"`
	// 字段含义：商品单价，单位为分。
	// 格式规则：int64。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：828800。
	UnitPrice int64 `json:"unit_price"`
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	// 字段含义：用户终端IP。
	// 格式规则：string[1,45]，支持 IPv4 和 IPv6 两种格式。
	// 业务规则：用户的客户端 IP。
	// 示例：14.23.150.211。
	PayerClientIp string `json:"payer_client_ip"`
	// 字段含义：商户端设备号。
	// 格式规则：string[1,32]。
	// 业务规则：商户端设备号，可传门店号或收银设备 ID。
	// 示例：13467007045764。
	DeviceId string `json:"device_id,omitempty"`
	// 字段含义：用户终端设备号。
	// 格式规则：string[1,32]。
	// 业务规则：安卓优先传 android_id（openudid），若没有则传 IMEI；iOS 优先传 IDFV，若没有 IDFV 则传 IDFA。
	PayerDeviceId string `json:"payer_device_id,omitempty"`
}

// SettleInfo 结算信息
type SettleInfo struct {
	// 字段含义：是否指定分账。
	// 格式规则：boolean。
	// 业务规则：传入 true 表示订单支付成功后可进行分账；传入 false 或不传时，订单收款成功后资金不会冻结，直接转入基本账户可用余额。
	// 示例：false。
	ProfitSharing bool `json:"profit_sharing,omitempty"`
}

// Transaction 交易信息
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
}

// TransactionAmount 交易金额信息
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

// TransactionPayer 付款方信息
type TransactionPayer struct {
	// 字段含义：用户标识。
	Openid string `json:"openid,omitempty"`
}

// PromotionDetail 营销信息
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

// PromotionGoodsDetail 营销明细信息
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

// DeductNotifyRequest 预约扣费通知请求参数
type DeductNotifyRequest struct {
	// 字段含义：委托代扣协议 ID。
	ContractId string `json:"contract_id,omitempty"`
	// 字段含义：直连商户号。
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：应用 ID。
	Appid string `json:"appid,omitempty"`
	// 字段含义：预约扣费金额信息。
	EstimatedAmount DeductAmount `json:"estimated_amount,omitempty"`
}

// DeductAmount 预约扣费金额信息
type DeductAmount struct {
	// 字段含义：预约扣费金额，单位为分。
	Amount int64 `json:"amount,omitempty"`
	// 字段含义：预约扣费币种。
	Currency string `json:"currency,omitempty"`
}

// PayApplyRequest 扣款申请请求参数。
type PayApplyRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商户在抖音开放平台申请的应用 ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]，只能是数字、大小写字母、下划线、中划线、星号。
	// 业务规则：商户系统内部订单号，同一商户号下唯一。
	// 示例：OUT_1669813156。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 字段含义：交易结束时间。
	// 格式规则：string[1,64]，遵循 RFC3339 标准格式。
	// 业务规则：该笔订单允许的最晚用户付款时间，超过后用户无法支付；若需继续扣款，需使用新的商户订单号发起扣款；传递时间需在请求时间 15 天以内，超过 15 天时抖音支付会自动调整为请求时间后的第 15 天。
	// 示例：2022-12-01T20:59:16+08:00。
	TimeExpire string `json:"time_expire,omitempty"`
	// 字段含义：委托代扣协议ID。
	// 格式规则：string[1,64]。
	// 业务规则：签约成功后抖音支付返回的委托代扣协议 ID。
	// 示例：MSN20230314112037389849955326013。
	ContractId string `json:"contract_id,omitempty"`
	// 字段含义：交易类型。
	// 格式规则：string[1,3]。
	// 业务规则：SGP 表示商户代扣；NPP 表示免密支付。
	// 示例：SGP。
	TradeType string `json:"trade_type,omitempty"`
	// 字段含义：商品描述。
	// 格式规则：string[1,127]。
	// 业务规则：商品信息描述，会展示在用户抖音钱包账单的“商品说明”内，需传递能真实代表商品信息的描述。
	// 示例：测试商品。
	Description string `json:"description,omitempty"`
	// 字段含义：通知地址。
	// 格式规则：URL，string[1,256]，必须为 HTTPS 地址，不允许携带查询串。
	// 业务规则：抖音支付通过该地址通知扣款结果。
	// 示例：https://www.mock.douyinpay.com。
	NotifyUrl string `json:"notify_url,omitempty"`
	// 字段含义：附加数据。
	// 格式规则：string[1,1024]。
	// 业务规则：在查询 API 和支付通知中原样返回，可作为自定义参数使用，实际情况下只有支付完成状态才会返回该字段。
	// 示例：自定义数据。
	Attach string `json:"attach,omitempty"`
	// 字段含义：优惠标记。
	// 格式规则：string[1,512]，JSON 字符串。
	// 业务规则：和抖音支付协商后可用。
	// 示例：""。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 字段含义：优惠信息。
	// 格式规则：object。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：{"cost_price":608800,"invoice_id":"dy123"}。
	Detail Detail `json:"detail,omitempty"`
	// 字段含义：订单金额。
	// 格式规则：object。
	// 业务规则：订单金额信息。
	// 示例：{"currency":"CNY","total":100}。
	Amount *Amount `json:"amount,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 业务规则：支付场景描述，免密支付产品需要传。
	// 示例：{"payer_client_ip":"14.23.150.211","device_id":"13467007045764"}。
	SceneInfo SceneInfo `json:"scene_info,omitempty"`
	// 字段含义：结算信息。
	// 格式规则：object。
	// 业务规则：订单分账标识在扣款时设置。
	// 示例：{"profit_sharing":false}。
	SettleInfo SettleInfo `json:"settle_info,omitempty"`
}

// PayApplyResponse 申请扣款响应参数
type PayApplyResponse struct {
	// 字段含义：业务结果，以回调或查单返回的扣款结果为准。
	ResultCode string `json:"result_code,omitempty"`
	// 字段含义：预支付交易会话标识。
	PrepayId string `json:"prepay_id,omitempty"`
}

// PartnerContractScheduleRequest 服务商预约扣费请求参数
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

// PartnerContractScheduleResponse 服务商预约扣费响应参数
type PartnerContractScheduleResponse struct {
	// 可扣费开始日期
	DeductStartDate string `json:"deduct_start_date,omitempty"`
	// 可扣费结束日期
	DeductEndDate string `json:"deduct_end_date,omitempty"`
	// 预约扣费金额
	ScheduleAmount Amount `json:"schedule_amount,omitempty"`
}

// PartnerContractScheduleQueryRequest 服务商预约扣费结果查询请求参数
type PartnerContractScheduleQueryRequest struct {
	// 代扣协议id
	ContractId string `json:"contract_id,omitempty"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
}

// PartnerContractScheduleQueryResponse 服务商预约扣费结果查询响应参数
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
