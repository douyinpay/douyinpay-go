package app

// PrepayRequest
type PrepayRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商户在抖音开放平台申请的移动应用 AppID，全局唯一，并确保该 AppID 与 mchid 有绑定关系。
	// 示例：awz9w2wncdof4ba6。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6000000000000001。
	Mchid string `json:"mchid"`
	// 字段含义：商品描述。
	// 格式规则：string[1,127]。
	// 业务规则：商品信息描述，会展示在用户抖音钱包账单的“商品说明”内，需真实代表商品信息。
	// 示例：测试商品。
	Description string `json:"description"`
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]，只能是数字、大小写字母、下划线、中划线、星号。
	// 业务规则：商户系统内部订单号，同一商户号下唯一；同一业务订单多次支付时需生成不同商户订单号。
	// 示例：OUT_1666688488。
	OutTradeNo string `json:"out_trade_no"`
	// 字段含义：交易结束时间。
	// 格式规则：遵循 RFC3339 标准格式。
	// 业务规则：用户可完成该笔订单支付的最后时限，超过后用户无法支付；传递时间需在下单时间 15 天以内。
	// 示例：2018-06-08T10:34:56+08:00。
	TimeExpire string `json:"time_expire,omitempty"`
	// 字段含义：附加数据。
	// 格式规则：string[1,1024]。
	// 业务规则：在查询 API 和支付通知中原样返回，可作为自定义参数使用，实际情况下只有支付完成状态才会返回该字段。
	// 示例：自定义数据。
	Attach string `json:"attach,omitempty"`
	// 字段含义：通知地址。
	// 格式规则：URL，string[1,256]，必须为 HTTPS 地址，不允许携带查询串。
	// 业务规则：抖音支付通过该地址通知支付结果。
	// 示例：https://www.mock.douyinpay.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：优惠标记。
	// 格式规则：string[1,512]，JSON 字符串。
	// 业务规则：和抖音支付协商后可用，可用于业务场景区分、个性化策略区分或指定优惠信息区分。
	// 示例：{"biz_scene":"","product_tag":"","assign_discounts":""}。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 字段含义：电子发票入口开放标识。
	// 格式规则：boolean。
	// 业务规则：为预留字段，传入 true 时支付成功消息和支付详情页将出现开票入口。
	SupportFapiao bool `json:"support_fapiao,omitempty"`
	// 字段含义：订单金额。
	// 格式规则：object。
	// 业务规则：订单金额信息。
	// 示例：{"total":100,"currency":"CNY"}。
	Amount *Amount `json:"amount"`
	// 字段含义：优惠信息。
	// 格式规则：object。
	// 业务规则：为预留字段，商户不需要传。
	Detail *Detail `json:"detail,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 业务规则：支付场景描述。
	// 示例：{"payer_client_ip":"14.23.150.211","device_id":"13467007045764"}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 字段含义：结算信息。
	// 格式规则：object。
	// 业务规则：设置订单是否支持分账。
	// 示例：{"profit_sharing":false}。
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
}

// PrepayResponse
type PrepayResponse struct {
	// 预支付交易会话标识
	PrepayId string `json:"prepay_id"`
}

// CloseOrderRequest
type CloseOrderRequest struct {
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]，只能是数字、大小写字母、下划线、中划线、星号。
	// 业务规则：商户系统内部订单号，同一商户号下唯一。
	// 示例：OUT_1666688488。
	OutTradeNo string `json:"out_trade_no"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6000000000000001。
	Mchid string `json:"mchid"`
}

// CloseRequest
type CloseRequest struct {
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6000000000000001。
	Mchid string `json:"mchid"`
}

// QueryOrderByIdRequest
type QueryOrderByIdRequest struct {
	// 字段含义：抖音支付订单号。
	// 格式规则：string。
	// 业务规则：抖音支付系统生成的订单号。
	// 示例：21000125010103000993845301326。
	TransactionId string `json:"transaction_id"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6000000000000001。
	Mchid string `json:"mchid"`
}

// QueryOrderByOutTradeNoRequest
type QueryOrderByOutTradeNoRequest struct {
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]，只能是数字、大小写字母、下划线、中划线、星号。
	// 业务规则：商户系统内部订单号，同一商户号下唯一。
	// 示例：OUT_1666688488。
	OutTradeNo string `json:"out_trade_no"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6000000000000001。
	Mchid string `json:"mchid"`
}

// Amount
type Amount struct {
	// 字段含义：总金额。
	// 格式规则：int，单位为分。
	// 业务规则：订单总金额。
	// 示例：100。
	Total int64 `json:"total"`
	// 字段含义：货币种类。
	// 格式规则：string[1,16]。
	// 业务规则：CNY 表示人民币，境内商户号仅支持人民币。
	// 示例：CNY。
	Currency string `json:"currency,omitempty"`
}

// Detail 优惠功能
type Detail struct {
	// 字段含义：订单原价。
	// 格式规则：int，单位为分。
	CostPrice int64 `json:"cost_price,omitempty"`
	// 字段含义：商品小票ID。
	// 格式规则：string[1,32]。
	InvoiceId string `json:"invoice_id,omitempty"`
	// 字段含义：单品列表。
	// 格式规则：array，至少传入 1 条。
	// 业务规则：为预留字段，商户不需要传；订单商品明细列表。
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// GoodsDetail
type GoodsDetail struct {
	// 字段含义：商户侧商品编码。
	// 格式规则：string[1,32]，由半角大小写字母、数字、中划线、下划线中的一种或几种组成。
	// 业务规则：预留字段。
	MerchantGoodsId string `json:"merchant_goods_id"`
	// 字段含义：抖音支付商品编码。
	// 格式规则：string[1,32]。
	// 业务规则：预留字段，抖音支付定义的统一商品编码，没有可不传。
	DouyinpayGoodsId string `json:"douyinpay_goods_id,omitempty"`
	// 字段含义：商品名称。
	// 格式规则：string[1,256]。
	// 业务规则：预留字段，商品的实际名称。
	GoodsName string `json:"goods_name,omitempty"`
	// 字段含义：商品数量。
	// 格式规则：int。
	// 业务规则：预留字段，用户购买的数量。
	Quantity int64 `json:"quantity"`
	// 字段含义：商品单价。
	// 格式规则：int，单位为分。
	// 业务规则：预留字段；如商户有优惠，需传输商户优惠后的单价。
	UnitPrice int64 `json:"unit_price"`
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	// 字段含义：用户终端IP。
	// 格式规则：string[1,45]，支持 IPv4 和 IPv6。
	// 业务规则：用户的客户端 IP。
	// 示例：14.23.150.211。
	PayerClientIp string `json:"payer_client_ip"`
	// 字段含义：商户端设备号。
	// 格式规则：string[1,32]。
	// 业务规则：预留字段，商户不需要传；可表示门店号或收银设备 ID。
	// 示例：13467007045764。
	DeviceId string `json:"device_id,omitempty"`
	// 字段含义：商户门店信息。
	// 格式规则：object。
	// 业务规则：预留字段，商户不需要传。
	StoreInfo *StoreInfo `json:"store_info,omitempty"`
}

// SettleInfo
type SettleInfo struct {
	// 字段含义：是否分账。
	// 格式规则：boolean。
	// 业务规则：传入 true 表示订单支付成功后可进行分账；不传或传 false 表示不需要分账。
	// 示例：false。
	ProfitSharing bool `json:"profit_sharing,omitempty"`
}

// StoreInfo 商户门店信息（预留字段）
type StoreInfo struct {
	// 字段含义：门店编号。
	// 格式规则：string[1,32]。
	// 业务规则：商户侧门店编号。
	Id string `json:"id"`
	// 字段含义：门店名称。
	// 格式规则：string[1,256]。
	// 业务规则：商户侧门店名称。
	Name string `json:"name,omitempty"`
	// 字段含义：地区编码。
	// 格式规则：string[1,32]。
	// 业务规则：详细请见抖音支付提供的地区编码文档。
	AreaCode string `json:"area_code,omitempty"`
	// 字段含义：详细地址。
	// 格式规则：string[1,512]。
	// 业务规则：详细的商户门店地址。
	Address string `json:"address,omitempty"`
}
