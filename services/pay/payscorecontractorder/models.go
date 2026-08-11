package payscorecontractorder

type PrepayRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商品描述。
	// 格式规则：string[1,127]。
	// 业务规则：商品信息描述，会展示在用户抖音钱包账单的"商品说明"内。商户需遵循"商品描述"参数传递建议，传递能真实代表商品信息的描述，不能超过127个字符。
	// 示例：测试商品。
	Description string `json:"description"`
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]，只能是数字、大小写字母_-*且在同一个商户号下唯一。
	// 业务规则：商户系统内部订单号。当同一业务订单需要进行多次支付时，需为每个支付请求生成独立的流水号，不能直接将业务订单号作为商户订单号使用。若订单已创建，商户订单号相同且其他参数有修改将报错。
	// 示例：OUT_1666688488。
	OutTradeNo string `json:"out_trade_no"`
	// 字段含义：交易结束时间。
	// 格式规则：string[1,64]，遵循RFC 3339标准格式：yyyy-MM-DDTHH:mm:ss+TIMEZONE。
	// 业务规则：用户能够完成该笔订单支付的最后时限，并非订单关闭的时间。若用户实际支付时间超过此时间，需使用新的商户订单号下单。传递的交易结束时间需在下单时间的15天以内，如超过15天会自动调整为下单时间后的第15天。
	// 示例：2018-06-08T10:34:56+08:00。
	TimeExpire string `json:"time_expire,omitempty"`
	// 字段含义：附加数据。
	// 格式规则：string[1,1024]。
	// 业务规则：在查询API和支付通知中原样返回，可作为自定义参数使用，实际情况下只有支付完成状态才会返回该字段。
	// 示例：自定义数据。
	Attach string `json:"attach,omitempty"`
	// 字段含义：通知地址。
	// 格式规则：URL，string[1,256]，必须为直接可访问的HTTPS地址，不允许携带查询串。
	// 业务规则：接收抖音支付异步通知回调的URL。
	// 示例：https://www.mock.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：交易类型。
	// 格式规则：string[1,16]。
	// 业务规则：取值：APP（APP支付）、JSAPI（JSAPI支付/小程序支付）、MWEB（H5支付）。
	// 示例：APP。
	TradeType string `json:"trade_type"`
	// 字段含义：用户标识。
	// 格式规则：string[1,128]。
	// 业务规则：trade_type=JSAPI时此参数必传，用户在商户appid下的唯一标识。
	// 示例：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o。
	Openid string `json:"openid,omitempty"`
	// 字段含义：优惠标记。
	// 格式规则：string[1,512]，JSON格式。
	// 业务规则：和抖音支付协商后可用。传参说明：业务场景区分可通过传入key值=biz_scene，value值为约定场景值；个性化策略区分可通过传入key值=product_tag，value值为约定参数值；指定优惠信息区分可通过传入key值=assign_discounts，value值为"抖音支付优惠查询接口"返回的"指定优惠信息"字段值。
	// 示例：{"biz_scene":"","product_tag":"","assign_discounts":""}。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 字段含义：订单金额。
	// 格式规则：object。
	// 业务规则：订单金额信息。
	// 示例：{"total":100,"currency":"CNY"}。
	Amount *Amount `json:"amount"`
	// 字段含义：优惠信息。
	// 格式规则：object。
	// 业务规则：为预留字段，商户不需要传。
	Detail *Detail `json:"detail,omitempty"`
	// 字段含义：签约信息。
	// 格式规则：object。
	// 业务规则：支付分签约并支付使用的签约参数。
	// 示例：{"contract_mchid":"6020230307605084","contract_appid":"awofz9bncda6w2w4","service_id":"001","out_contract_code":"100001256","contract_notify_url":"https://yoursite.com"}。
	ContractInfo *ContractInfo `json:"contract_info,omitempty"`
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
	// 字段含义：电子发票入口开放标识。
	// 格式规则：boolean。
	// 业务规则：为预留字段，商户不需要传。
	SupportFapiao *bool `json:"support_fapiao,omitempty"`
}

// PrepayResponse
type PrepayResponse struct {
	// 字段含义：预支付交易会话标识。
	// trade_type为APP和JSAPI时返回，有效期为2小时，失效后需要重新请求该接口以获取新的prepay_id。
	PrepayId string `json:"prepay_id"`
	// 字段含义：支付跳转链接。
	// trade_type为MWEB时返回，有效期为5分钟，超过有效期需重新请求下单接口以获取新的h5_url。
	H5Url string `json:"h5_url,omitempty"`
}

// CloseOrderRequest
type CloseOrderRequest struct {
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]，只能是数字、大小写字母_-*且在同一个商户号下唯一。
	// 业务规则：商户系统内部订单号，同一商户号下唯一。
	// 示例：OUT_1666688488。
	OutTradeNo string `json:"out_trade_no"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
}

// CloseRequest
type CloseRequest struct {
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
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
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
}

// QueryOrderByOutTradeNoRequest
type QueryOrderByOutTradeNoRequest struct {
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]，只能是数字、大小写字母_-*且在同一个商户号下唯一。
	// 业务规则：商户系统内部订单号，同一商户号下唯一。
	// 示例：OUT_1666688488。
	OutTradeNo string `json:"out_trade_no"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
}

// Amount
type Amount struct {
	// 字段含义：订单总金额，单位为分。
	// 格式规则：int。
	// 业务规则：订单金额信息中的总金额。
	// 示例：100。
	Total int64 `json:"total"`
	// 字段含义：货币种类。
	// 格式规则：string[1,16]。
	// 业务规则：CNY：人民币，境内商户号仅支持人民币。
	// 示例：CNY。
	Currency string `json:"currency,omitempty"`
}

// Detail 优惠功能（预留字段）
type Detail struct {
	// 字段含义：订单原价。
	// 格式规则：int，单位为分。
	// 业务规则：为预留字段，商户不需要传。
	CostPrice int64 `json:"cost_price,omitempty"`
	// 字段含义：商品小票ID。
	// 格式规则：string[1,32]。
	// 业务规则：为预留字段，商户不需要传。
	InvoiceId string `json:"invoice_id,omitempty"`
	// 字段含义：单品列表。
	// 格式规则：array。
	// 业务规则：为预留字段，商户不需要传。
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// GoodsDetail（预留字段）
type GoodsDetail struct {
	// 字段含义：商户侧商品编码。
	// 格式规则：string[1,32]。
	// 业务规则：预留字段，商户不需要传。
	MerchantGoodsId string `json:"merchant_goods_id"`
	// 字段含义：抖音支付商品编码。
	// 格式规则：string[1,32]。
	// 业务规则：预留字段，商户不需要传。抖音支付定义的统一商品编码，没有可不传。
	DouyinpayGoodsId string `json:"douyinpay_goods_id,omitempty"`
	// 字段含义：商品名称。
	// 格式规则：string[1,256]。
	// 业务规则：预留字段，商户不需要传。商品的实际名称。
	GoodsName string `json:"goods_name,omitempty"`
	// 字段含义：商品数量。
	// 格式规则：int。
	// 业务规则：预留字段，商户不需要传。用户购买的数量。
	Quantity int64 `json:"quantity"`
	// 字段含义：商品单价，单位为分。
	// 格式规则：int。
	// 业务规则：预留字段，商户不需要传。如果商户有优惠，需传输商户优惠后的单价。
	UnitPrice int64 `json:"unit_price"`
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	// 字段含义：用户终端IP。
	// 格式规则：string[1,45]，支持IPv4和IPv6两种格式。
	// 业务规则：用户的客户端IP地址。
	// 示例：14.23.150.211。
	PayerClientIp string `json:"payer_client_ip"`
	// 字段含义：商户端设备号。
	// 格式规则：string[1,32]。
	// 业务规则：商户端设备号（门店号或收银设备ID）。
	// 示例：13467007045764。
	DeviceId string `json:"device_id,omitempty"`
	// 字段含义：用户终端设备号。
	// 格式规则：string[1,45]。
	// 业务规则：安卓：优先传android_id（openudid），若没有，传IMEI；IOS：优先IDFV，若没有IDFV，传IDFA。
	PayerDeviceId string `json:"payer_device_id,omitempty"`
}

// SettleInfo
type SettleInfo struct {
	// 字段含义：是否分账。
	// 格式规则：boolean。
	// 业务规则：传入true表示在订单支付成功后可进行分账操作。需分账时订单收款成功后资金将被冻结并转入基本账户的不可用余额，可通过请求分账API分配给其他商户或用户；完成分账后可通过完结分账接口解冻剩余资金，或在支付成功30天后自动解冻。传入false或不传（默认false）时订单收款成功后资金直接转入基本账户的可用余额。
	// 示例：false。
	ProfitSharing bool `json:"profit_sharing,omitempty"`
}

type ContractInfo struct {
	// 字段含义：签约商户号。
	// 格式规则：string[1,32]。
	// 业务规则：签约商户号，必须与mchid一致。
	// 示例：6020230307605084。
	ContractMchId string `json:"contract_mchid"`
	// 字段含义：签约应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音开放平台生成的应用ID，全局唯一，必须与appid一致。
	// 示例：awofz9bncda6w2w4。
	ContractAppId string `json:"contract_appid"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，配置商户和场景维度信息。
	// 示例：001。
	ServiceId *string `json:"service_id"`
	// 字段含义：签约协议号。
	// 格式规则：string[1,64]，只能是数字、大小写字母。
	// 业务规则：商户侧的签约协议号，由商户生成。
	// 示例：100001256。
	OutContractCode string `json:"out_contract_code"`
	// 字段含义：请求序列号。
	// 格式规则：int64。
	// 业务规则：暂不开放。商户请求签约时的序列号，要求唯一性。
	RequestSerial int64 `json:"request_serial,omitempty"`
	// 字段含义：用户账户展示名称。
	// 格式规则：string[1,64]。
	// 业务规则：暂不开放。签约用户的名称，用于页面展示。
	ContractDisplayAccount string `json:"contract_display_account,omitempty"`
	// 字段含义：签约结果通知地址。
	// 格式规则：URL，string[1,256]，必须以https开头，不能携带参数。
	// 业务规则：签约信息回调通知的URL，必须为外网可访问的URL。
	// 示例：https://yoursite.com。
	ContractNotifyUrl string `json:"contract_notify_url"`
	// 字段含义：签约拓展业务参数。
	// 格式规则：string[1,512]，JSON格式。
	// 业务规则：签约拓展参数，仅与抖音支付线下约定后使用。
	ContractExt string `json:"contract_ext,omitempty"`
}
