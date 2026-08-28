package contractorder

type ApiCreateServiceOrderRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一。此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：服务信息。
	// 格式规则：string[1,20]。
	// 业务规则：用于介绍本订单所提供的服务。
	// 示例：某某酒店。
	ServiceIntroduction string `json:"service_introduction"`
	// 字段含义：商户服务订单号。
	// 格式规则：string[1,32]，只能是数字、大小写字母、下划线、中划线、星号。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务风险金。
	// 格式规则：object。
	// 示例：{"name":"ESTIMATE_ORDER_COST","amount":10000,"description":"预估订单费用"}。
	RiskFund *RiskFund `json:"risk_fund"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 业务规则：后付费项目列表，最多包含100条付费项目。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：后付费商户优惠。
	// 格式规则：array。
	// 业务规则：商户优惠列表，最多包含30条商户优惠。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 业务规则：支付场景描述。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 字段含义：是否需要用户确认。
	// 格式规则：boolean。
	// 业务规则：【需确认模式】必须传true；【免确认模式】可不传或传false，若service_id不支持免确认，则报错返回。
	// 示例：true。
	NeedUserConfirm bool `json:"need_user_confirm,omitempty"`
	// 字段含义：用户标识。
	// 格式规则：string[1,128]。
	// 业务规则：用户在直连商户appid下的唯一标识。
	// 示例：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o。
	OpenId string `json:"openid,omitempty"`
	// 字段含义：附加数据。
	// 格式规则：string[1,256]。
	// 业务规则：商户自定义数据。
	// 示例：自定义数据。
	Attach string `json:"attach,omitempty"`
	// 字段含义：商户接收回调通知的地址。
	// 格式规则：string[1,256]。
	// 业务规则：必须为https地址。请确保回调URL是外部可正常访问的，且不能携带后缀参数，否则可能导致商户无法接收到抖音的回调通知信息。
	// 示例：https://www.bytedance.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：优惠标记。
	// 格式规则：string[1,512]，json格式。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{"product_tag":"xxxx","biz_scene":"aaaa","assign_discounts":"11091104_1442"}。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 字段含义：扩展参数。
	// 格式规则：string[1,2048]，JSON格式。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{}。
	ExtInfo string `json:"ext_info,omitempty"`
}

type ApiCreateServiceOrderResponse struct {
	// 字段含义：应用ID。
	// 格式规则：string。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一。此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：服务描述。
	// 格式规则：string。
	// 业务规则：服务描述，用于介绍本订单所提供的服务。
	// 示例：某某酒店。
	ServiceIntroduction string `json:"service_introduction"`
	// 字段含义：商户订单号。
	// 格式规则：string。
	// 业务规则：商户系统内部服务订单号。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：抖音支付服务订单号。
	// 格式规则：string。
	// 业务规则：每个抖音支付服务订单号与商户号下对应的商户服务订单号一一对应。
	// 示例：10050100220208060900000023310957。
	OrderId string `json:"order_id"`
	// 字段含义：服务订单状态。
	// 格式规则：string。
	// 业务规则：枚举值：CREATED: 商户已创建服务订单；DOING: 服务订单进行中；DONE: 服务订单完成；REVOKED: 商户取消服务订单；EXPIRED: 服务订单已失效，“商户已创建服务订单”状态超过7天未变动，则订单失效。
	// 示例：DOING。
	State string `json:"state"`
	// 字段含义：服务订单状态描述。
	// 格式规则：string。
	// 业务规则：对服务订单“进行中”状态的附加说明：MCH_COMPLETE: 商户完结；USER_PAYING: 用户支付中；USER_CONFIRM: 用户已确认。
	// 示例：USER_PAYING。
	StateDescription string `json:"state_description"`
	// 字段含义：服务风险金。
	// 格式规则：object。
	// 示例：{"name":"ESTIMATE_ORDER_COST","amount":10000,"description":"预估订单费用"}。
	RiskFund *RiskFund `json:"risk_fund"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 业务规则：后付费项目列表，最多包含100条付费项目。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：后付费商户优惠。
	// 格式规则：array。
	// 业务规则：商户优惠列表，最多包含30条商户优惠。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 业务规则：支付场景描述。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 字段含义：是否需要用户确认。
	// 格式规则：boolean。
	// 业务规则：【需确认模式】必须传true；【免确认模式】可不传或传false，若service_id不支持免确认，则报错返回。
	// 示例：true。
	NeedUserConfirm *bool `json:"need_user_confirm,omitempty"`
	// 字段含义：先享后付申请token。
	// 格式规则：string。
	// 业务规则：用于拉起抖音支付APP的数据，需确认模式下有值。
	// 示例：5778aadY9nlt1234XixCkFIGYnV2V。
	PayscoreApplyToken *string `json:"payscore_apply_token,omitempty"`
	// 字段含义：优惠标记。
	// 格式规则：string。
	// 业务规则：抖音支付优惠标记，json格式，和抖音支付协商后可用。
	// 示例：{}。
	GoodsTag *string `json:"goods_tag,omitempty"`
	// 字段含义：用户标识。
	// 格式规则：string。
	// 业务规则：用户在直连商户appid下的唯一标识。
	// 示例：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o。
	OpenId *string `json:"openid,omitempty"`
	// 字段含义：附加数据。
	// 格式规则：string。
	// 示例：自定义数据。
	Attach *string `json:"attach,omitempty"`
	// 字段含义：通知地址
	// 格式规则：string。
	// 业务规则：商户接收回调通知的地址 有效性：1. HTTPS；2. 不允许携带查询串。
	// 示例：https://www.bytedance.com。
	NotifyUrl string `json:"notify_url"`
}

// 收款信息
type Collection struct {
	// 字段含义：收款状态。
	// 格式规则：string。
	// 业务规则：枚举值：USER_PAYING：待支付；USER_PAID：已支付。
	// 示例：USER_PAYING。
	State string `json:"state"`
	// 字段含义：总收款金额。
	// 格式规则：int64。
	// 业务规则：单位分。
	// 示例：100。
	TotalAmount int64 `json:"total_amount"`
	// 字段含义：待收金额。
	// 格式规则：int64。
	// 业务规则：单位为分。
	// 示例：100。
	PayingAmount int64 `json:"paying_amount"`
	// 字段含义：已收金额。
	// 格式规则：int64。
	// 业务规则：单位为分。
	// 示例：100。
	PaidAmount int64 `json:"paid_amount"`
	// 字段含义：收款明细列表。
	// 格式规则：array。
	// 示例：[{"transaction_id":"TP2022101314262940644982204453","amount":100,"paid_type":"JSAPI","paid_time":"20220208060910"}]。
	Details []*CollectionDetail `json:"details"`
}

type CollectionDetail struct {
	// 字段含义：抖音支付交易单号。
	// 格式规则：string。
	// 业务规则：抖音支付交易单号，等于普通支付接口中的transaction_id，可以使用该订单号进行查询订单、申请退款操作。只有单据状态为USER_PAID，且收款成功渠道为抖音支付渠道, 收款金额大于0，才会返回该交易单号。
	// 示例：1001。
	TransactionId string `json:"transaction_id"`
	// 字段含义：单笔收款金额。
	// 格式规则：int64。
	// 业务规则：单位分。
	// 示例：100。
	Amount int64 `json:"amount"`
	// 字段含义：收款成功渠道。
	// 格式规则：string。
	// 业务规则：枚举值：抖音支付：DOUYINPAY；商户渠道：MCH。
	// 示例：DOUYINPAY。
	PaidType string `json:"paid_type"`
	// 字段含义：收款成功时间。
	// 格式规则：string，支持两种格式：yyyyMMddHHmmss和yyyyMMdd。
	// 示例：20220208060910。
	PaidTime string `json:"paid_time"`
	// 字段含义：收款银行。
	// 格式规则：string。
	// 业务规则：银行类型，采用字符串类型的银行标识。默认不返回，若需获取具体信息请联系抖音支付运营。
	// 示例：ICBC。
	BankType string `json:"bank_type"`
	// 字段含义：优惠信息。
	// 格式规则：array。
	// 业务规则：享受优惠时返回该字段。
	// 示例：[{"coupon_id":"109519","name":"满20减1元","scope":"GLOBAL","type":"CASH","amount":100,"currency":"CNY"}]。
	PromotionDetail []*PromotionDetail `json:"promotion_detail"`
}

type PromotionDetail struct {
	// 字段含义：券ID。
	// 格式规则：string。
	// 示例：1001。
	CouponId string `json:"coupon_id"`
	// 字段含义：优惠名称。
	// 格式规则：string。
	// 示例：服务费。
	Name string `json:"name"`
	// 字段含义：优惠范围。
	// 格式规则：string。
	// 业务规则：枚举值：GLOBAL：全场优惠；SINGLE：单品优惠。
	// 示例：GLOBAL。
	Scope string `json:"scope"`
	// 字段含义：优惠类型。
	// 格式规则：string。
	// 业务规则：枚举值：CASH：充值；NOCASH：免充值。
	// 示例：Order_Paid。
	Type string `json:"type"`
	// 字段含义：优惠券面额。
	// 格式规则：int64。
	// 业务规则：单位分。
	// 示例：100。
	Amount int64 `json:"amount"`
	// 字段含义：活动ID。
	// 格式规则：string。
	// 示例：1001。
	StockId string `json:"stock_id"`
	// 字段含义：抖音支付出资。
	// 格式规则：int64。
	// 业务规则：单位分。
	// 示例：100。
	DouyinpayContribute int64 `json:"douyinpay_contribute"`
	// 字段含义：商户出资。
	// 格式规则：int64。
	// 业务规则：单位分。
	// 示例：100。
	MerchantContribute int64 `json:"merchant_contribute"`
	// 字段含义：其他出资。
	// 格式规则：int64。
	// 业务规则：单位分。
	// 示例：100。
	OtherContribute int64 `json:"other_contribute"`
	// 字段含义：优惠币种。
	// 格式规则：string。
	// 业务规则：CNY：人民币，境内商户号仅支持人民币。
	// 示例：CNY。
	Currency string `json:"currency"`
	// 字段含义：商品列表。
	// 格式规则：array。
	// 业务规则：预留字段。
	// 示例：[{"goods_id":"D123456","quantity":1,"unit_price":10000,"discount_amount":0,"goods_remark":"商品备注"}]。
	GoodsDetail []*GoodsDetail `json:"goods_detail"`
}

type GoodsDetail struct {
	// 字段含义：商品编码。
	// 格式规则：string。
	// 示例：1001。
	GoodsId string `json:"goods_id"`
	// 字段含义：商品数量。
	// 格式规则：int32。
	// 示例：100。
	Quantity int32 `json:"quantity"`
	// 字段含义：商品价格。
	// 格式规则：int64。
	// 业务规则：单位分。
	// 示例：100。
	UnitPrice int64 `json:"unit_price"`
	// 字段含义：商品优惠金额。
	// 格式规则：int64。
	// 业务规则：单位分。
	// 示例：100。
	DiscountAmount int64 `json:"discount_amount"`
	// 字段含义：商品备注。
	// 格式规则：string。
	// 示例：商品备注。
	GoodsRemark string `json:"goods_remark"`
}

// RiskFund 服务风险金
type RiskFund struct {
	// 字段含义：风险金名称。
	// 格式规则：string[1,30]。
	// 业务规则：枚举值：DEPOSIT：押金；ADVANCE：预付款；CASH_DEPOSIT：保证金；ESTIMATE_ORDER_COST：预估订单费用。
	// 示例：服务费。
	Name string `json:"name"`
	// 字段含义：风险金额。
	// 格式规则：int64。
	// 业务规则：1、数字，必须>0（单位分）2、风险金额≤每个服务ID的风险金额上限 3、当商户优惠字段为空时，付费项目总金额≤服务ID的风险金额上限 （未填写金额的付费项目，视为该付费项目金额为0）。
	// 示例：100。
	Amount int64 `json:"amount,omitempty"`
	// 字段含义：风险说明。
	// 格式规则：string[1,30]。
	// 示例：酒店预估缴纳费用。
	Description string `json:"description,omitempty"`
}

// PostItem 后付费信息
type PostItem struct {
	// 字段含义：付费名称。
	// 格式规则：string[1,20]。
	// 示例：服务费。
	Name string `json:"name"`
	// 字段含义：付费金额。
	// 格式规则：int64。
	// 业务规则：此付费项目总金额，必须≥0（单位分），等于0时代表不需要扣费。如果填写了name（付费名称），amount或description必须填写其一或都填。
	// 示例：100。
	Amount int64 `json:"amount,omitempty"`
	// 字段含义：付费说明。
	// 格式规则：string[1,30]。
	// 业务规则：描述计费规则。
	// 示例：服务费：100/小时。
	Description string `json:"description,omitempty"`
	// 字段含义：付费数量。
	// 格式规则：int64。
	// 示例：100。
	Count int64 `json:"count,omitempty"`
}

// TimeRange 服务时间范围
type TimeRange struct {
	// 字段含义：服务开始时间。
	// 格式规则：string[14]，格式为：yyyyMMddHHmmss。
	// 业务规则：用户下单时确认的服务开始时间（比如用户今天下单，明天开始接受服务，这里指的是明天的服务开始时间）。
	// 示例：20220208060910。
	StartTime string `json:"start_time,omitempty"`
	// 字段含义：服务开始时间备注。
	// 格式规则：string[1,20]。
	// 业务规则：服务开始时间有填时，可填写服务开始时间备注。
	// 示例：开始租借日期。
	StartTimeRemark string `json:"start_time_remark,omitempty"`
	// 字段含义：服务结束时间。
	// 格式规则：string[14]，格式为：yyyyMMddHHmmss。
	// 业务规则：用户享受服务的完成时间。
	// 示例：20220208060910。
	EndTime string `json:"end_time,omitempty"`
	// 字段含义：服务结束时间备注。
	// 格式规则：string[1,20]。
	// 业务规则：服务结束时间有填时，可填写服务结束时间备注。
	// 示例：结束租借日期。
	EndTimeRemark string `json:"end_time_remark,omitempty"`
}

// Location 位置信息
type Location struct {
	// 字段含义：服务开始地点。
	// 格式规则：string[1,50]。
	// 示例：美味餐厅。
	StartLocation string `json:"start_location,omitempty"`
	// 字段含义：服务结束地点。
	// 格式规则：string[1,50]。
	// 示例：美味餐厅。
	EndLocation string `json:"end_location,omitempty"`
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	// 字段含义：用户终端IP。
	// 格式规则：string[1,45]，支持IPv4和IPv6两种格式的IP地址。
	// 示例：14.23.150.211。
	ClientIp string `json:"client_ip"`
	// 字段含义：商户端设备号。
	// 格式规则：string[1,32]。
	// 业务规则：门店号或收银设备ID。
	// 示例：13467007045764。
	DeviceId string `json:"device_id,omitempty"`
	// 字段含义：门店信息。
	// 格式规则：object。
	// 示例：{"id":"1089","name":"辉煌购物中心店","area_code":"100089","address":"北京市海淀区辉煌购物中心5层505"}。
	StoreInfo *StoreInfo `json:"store_info,omitempty"`
}

// 门店信息（预留字段）
type StoreInfo struct {
	// 字段含义：门店编号。
	// 格式规则：string[1,32]。
	// 示例：1089。
	Id string `json:"id"`
	// 字段含义：门店名称。
	// 格式规则：string[1,256]。
	// 示例：辉煌购物中心店。
	Name string `json:"name"`
	// 字段含义：地区编码。
	// 格式规则：string[1,32]。
	// 示例：100089。
	AreaCode string `json:"area_code"`
	// 字段含义：详细的商户门店地址。
	// 格式规则：string[1,512]。
	// 示例：北京市海淀区辉煌购物中心5层505。
	Address string `json:"address"`
}

// ApiSynchronizeServiceOrderInfoRequest 同步服务订单信息请求参数
type ApiSynchronizeServiceOrderInfoRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string[1,32]。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：同步类型。
	// 格式规则：string[1,20]。
	// 业务规则：枚举值：订单已支付：ORDER_PAID。
	// 示例：ORDER_PAID。
	Type   string `json:"type"` // Order_Paid-订单已支付
	Detail struct {
		// 字段含义：收款成功时间。
		// 格式规则：string[14]。
		// 示例：20220208060910。
		PaidTime string `json:"paid_time"`
	} `json:"detail"` // 同步详情
}

// ApiSynchronizeServiceOrderInfoResponse 同步服务订单信息响应参数
type ApiSynchronizeServiceOrderInfoResponse struct {
	// 字段含义：应用ID。
	// 格式规则：string。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string。
	// 业务规则：直连商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：服务信息。
	// 格式规则：string。
	// 示例：某某酒店。
	ServiceIntroduction string `json:"service_introduction"`
	// 字段含义：总金额。
	// 格式规则：int64。
	// 业务规则：1. 金额：数字，必须≥0（单位：分） 2. 总金额=后付费项目金额之和-后付费商户优惠项目金额之和，且小于等于服务风险金额。取消订单时，该字段必须为0 。
	// 示例：100。
	TotalAmount int64 `json:"total_amount"`
	// 字段含义：商户接收回调通知的地址。
	// 格式规则：string。
	// 示例：https://www.bytedance.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：商户数据包。
	// 格式规则：string。
	// 业务规则：商户自定义数据，回调时原样返回。
	// 示例：自定义数据。
	Attach string `json:"attach"`
	// 字段含义：优惠标记。
	// 格式规则：string。
	// 示例：{}。
	GoodsTag string `json:"goods_tag"`

	// 字段含义：抖音支付服务订单号。
	// 格式规则：string。
	// 业务规则：每个抖音支付服务订单号与商户号下的商户服务订单号一一对应。
	// 示例：10050100220208060900000023310957。
	OrderId string `json:"order_id"`
	// 字段含义：服务订单状态。
	// 格式规则：string。
	// 业务规则：表示当前单据状态，枚举值：CREATED: 商户已创建服务订单；DOING: 服务订单进行中；DONE: 服务订单完成；REVOKED: 商户取消服务订单；EXPIRED: 服务订单已失效。
	// 示例：DOING。
	State string `json:"state"`
	// 字段含义：服务订单状态描述。
	// 格式规则：string。
	// 业务规则：对服务订单“进行中(DOING)”状态的附加说明：MCH_COMPLETE: 商户完结；USER_PAYING: 用户支付中；USER_CONFIRM: 用户确认。
	// 示例：USER_PAYING。
	StateDescription string `json:"state_description"`

	// 字段含义：用户标识。
	// 格式规则：string。
	// 业务规则：用户在直连商户appid下的唯一标识。
	// 示例：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o。
	OpenId string `json:"openid"`
	// 字段含义：商户协议号。
	// 格式规则：string。
	// 业务规则：商户侧生成的协议号，在同一个商户号下唯一。
	// 示例：out3B2B64CC652AE663。
	AuthorizationCode string `json:"authorization_code"`
	// 字段含义：收款信息。
	// 格式规则：object。
	// 业务规则：收款成功后，展示具体的收款信息（非0元完结时返回）。
	// 示例：{"state":"PROCESSING","total_amount":10000,"paying_amount":10000,"paid_amount":0,"details":[]}。
	Collection *Collection `json:"collection"`
	// 字段含义：服务风险金。
	// 格式规则：object。
	// 示例：{"name":"ESTIMATE_ORDER_COST","amount":10000,"description":"预估订单费用"}。
	RiskFund *RiskFund `json:"risk_fund,omitempty"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：后付费商户优惠。
	// 格式规则：array。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
	// 字段含义：支付场景描述。
	// 格式规则：object。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
}

// ApiQueryServiceOrderRequest 查询服务订单信息请求参数
type ApiQueryServiceOrderRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string[1,32]。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
}

// ApiQueryServiceOrderResponse 查询服务订单信息响应参数
type ApiQueryServiceOrderResponse struct {
	// 字段含义：应用ID。
	// 格式规则：string。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string。
	// 业务规则：直连商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：商户协议号。
	// 格式规则：string。
	// 业务规则：商户侧生成的协议号，在同一个商户号下唯一。
	// 示例：out3B2B64CC652AE663。
	AuthorizationCode string `json:"authorization_code"`
	// 字段含义：用户标识。
	// 格式规则：string。
	// 业务规则：用户在直连商户appid下的唯一标识。
	// 示例：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o。
	OpenId string `json:"openid"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：服务信息。
	// 格式规则：string。
	// 业务规则：用于介绍本订单所提供的服务。
	// 示例：某某酒店。
	ServiceIntroduction string `json:"service_introduction"`
	// 字段含义：总金额。
	// 格式规则：int64。
	// 业务规则：1. 金额：数字，必须≥0（单位：分） 2. 总金额=后付费项目金额之和-后付费商户优惠项目金额之和，且小于等于订单风险金额。取消订单时，该字段必须为0 。
	// 示例：100。
	TotalAmount int64 `json:"total_amount"`
	// 字段含义：业务接入时分配，配置商户和场景维度信息。
	// 格式规则：string。
	// 示例：https://www.bytedance.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：附加数据包。
	// 格式规则：string。
	// 示例：自定义数据。
	Attach string `json:"attach,omitempty"`
	// 字段含义：抖音支付服务订单号。
	// 格式规则：string。
	// 业务规则：每个抖音支付服务订单号与商户号下的商户服务订单号一一对应。
	// 示例：10050100220208060900000023310957。
	OrderId string `json:"order_id"`
	// 字段含义：服务订单状态。
	// 格式规则：string。
	// 业务规则：表示当前单据状态。枚举值：CREATED: 商户已创建服务订单；DOING: 服务订单进行中；DONE: 服务订单完成；REVOKED: 商户取消服务订单；EXPIRED: 服务订单已失效。
	// 示例：DOING。
	State string `json:"state"`
	// 字段含义：服务订单状态描述。
	// 格式规则：string。
	// 业务规则：对服务订单“进行中”状态的附加说明：MCH_COMPLETE: 商户完结；USER_PAYING: 用户支付中；USER_CONFIRM: 用户已确认。
	// 示例：USER_PAYING。
	StateDescription string `json:"state_description"`
	// 字段含义：服务风险金。
	// 格式规则：object。
	// 示例：{"name":"ESTIMATE_ORDER_COST","amount":10000,"description":"预估订单费用"}。
	RiskFund *RiskFund `json:"risk_fund"`
	// 字段含义：收款信息。
	// 格式规则：object。
	// 业务规则：收款成功后，展示具体的收款信息（非0元完结时返回）。
	// 示例：{"state":"PROCESSING","total_amount":10000,"paying_amount":10000,"paid_amount":0,"details":[]}。
	Collection *Collection `json:"collection"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：后付费商户优惠。
	// 格式规则：array。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
}

// ApiCancelServiceOrderRequest 查询服务订单信息请求参数
type ApiCancelServiceOrderRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string[1,32]。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：取消原因。
	// 格式规则：string[1,50]。
	// 示例：服务费：100/小时。
	Reason string `json:"reason,omitempty"`
}

// ApiCancelServiceOrderResponse 取消服务订单信息响应参数
type ApiCancelServiceOrderResponse struct {
	// 字段含义：应用ID。
	// 格式规则：string。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string。
	// 业务规则：直连商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，配置商户和场景维度信息 。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：抖音支付服务订单号。
	// 格式规则：string。
	// 业务规则：每个抖音支付服务订单号与商户号下的商户服务订单号一一对应。
	// 示例：10050100220208060900000023310957。
	OrderId string `json:"order_id"`
}

// ApiModifyAmountRequest 服务订单改价请求参数
type ApiModifyAmountRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string[1,32]，只能是数字、大小写字母_-*。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：总金额。
	// 格式规则：int64。
	// 业务规则：必须≥0（单位：分）。
	// 示例：100。
	TotalAmount int64 `json:"total_amount"`
	// 字段含义：后付费项目。
	// 业务规则：后付费项目列表，最多包含100条付费项目。
	// 格式规则：array。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：后付费商户优惠。
	// 业务规则：商户优惠列表，最多包含30条商户优惠。
	// 格式规则：array。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：改价原因。
	// 格式规则：string[1,50]。
	// 示例：修改金额。
	Reason string `json:"reason,omitempty"`
}

// ApiModifyAmountResponse 服务订单改价响应参数
type ApiModifyAmountResponse struct {
	// 字段含义：应用ID。
	// 格式规则：string。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string。
	// 业务规则：直连商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：抖音支付服务订单号。
	// 格式规则：string。
	// 业务规则：每个抖音支付服务订单号与商户号下的商户服务订单号一一对应。
	// 示例：10050100220208060900000023310957。
	OrderId string `json:"order_id"`
	// 字段含义：服务信息。
	// 格式规则：string。
	// 业务规则：服务信息，用于介绍本订单所提供的服务。
	// 示例：某某酒店。
	ServiceIntroduction string `json:"service_introduction"`
	// 字段含义：总金额。
	// 格式规则：int64。
	// 业务规则：1. 金额：数字，必须≥0（单位：分） 2. 总金额=后付费项目金额之和-后付费商户优惠项目金额之和，且小于等于服务风险金额。取消订单时，该字段必须为0。
	// 示例：100。
	TotalAmount int64 `json:"total_amount"`
	// 字段含义：商户接收回调通知的地址。
	// 格式规则：string。
	// 示例：https://www.bytedance.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：商户数据包。
	// 格式规则：string。
	// 示例：自定义数据。
	Attach string `json:"attach,omitempty"`
	// 字段含义：服务订单状态。
	// 格式规则：string。
	// 业务规则：表示当前单据状态。枚举值：CREATED: 商户已创建服务订单；DOING: 服务订单进行中；DONE: 服务订单完成；REVOKED: 商户取消服务订单；EXPIRED: 服务订单已失效。
	// 示例：DOING。
	State string `json:"state"`
	// 字段含义：订单状态说明。
	// 格式规则：string。
	// 业务规则：对服务订单“进行中”状态的附加说明：MCH_COMPLETE: 商户完结；USER_PAYING: 用户支付中；USER_CONFIRM: 用户确认。
	// 示例：USER_PAYING。
	StateDescription string `json:"state_description"`
	// 字段含义：服务风险金。
	// 格式规则：object。
	// 示例：{"name":"ESTIMATE_ORDER_COST","amount":10000,"description":"预估订单费用"}。
	RiskFund *RiskFund `json:"risk_fund"`
	// 字段含义：收款信息。
	// 格式规则：object。
	// 示例：{"state":"PROCESSING","total_amount":10000,"paying_amount":10000,"paid_amount":0,"details":[]}。
	Collection *Collection `json:"collection"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：后付费商户优惠。
	// 格式规则：array。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
}

// ApiCompleteServiceOrderRequest 完结服务订单请求参数
type ApiCompleteServiceOrderRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一。此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string[1,32]，只能是数字、大小写字母_-*。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：总金额。
	// 格式规则：int64。
	// 业务规则：1. 金额：数字，必须≥0（单位：分） 2. 总金额=后付费项目金额之和-后付费商户优惠项目金额之和，且小于等于服务风险金额。取消订单时，该字段必须为0 。
	// 示例：100。
	TotalAmount int64 `json:"total_amount"`
	// 字段含义：商户数据包。
	// 格式规则：string[1,256]。
	// 示例：自定义数据。
	Attach string `json:"attach,omitempty"`
	// 字段含义：优惠标记。
	// 格式规则：string[1,512]，json格式。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{}。
	GoodsTag string `json:"goods_tag"`

	// 字段含义：后付费项目。
	// 格式规则：array。
	// 业务规则：后付费项目列表，最多包含100条付费项目。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：后付费商户优惠。
	// 格式规则：array。
	// 业务规则：商户优惠列表，最多包含30条商户优惠。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
}

// ApiCompleteServiceOrderResponse 完结服务订单响应参数
type ApiCompleteServiceOrderResponse struct {
	// 字段含义：应用ID。
	// 格式规则：string。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一。此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string。
	// 业务规则：直连商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：抖音支付服务订单号。
	// 格式规则：string。
	// 业务规则：每个抖音支付服务订单号与商户号下对应的商户服务订单号一一对应。
	// 示例：10050100220208060900000023310957。
	OrderId string `json:"order_id"`
	// 字段含义：服务信息。
	// 格式规则：string。
	// 业务规则：服务信息，用于介绍本订单所提供的服务。
	// 示例：某某酒店。
	ServiceIntroduction string `json:"service_introduction"`
	// 字段含义：总金额。
	// 格式规则：int64。
	// 业务规则：1. 金额：数字，必须≥0（单位：分） 2. 总金额=后付费项目金额之和-后付费商户优惠项目金额之和，且小于等于服务风险金额。取消订单时，该字段必须为0 。
	// 示例：100。
	TotalAmount int64 `json:"total_amount"`
	// 字段含义：商户数据包。
	// 格式规则：string。
	// 示例：自定义数据。
	Attach string `json:"attach,omitempty"`
	// 字段含义：服务订单状态。
	// 格式规则：string。
	// 业务规则：表示当前单据状态。枚举值：CREATED: 商户已创建服务订单；DOING: 服务订单进行中；DONE: 服务订单完成；REVOKED: 商户取消服务订单；EXPIRED: 服务订单已失效。
	// 示例：DOING。
	State string `json:"state"`
	// 字段含义：优惠标记。
	// 格式规则：string。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{}。
	GoodsTag string `json:"goods_tag"`
	// 字段含义：服务订单状态描述。
	// 格式规则：string。
	// 业务规则：DOING状态的附加说明，枚举值：	// 业务规则：对服务订单“进行中”状态的附加说明：MCH_COMPLETE: 商户完结；USER_PAYING: 用户支付中；USER_CONFIRM: 用户已确认。MCH_COMPLETE表示商户完结；USER_PAYING表示用户支付中；USER_CONFIRM表示用户已确认。
	// 示例：USER_PAYING。
	StateDescription string `json:"state_description"`
	// 字段含义：服务风险金。
	// 格式规则：object。
	// 示例：{"name":"ESTIMATE_ORDER_COST","amount":10000,"description":"预估订单费用"}。
	RiskFund *RiskFund `json:"risk_fund"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：后付费商户优惠。
	// 格式规则：array。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
}

// ApiCloseCreditServiceRequest 解除用户授权关系请求参数
type ApiCloseCreditServiceRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，配置商户和场景维度信息 。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：商户协议号。
	// 格式规则：string[1,64]。
	// 业务规则：商户侧生成的协议号，在同一个商户号下唯一。
	// 示例：out3B2B64CC652AE663。
	AuthorizationCode string `json:"authorization_code"`
	// 字段含义：解约原因。
	// 格式规则：string[1,50]。
	// 示例：服务费：100/小时。
	Reason string `json:"reason"`
}

// ApiCloseCreditServiceResponse 解除用户授权关系响应参数
type ApiCloseCreditServiceResponse struct {
}

type ApiServiceOrderPayRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：商户订单号。
	// 格式规则：string。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
}

type ApiServiceOrderPayResponse struct {
	// 字段含义：应用ID。
	// 格式规则：string。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string。
	// 业务规则：商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：商户订单号。
	// 格式规则：string。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：抖音支付服务订单号。
	// 格式规则：string。
	// 业务规则：抖音支付服务订单号。
	// 示例：每个抖音支付服务订单号与商户号下的商户服务订单号一一对应。
	OrderId string `json:"order_id"`
}

type ApiCreditSrvSignApplyRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，配置商户和场景维度信息 。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：商户协议号。
	// 格式规则：string[1,64]。
	// 业务规则：商户侧生成的协议号，在同一个商户号下唯一。
	// 示例：out3B2B64CC652AE663。
	AuthorizationCode string `json:"authorization_code"`
	// 字段含义：通知地址。
	// 格式规则：string[1,255]，必须为https协议，必须为直接可访问的url，不能携带参数。
	// 业务规则：商户接收授权开启/解除回调通知的地址（用户从商户侧发起解约或用户在抖音侧解约，都会以该地址通知商户）。
	// 示例：https://www.bytedance.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：优惠标记。
	// 格式规则：string[1,256]，JSON字符串格式。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{}。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 字段含义：扩展参数。
	// 格式规则：string[1,2048]，JSON字符串格式。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{}。
	ExtInfo string `json:"ext_info,omitempty"`
}

type ApiCreditSrvSignApplyResponse struct {
	// 字段含义：先享后付申请token。
	// 格式规则：string。
	// 业务规则：用于拉起抖音/抖音极速版客户端。
	// 示例：5778aadY9nlt1234XixCkFIGYnV2V。
	PayscoreApplyToken string `json:"payscore_apply_token"`
}

type ApiCreditSrvSignQueryRequest struct {
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：商户协议号。
	// 格式规则：string[1,64]。
	// 业务规则：商户侧生成的协议号，在同一个商户号下唯一。
	// 示例：out3B2B64CC652AE663。
	AuthorizationCode string `json:"authorization_code"`
}

type ApiCreditSrvSignQueryResponse struct {
	// 字段含义：应用ID。
	// 格式规则：string。
	// 业务规则：商户在抖音开放平台申请的应用ID，全局唯一；此处请填写移动应用类型的AppID，并确保该AppID与mchid有绑定关系。
	// 示例：awofz9bncda6w2w4。
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	Mchid string `json:"mchid"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：用户标识。
	// 格式规则：string。
	// 业务规则：户在直连商户appid下的唯一标识。
	// 示例：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o。
	OpenId string `json:"openid"`
	// 字段含义：商户协议号。
	// 格式规则：string。
	// 业务规则：商户侧生产的外部协议号，在同一个商户号下唯一。
	// 示例：out3B2B64CC652AE663。
	AuthorizationCode string `json:"authorization_code"`
	// 字段含义：授权状态。
	// 格式规则：string。
	// 业务规则：标识用户授权服务情况：UNAVAILABLE: 用户未授权服务；AVAILABLE: 用户已授权服务。
	// 示例：UNAVAILABLE。
	AuthorizationState string `json:"authorization_state"`
	// 字段含义：解除授权时间。
	// 格式规则：string，遵循rfc3339标准格式，格式为yyyy-MM-DDTHH:mm:ss.sss+TIMEZONE。
	// 示例：20220208060910。
	CancelAuthorizationTime string `json:"cancel_authorization_time"`
	// 字段含义：授权成功时间。
	// 格式规则：string，遵循rfc3339标准格式，格式为yyyy-MM-DDTHH:mm:ss.sss+TIMEZONE。
	// 示例：20220208060910。
	AuthorizationSuccessTime string `json:"authorization_success_time"`
}

// ApiPartnerCreateServiceOrderRequest 服务商创建服务订单请求参数。
type ApiPartnerCreateServiceOrderRequest struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：商户服务单号。
	// 格式规则：string[1,32]，只能是数字、大小写字母_-*。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：服务信息。
	// 格式规则：string[1,20]。
	// 业务规则：服务信息，用于介绍本订单所提供的服务。
	// 示例：某某酒店。
	ServiceIntroduction string `json:"service_introduction"`
	// 字段含义：商户协议号。
	// 格式规则：string[1,64]。
	// 业务规则：商户侧生成的协议号，在同一个商户号下唯一，免确认模式下必传。
	// 示例：1736173225954193889。
	AuthorizationCode string `json:"authorization_code"`
	// 字段含义：商户数据包。
	// 格式规则：string[1,1024]。
	// 示例：{"out_product_category":"BATTERY_CHANGE"}。
	Attach string `json:"attach,omitempty"`
	// 字段含义：通知地址。
	// 格式规则：string[1,256]，必须为https地址。请确保回调URL是外部可正常访问的，且不能携带后缀参数。
	// 示例：https://www.bytedance.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：优惠标记。
	// 格式规则：string[1,512]，json格式。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{"product_tag":"xxxx","biz_scene":"aaaa"}。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 字段含义：服务风险金。
	// 格式规则：object。
	// 示例：{"name":"ESTIMATE_ORDER_COST","amount":10000,"description":"预估订单费用"}。
	RiskFund *RiskFund `json:"risk_fund"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：商户优惠。
	// 格式规则：array。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：实际服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 字段含义：是否需要用户确认。
	// 格式规则：boolean。
	// 业务规则：【需确认模式】必须传true，【免确认模式】可不传或传false，若service_id不支持免确认，则报错返回。
	// 示例：true。
	NeedUserConfirm bool `json:"need_user_confirm,omitempty"`
	// 字段含义：扩展参数。
	// 格式规则：string，json字符串。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{"A":"a"}。
	ExtInfo string `json:"ext_info,omitempty"`
}

// ApiPartnerCreateServiceOrderResponse 服务商创建服务订单响应参数。
type ApiPartnerCreateServiceOrderResponse struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户商户号。
	// 格式规则：string。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：商户服务单号。
	// 格式规则：string，只能是数字、大小写字母_-*。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：服务信息。
	// 格式规则：string。
	// 业务规则：服务信息，用于介绍本订单所提供的服务。
	// 示例：某某酒店。
	ServiceIntroduction string `json:"service_introduction"`
	// 字段含义：服务风险金。
	// 格式规则：object。
	// 示例：{"name":"ESTIMATE_ORDER_COST","amount":10000,"description":"预估订单费用"}。
	RiskFund *RiskFund `json:"risk_fund"`
	// 字段含义：抖音支付服务订单号。
	// 格式规则：string。
	// 业务规则：每个抖音支付服务订单号与商户号下对应的商户服务订单号一一对应。
	// 示例：10050100220208060900000023310957。
	OrderId string `json:"order_id"`
	// 字段含义：服务订单状态。
	// 格式规则：string。
	// 业务规则：枚举值：CREATED表示商户已创建服务订单；DOING表示服务订单进行中；DONE表示服务订单完成；REVOKED表示商户取消服务订单；EXPIRED表示服务订单已失效。
	// 示例：DOING。
	State string `json:"state"`
	// 字段含义：订单状态说明。
	// 格式规则：string。
	// 业务规则：对服务订单“进行中”状态的附加说明：MCH_COMPLETE: 商户完结；USER_PAYING: 用户支付中；USER_CONFIRM: 用户已确认。
	// 示例：USER_PAYING。
	StateDescription string `json:"state_description"`
	// 字段含义：服务商商户下用户标识。
	// 格式规则：string。
	// 业务规则："用户在服务商户对应appid下的唯一标识。传入sp_appid，未传入sub_appid时返回"。
	// 示例：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o 。
	SpOpenid string `json:"sp_openid"`
	// 字段含义：子商户下用户标识。
	// 格式规则：string。
	// 业务规则：用户在子商户对应appid下的唯一标识，sub_appid传入时返回。
	// 示例：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o。
	SubOpenid string `json:"sub_openid"`
	// 字段含义：商户协议号。
	// 格式规则：string。
	// 业务规则：商户侧生成的协议号，在同一个商户号下唯一。
	// 示例：1736173225954193889。
	AuthorizationCode string `json:"authorization_code"`
	// 字段含义：商户数据包。
	// 格式规则：string。
	// 示例：{"out_product_category":"BATTERY_CHANGE"}。
	Attach string `json:"attach"`
	// 字段含义：通知地址。
	// 格式规则：string，必须为https地址。请确保回调URL是外部可正常访问的，且不能携带后缀参数。
	// 示例：https://www.bytedance.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：商户优惠。
	// 格式规则：array。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：实际服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 字段含义：是否需要用户确认。
	// 格式规则：boolean。
	// 业务规则：【需确认模式】必须传true，【免确认模式】可不传或传false，若service_id不支持免确认，则报错返回。
	// 示例：true。
	NeedUserConfirm bool `json:"need_user_confirm,omitempty"`
	// 字段含义：优惠标记。
	// 格式规则：string，json格式。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{"product_tag":"xxxx","biz_scene":"aaaa"}。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 字段含义：跳转抖音支付token。
	// 格式规则：string。
	// 示例：5778aadY9nlt1234XixCkFIGYnV2V。
	PayscoreApplyToken string `json:"payscore_apply_token,omitempty"`
}

// ApiPartnerCompleteServiceOrderRequest 服务商完结服务订单请求参数。
type ApiPartnerCompleteServiceOrderRequest struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：商户服务单号。
	// 格式规则：string[1,32]，只能是数字、大小写字母_-*。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：总金额。
	// 格式规则：int64。
	// 业务规则：1. 金额：数字，必须≥0（单位：分） 2. 总金额 =（完结付费项目1…+完结付费项目n）-（完结商户优惠项目1…+完结商户优惠项目n）。
	// 示例：100。
	TotalAmount int64 `json:"total_amount"`
	// 字段含义：商户数据包。
	// 格式规则：string[1,1024]。
	// 示例：{"out_product_category":"BATTERY_CHANGE"}。
	Attach string `json:"attach"`
	// 字段含义：优惠标记。
	// 格式规则：string[1,512]，json格式。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{"product_tag":"xxxx","biz_scene":"aaaa"}。
	GoodsTag string `json:"goods_tag"`
	// 字段含义：支付渠道信息。
	// 格式规则：object。
	// 业务规则：渠道信息，可以指定优先、指定必用渠道类型。
	// 示例：{"preset_channel":[{"channel_code":"OUTSIDE_MC","channel_id":"M2025042914432001250054700","channel_amount":100,"channel_ext_info":""}]}。
	ChannelInfo *ChannelInfo `json:"channel_info,omitempty"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：商户优惠。
	// 格式规则：array。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：实际服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
}

// ApiPartnerCompleteServiceOrderResponse 服务商完结服务订单响应参数。
type ApiPartnerCompleteServiceOrderResponse struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户商户号。
	// 格式规则：string。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：商户服务单号。
	// 格式规则：string，只能是数字、大小写字母_-*。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：服务信息。
	// 格式规则：string。
	// 业务规则：服务信息，用于介绍本订单所提供的服务。
	// 示例：某某酒店。
	ServiceIntroduction string `json:"service_introduction"`
	// 字段含义：抖音支付服务订单号。
	// 格式规则：string。
	// 业务规则：每个抖音支付服务订单号与商户号下对应的商户服务订单号一一对应。
	// 示例：10050100220208060900000023310957。
	OrderId string `json:"order_id"`
	// 字段含义：服务订单状态。
	// 格式规则：string。
	// 业务规则：枚举值：CREATED表示商户已创建服务订单；DOING表示服务订单进行中；DONE表示服务订单完成；REVOKED表示商户取消服务订单；EXPIRED表示服务订单已失效。
	// 示例：DOING。
	State string `json:"state"`
	// 字段含义：订单状态说明。
	// 格式规则：string。
	// 业务规则：对服务订单“进行中”状态的附加说明：MCH_COMPLETE: 商户完结；USER_PAYING: 用户支付中；USER_CONFIRM: 用户已确认。
	// 示例：USER_PAYING。
	StateDescription string `json:"state_description"`
	// 字段含义：服务风险金。
	// 格式规则：object。
	// 示例：{"name":"ESTIMATE_ORDER_COST","amount":10000,"description":"预估订单费用"}。
	RiskFund *RiskFund `json:"risk_fund,omitempty"`
	// 字段含义：订单总金额。
	// 格式规则：int64。
	// 业务规则：1. 金额：数字，必须≥0（单位：分） 2. 总金额 =（完结付费项目1…+完结付费项目n）-（完结商户优惠项目1…+完结商户优惠项目n）。
	// 示例：10000。
	TotalAmount int64 `json:"total_amount"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：商户优惠。
	// 格式规则：array。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：实际服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 字段含义：优惠标记。
	// 格式规则：string，json格式。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{"product_tag":"xxxx","biz_scene":"aaaa"}。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 字段含义：商户数据包。
	// 格式规则：string。
	// 示例：{"out_product_category":"BATTERY_CHANGE"}。
	Attach string `json:"attach,omitempty"`
}

// ChannelInfo 支付渠道信息
type ChannelInfo struct {
	// 字段含义：指定渠道列表。
	// 格式规则：array。
	// 示例：[{"channel_code":"OUTSIDE_MC","channel_id":"M2025042914432001250054700","channel_amount":100,"channel_ext_info":""}]。
	PresetChannel []*PresetChannel `json:"preset_channel,omitempty"`
}

// PresetChannel 指定渠道列表
type PresetChannel struct {
	// 字段含义：指定渠道。
	// 格式规则：string[1,128]。
	// 业务规则：指定必用渠道名称，上传与抖音支付约定的值。
	// 示例：OUTSIDE_MC。
	ChannelCode string `json:"channel_code"`
	// 字段含义：指定渠道ID。
	// 格式规则：string[1,64]。
	// 业务规则：指定必用渠道ID，上传与抖音支付约定的渠道。
	// 示例：HLQXK。
	ChannelId string `json:"channel_id"`
	// 字段含义：指定渠道金额。
	// 格式规则：int64。
	// 示例：100。
	ChannelAmount int64 `json:"channel_amount"`
	// 字段含义：扩展信息。
	// 格式规则：string。
	// 示例：""。
	ChannelExtInfo string `json:"channel_ext_info"`
}

// ApiPartnerQueryServiceOrderRequest 服务商查询服务订单请求参数。
type ApiPartnerQueryServiceOrderRequest struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：商户服务单号。
	// 格式规则：string[1,32]，只能是数字、大小写字母_-*。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
}

// ApiPartnerQueryServiceOrderResponse 服务商查询服务订单响应参数。
type ApiPartnerQueryServiceOrderResponse struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户商户号。
	// 格式规则：string。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：商户服务单号。
	// 格式规则：string，只能是数字、大小写字母_-*。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：服务信息。
	// 格式规则：string。
	// 业务规则：服务信息，用于介绍本订单所提供的服务。
	// 示例：某某酒店。
	ServiceIntroduction string `json:"service_introduction"`
	// 字段含义：抖音支付服务订单号。
	// 格式规则：string。
	// 业务规则：每个抖音支付服务订单号与商户号下对应的商户服务订单号一一对应。
	// 示例：10050100220208060900000023310957。
	OrderId string `json:"order_id"`
	// 字段含义：服务订单状态。
	// 格式规则：string。
	// 业务规则：枚举值：CREATED表示商户已创建服务订单；DOING表示服务订单进行中；DONE表示服务订单完成；REVOKED表示商户取消服务订单；EXPIRED表示服务订单已失效。
	// 示例：DOING。
	State string `json:"state"`
	// 字段含义：订单状态说明。
	// 格式规则：string。
	// 业务规则：对服务订单“进行中”状态的附加说明：MCH_COMPLETE: 商户完结；USER_PAYING: 用户支付中；USER_CONFIRM: 用户已确认。
	// 示例：USER_PAYING。
	StateDescription string `json:"state_description"`
	// 字段含义：服务风险金。
	// 格式规则：object。
	// 示例：{"name":"ESTIMATE_ORDER_COST","amount":10000,"description":"预估订单费用"}。
	RiskFund *RiskFund `json:"risk_fund,omitempty"`
	// 字段含义：订单总金额。
	// 格式规则：int64。
	// 业务规则：1. 金额：数字，必须≥0（单位：分） 2. 总金额 =（完结付费项目1…+完结付费项目n）-（完结商户优惠项目1…+完结商户优惠项目n）。
	// 示例：10000。
	TotalAmount int64 `json:"total_amount"`
	// 字段含义：服务商商户下用户标识。
	// 格式规则：string。
	// 业务规则：用户在服务商户对应appid下的唯一标识。传入sp_appid，未传入sub_appid时返回。
	// 示例：。
	SpOpenid string `json:"sp_openid"`
	// 字段含义：子商户下用户标识。
	// 格式规则：string。
	// 业务规则：用户在子商户对应appid下的唯一标识，sub_appid传入时返回。
	// 示例：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o。
	SubOpenid string `json:"sub_openid"`
	// 字段含义：商户协议号。
	// 格式规则：string。
	// 业务规则：商户侧生成的协议号，在同一个商户号下唯一。
	// 示例：1736173225954193889。
	AuthorizationCode string `json:"authorization_code"`
	// 字段含义：商户数据包。
	// 格式规则：string。
	// 示例：{"out_product_category":"BATTERY_CHANGE"}。
	Attach string `json:"attach"`
	// 字段含义：通知地址。
	// 格式规则：string。
	// 示例：https://www.bytedance.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：商户优惠。
	// 格式规则：array。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：收款信息。
	// 格式规则：object。
	// 示例：{"state":"PROCESSING","total_amount":10000,"paying_amount":10000,"paid_amount":0,"details":[]}。
	Collection *Collection `json:"collection,omitempty"`
	// 字段含义：实际服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
}

// ApiPartnerCancelServiceOrderRequest 服务商取消服务订单请求参数。
type ApiPartnerCancelServiceOrderRequest struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：商户服务单号。
	// 格式规则：string[1,32]，只能是数字、大小写字母_-*。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：取消原因。
	// 格式规则：string[1,50]。
	// 示例：用户取消。
	Reason string `json:"reason"`
}

// ApiPartnerCancelServiceOrderResponse 服务商取消服务订单响应参数。
type ApiPartnerCancelServiceOrderResponse struct {
	// 字段含义：服务订单状态。
	// 格式规则：string。
	// 业务规则：枚举值：CREATED表示商户已创建服务订单；DOING表示服务订单进行中；DONE表示服务订单完成；REVOKED表示商户取消服务订单；EXPIRED表示服务订单已失效。
	// 示例：DOING。
	State string `json:"state"`
	// 字段含义：抖音支付服务订单号。
	// 格式规则：string。
	// 业务规则：每个抖音支付服务订单号与商户号下对应的商户服务订单号一一对应。
	// 示例：10050100220208060900000023310957。
	OrderId string `json:"order_id"`
	// 字段含义：订单状态说明。
	// 格式规则：string。
	// 业务规则：DOING状态的附加说明，枚举值：MCH_COMPLETE表示商户完结；USER_PAYING表示用户支付中；USER_CONFIRM表示用户已确认。
	// 示例：USER_PAYING。
	StateDescription string `json:"state_description"`
}

// ApiPartnerSynchronizeServiceOrderInfoRequest 服务商同步服务订单信息请求参数。
type ApiPartnerSynchronizeServiceOrderInfoRequest struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：商户服务单号。
	// 格式规则：string[1,32]，只能是数字、大小写字母_-*。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：同步类型。
	// 格式规则：string[1,20]。
	// 业务规则：枚举值：订单已支付-ORDER_PAID。
	// 示例：ORDER_PAID。
	Type string `json:"type"`
	// 同步内容信息详情
	Detail struct {
		// 字段含义：收款完成时间。
		// 格式规则：string[14]。
		// 示例：20220208060910。
		PaidTime string `json:"paid_time"`
	} `json:"detail"`
}

// ApiPartnerSynchronizeServiceOrderInfoResponse 服务商同步服务订单信息响应参数。
type ApiPartnerSynchronizeServiceOrderInfoResponse struct{}

// ApiPartnerModifyAmountRequest 服务商修改订单金额请求参数。
type ApiPartnerModifyAmountRequest struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：商户服务单号。
	// 格式规则：string[1,32]，只能是数字、大小写字母_-*。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：总金额。
	// 格式规则：int64。
	// 业务规则：1. 金额：数字，必须≥0（单位：分） 2. 总金额 =（完结付费项目1…+完结付费项目n）-（完结商户优惠项目1…+完结商户优惠项目n）。
	// 示例：10000。
	TotalAmount int64 `json:"total_amount"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：商户优惠。
	// 格式规则：array。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：修改原因。
	// 格式规则：string[1,50]。
	// 业务规则：说明服务订单金额下调的原因。
	// 示例：用户取消。
	Reason string `json:"reason,omitempty"`
}

// ApiPartnerModifyAmountResponse 服务商修改订单金额响应参数。
type ApiPartnerModifyAmountResponse struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户商户号。
	// 格式规则：string。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：商户服务单号。
	// 格式规则：string，只能是数字、大小写字母_-*。
	// 业务规则：商户系统内部服务订单号，在同一个商户号下唯一。
	// 示例：OUT_1666688488。
	OutOrderNo string `json:"out_order_no"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：服务信息。
	// 格式规则：string。
	// 业务规则：服务信息，用于介绍本订单所提供的服务。
	// 示例：某某酒店。
	ServiceIntroduction string `json:"service_introduction"`
	// 字段含义：抖音支付服务订单号。
	// 格式规则：string。
	// 业务规则：每个抖音支付服务订单号与商户号下对应的商户服务订单号一一对应。
	// 示例：10050100220208060900000023310957。
	OrderId string `json:"order_id"`
	// 字段含义：服务订单状态。
	// 格式规则：string。
	// 业务规则：枚举值：CREATED表示商户已创建服务订单；DOING表示服务订单进行中；DONE表示服务订单完成；REVOKED表示商户取消服务订单；EXPIRED表示服务订单已失效。
	// 示例：DOING。
	State string `json:"state"`
	// 字段含义：订单状态说明。
	// 格式规则：string。
	// 业务规则：对服务订单“进行中”状态的附加说明：MCH_COMPLETE: 商户完结；USER_PAYING: 用户支付中；USER_CONFIRM: 用户已确认。
	// 示例：USER_PAYING。
	StateDescription string `json:"state_description"`
	// 字段含义：服务风险金。
	// 格式规则：object。
	// 示例：{"name":"ESTIMATE_ORDER_COST","amount":10000,"description":"预估订单费用"}。
	RiskFund *RiskFund `json:"risk_fund,omitempty"`
	// 字段含义：订单总金额。
	// 格式规则：int64。
	// 业务规则：1. 金额：数字，必须≥0（单位：分） 2. 总金额 =（完结付费项目1…+完结付费项目n）-（完结商户优惠项目1…+完结商户优惠项目n）。
	// 示例：10000。
	TotalAmount int64 `json:"total_amount"`
	// 字段含义：实际服务时间段。
	// 格式规则：object。
	// 示例：{"start_time":"20220208060910","end_time":"20220208060920","start_time_remark":"备注1","end_time_remark":"备注2"}。
	TimeRange *TimeRange `json:"time_range,omitempty"`
	// 字段含义：服务位置。
	// 格式规则：object。
	// 示例：{"start_location":"美味餐厅","end_location":"美味餐厅"}。
	Location *Location `json:"location,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 字段含义：商户数据包。
	// 格式规则：string。
	// 示例：{"out_product_category":"BATTERY_CHANGE"}。
	Attach string `json:"attach"`
	// 字段含义：通知地址。
	// 格式规则：string。
	// 示例：https://www.bytedance.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：后付费项目。
	// 格式规则：array。
	// 示例：[{"name":"出行费用","amount":4000,"description":"美味餐厅","count":1}]。
	PostPayments []*PostItem `json:"post_payments,omitempty"`
	// 字段含义：商户优惠。
	// 格式规则：array。
	// 示例：[{"name":"满20减1元","amount":100,"description":"美味餐厅","count":1}]。
	PostDiscounts []*PostItem `json:"post_discounts,omitempty"`
	// 字段含义：收款信息。
	// 格式规则：object。
	// 示例：{"state":"PROCESSING","total_amount":10000,"paying_amount":10000,"paid_amount":0,"details":[]}。
	Collection *Collection `json:"collection,omitempty"`
}

// ApiPartnerCreditSrvSignApplyRequest 服务商申请先享后付授权请求参数。
type ApiPartnerCreditSrvSignApplyRequest struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：商户协议号。
	// 格式规则：string[1,64]。
	// 业务规则：商户侧生成的协议号，在同一个商户号下唯一。
	// 示例：1736173225954193889。
	AuthorizationCode string `json:"authorization_code"`
	// 字段含义：通知地址。
	// 格式规则：string[1,255]，必须为https地址。请确保回调URL是外部可正常访问的，且不能携带后缀参数。
	// 示例：https://www.bytedance.com。
	NotifyUrl string `json:"notify_url"`
	// 字段含义：商户数据包。
	// 格式规则：string[1,1024]。
	// 示例：{"out_product_category":"BATTERY_CHANGE"}。
	Attach string `json:"attach"`
	// 字段含义：优惠标记。
	// 格式规则：string[1,512]，json格式。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{"product_tag":"xxxx","biz_scene":"aaaa"}。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 字段含义：场景信息。
	// 格式规则：object。
	// 示例：{"client_ip":"14.23.150.211","device_id":"13467007045764","store_info":{"id":"1089"}}。
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 字段含义：扩展参数。
	// 格式规则：string[1,2048]，json字符串。
	// 业务规则：和抖音支付协商后可用。
	// 示例：{"A":"a"}。
	ExtInfo string `json:"ext_info,omitempty"`
}

// ApiPartnerCreditSrvSignApplyResponse 服务商申请先享后付授权响应参数。
type ApiPartnerCreditSrvSignApplyResponse struct {
	// 字段含义：先享后付申请token。
	// 格式规则：string。
	// 示例：5778aadY9nlt1234XixCkFIGYnV2V。
	PayscoreApplyToken string `json:"payscore_apply_token"`
}

// ApiPartnerCreditSrvSignQueryRequest 服务商查询用户授权记录请求参数。
type ApiPartnerCreditSrvSignQueryRequest struct {
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：商户协议号。
	// 格式规则：string[1,64]。
	// 业务规则：商户侧生成的协议号，在同一个商户号下唯一。
	// 示例：1736173225954193889。
	AuthorizationCode string `json:"authorization_code"`
}

// ApiPartnerCreditSrvSignQueryResponse 服务商查询用户授权记录响应参数。
type ApiPartnerCreditSrvSignQueryResponse struct {
	// 字段含义：服务商商户号。
	// 格式规则：string。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：服务商应用ID。
	// 格式规则：string。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：子商户商户号。
	// 格式规则：string。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：商户协议号。
	// 格式规则：string。
	// 业务规则：商户侧生成的协议号，在同一个商户号下唯一。
	// 示例：1736173225954193889。
	AuthorizationCode string `json:"authorization_code"`
	// 字段含义：服务ID。
	// 格式规则：string。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：服务商商户下用户标识。
	// 格式规则：string。
	// 业务规则：用户在服务商户对应appid下的唯一标识。传入sp_appid，未传入sub_appid时返回。
	// 示例：示例：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o。。
	SpOpenid string `json:"sp_openid"`
	// 字段含义：子商户下用户标识。
	// 格式规则：string。
	// 业务规则：用户在子商户对应appid下的唯一标识，sub_appid传入时返回。
	// 示例：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o。
	SubOpenid string `json:"sub_openid"`
	// 字段含义：授权状态。
	// 格式规则：string。
	// 业务规则：标识用户授权服务情况： UNAVAILABLE: 用户未授权服务 AVAILABLE: 用户已授权服务。
	// 示例：AVAILABLE。
	AuthorizationState string `json:"authorization_state"`
	// 字段含义：解除授权时间。
	// 格式规则：string，遵循rfc3339标准格式，格式为yyyy-MM-DDTHH:mm:ss.sss+TIMEZONE 。
	// 示例："2026-05-28T13:22:45.120+08:00"。
	CancelAuthorizationTime string `json:"cancel_authorization_time"`
	// 字段含义：授权成功时间。
	// 格式规则：string，遵循rfc3339标准格式，格式为yyyy-MM-DDTHH:mm:ss.sss+TIMEZONE 。
	// 示例：2026-05-22T13:22:45.120+08:00。
	AuthorizationSuccessTime string `json:"authorization_success_time"`
}

// ApiPartnerCloseCreditServiceRequest 服务商解除用户授权关系请求参数。
type ApiPartnerCloseCreditServiceRequest struct {
	// 字段含义：服务商应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SpAppid string `json:"sp_appid"`
	// 字段含义：服务商商户号。
	// 格式规则：string[1,32]。
	// 业务规则：服务商的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084。
	SpMchid string `json:"sp_mchid"`
	// 字段含义：子商户应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：由抖音支付生成的应用ID，全局唯一。
	// 示例：awofz9bncda6w2w4。
	SubAppid string `json:"sub_appid"`
	// 字段含义：子商户商户号。
	// 格式规则：string[1,32]。
	// 业务规则：子商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605085。
	SubMchid string `json:"sub_mchid"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：业务接入时分配，用于配置商户和场景维度信息。
	// 示例：101。
	ServiceId string `json:"service_id"`
	// 字段含义：商户协议号。
	// 格式规则：string[1,64]。
	// 业务规则：商户侧生成的协议号，在同一个商户号下唯一。
	// 示例：1736173225954193889。
	AuthorizationCode string `json:"authorization_code"`
	// 字段含义：解约原因。
	// 格式规则：string[1, 50]。
	// 示例：用户取消。
	Reason string `json:"reason,omitempty"`
}

// ApiPartnerCloseCreditServiceResponse 服务商解除用户授权关系响应参数。
type ApiPartnerCloseCreditServiceResponse struct{}
