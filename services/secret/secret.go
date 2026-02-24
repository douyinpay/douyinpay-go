package secret

var (
	// 服务商JSAPI 测试环境配置
	JsapiSpMchID       = "" // 服务商商户号
	JsapiSpAppID       = "" // 服务商应用ID
	JsapiSubMchID      = "" // 子商户号
	JsapiSubAppID      = "" // 子商户应用ID
	JsapiOpenID        = "" // 用户OpenID
	JsapiOutTradeNo    = "" // 订单号
	JsapiTransactionID = "" // 交易订单号

	// merchantSerialNo    = ""
	JsapiMerchantSerialNo    = "" // 商户证书序列号
	JsapiMerchantPrivateKey  = "" // 商户私钥
	JsapiPlatformCertificate = ""

	// 直连商户、服务商 退款 测试环境配置
	RefundMchID         = "" // 商户号
	RefundAppID         = ""
	RefundSubMchID      = "" // 子商户号
	RefundTransactionID = "" // 交易订单号
	RefundOutRefundNo   = ""

	// 商户AppId
	RefundMerchantSerialNo    = "" // 商户证书序列号
	RefundMerchantPrivateKey  = "" // 商户私钥
	RefundPlatformCertificate = ""

	// 直联商户代扣
	DeductMchID               = "" // RSA商户号
	DeductAppID               = "" // RSA商户AppId
	DeductMerchantSerialNo    = "" // RSA商户证书序列号
	DeductMerchantPrivateKey  = ""
	DeductPlatformCertificate = ""
	DeductOutContractCode     = ""
	DeductPlanID              = ""
	DeductContractID          = ""

	DeductDeletePlanID            = ""
	DeductDeleteOutContractCode   = ""
	DeductPayApplyOutContractCode = ""
	DeductPayApplyContractID      = ""

	DeductNotifyAppID      = ""
	DeductNotifyContractID = ""

	// 服务商代扣
	MchID              = ""
	AppID              = ""
	MerchantSerialNo   = ""
	MerchantPrivateKey = ""
	// platformCertificate = ""
	PlatformCertificate   = ""
	SubMchID              = ""
	LocalAppID            = ""
	LocalOutContractCode  = ""
	LocalPlanID           = 0
	LocalPlanIdStr        = ""
	LocalOpenID           = ""
	LocalSubOpenID        = ""
	LocalContractID       = ""
	LocalOutContractCode2 = ""

	LocalContractIDQuery = ""

	PayMchID               = "" // RSA商户号
	PayAppID               = "" // RSA商户AppId
	PayMerchantSerialNo    = "" // RSA商户证书序列号
	PayMerchantPrivateKey  = ""
	PayPlatformCertificate = ""
	// mchID               = "6020230307605084"
	PayOutTradeNo    = ""
	PayTransactionId = ""
	PayQueryTradeNo  = ""
)
