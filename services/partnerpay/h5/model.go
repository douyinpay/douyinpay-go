package partnerh5

type PrepayRequest struct {
	SpMchid       string      `json:"sp_mchid,omitempty"`       // 商户号
	SubMchid      string      `json:"sub_mchid,omitempty"`      // 子商户号
	SpAppid       string      `json:"sp_appid,omitempty"`       // 服务商Appid
	SubAppid      string      `json:"sub_appid,omitempty"`      // 子商户Appid
	Description   string      `json:"description,omitempty"`    // 商品描述
	OutTradeNo    string      `json:"out_trade_no,omitempty"`   // 商户订单号
	TimeExpire    string      `json:"time_expire,omitempty"`    // 交易结束时间
	Attach        string      `json:"attach,omitempty"`         // 附加数据
	NotifyUrl     string      `json:"notify_url,omitempty"`     // 通知地址
	GoodsTag      string      `json:"goods_tag,omitempty"`      // 优惠标记
	SupportFapiao bool        `json:"support_fapiao,omitempty"` // 电子发票入口开放标识
	Amount        *Amount     `json:"amount,omitempty"`         // 订单金额
	Detail        *Detail     `json:"detail,omitempty"`         // 优惠信息
	SceneInfo     *SceneInfo  `json:"scene_info,omitempty"`     // 场景信息
	SettleInfo    *SettleInfo `json:"settle_info,omitempty"`    // 结算信息
}

type Amount struct {
	Total    int64  `json:"total,omitempty"`    // 订单总金额
	Currency string `json:"currency,omitempty"` // 货币类型
}

type Detail struct {
	CostPrice   int64          `json:"cost_price,omitempty"`   // 订单原价
	InvoiceId   string         `json:"invoice_id,omitempty"`   // 发票ID
	GoodsDetail []*GoodsDetail `json:"goods_detail,omitempty"` // 单品列表
}

type GoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id,omitempty"`  // 商户侧商品编码
	DouyinpayGoodsId string `json:"douyinpay_goods_id,omitempty"` // 抖音侧商品编码
	GoodsName        string `json:"goods_name,omitempty"`         // 商品名称
	Quantity         int64  `json:"quantity,omitempty"`           // 商品数量
	UnitPrice        int64  `json:"unit_price,omitempty"`         // 商品单价
}

type SceneInfo struct {
	PayerClientIp string     `json:"payer_client_ip,omitempty"` //  用户的客户端IP
	DeviceId      string     `json:"device_id,omitempty"`       //  商户端设备号
	PayerDeviceId string     `json:"payer_device_id,omitempty"`
	StoreInfo     *StoreInfo `json:"store_info,omitempty"` //  门店信息
	H5Info        *H5Info    `json:"h5_info,omitempty"`
}

type StoreInfo struct {
	Id       string `json:"id,omitempty"`        //  门店编号
	Name     string `json:"name,omitempty"`      //  门店名称
	AreaCode string `json:"area_code,omitempty"` //  地区编码
	Address  string `json:"address,omitempty"`   //  详细地址
}

type H5Info struct {
	Type        string `json:"type,omitempty"`         // 场景类型
	AppName     string `json:"app_name,omitempty"`     //  门店名称
	AppUrl      string `json:"app_url,omitempty"`      // 网站url
	BundleID    string `json:"bundle_id,omitempty"`    // iOS 平台 BundleID
	PackageName string `json:"package_name,omitempty"` // Android平台PackageName
}

type SettleInfo struct {
	ProfitSharing bool `json:"profit_sharing,omitempty"` // 是否开启分账
}

type PrepayResponse struct {
	H5_url string `json:"h5_url,omitempty"` // 为H5拉起支付的中间页面
}

// CloseOrderRequest
type CloseOrderRequest struct {
	OutTradeNo string `json:"out_trade_no"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
}

// CloseRequest
type CloseRequest struct {
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
}

// QueryOrderByIdRequest
type QueryOrderByIdRequest struct {
	//支付订单号
	TransactionId string `json:"transaction_id"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
}

// QueryOrderByOutTradeNoRequest
type QueryOrderByOutTradeNoRequest struct {
	//支付订单号
	OutTradeNo string `json:"out_trade_no"`
	// 服务商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户号
	SubMchid string `json:"sub_mchid"`
}
