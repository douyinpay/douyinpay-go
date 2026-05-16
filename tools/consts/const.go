package consts

import "time"

// 抖音支付API请求地址
const (
	DouyinPayServer              = "https://api.douyinpay.com"
	DouyPayServer                = "https://api.doupay.com"
	AppPrepayPath                = "/v1/trade/transactions/app"
	ClosePath                    = "/v1/trade/transactions/out-trade-no/{out_trade_no}/close"
	QueryByIdPath                = "/v1/trade/transactions/id/{transaction_id}"
	QueryByOutTradeNoPath        = "/v1/trade/transactions/out-trade-no/{out_trade_no}"
	RefundPath                   = "/v1/trade/refund/domestic/refunds"
	RefundQueryByOutTradeNoPath  = "/v1/trade/refund/domestic/refunds/{out_refund_no}"
	BillApplyPath                = "/v1/bill/billapply"
	ApplyTradeBillPath           = "/v1/bill/tradebill"
	ApplyFundFlowBillPath        = "/v1/bill/fundflowbill"
	ApplySplitBillPath           = "/v1/bill/splitbill"
	H5PrepayPath                 = "/v1/trade/transactions/h5"
	ContractPrepayPath           = "/v1/trade/transactions/contractorder"
	DeductPath                   = "/v1/deduct/payapply"       // 申请扣款
	QueryContractPath            = "/v1/member/querycontract"  // 代扣协议查询
	DeleteContractPath           = "/v1/member/deletecontract" // 代扣协议解约
	JsapiPrepayPath              = "/v1/trade/transactions/jsapi"
	SplitFundPath                = "/v1/trade/profitsharing/orders"
	QuerySplitFundPath           = "/v1/trade/profitsharing/orders/{out_trade_no}"
	ReturnSplitFundPath          = "/v1/trade/profitsharing/return-orders"
	QueryReturnSplitFundPath     = "/v1/trade/profitsharing/return-orders/{out_return_no}"
	FinishSplitFundPath          = "/v1/trade/profitsharing/finish-orders"
	QueryUnsplitFundPath         = "/v1/trade/profitsharing/order/{transaction_id}/amounts"
	AddSplitReceiverPath         = "/v1/trade/profitsharing/receivers/add"
	DeleteSplitReceiverPath      = "/v1/trade/profitsharing/receivers/delete"
	NativePrepayPath             = "/v1/trade/transactions/native"
	CreditContractPrepayPath     = "/v1/trade/transactions/payscorecontractorder"
	PartnerJsapiPrepayPath       = "/v1/trade/partner/transactions/jsapi"                             //服务商jsapi下单
	PartnerNativePrepayPath      = "/v1/trade/partner/transactions/native"                            //服务商native下单
	PartnerClosePath             = "/v1/trade/partner/transactions/out-trade-no/{out_trade_no}/close" //服务商交易关单
	PartnerQueryByIdPath         = "/v1/trade/partner/transactions/id/{transaction_id}"
	PartnerQueryByOutTradeNoPath = "/v1/trade/partner/transactions/out-trade-no/{out_trade_no}"

	CreateServiceOrderPath   = "/v1/payscore/serviceorder/create"                                           // 创建服务订单
	SyncServiceOrderPath     = "/v1/payscore/serviceorder/{out_order_no}/sync"                              // 同步服务订单信息
	ServiceOrderPayPath      = "/v1/payscore/serviceorder/{out_order_no}/pay"                               // 商户发起催收扣款
	QueryServiceOrderPath    = "/v1/payscore/serviceorder/query"                                            // 查询服务订单信息
	CancelServiceOrderPath   = "/v1/payscore/serviceorder/cancel"                                           // 取消服务订单
	ModifyServiceOrderPath   = "/v1/payscore/serviceorder/modify"                                           // 服务订单改价
	CompleteServiceOrderPath = "/v1/payscore/serviceorder/complete"                                         // 完结服务订单
	CreditSrvSignApplyPath   = "/v1/payscore/permissions"                                                   // 申请服务授权
	CreditSrvSignQueryPath   = "/v1/payscore/permissions/authorization-code/{authorization_code}"           // 查询用户授权记录
	CloseCreditServicePath   = "/v1/payscore/permissions/authorization-code/{authorization_code}/terminate" // 解除用户授权关系
	// 直联商户
	PreentrustWebPath = "/v1/agreementauth/preentrustweb" // App纯签约下单
	H5EntrustwebPath  = "/v1/agreementauth/h5entrustweb"  // H5独立签约
	DeductNotifyPath  = "/v1/agreementauth/deductNotify"  // 预约扣费通知
	PayApplyPath      = "/v1/deduct/payapply"             // 申请扣款
	// 服务商接口
	PartnerContractOrderPath         = "/v1/trade/partner/transactions/contractorder"                                                          // 签约中下单
	PartnerPayApplyPath              = "/v1/trade/partner/deduct/payapply"                                                                     // 申请扣款
	PartnerQueryContractPath         = "/v1/agreementauth/partner/contracts/plan-id/{plan_id}/out-contract-code/{out_contract_code}"           // 查询签约关系
	PartnerTerminateContractPath     = "/v1/agreementauth/partner/contracts/plan-id/{plan_id}/out-contract-code/{out_contract_code}/terminate" // 协议解约
	PartnerContractSchedulePath      = "/v1/agreementauth/partner/schedules/contract-id/{contract_id}/schedule"                                // 预约扣费
	PartnerContractScheduleQueryPath = "/v1/agreementauth/partner/schedules/contract-id/{contract_id}"                                         // 预约扣费结果查询

	// 下载平台证书列表
	DownloadCertificatesPath = "/v1/merchant/certificates/getPlatformCertificates"
)

// HTTP 请求报文 Header 相关常量
const (
	Authorization     = "Authorization"       // Header 中的 Authorization 字段
	Accept            = "Accept"              // Header 中的 Accept 字段
	ContentType       = "Content-Type"        // Header 中的 ContentType 字段
	ContentLength     = "Content-Length"      // Header 中的 ContentLength 字段
	UserAgent         = "User-Agent"          // Header 中的 UserAgent 字段
	DouyinpaySerial   = "Douyinpay-Serial"    // Header 中的 Douyinpay-Serial 字段
	DouyinpaySdkAgent = "Douyinpay-Sdk-Agent" // Header 中的 Douyinpay-Sdk-Agent 字段
)

// 常用 ContentType
const (
	ApplicationJSON = "application/json"
)

// 请求报文签名相关常量
const (
	SignatureMessageFormat = "%s\n%s\n%d\n%s\n%s\n" // 数字签名原文格式
	VerifyMessageFormat    = "%d\n%s\n%s\n"         // 验签原文格式
	// 请求头中的 Authorization 拼接格式
	HeaderAuthorizationFormat = "%s mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\""
	EncryptTypeAES256GCM      = "AEAD-AES-256-GCM"
	EncryptTypeSM4            = "SM4"
)

// HTTP 应答报文 Header 相关常量
const (
	DouyinPayTimestamp = "Douyinpay-Timestamp" // 抖音支付回包时间戳
	DouyinPayNonce     = "Douyinpay-Nonce"     // 抖音支付回包随机字符串
	DouyinPaySignature = "Douyinpay-Signature" // 抖音支付回包签名信息
	DouyinPaySerial    = "Douyinpay-Serial"    // 抖音支付回包平台序列号
	RequestID          = "Request-Id"          // 抖音支付回包请求ID
)

// 时间相关常量
const (
	FiveMinute                  = 5 * 60           // 回包校验最长时间（秒）
	DefaultTimeout              = 30 * time.Second // HTTP 请求默认超时时间
	TwentyFourHourAndFourMinute = (24*60 + 4) * 60 // 回调校验最长时间（秒）
)

// 加密类型常量
const (
	CRYPTO_TYPE_RSA = "RSA"
	CRYPTO_TYPE_SM2 = "SM2"
)

// 交易类型
const (
	TradeTypeApp       = "APP"    // APP支付
	TradeTypeJsapi     = "JSAPI"  // JSAPI支付/小程序支付
	TradeTypeNative    = "NATIVE" // 扫码支付（仅线上场景）
	TradeTypeMweb      = "MWEB"   // H5支付
	TradeTypeCyclePay  = "CCP"    // 周期扣款
	TradeTypeSinglePay = "SGP"    // 商户代扣
	TradeTypeNoPwdPay  = "NPP"    // 免密支付
)
