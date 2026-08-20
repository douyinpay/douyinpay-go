package contractorder

type ApiCreateServiceOrderRequest struct {
	// 应用ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 服务描述
	ServiceIntroduction string `json:"service_introduction"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 服务风险金
	RiskFund *RiskFund `json:"risk_fund"`
	// 后付费项目
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 后付费商户优惠
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 服务时间段
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 服务位置
	Location *Location `json:"location,omitempty"`
	// 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 是否需要用户确认
	NeedUserConfirm bool `json:"need_user_confirm,omitempty"`
	// 用户标识
	OpenId string `json:"openid,omitempty"`
	// 附加数据
	Attach string `json:"attach,omitempty"`
	// 通知地址
	NotifyUrl string `json:"notify_url"`
	// 优惠标记
	GoodsTag string `json:"goods_tag,omitempty"`
	// 扩展参数
	ExtInfo string `json:"ext_info,omitempty"`
}

type ApiCreateServiceOrderResponse struct {
	// 应用ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 服务描述
	ServiceIntroduction string `json:"service_introduction"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 抖音支付服务订单号
	OrderId string `json:"order_id"`
	// 服务订单状态
	State string `json:"state"`
	// 服务订单状态描述
	StateDescription string `json:"state_description"`
	// 服务风险金
	RiskFund *RiskFund `json:"risk_fund"`
	// 后付费项目
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 后付费商户优惠
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 服务时间段
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 服务位置
	Location *Location `json:"location,omitempty"`
	// 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 是否需要用户确认
	NeedUserConfirm *bool `json:"need_user_confirm,omitempty"`
	// 先享后付申请token
	PayscoreApplyToken *string `json:"payscore_apply_token,omitempty"`
	// 优惠标记
	GoodsTag *string `json:"goods_tag,omitempty"`
	// 用户标识
	OpenId *string `json:"openid,omitempty"`
	// 附加数据
	Attach *string `json:"attach,omitempty"`
	// 通知地址 有效性：1. HTTPS；2. 不允许携带查询串。
	NotifyUrl string `json:"notify_url"`
}

// 收款信息
type Collection struct {
	// 收款状态
	State string `json:"state"`
	// 总收款金额
	TotalAmount int64 `json:"total_amount"`
	// 待收金额
	PayingAmount int64 `json:"paying_amount"`
	// 已收金额
	PaidAmount int64 `json:"paid_amount"`
	// 收款明细列表
	Details []*CollectionDetail `json:"details"`
}

type CollectionDetail struct {
	// 抖音支付交易单号
	TransactionId string `json:"transaction_id"`
	// 单笔收款金额
	Amount int64 `json:"amount"`
	// 收款成功渠道
	PaidType string `json:"paid_type"`
	// 收款成功时间
	PaidTime string `json:"paid_time"`
	// 收款银行
	BankType string `json:"bank_type"`
	// 优惠信息
	PromotionDetail []*PromotionDetail `json:"promotion_detail"`
}

type PromotionDetail struct {
	// 券ID
	CouponId string `json:"coupon_id"`
	// 优惠名称
	Name string `json:"name"`
	// 优惠范围
	Scope string `json:"scope"`
	// 优惠类型
	Type string `json:"type"`
	// 优惠券面额
	Amount int64 `json:"amount"`
	// 活动ID
	StockId string `json:"stock_id"`
	// 抖音支付出资
	DouyinpayContribute int64 `json:"douyinpay_contribute"`
	// 商户出资
	MerchantContribute int64 `json:"merchant_contribute"`
	// 其他出资
	OtherContribute int64 `json:"other_contribute"`
	// 优惠币种
	Currency string `json:"currency"`
	// 商品列表
	GoodsDetail []*GoodsDetail `json:"goods_detail"`
}

type GoodsDetail struct {
	// 商品编码
	GoodsId string `json:"goods_id"`
	// 商品数量
	Quantity int32 `json:"quantity"`
	// 商品价格
	UnitPrice int64 `json:"unit_price"`
	// 商品优惠金额
	DiscountAmount int64 `json:"discount_amount"`
	// 商品备注
	GoodsRemark string `json:"goods_remark"`
}

// RiskFund 服务风险金
type RiskFund struct {
	// 风险金名称
	Name string `json:"name"`
	// 风险金额
	Amount int64 `json:"amount,omitempty"`
	// 风险说明
	Description string `json:"description,omitempty"`
}

// PostItem 后付费信息
type PostItem struct {
	// 付费名称
	Name string `json:"name"`
	// 付费金额
	Amount int64 `json:"amount,omitempty"`
	// 付费说明
	Description string `json:"description,omitempty"`
	// 付费数量
	Count int64 `json:"count,omitempty"`
}

// TimeRange 服务时间范围
type TimeRange struct {
	// 服务开始时间
	StartTime string `json:"start_time,omitempty"`
	// 服务开始时间备注
	StartTimeRemark string `json:"start_time_remark,omitempty"`
	// 服务结束时间
	EndTime string `json:"end_time,omitempty"`
	// 服务结束时间备注
	EndTimeRemark string `json:"end_time_remark,omitempty"`
}

// Location 位置信息
type Location struct {
	// 服务开始地点
	StartLocation string `json:"start_location,omitempty"`
	// 服务结束地点
	EndLocation string `json:"end_location,omitempty"`
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	// 用户终端IP
	ClientIp string `json:"client_ip"`
	// 商户端设备号（预留字段）
	DeviceId string `json:"device_id,omitempty"`
	// 门店信息（预留字段）
	StoreInfo *StoreInfo `json:"store_info,omitempty"`
}

// 门店信息（预留字段）
type StoreInfo struct {
	// 门店编号
	Id string `json:"id"`
	// 门店名称
	Name string `json:"name"`
	// 地区编码
	AreaCode string `json:"area_code"`
	// 详细的商户门店地址
	Address string `json:"address"`
}

// ApiSynchronizeServiceOrderInfoRequest 同步服务订单信息请求参数
type ApiSynchronizeServiceOrderInfoRequest struct {
	// 应用ID
	Appid string `json:"appid"`
	// 商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 同步类型
	Type   string `json:"type"` // Order_Paid-订单已支付
	Detail struct {
		// 收款完成时间
		PaidTime string `json:"paid_time"`
	} `json:"detail"` // 同步详情
}

// ApiSynchronizeServiceOrderInfoResponse 同步服务订单信息响应参数
type ApiSynchronizeServiceOrderInfoResponse struct {
	// 应用ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 服务描述
	ServiceIntroduction string `json:"service_introduction"`
	// 交易金额
	TotalAmount int64 `json:"total_amount"`
	// 通知地址 有效性：1. HTTPS；2. 不允许携带查询串。
	NotifyUrl string `json:"notify_url"`
	// 附加数据
	Attach string `json:"attach"`
	// 优惠标记
	GoodsTag string `json:"goods_tag"`

	// 抖音支付服务订单号
	OrderId string `json:"order_id"`
	// 服务订单状态
	State string `json:"state"`
	// 服务订单状态描述
	StateDescription string `json:"state_description"`

	OpenId            string      `json:"openid"`
	AuthorizationCode string      `json:"authorization_code"`
	Collection        *Collection `json:"collection"`
	RiskFund          *RiskFund   `json:"risk_fund,omitempty"`
	PostPayments      []*PostItem `json:"post_payments,omitempty"`
	PostDiscounts     []*PostItem `json:"post_discounts,omitempty"`
	TimeRange         *TimeRange  `json:"time_range,omitempty"`
	Location          *Location   `json:"location,omitempty"`
	SceneInfo         *SceneInfo  `json:"scene_info,omitempty"`
}

// ApiQueryServiceOrderRequest 查询服务订单信息请求参数
type ApiQueryServiceOrderRequest struct {
	// 应用ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
}

// ApiQueryServiceOrderResponse 查询服务订单信息响应参数
type ApiQueryServiceOrderResponse struct {
	// 应用ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 商户协议号
	AuthorizationCode string `json:"authorization_code"`
	OpenId            string `json:"openid"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 服务描述
	ServiceIntroduction string `json:"service_introduction"`
	// 交易金额
	TotalAmount int64 `json:"total_amount"`
	// 通知地址 有效性：1. HTTPS；2. 不允许携带查询串。
	NotifyUrl string `json:"notify_url"`
	// 附加数据
	Attach string `json:"attach,omitempty"`
	// 抖音支付服务订单号
	OrderId string `json:"order_id"`
	// 服务订单状态
	State string `json:"state"`
	// 服务订单状态描述
	StateDescription string      `json:"state_description"`
	RiskFund         *RiskFund   `json:"risk_fund"`
	Collection       *Collection `json:"collection"`
	PostPayments     []*PostItem `json:"post_payments,omitempty"`
	PostDiscounts    []*PostItem `json:"post_discounts,omitempty"`
	TimeRange        *TimeRange  `json:"time_range,omitempty"`
	Location         *Location   `json:"location,omitempty"`
	SceneInfo        *SceneInfo  `json:"scene_info,omitempty"`
}

// ApiCancelServiceOrderRequest 查询服务订单信息请求参数
type ApiCancelServiceOrderRequest struct {
	// 应用ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 取消原因
	Reason string `json:"reason,omitempty"`
}

// ApiCancelServiceOrderResponse 取消服务订单信息响应参数
type ApiCancelServiceOrderResponse struct {
	// 应用ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 抖音支付服务订单号
	OrderId string `json:"order_id"`
}

// ApiModifyAmountRequest 服务订单改价请求参数
type ApiModifyAmountRequest struct {
	// 应用ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 交易金额
	TotalAmount   int64       `json:"total_amount"`
	PostPayments  []*PostItem `json:"post_payments,omitempty"`
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	Reason        string      `json:"reason,omitempty"`
}

// ApiModifyAmountResponse 服务订单改价响应参数
type ApiModifyAmountResponse struct {
	// 应用ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 抖音支付服务订单号
	OrderId string `json:"order_id"`
	// 服务描述
	ServiceIntroduction string `json:"service_introduction"`
	// 交易金额
	TotalAmount int64 `json:"total_amount"`
	// 通知地址 有效性：1. HTTPS；2. 不允许携带查询串。
	NotifyUrl string `json:"notify_url"`
	// 附加数据
	Attach string `json:"attach,omitempty"`
	// 服务订单状态
	State string `json:"state"`
	// 服务订单状态描述
	StateDescription string      `json:"state_description"`
	RiskFund         *RiskFund   `json:"risk_fund"`
	Collection       *Collection `json:"collection"`
	PostPayments     []*PostItem `json:"post_payments,omitempty"`
	PostDiscounts    []*PostItem `json:"post_discounts,omitempty"`
	TimeRange        *TimeRange  `json:"time_range,omitempty"`
	Location         *Location   `json:"location,omitempty"`
}

// ApiCompleteServiceOrderRequest 完结服务订单请求参数
type ApiCompleteServiceOrderRequest struct {
	// 应用ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 交易金额
	TotalAmount int64 `json:"total_amount"`
	// 附加数据
	Attach string `json:"attach,omitempty"`
	// 优惠标记
	GoodsTag string `json:"goods_tag"`

	PostPayments  []*PostItem `json:"post_payments,omitempty"`
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	TimeRange     *TimeRange  `json:"time_range,omitempty"`
	Location      *Location   `json:"location,omitempty"`
	SceneInfo     *SceneInfo  `json:"scene_info,omitempty"`
}

// ApiCompleteServiceOrderResponse 完结服务订单响应参数
type ApiCompleteServiceOrderResponse struct {
	// 应用ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 抖音支付服务订单号
	OrderId string `json:"order_id"`
	// 服务描述
	ServiceIntroduction string `json:"service_introduction"`
	// 交易金额
	TotalAmount int64 `json:"total_amount"`
	// 附加数据
	Attach string `json:"attach,omitempty"`
	// 服务订单状态
	State string `json:"state"`
	// 优惠标记
	GoodsTag string `json:"goods_tag"`
	// 服务订单状态描述
	StateDescription string      `json:"state_description"`
	RiskFund         *RiskFund   `json:"risk_fund"`
	PostPayments     []*PostItem `json:"post_payments,omitempty"`
	PostDiscounts    []*PostItem `json:"post_discounts,omitempty"`
	TimeRange        *TimeRange  `json:"time_range,omitempty"`
	Location         *Location   `json:"location,omitempty"`
	SceneInfo        *SceneInfo  `json:"scene_info,omitempty"`
}

// ApiCloseCreditServiceRequest 解除用户授权关系请求参数
type ApiCloseCreditServiceRequest struct {
	// 应用ID
	Appid string `json:"appid"`
	// 直连商户号
	Mchid string `json:"mchid"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 协议号
	AuthorizationCode string `json:"authorization_code"`
	// 解约原因
	Reason string `json:"reason"`
}

// ApiCloseCreditServiceResponse 解除用户授权关系响应参数
type ApiCloseCreditServiceResponse struct {
}

type ApiServiceOrderPayRequest struct {
	// 应用ID
	Appid string `json:"appid"`
	// 商户号
	Mchid string `json:"mchid"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
}

type ApiServiceOrderPayResponse struct {
	// 应用ID
	Appid string `json:"appid"`
	// 商户号
	Mchid string `json:"mchid"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 商户订单号
	OutOrderNo string `json:"out_order_no"`
	// 抖音支付服务订单号
	OrderId string `json:"order_id"`
}

type ApiCreditSrvSignApplyRequest struct {
	// 应用ID
	Appid string `json:"appid"`
	// 商户号
	Mchid string `json:"mchid"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 商户协议号
	AuthorizationCode string `json:"authorization_code"`
	// 通知地址
	NotifyUrl string `json:"notify_url"`
	// 优惠标记
	GoodsTag string `json:"goods_tag,omitempty"`
	// 扩展参数
	ExtInfo string `json:"ext_info,omitempty"`
}

type ApiCreditSrvSignApplyResponse struct {
	PayscoreApplyToken string `json:"payscore_apply_token"`
}

type ApiCreditSrvSignQueryRequest struct {
	// 商户号
	Mchid string `json:"mchid"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 商户协议号
	AuthorizationCode string `json:"authorization_code"`
}

type ApiCreditSrvSignQueryResponse struct {
	// 应用ID
	Appid string `json:"appid"`
	// 商户号
	Mchid string `json:"mchid"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 用户标识
	OpenId string `json:"openid"`
	// 商户协议号
	AuthorizationCode string `json:"authorization_code"`
	// 授权状态
	AuthorizationState string `json:"authorization_state"`
	// 解除授权时间
	CancelAuthorizationTime string `json:"cancel_authorization_time"`
	// 授权成功时间
	AuthorizationSuccessTime string `json:"authorization_success_time"`
}

// ApiPartnerCreateServiceOrderRequest 服务商创建服务订单请求参数。
type ApiPartnerCreateServiceOrderRequest struct {
	// 服务商应用ID
	SpAppid string `json:"sp_appid,omitempty"`
	// 服务商商户号
	SpMchid string `json:"sp_mchid,omitempty"`
	// 子商户应用ID
	SubAppid string `json:"sub_appid,omitempty"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid,omitempty"`
	// 商户服务单号
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 服务ID
	ServiceId string `json:"service_id,omitempty"`
	// 服务信息
	ServiceIntroduction string `json:"service_introduction,omitempty"`
	// 商户协议号
	AuthorizationCode string `json:"authorization_code,omitempty"`
	// 商户数据包
	Attach string `json:"attach,omitempty"`
	// 通知地址
	NotifyUrl string `json:"notify_url,omitempty"`
	// 优惠标记
	GoodsTag string `json:"goods_tag,omitempty"`
	// 服务风险金
	RiskFund *RiskFund `json:"risk_fund,omitempty"`
	// 后付费项目
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 商户优惠
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 实际服务时间段
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 服务位置
	Location *Location `json:"location,omitempty"`
	// 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 是否需要用户确认
	NeedUserConfirm bool `json:"need_user_confirm,omitempty"`
	// 扩展参数
	ExtInfo string `json:"ext_info,omitempty"`
}

// ApiPartnerCreateServiceOrderResponse 服务商创建服务订单响应参数。
type ApiPartnerCreateServiceOrderResponse struct {
	// 服务商应用ID
	SpAppid string `json:"sp_appid"`
	// 服务商商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户应用ID
	SubAppid string `json:"sub_appid"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid"`
	// 商户服务单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 服务信息
	ServiceIntroduction string `json:"service_introduction"`
	// 服务风险金
	RiskFund *RiskFund `json:"risk_fund,omitempty"`
	// 抖音支付服务订单号
	OrderId string `json:"order_id"`
	// 服务订单状态
	State string `json:"state"`
	// 订单状态说明
	StateDescription string `json:"state_description"`
	// 服务商商户下用户标识
	SpOpenid string `json:"sp_openid"`
	// 子商户下用户标识
	SubOpenid string `json:"sub_openid"`
	// 商户协议号
	AuthorizationCode string `json:"authorization_code"`
	// 商户数据包
	Attach string `json:"attach"`
	// 通知地址
	NotifyUrl string `json:"notify_url"`
	// 后付费项目
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 商户优惠
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 实际服务时间段
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 服务位置
	Location *Location `json:"location,omitempty"`
	// 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 是否需要用户确认
	NeedUserConfirm bool `json:"need_user_confirm,omitempty"`
	// 优惠标记
	GoodsTag string `json:"goods_tag,omitempty"`
	// 跳转抖音支付token
	PayscoreApplyToken string `json:"payscore_apply_token,omitempty"`
}

// ApiPartnerCompleteServiceOrderRequest 服务商完结服务订单请求参数。
type ApiPartnerCompleteServiceOrderRequest struct {
	// 服务商应用ID
	SpAppid string `json:"sp_appid,omitempty"`
	// 服务商商户号
	SpMchid string `json:"sp_mchid,omitempty"`
	// 子商户应用ID
	SubAppid string `json:"sub_appid,omitempty"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid,omitempty"`
	// 商户服务单号
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 服务ID
	ServiceId string `json:"service_id,omitempty"`
	// 总金额，单位为分
	TotalAmount int64 `json:"total_amount,omitempty"`
	// 商户数据包
	Attach string `json:"attach,omitempty"`
	// 优惠标记
	GoodsTag string `json:"goods_tag,omitempty"`
	// 支付渠道信息
	ChannelInfo *ChannelInfo `json:"channel_info,omitempty"`
	// 后付费项目
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 商户优惠
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 实际服务时间段
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 服务位置
	Location *Location `json:"location,omitempty"`
	// 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
}

// ApiPartnerCompleteServiceOrderResponse 服务商完结服务订单响应参数。
type ApiPartnerCompleteServiceOrderResponse struct {
	// 服务商应用ID
	SpAppid string `json:"sp_appid"`
	// 服务商商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户应用ID
	SubAppid string `json:"sub_appid"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid"`
	// 商户服务单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 服务信息
	ServiceIntroduction string `json:"service_introduction"`
	// 抖音支付服务订单号
	OrderId string `json:"order_id"`
	// 服务订单状态
	State string `json:"state"`
	// 订单状态说明
	StateDescription string `json:"state_description"`
	// 服务风险金
	RiskFund *RiskFund `json:"risk_fund,omitempty"`
	// 订单总金额
	TotalAmount int64 `json:"total_amount"`
	// 后付费项目
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 商户优惠
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 实际服务时间段
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 服务位置
	Location *Location `json:"location,omitempty"`
	// 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 优惠标记
	GoodsTag string `json:"goods_tag,omitempty"`
	// 商户数据包
	Attach string `json:"attach,omitempty"`
}

// ChannelInfo 支付渠道信息
type ChannelInfo struct {
	// 指定渠道列表
	PresetChannel []*PresetChannel `json:"preset_channel,omitempty"`
}

// PresetChannel 指定渠道列表
type PresetChannel struct {
	// 指定渠道
	ChannelCode string `json:"channel_code,omitempty"`
	// 指定渠道ID
	ChannelId string `json:"channel_id,omitempty"`
	// 指定渠道金额
	ChannelAmount int64 `json:"channel_amount,omitempty"`
	// 扩展信息
	ChannelExtInfo string `json:"channel_ext_info,omitempty"`
}

// ApiPartnerQueryServiceOrderRequest 服务商查询服务订单请求参数。
type ApiPartnerQueryServiceOrderRequest struct {
	// 服务商应用ID
	SpAppid string `json:"sp_appid,omitempty"`
	// 服务商商户号
	SpMchid string `json:"sp_mchid,omitempty"`
	// 子商户应用ID
	SubAppid string `json:"sub_appid,omitempty"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid,omitempty"`
	// 商户服务单号
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 服务ID
	ServiceId string `json:"service_id,omitempty"`
}

// ApiPartnerQueryServiceOrderResponse 服务商查询服务订单响应参数。
type ApiPartnerQueryServiceOrderResponse struct {
	// 服务商应用ID
	SpAppid string `json:"sp_appid"`
	// 服务商商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户应用ID
	SubAppid string `json:"sub_appid"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid"`
	// 商户服务单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 服务信息
	ServiceIntroduction string `json:"service_introduction"`
	// 抖音支付服务订单号
	OrderId string `json:"order_id"`
	// 服务订单状态
	State string `json:"state"`
	// 订单状态说明
	StateDescription string `json:"state_description"`
	// 服务风险金
	RiskFund *RiskFund `json:"risk_fund,omitempty"`
	// 订单总金额
	TotalAmount int64 `json:"total_amount"`
	// 服务商商户下用户标识
	SpOpenid string `json:"sp_openid"`
	// 子商户下用户标识
	SubOpenid string `json:"sub_openid"`
	// 商户协议号
	AuthorizationCode string `json:"authorization_code"`
	// 商户数据包
	Attach string `json:"attach"`
	// 通知地址
	NotifyUrl string `json:"notify_url"`
	// 后付费项目
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 后付费商户优惠
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 收款信息
	Collection *Collection `json:"collection,omitempty"`
	// 服务时间范围
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 服务位置
	Location *Location `json:"location,omitempty"`
	// 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
}

// ApiPartnerCancelServiceOrderRequest 服务商取消服务订单请求参数。
type ApiPartnerCancelServiceOrderRequest struct {
	// 服务商应用ID
	SpAppid string `json:"sp_appid,omitempty"`
	// 服务商商户号
	SpMchid string `json:"sp_mchid,omitempty"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid,omitempty"`
	// 商户服务单号
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 服务ID
	ServiceId string `json:"service_id,omitempty"`
	// 取消原因
	Reason string `json:"reason,omitempty"`
}

// ApiPartnerCancelServiceOrderResponse 服务商取消服务订单响应参数。
type ApiPartnerCancelServiceOrderResponse struct {
	// 服务订单状态
	State string `json:"state"`
	// 抖音支付服务订单号
	OrderId string `json:"order_id"`
	// 订单状态说明
	StateDescription string `json:"state_description"`
}

// ApiPartnerSynchronizeServiceOrderInfoRequest 服务商同步服务订单信息请求参数。
type ApiPartnerSynchronizeServiceOrderInfoRequest struct {
	// 服务商应用ID
	SpAppid string `json:"sp_appid,omitempty"`
	// 服务商商户号
	SpMchid string `json:"sp_mchid,omitempty"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid,omitempty"`
	// 商户服务单号
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 服务ID
	ServiceId string `json:"service_id,omitempty"`
	// 同步类型
	Type string `json:"type,omitempty"`
	// 同步内容信息详情
	Detail struct {
		// 收款完成时间
		PaidTime string `json:"paid_time"`
	} `json:"detail"`
}

// ApiPartnerSynchronizeServiceOrderInfoResponse 服务商同步服务订单信息响应参数。
type ApiPartnerSynchronizeServiceOrderInfoResponse struct{}

// ApiPartnerModifyAmountRequest 服务商修改订单金额请求参数。
type ApiPartnerModifyAmountRequest struct {
	// 服务商应用ID
	SpAppid string `json:"sp_appid,omitempty"`
	// 服务商商户号
	SpMchid string `json:"sp_mchid,omitempty"`
	// 子商户应用ID
	SubAppid string `json:"sub_appid,omitempty"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid,omitempty"`
	// 商户服务单号
	OutOrderNo string `json:"out_order_no,omitempty"`
	// 服务ID
	ServiceId string `json:"service_id,omitempty"`
	// 总金额
	TotalAmount int64 `json:"total_amount,omitempty"`
	// 后付费项目
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 商户优惠
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 修改原因
	Reason string `json:"reason,omitempty"`
}

// ApiPartnerModifyAmountResponse 服务商修改订单金额响应参数。
type ApiPartnerModifyAmountResponse struct {
	// 服务商应用ID
	SpAppid string `json:"sp_appid"`
	// 服务商商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户应用ID
	SubAppid string `json:"sub_appid"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid"`
	// 商户服务单号
	OutOrderNo string `json:"out_order_no"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 服务信息
	ServiceIntroduction string `json:"service_introduction"`
	// 抖音支付服务订单号
	OrderId string `json:"order_id"`
	// 服务订单状态
	State string `json:"state"`
	// 服务订单状态描述
	StateDescription string `json:"state_description"`
	// 服务风险金
	RiskFund *RiskFund `json:"risk_fund,omitempty"`
	// 总金额
	TotalAmount int64 `json:"total_amount"`
	// 实际服务时间段
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 服务位置
	Location *Location `json:"location,omitempty"`
	// 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 商户数据包
	Attach string `json:"attach"`
	// 通知地址
	NotifyUrl string `json:"notify_url"`
	// 后付费项目
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 商户优惠
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 收款信息
	Collection *Collection `json:"collection,omitempty"`
}

// ApiPartnerCreditSrvSignApplyRequest 服务商申请先享后付授权请求参数。
type ApiPartnerCreditSrvSignApplyRequest struct {
	// 服务商应用ID
	SpAppid string `json:"sp_appid"`
	// 服务商商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户应用ID。
	SubAppid string `json:"sub_appid,omitempty"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 商户协议号
	AuthorizationCode string `json:"authorization_code"`
	// 通知地址
	NotifyUrl string `json:"notify_url,omitempty"`
	// 商户数据包
	Attach string `json:"attach"`
	// 优惠标记
	GoodsTag string `json:"goods_tag,omitempty"`
	// 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 扩展参数
	ExtInfo string `json:"ext_info,omitempty"`
}

// ApiPartnerCreditSrvSignApplyResponse 服务商申请先享后付授权响应参数。
type ApiPartnerCreditSrvSignApplyResponse struct {
	// 先享后付申请token
	PayscoreApplyToken string `json:"payscore_apply_token"`
}

// ApiPartnerCreditSrvSignQueryRequest 服务商查询用户授权记录请求参数。
type ApiPartnerCreditSrvSignQueryRequest struct {
	// 服务商商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 商户协议号
	AuthorizationCode string `json:"authorization_code"`
}

// ApiPartnerCreditSrvSignQueryResponse 服务商查询用户授权记录响应参数。
type ApiPartnerCreditSrvSignQueryResponse struct {
	// 服务商商户号
	SpMchid string `json:"sp_mchid"`
	// 服务商应用ID
	SpAppid string `json:"sp_appid"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid"`
	// 子商户应用ID
	SubAppid string `json:"sub_appid"`
	// 商户协议号
	AuthorizationCode string `json:"authorization_code"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 服务商商户下用户标识
	SpOpenid string `json:"sp_openid"`
	// 子商户下用户标识
	SubOpenid string `json:"sub_openid"`
	// 授权状态
	AuthorizationState string `json:"authorization_state"`
	// 解除授权时间
	CancelAuthorizationTime string `json:"cancel_authorization_time"`
	// 授权成功时间
	AuthorizationSuccessTime string `json:"authorization_success_time"`
}

// ApiPartnerCloseCreditServiceRequest 服务商解除用户授权关系请求参数。
type ApiPartnerCloseCreditServiceRequest struct {
	// 服务商应用ID
	SpAppid string `json:"sp_appid"`
	// 服务商商户号
	SpMchid string `json:"sp_mchid"`
	// 子商户应用ID
	SubAppid string `json:"sub_appid,omitempty"`
	// 子商户商户号
	SubMchid string `json:"sub_mchid"`
	// 服务ID
	ServiceId string `json:"service_id"`
	// 商户协议号
	AuthorizationCode string `json:"authorization_code"`
	// 解约原因
	Reason string `json:"reason,omitempty"`
}

// ApiPartnerCloseCreditServiceResponse 服务商解除用户授权关系响应参数。
type ApiPartnerCloseCreditServiceResponse struct{}
