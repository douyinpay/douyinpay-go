package contract

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
	// 活动ID
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

type PartnerContractOrderRequest struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：服务商在抖音开放平台申请的应用 ID，全局唯一，需确保该 sp_appid 与 sp_mchid 有绑定关系。
	// 示例：awofz9bncda6w200。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商商户号，由抖音支付生成并下发。
	// 示例：6020230301343900。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：子商户在抖音开放平台申请的应用 AppID，全局唯一；APP 场景下需使用应用属性为移动应用的 AppID，JSAPI/H5 场景下需使用应用属性为网站应用的 AppID，并确保该 sub_appid 与 sub_mchid 有绑定关系；若 sub_openid 有传则 sub_appid 必填且需与 sub_openid 对应。
	// 示例：awofz9bncda6x700。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605000。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：商品描述。
	// 格式规则：string[1,127]。
	// 业务规则：商品信息描述，会展示在用户抖音钱包账单的"商品说明"内，需传递能真实代表商品信息的描述。
	// 示例：测试商品。
	Description string `json:"description,omitempty"`
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]，只能是数字、大小写字母、下划线、中划线、星号。
	// 业务规则：商户系统内部订单号，同一服务商商户号下唯一。
	// 示例：OUT_1666688488。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 字段含义：交易结束时间。
	// 格式规则：string[1,64]，遵循 RFC3339 标准格式。
	// 业务规则：用户可完成该笔订单支付的最后时限，并非订单关闭的时间；超过后用户无法支付，若需继续支付需使用新的商户订单号下单；传递时间需在请求时间 15 天以内，超过 15 天时抖音支付会自动调整为请求时间后的第 15 天。
	// 示例：2018-06-08T10:34:56+08:00。
	TimeExpire string `json:"time_expire,omitempty"`
	// 字段含义：交易类型。
	// 格式规则：string[1,16]。
	// 业务规则：取值如下：JSAPI、APP、MWEB。
	// 示例：APP。
	TradeType string `json:"trade_type,omitempty"`
	// 字段含义：附加数据。
	// 格式规则：string[1,1024]。
	// 业务规则：在查询 API 和支付通知中原样返回，可作为自定义参数使用，实际情况下只有支付完成状态才会返回该字段。
	// 示例：自定义数据。
	Attach string `json:"attach,omitempty"`
	// 字段含义：通知地址。
	// 格式规则：URL，string[1,256]，必须为 HTTPS 地址，不允许携带查询串。
	// 业务规则：抖音支付通过该地址通知支付结果。
	// 示例：https://www.mock.douyinpay.com。
	NotifyUrl string `json:"notify_url,omitempty"`
	// 字段含义：优惠标记。
	// 格式规则：string[1,512]，JSON 字符串。
	// 业务规则：和抖音支付协商后可用；可通过 biz_scene 区分业务场景、product_tag 指定个性化策略、assign_discounts 指定优惠信息。
	// 示例：{"biz_scene":"","product_tag":"","assign_discounts":""}。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 字段含义：电子发票入口开放标识。
	// 格式规则：boolean。
	// 业务规则：为预留字段，服务商不需要传。
	SupportFapiao bool `json:"support_fapiao,omitempty"`
	// 字段含义：订单金额。
	// 格式规则：object。
	// 业务规则：订单金额信息。
	// 示例：{"currency":"CNY","total":100}。
	Amount *Amount `json:"amount"`
	// 字段含义：支付者。
	// 格式规则：object。
	// 业务规则：支付者信息；JSAPI 场景下 sp_openid、sub_openid 至少填一个。
	// 示例：{"sp_openid":"897ae8bd9f194107-9cb3-85f5672037de","sub_openid":"823ae8bd9f893402-9cb3-85f8794657ea"}。
	Payer *Payer `json:"payer"`
	// 字段含义：优惠信息。
	// 格式规则：object。
	// 业务规则：为预留字段，服务商不需要传。
	Detail *Detail `json:"detail,omitempty"`
	// 商品列表
	//GoodsDetail []GoodsDetail `json:"goods_detail"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 业务规则：支付场景描述。
	// 示例：{"payer_client_ip":"14.23.150.211","device_id":"13467007045764"}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 商户门店信息（预留字段）
	//StoreInfo *StoreInfo `json:"store_info"`
	// 字段含义：结算信息。
	// 格式规则：object。
	// 业务规则：订单分账标识在下单时设置；传入 true 表示订单支付成功后可进行分账，传入 false 或不传时资金不会冻结，直接转入基本账户可用余额。
	// 示例：{"profit_sharing":false}。
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
	// 字段含义：签约信息。
	// 格式规则：object。
	// 业务规则：服务商签约并支付下单使用的签约参数。
	// 示例：{"contract_mchid":"6020230307605084","contract_appid":"awofz9bncda6x7w8","plan_id":"48","out_contract_code":"100001258","request_serial":1,"contract_display_account":"测试账号","contract_notify_url":"https://www.mock.com"}。
	ContractInfo *ContractInfo `json:"contract_info,omitempty"`
}

// ContractPayer 支付者
type Payer struct {
	// 字段含义：用户服务标识。
	// 格式规则：string[1,128]。
	// 业务规则：用户在服务商 AppID 下的唯一标识；JSAPI 场景下 sp_openid、sub_openid 至少填一个。
	// 示例：897ae8bd9f194107-9cb3-85f5672037de。
	SpOpenid string `json:"sp_openid"`
	// 字段含义：用户子标识。
	// 格式规则：string[1,128]。
	// 业务规则：用户在子商户 AppID 下的唯一标识；若传 sub_openid，则 sub_appid 必填。
	// 示例：823ae8bd9f893402-9cb3-85f8794657ea。
	SubOpenid string `json:"sub_openid"`
}

// Amount
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
	// 业务规则：为预留字段，服务商不需要传。
	CostPrice int64 `json:"cost_price"`
	// 字段含义：商品小票ID。
	// 格式规则：string[1,32]。
	// 业务规则：为预留字段，服务商不需要传。
	InvoiceId string `json:"invoice_id,omitempty"`
	// 字段含义：单品列表。
	// 格式规则：array。
	// 业务规则：为预留字段，服务商不需要传。
	GoodsDetail []GoodsDetail `json:"goods_detail"`
}

type PayApplyDetail struct {
	// 字段含义：订单原价。
	// 格式规则：int64。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：608800。
	CostPrice int64 `json:"cost_price"`
	// 字段含义：商品小票ID。
	// 格式规则：string[1,32]。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：dy123。
	InvoiceId string `json:"invoice_id,omitempty"`
	// 字段含义：单品列表。
	// 格式规则：array。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：[{"goods_name":"iPhoneX 256G","merchant_goods_id":"ABC","quantity":1,"unit_price":828800,"douyinpay_goods_id":"1001"}]。
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// GoodsDetail
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
	// 示例：1。
	Quantity int64 `json:"quantity,omitempty"`
	// 字段含义：商品单价，单位为分。
	// 格式规则：int64。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：828800。
	UnitPrice int64 `json:"unit_price,omitempty"`
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	// 字段含义：用户终端IP。
	// 格式规则：string[1,45]，支持 IPv4 和 IPv6 两种格式。
	// 业务规则：用户的客户端 IP。
	// 示例：14.23.150.211。
	PayerClientIp string `json:"payer_client_ip,omitempty"`
	// 字段含义：商户端设备号。
	// 格式规则：string[1,32]。
	// 业务规则：商户端设备号，可传门店号或收银设备 ID。
	// 示例：13467007045764。
	DeviceId string `json:"device_id,omitempty"`
	// 字段含义：用户终端设备号。
	// 格式规则：string[1,45]。
	// 业务规则：安卓优先传 android_id（openudid），若没有则传 IMEI；iOS 优先传 IDFV，若没有 IDFV 则传 IDFA。
	// 示例：a0e4b456-c9e5-3783-a422。
	PayerDeviceId string `json:"payer_device_id,omitempty"`
	// 字段含义：商户门店信息。
	StoreInfo *StoreInfo `json:"store_info,omitempty"`
}

// SettleInfo
type SettleInfo struct {
	// 字段含义：是否指定分账。
	// 格式规则：boolean。
	// 业务规则：传入 true 表示订单支付成功后可进行分账；传入 false 或不传时，订单收款成功后资金不会冻结，直接转入基本账户可用余额。
	// 示例：false。
	ProfitSharing bool `json:"profit_sharing,omitempty"`
}

type ContractInfo struct {
	// 字段含义：签约商户号。
	// 格式规则：string[1,32]。
	// 业务规则：签约商户号，必须与 sub_mchid 一致。
	// 示例：6020230307605084。
	ContractMchId string `json:"contract_mchid,omitempty"`
	// 字段含义：签约应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用 ID，全局唯一，必须与 sub_appid 一致。
	// 示例：awofz9bncda6x7w8。
	ContractAppId string `json:"contract_appid,omitempty"`
	// 字段含义：模板ID。
	// 格式规则：string[1,32]。
	// 业务规则：模板 ID，联系抖音支付运营申请。
	// 示例：48。
	PlanId string `json:"plan_id,omitempty"`
	// 字段含义：商户侧签约协议号。
	// 格式规则：string[1,64]，只能是数字、大小写字母的描述。
	// 业务规则：商户侧的签约协议号，由商户生成。
	// 示例：100001258。
	OutContractCode string `json:"out_contract_code,omitempty"`
	// 字段含义：请求序列号。
	// 格式规则：int64，禁止使用 0 开头，范围不能超过 int64 的最大值（9223372036854775807）。
	// 业务规则：商户请求签约时的序列号，要求唯一性，主要用于排序，不作为查询条件。
	// 示例：1。
	RequestSerial int64 `json:"request_serial,omitempty"`
	// 字段含义：用户账户展示名称。
	// 格式规则：string[1,64]，不支持 UTF8 非 3 字节编码的字符（如表情符号）。
	// 业务规则：签约用户的名称，用于页面展示。
	// 示例：测试账号。
	ContractDisplayAccount string `json:"contract_display_account,omitempty"`
	// 字段含义：签约信息通知地址。
	// 格式规则：URL，string[1,256]，必须以 https 开头，不允许携带参数。
	// 业务规则：签约信息回调通知的 URL，必须为外网可访问的 URL。
	// 示例：https://www.mock.com。
	ContractNotifyUrl string `json:"contract_notify_url,omitempty"`
	// 字段含义：签约拓展业务参数。
	// 格式规则：string[1,512]，JSON 格式。
	// 业务规则：仅与抖音支付线下约定后使用。
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
	// 字段含义：服务商应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：服务商在抖音开放平台申请的应用 ID，全局唯一。
	// 示例：awofz9bncda6w2w3。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605083。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：特约商户应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：特约商户在抖音开放平台申请的应用 ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：特约商户号。
	// 格式规则：string[1,32]。
	// 业务规则：特约商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]，只能是数字、大小写字母、下划线、中划线、星号。
	// 业务规则：商户系统内部订单号，同一商户号下唯一。
	// 示例：OUT_1669813156。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 字段含义：交易结束时间。
	// 格式规则：string[1,64]，遵循 RFC3339 标准格式。
	// 业务规则：该笔订单允许的最晚用户付款时间，超过后用户无法支付；若需继续扣款，服务商需使用新的商户订单号发起扣款；传递时间需在请求时间 15 天以内，超过 15 天时抖音支付会自动调整为请求时间后的第 15 天。
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
	Detail *PayApplyDetail `json:"detail,omitempty"`
	// 商品列表
	// GoodsDetail []PayApplyDetail `json:"goods_detail,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 业务规则：支付场景描述，免密支付产品需要传。
	// 示例：{"payer_client_ip":"14.23.150.211","device_id":"13467007045764"}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 字段含义：结算信息。
	// 格式规则：object。
	// 业务规则：订单分账标识在扣款时设置。
	// 示例：{"profit_sharing":false}。
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
	// 字段含义：订单金额。
	// 格式规则：object。
	// 业务规则：订单金额信息。
	// 示例：{"currency":"CNY","total":100}。
	Amount *Amount `json:"amount"`
}

type PartnerPayApplyResponse struct {
	// 字段含义：业务结果，以回调或查单返回的扣款结果为准。
	ResultCode string `json:"result_code"`
}
