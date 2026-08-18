package cashier

// PrePayConsultRequest 前置咨询请求
type PrePayConsultRequest struct {
	// 字段含义：应用ID。
	// 格式规则：string[1,32]。
	// 业务规则：商家入驻抖音开放平台时分配的应用 ID，需与商户号有绑定关系。
	// 示例：awofz9bncda6w2w4
	Appid string `json:"appid"`
	// 字段含义：直连商户号。
	// 格式规则：string[1,32]。
	// 业务规则：商家入驻抖音支付商家平台分配的商户号。
	// 示例：6020230307605084
	Mchid string `json:"mchid"`
	// 字段含义：商户订单号。
	// 格式规则：string[6,32]，只能是数字、大小写字母、_、-、*。
	// 业务规则：商户系统内部订单号，在同一商户号下唯一。
	// 示例：OUT_1666688488
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 字段含义：服务ID。
	// 格式规则：string[1,32]。
	// 业务规则：先享后付业务接入时分配，配置商户和场景维度信息；咨询先享后付渠道时必传。
	// 示例：100001
	ServiceId string `json:"service_id,omitempty"`
	// 字段含义：营销优惠标记。
	// 格式规则：string[1,512]，键值对类型的 JSON 数据序列化后的字符串。
	// 业务规则：用于营销差异化展示的标记信息，需与抖音支付协商后传递；可传入业务场景 biz_scene、个性化策略 product_tag。
	// 示例：{"biz_scene":"pre_consult","product_tag":"default"}
	GoodsTag string `json:"goods_tag,omitempty"`
	// 字段含义：订单总金额。
	// 格式规则：string[1,11]。
	// 业务规则：单位为分；若需要查询用户营销（例如获取 operation_tip），本字段为必传字段；有订单金额且参与计算的优惠金额大于 0 时，返回满足条件的营销信息。
	// 示例：2000
	TotalAmount string `json:"total_amount,omitempty"`
	// 字段含义：不参与优惠计算订单金额。
	// 格式规则：string[1,11]。
	// 业务规则：单位为分，表示订单中不参与优惠计算的金额。
	// 示例：100
	UndiscountableAmount string `json:"undiscountable_amount,omitempty"`
	// 字段含义：支付产品。
	// 格式规则：字符串数组，暂时只支持传一个支付产品。
	// 业务规则：若需要查询用户营销（例如获取 operation_tip），本字段为必传字段；取值需与 commerical_product_code、trade_type 按官方文档映射表组合传入。
	// 示例：["NormalPay"]
	ProductCode []string `json:"product_code,omitempty"`
	// 字段含义：商业产品码。
	// 格式规则：string，暂时只支持传一个商业产品码。
	// 业务规则：商家和抖音支付签约的产品码；若需要查询用户营销（例如获取 operation_tip），本字段为必传字段。
	// 示例：CO_PAY_APP
	CommericalProductCode string `json:"commerical_product_code,omitempty"`
	// 字段含义：交易类型。
	// 格式规则：string。
	// 业务规则：当前订单的交易类型；若需要查询用户营销（例如获取 operation_tip），本字段为必传字段；取值需与 product_code、commerical_product_code 按官方文档映射表组合传入。
	// 示例：APP
	TradeType string `json:"trade_type,omitempty"`
	// 字段含义：签约模板号。
	// 格式规则：string。
	// 业务规则：代扣签约的模板 ID，商户接入时由支付系统分配；签约、代扣类咨询请求可传入。
	// 示例：2420
	TemplateId string `json:"template_id,omitempty"`
	// 字段含义：商品列表信息。
	// 格式规则：对象数组。
	// 业务规则：订单包含的商品列表信息；传入的商品数量与商品单价乘积总和不可超过订单金额，即 sum{quantity*unit_price} <= total_amount，不满足时返回参数错误。
	// 示例：[{"merchant_goods_id":"app-01","goods_name":"ipad","quantity":2,"unit_price":2000}]
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
	// 字段含义：用户唯一标识。
	// 格式规则：string[1,64]。
	// 业务规则：openid 是用户在应用下的唯一用户标识；openid、加密手机号、设备号能获取到任意一项则真实上送，三项均无法获取时支持全部留空；推荐优先传 openid 以获得最佳接口性能，传入后以 openid 为最高优先级查询用户身份；三项均未传时按抖音新用户查询新用户营销信息。
	// 示例：V3WvSshYq9wWnB
	Openid string `json:"openid,omitempty"`
	// 字段含义：设备号。
	// 格式规则：string[1,64]。
	// 业务规则：设备号类型由 device_type 字段指定；openid、加密手机号、设备号能获取到任意一项则真实上送；当 openid 无法获取时，推荐同时传入 device_id 和 blind_mobile_list 以提升匹配精度。
	// 示例：14b07957e368d91
	DeviceId string `json:"device_id,omitempty"`
	// 字段含义：设备号类型。
	// 格式规则：string[1,16]，枚举字符串。
	// 业务规则：与设备号 device_id 字段组合使用；OAID 表示 OAID，IDFA 表示 IDFA，CAID 表示 CAID（若有多个取最新的）。
	// 示例：OAID
	DeviceType string `json:"device_type,omitempty"`
	// 字段含义：手机号列表。
	// 格式规则：字符串数组，使用 SHA256 算法盲化后的手机号，目前最多支持同时查询两个手机号。
	// 业务规则：仅支持境内手机号；openid、加密手机号、设备号能获取到任意一项则真实上送；当 openid 无法获取时，推荐同时传入 device_id 和 blind_mobile_list 以提升匹配精度。
	// 示例：["66d0fba82f83396b8c37c47e151f8076a479064eccd78517b604646040e8fcfd"]
	BlindMobileList []string `json:"blind_mobile_list,omitempty"`
	// 字段含义：手机号加密方式。
	// 格式规则：string[1,32]，枚举字符串。
	// 业务规则：与手机号列表 blind_mobile_list 组合使用，目前仅支持 SHA256 算法。
	// 示例：SHA256
	EncryptType string `json:"encrypt_type,omitempty"`
	// 字段含义：拓展字段。
	// 格式规则：string[1,1024]，键值对类型的 JSON 数据。
	// 业务规则：用于传递拓展信息，需与抖音支付协商后传递。
	// 示例：
	// TODO: 官方文档未提供标准示例值，暂留空。
	ExtInfo string `json:"ext_info,omitempty"`
}

// GoodsDetail 商品列表信息
type GoodsDetail struct {
	// 字段含义：商户侧商品编码。
	// 格式规则：string[1,32]。
	// 业务规则：商户系统内部的商品编码。
	// 示例：app-01
	MerchantGoodsId string `json:"merchant_goods_id"`
	// 字段含义：抖音支付商品编码。
	// 格式规则：string[1,32]。
	// 业务规则：抖音支付侧的商品编码，选填。
	// 示例：
	// TODO: 官方文档未提供标准示例值，暂留空。
	DouyinpayGoodsId string `json:"douyinpay_goods_id,omitempty"`
	// 字段含义：商品名称。
	// 格式规则：string[1,256]。
	// 业务规则：商品的名称，选填。
	// 示例：ipad
	GoodsName string `json:"goods_name,omitempty"`
	// 字段含义：商品数量。
	// 格式规则：int。
	// 业务规则：与商品单价共同参与订单金额校验。
	// 示例：2
	Quantity int32 `json:"quantity"`
	// 字段含义：商品单价。
	// 格式规则：int。
	// 业务规则：单位为分；与商品数量共同参与订单金额校验。
	// 示例：2000
	UnitPrice int64 `json:"unit_price"`
}

// PrePayConsultResponse 前置咨询响应
type PrePayConsultResponse struct {
	// 字段含义：渠道信息列表，包含产品、优惠信息，只有查询用户成功才返回对应的查询内容。
	ChannelInfoList []ChannelInfo `json:"channel_info_list"`
}

// ChannelInfo 渠道信息
type ChannelInfo struct {
	// 字段含义：加密手机号，即请求参数中传入的手机号。
	BlindMobile string `json:"blind_mobile"`
	// 字段含义：渠道唯一索引。
	ChannelIndex string `json:"channel_index"`
	// 字段含义：支付渠道名称。
	ChannelName string `json:"channel_name"`
	// 字段含义：支付渠道是否可用，true 表示可用，false 表示不可用。
	ChannelEnable bool `json:"channel_enable"`
	// 字段含义：当次咨询匹配到的营销内容，商户可直接取值展示；需先与抖音支付行业运营沟通，由行业运营配置对应的前置咨询策略及营销活动后才会返回。
	OperationInfo *PrePayOperationInfo `json:"operation_info"`
	// 字段含义：指定优惠信息，商户无需关注内容；抖音支付渠道在调用下单接口时需将取值放在优惠标记 goods_tag 中透传带入，先享后付渠道在调用先享后付相关接口时需将取值放在 ext_info 中透传带入，key 均为 assign_discounts。
	AssignDiscounts string `json:"assign_discounts"`
	// 字段含义：扩展信息，JSON 格式，例如返回人群标签信息 biz_tag_list，其中 tag_code 为与抖音支付线下约定的人群标签编码。
	ExtInfo string `json:"ext_info"`
}

// PrePayOperationInfo 营销模型
type PrePayOperationInfo struct {
	// 字段含义：支付产品，即请求中传入的产品编码。
	ProductCode string `json:"product_code"`
	// 字段含义：运营展示数据，可用于展示营销内容的文案组合。
	ViewData *PrePayOperationInfoViewData `json:"view_data"`
}

// PrePayOperationInfoViewData 运营展示数据
type PrePayOperationInfoViewData struct {
	// 字段含义：运营文案描述，可用于展示营销内容的具体文案。
	OperationTip string `json:"operation_tip"`
	// 字段含义：固定立减金额，单位为元，只有当立减金额最小值与最大值相等时返回，否则为空。
	OperationAmount string `json:"operation_amount"`
	// 字段含义：立减金额最小值，单位为元，固定立减和随机立减均有值。
	OperationMinAmount string `json:"operation_min_amount"`
	// 字段含义：立减金额最大值，单位为元，固定立减和随机立减均有值。
	OperationMaxAmount string `json:"operation_max_amount"`
	// 字段含义：立减金额单位，默认为“元”。
	OperationUnit string `json:"operation_unit"`
}
