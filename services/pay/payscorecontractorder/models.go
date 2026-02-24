package payscorecontractorder

type PrepayRequest struct {
	// 公众号ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
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
	//交易类型
	TradeType string `json:"trade_type"`
	//用户标识
	Openid string `json:"openid,omitempty"`
	// 商品标记，代金券或立减优惠功能的参数。
	GoodsTag     string        `json:"goods_tag,omitempty"`
	Amount       *Amount       `json:"amount"`
	Detail       *Detail       `json:"detail,omitempty"`
	ContractInfo *ContractInfo `json:"contract_info,omitempty"`
	SceneInfo    *SceneInfo    `json:"scene_info,omitempty"`
	SettleInfo   *SettleInfo   `json:"settle_info,omitempty"`
	//是否支持发票
	SupportFapiao *bool `json:"support_fapiao,omitempty"`
}

// PrepayResponse
type PrepayResponse struct {
	// 预支付交易会话标识
	PrepayId string `json:"prepay_id"`
}

// CloseOrderRequest
type CloseOrderRequest struct {
	OutTradeNo string `json:"out_trade_no"`
	Mchid      string `json:"mchid"`
}

// CloseRequest
type CloseRequest struct {
	Mchid string `json:"mchid"`
}

// QueryOrderByIdRequest
type QueryOrderByIdRequest struct {
	TransactionId string `json:"transaction_id"`
	Mchid         string `json:"mchid"`
}

// QueryOrderByOutTradeNoRequest
type QueryOrderByOutTradeNoRequest struct {
	OutTradeNo string `json:"out_trade_no"`
	Mchid      string `json:"mchid"`
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
	DeviceId string `json:"device_id,omitempty"`
	// 用户终端设备号
	PayerDeviceId string `json:"payer_device_id,omitempty"`
}
type H5Info struct {
	// 场景类型 示例值：iOS, Android, Wap
	Type string `json:"type"`
	// 应用名称
	AppName string `json:"app_name"`
	// 网站URL
	AppUrl string `json:"app_url"`
	// iOS平台BundleID
	BundleID string `json:"bundle_id"`
	// Android平台PackageName
	PackageName string `json:"package_name"`
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

type ContractInfo struct {
	//签约商户号
	ContractMchId string `json:"contract_mchid"`
	//签约appid
	ContractAppId string `json:"contract_appid"`
	//签约协议号
	OutContractCode string `json:"out_contract_code"`
	//请求序列号
	RequestSerial int64 `json:"request_serial,omitempty"`
	//用户账户展示名称
	ContractDisplayAccount string `json:"contract_display_account,omitempty"`
	//签约信息通知url
	ContractNotifyUrl string `json:"contract_notify_url"`
	//签约扩展业务参数
	ContractExt string `json:"contract_ext,omitempty"`
	//服务ID
	ServiceId *string `json:"service_id"`
}
