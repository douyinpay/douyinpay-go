package partnerh5

type PrepayRequest struct {
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商商户号，由抖音支付生成并下发。
	// 示例：6020230301343000。
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：子商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605000。
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：服务商应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：服务商在抖音开放平台申请的应用 ID，全局唯一，需确保该 sp_appid 与 sp_mchid 有绑定关系。
	// 示例：awofz9bncda6w000。
	SpAppid string `json:"sp_appid,omitempty"`
	// 字段含义：子商户应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：子商户在抖音开放平台申请的应用 ID，全局唯一，需确保该 sub_appid 与 sub_mchid 有绑定关系。
	// 示例：awofz9bncda6x000。
	SubAppid string `json:"sub_appid,omitempty"`
	// 字段含义：商品描述。
	// 格式规则：string[1,127]。
	// 业务规则：商品信息描述，会展示在用户抖音钱包账单的"商品说明"内，需传递能真实代表商品信息的描述。
	// 示例：测试商品。
	Description string `json:"description,omitempty"`
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]，只能是数字、大小写字母、下划线、中划线、星号。
	// 业务规则：服务商系统内部订单号，同一服务商商户号下唯一。
	// 示例：OUT_1666688488。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 字段含义：交易结束时间。
	// 格式规则：string[1,64]，遵循 RFC3339 标准格式。
	// 业务规则：用户可完成该笔订单支付的最后时限，并非订单关闭的时间；超过后用户无法支付，若需继续支付需使用新的商户订单号下单；传递时间需在请求时间 15 天以内，超过 15 天时抖音支付会自动调整为请求时间后的第 15 天。
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
	NotifyUrl string `json:"notify_url,omitempty"`
	// 字段含义：优惠标记。
	// 格式规则：string[1,512]，JSON 字符串。
	// 业务规则：和抖音支付协商后可用；可通过 biz_scene 区分业务场景、product_tag 指定个性化策略、assign_discounts 指定优惠信息。
	// 示例：{"biz_scene":"","product_tag":"","assign_discounts":""}。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 字段含义：电子发票入口开放标识。
	// 格式规则：boolean。
	// 业务规则：传入true时，支付成功消息和支付详情页将出现开票入口。需要在抖音支付商户平台或抖音公众平台开通电子发票功能，传此字段才可生效。
	SupportFapiao bool `json:"support_fapiao,omitempty"`
	// 字段含义：订单金额。
	// 格式规则：object。
	// 业务规则：订单金额信息。
	// 示例：{"total":100,"currency":"CNY"}。
	Amount *Amount `json:"amount,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 业务规则：支付场景描述。
	// 示例：{"payer_client_ip":"14.23.150.211","device_id":"13467007045764"}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 字段含义：优惠信息。
	// 格式规则：object。
	// 业务规则：为预留字段，服务商不需要传。
	Detail *Detail `json:"detail,omitempty"`
	// 字段含义：结算信息。
	// 格式规则：object。
	// 业务规则：订单分账标识在下单时设置；传入 true 表示订单支付成功后可进行分账，传入 false 或不传时资金不会冻结，直接转入基本账户可用余额。
	// 示例：{"profit_sharing":false}。
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
}

type Amount struct {
	// 字段含义：订单总金额，单位为分。
	// 格式规则：int64。
	// 业务规则：订单金额信息中的总金额。
	// 示例：100。
	Total int64 `json:"total,omitempty"`
	// 字段含义：货币类型。
	// 格式规则：string[1,16]。
	// 业务规则：CNY 表示人民币，境内商户号仅支持人民币。
	// 示例：CNY。
	Currency string `json:"currency,omitempty"`
}

type SceneInfo struct {
	// 字段含义：用户终端IP。
	// 格式规则：string[1,45]，支持 IPv4 和 IPv6 两种格式。
	// 业务规则：用户的客户端 IP。
	// 示例：14.23.150.211。
	PayerClientIp string `json:"payer_client_ip,omitempty"`
	// 字段含义：商户端设备号。
	// 格式规则：string[1,32]。
	// 业务规则：为预留字段，商户不需要传。
	// 示例：13467007045764。
	DeviceId string `json:"device_id,omitempty"`
	// 字段含义：用户终端设备号。
	// 格式规则：string[1,32]。
	// 业务规则：安卓优先传 android_id（openudid），若没有则传 IMEI；iOS 优先传 IDFV，若没有 IDFV 则传 IDFA。
	// 示例：a0e4b456-c9e5-3783-a001。
	PayerDeviceId string `json:"payer_device_id,omitempty"`
	// 字段含义：商户门店信息。
	// 格式规则：object。
	// 业务规则：为预留字段，商户不需要传。
	StoreInfo *StoreInfo `json:"store_info,omitempty"`
	// 字段含义：H5场景信息。
	// 格式规则：object。
	// 业务规则：H5 支付场景信息。
	// 示例：{"type":"Wap","app_name":"抖音","app_url":"https://douyinpay.com/"}。
	H5Info *H5Info `json:"h5_info,omitempty"`
}

type StoreInfo struct {
	// 字段含义：商户侧门店编号。
	// 格式规则：string[1,32]。
	// 业务规则：为预留字段，商户不需要传。
	Id string `json:"id,omitempty"`
	// 字段含义：商户侧门店名称。
	// 格式规则：string[1,256]。
	// 业务规则：为预留字段，商户不需要传。
	Name string `json:"name,omitempty"`
	// 字段含义：地区编码。
	// 格式规则：string[1,32]。
	// 业务规则：为预留字段，商户不需要传。
	AreaCode string `json:"area_code,omitempty"`
	// 字段含义：详细的商户门店地址。
	// 格式规则：string[1,512]。
	// 业务规则：为预留字段，商户不需要传。
	Address string `json:"address,omitempty"`
}

type H5Info struct {
	// 字段含义：场景类型。
	// 格式规则：string[1,32]。
	// 业务规则：取值如下：iOS、Android、Wap、HarmonyOS。
	// 示例：Wap。
	Type string `json:"type"`
	// 字段含义：应用名称。
	// 格式规则：string[1,64]。
	// 业务规则：H5 场景对应的应用名称。
	// 示例：抖音。
	AppName string `json:"app_name,omitempty"`
	// 字段含义：网站URL。
	// 格式规则：URL，string[1,128]，必须为 HTTPS 地址，不允许携带查询串。
	// 业务规则：H5 场景对应的网站 URL。
	// 示例：https://douyinpay.com/。
	AppUrl string `json:"app_url,omitempty"`
	// 字段含义：iOS平台BundleID。
	// 格式规则：string[1,128]。
	// 业务规则：iOS 平台 BundleID。
	// 示例：com.test.testiOS。
	BundleID string `json:"bundle_id,omitempty"`
	// 字段含义：Android平台PackageName。
	// 格式规则：string[1,128]。
	// 业务规则：Android 平台 PackageName。
	// 示例：com.test.testAndroid。
	PackageName string `json:"package_name,omitempty"`
}

type SettleInfo struct {
	// 字段含义：是否分账。
	// 格式规则：boolean。
	// 业务规则：传入 true 表示订单支付成功后可进行分账；传入 false 或不传时，订单收款成功后资金不会冻结，直接转入基本账户可用余额。
	// 示例：false。
	ProfitSharing bool `json:"profit_sharing,omitempty"`
}

type PrepayResponse struct {
	// 字段含义：H5支付跳转链接。
	H5_url string `json:"h5_url,omitempty"`
}

// CloseOrderRequest
type CloseOrderRequest struct {
	// 字段含义：商户订单号。
	// 格式规则：string。
	// 业务规则：商户系统内部订单号，同一子商户号下唯一。
	OutTradeNo string `json:"out_trade_no"`
	// 字段含义：服务商商户号。
	// 格式规则：string。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户号。
	// 格式规则：string。
	// 业务规则：服务商模式下的子商户号。
	SubMchid string `json:"sub_mchid"`
}

// CloseRequest
type CloseRequest struct {
	// 字段含义：服务商商户号。
	// 格式规则：string。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户号。
	// 格式规则：string。
	// 业务规则：服务商模式下的子商户号。
	SubMchid string `json:"sub_mchid"`
}

// QueryOrderByIdRequest
type QueryOrderByIdRequest struct {
	// 字段含义：抖音支付订单号。
	// 格式规则：string。
	// 业务规则：抖音支付系统生成的订单号。
	TransactionId string `json:"transaction_id"`
	// 字段含义：服务商商户号。
	// 格式规则：string。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户号。
	// 格式规则：string。
	// 业务规则：服务商模式下的子商户号。
	SubMchid string `json:"sub_mchid"`
}

// QueryOrderByOutTradeNoRequest
type QueryOrderByOutTradeNoRequest struct {
	// 字段含义：商户订单号。
	// 格式规则：string。
	// 业务规则：商户系统内部订单号，同一子商户号下唯一。
	OutTradeNo string `json:"out_trade_no"`
	// 字段含义：服务商商户号。
	// 格式规则：string。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户号。
	// 格式规则：string。
	// 业务规则：服务商模式下的子商户号。
	SubMchid string `json:"sub_mchid"`
}

// Detail 优惠信息（预留字段）
type Detail struct {
	CostPrice   int64          `json:"cost_price,omitempty"`   // 订单原价
	InvoiceId   string         `json:"invoice_id,omitempty"`   // 发票ID
	GoodsDetail []*GoodsDetail `json:"goods_detail,omitempty"` // 单品列表
}

// GoodsDetail 单品信息（预留字段）
type GoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id,omitempty"`  // 商户侧商品编码
	DouyinpayGoodsId string `json:"douyinpay_goods_id,omitempty"` // 抖音侧商品编码
	GoodsName        string `json:"goods_name,omitempty"`         // 商品名称
	Quantity         int64  `json:"quantity,omitempty"`           // 商品数量
	UnitPrice        int64  `json:"unit_price,omitempty"`         // 商品单价
}
