package indirectmerchant

import "fmt"

// ApplyRequest 间连商户入驻申请请求（create_apply）。
// 从业机构提交商户资料入驻成为抖音支付间连商户。
// 官方文档：https://partner.douyinpay.com/wiki/682c7a8e82b07604fd4deccb/682d90a7f6171b04e93c7b83
type ApplyRequest struct {
	// OutIndirectMerchantID 间连商户外部商户号，从业机构在清算机构侧登记的间连商户唯一标识。
	OutIndirectMerchantID string `json:"out_indirect_merchant_id,omitempty"`
	// OutPaymentInstID 从业机构外部商户号，从业机构在清算机构侧登记的商户号。
	OutPaymentInstID string `json:"out_payment_inst_id,omitempty"`
	// OutChannelSPID 渠道商外部商户号，渠道商在清算机构侧登记的商户号。
	OutChannelSPID string `json:"out_channel_sp_id,omitempty"`
	// OutPaymentInstSubID 子从业机构外部商户号，支持银联下同一从业机构按地区拆分子从业机构的场景，
	// 子从业机构号需与从业机构号（OutPaymentInstID）前 4 位一致，后续交易需使用该子从业机构。
	OutPaymentInstSubID string `json:"out_payment_inst_sub_id,omitempty"`
	// BusinessCode 业务申请编号，由数字、小写字母或下划线组成的自定义唯一编号，
	// 每个编号对应一个申请单，审核通过后生成一个抖音支付商户号。
	BusinessCode string `json:"business_code,omitempty"`
	// SubjectInfo 主体资料，商家的营业执照/登记证书、经营者/法人证件等信息。
	SubjectInfo *SubjectInfo `json:"subject_info,omitempty"`
	// IdentificationInfo 经营者/法人身份证件信息，其中证件持有人姓名、证件号码为敏感字段。
	IdentificationInfo *IdentificationInfo `json:"identification_info,omitempty"`
	// BusinessInfo 经营信息，商家的经营业务、售卖商品/提供服务场景信息。
	BusinessInfo *BusinessInfo `json:"business_info,omitempty"`
	// AdditionInfo 补充信息，根据审核情况额外要求商家提供的补充资料。
	AdditionInfo *AdditionInfo `json:"addition_info,omitempty"`
	// IndirectMerchantID 抖音支付间连商户号，传入已有商户号时会校验一致性，未传则按新入驻创建。
	IndirectMerchantID string `json:"indirect_merchant_id,omitempty"`
}

// CreateApplyRequest 是入驻申请请求的别名，与官方接口命名保持一致。
type CreateApplyRequest = ApplyRequest

// Validate 校验入驻申请的商户定位组合与业务申请编号。
func (r ApplyRequest) Validate() error {
	if r.OutIndirectMerchantID == "" || r.OutPaymentInstID == "" {
		return fmt.Errorf("out_indirect_merchant_id and out_payment_inst_id are required")
	}
	if r.BusinessCode == "" {
		return fmt.Errorf("business_code is required")
	}
	return nil
}

// SubjectInfo 主体资料。
type SubjectInfo struct {
	// SubjectType 主体类型，需与营业执照/登记证书一致，枚举包括 SUBJECT_TYPE_INDIVIDUAL（个体户）、
	// SUBJECT_TYPE_ENTERPRISE（企业）、SUBJECT_TYPE_GOVERNMENT（政府机关）、SUBJECT_TYPE_INSTITUTIONS（事业单位）、
	// SUBJECT_TYPE_OTHERS（其他组织）、SUBJECT_TYPE_MICRO（小微商户）。
	SubjectType string `json:"subject_type,omitempty"`
	// CertificateInfo 主体证件信息，主体为企业/个体户/政府机关/事业单位/其他组织时必填，小微商户无需填写。
	CertificateInfo *CertificateInfo `json:"certificate_info,omitempty"`
}

// CertificateInfo 主体证件信息。
type CertificateInfo struct {
	// CertNo 证件编号，营业执照/登记证书上的统一社会信用代码。
	CertNo string `json:"cert_no,omitempty"`
	// MerchantName 商户名称，对应证件的主体名称，小微商户使用法人名称。
	MerchantName string `json:"merchant_name,omitempty"`
	// CertCopy 证件图片，通过图片上传接口预先上传生成的 MediaID。
	CertCopy string `json:"cert_copy,omitempty"`
	// RegisterAddress 注册地址，营业执照/登记证书的注册地址。
	RegisterAddress string `json:"register_address,omitempty"`
	// LegalPersonName 经营者/法人姓名，营业执照的经营者/法定代表人姓名。
	LegalPersonName string `json:"legal_person_name,omitempty"`
	// PeriodBegin 有效期限开始日期，YYYY-MM-DD 格式。
	PeriodBegin string `json:"period_begin,omitempty"`
	// PeriodEnd 有效期限结束日期，YYYY-MM-DD 格式或"长期"。
	PeriodEnd string `json:"period_end,omitempty"`
}

// IdentificationInfo 经营者/法人身份证件信息。
type IdentificationInfo struct {
	// IdentificationType 证件类型，枚举包括 IDENTIFICATION_TYPE_IDCARD（中国大陆居民身份证）、
	// IDENTIFICATION_TYPE_OVERSEA_PASSPORT（护照）、IDENTIFICATION_TYPE_HONGKONG_PASSPORT、
	// IDENTIFICATION_TYPE_MACAO_PASSPORT、IDENTIFICATION_TYPE_TAIWAN_PASSPORT、
	// IDENTIFICATION_TYPE_FOREIGN_RESIDENT、IDENTIFICATION_TYPE_HONGKONG_MACAO_RESIDENT、
	// IDENTIFICATION_TYPE_TAIWAN_RESIDENT。
	IdentificationType string `json:"identification_type,omitempty"`
	// IdentificationName 证件持有人姓名，为敏感字段。
	// 上送时使用抖音支付平台证书公钥加密，下行时使用服务商私钥解密；
	// SDK 在 CreateApply 发送前自动加密、在 QueryMerchant 返回后自动解密。
	IdentificationName string `json:"identification_name,omitempty"`
	// IdentificationNumber 证件持有人证件号码，为敏感字段。
	// 上送时使用抖音支付平台证书公钥加密，下行时使用服务商私钥解密；
	// SDK 在 CreateApply 发送前自动加密、在 QueryMerchant 返回后自动解密。
	IdentificationNumber string `json:"identification_number,omitempty"`
	// PeriodBegin 证件持有人有效期开始时间，YYYY-MM-DD 格式。
	PeriodBegin string `json:"period_begin,omitempty"`
	// PeriodEnd 证件持有人有效期结束时间，YYYY-MM-DD 格式或"长期"。
	PeriodEnd string `json:"period_end,omitempty"`
	// IdentificationFrontCopy 证件正面照片的 MediaID，身份证为人像面。
	IdentificationFrontCopy string `json:"identification_front_copy,omitempty"`
	// IdentificationBackCopy 证件背面照片的 MediaID，身份证为国徽面。
	IdentificationBackCopy string `json:"identification_back_copy,omitempty"`
}

// BusinessInfo 经营信息。
type BusinessInfo struct {
	// MerchantShortName 商户简称，在支付完成页向买家展示。
	MerchantShortName string `json:"merchant_short_name,omitempty"`
	// ServicePhone 客服电话，将在交易记录中向买家展示。
	ServicePhone string `json:"service_phone,omitempty"`
	// MCC 所属行业，间连商户在清算机构侧登记的行业 MCC 编码。
	MCC string `json:"mcc,omitempty"`
	// SpecialIndustryPics 特殊资质图片的 MediaID 列表，支持 1-5 张图片。
	SpecialIndustryPics []string `json:"special_industry_pics,omitempty"`
}

// AdditionInfo 补充信息。
type AdditionInfo struct {
	// AdditionPics 补充材料图片的 MediaID 列表，最多 5 张图片。
	AdditionPics []string `json:"addition_pics,omitempty"`
	// AdditionRemark 补充说明，文案不超过 100 字。
	AdditionRemark string `json:"addition_remark,omitempty"`
}

// Locator 间连商户定位信息，用于查询/绑定/修改等接口。
// 支持两种定位方式二选一：外部定位（OutIndirectMerchantID + OutPaymentInstID，可选 OutChannelSPID）
// 或抖音支付间连商户号（IndirectMerchantID）。
type Locator struct {
	// OutIndirectMerchantID 间连商户外部商户号。
	OutIndirectMerchantID string `json:"out_indirect_merchant_id,omitempty"`
	// OutPaymentInstID 从业机构外部商户号。
	OutPaymentInstID string `json:"out_payment_inst_id,omitempty"`
	// OutChannelSPID 渠道商外部商户号，仅能与外部定位方式组合使用。
	OutChannelSPID string `json:"out_channel_sp_id,omitempty"`
	// IndirectMerchantID 抖音支付间连商户号。
	IndirectMerchantID string `json:"indirect_merchant_id,omitempty"`
}

// Validate 校验商户定位信息必须且仅能提供一种定位方式。
func (r Locator) Validate() error {
	external := r.OutIndirectMerchantID != "" || r.OutPaymentInstID != ""
	if external && (r.OutIndirectMerchantID == "" || r.OutPaymentInstID == "") {
		return fmt.Errorf("external merchant locator requires both identifiers")
	}
	if external == (r.IndirectMerchantID != "") {
		return fmt.Errorf("exactly one merchant locator is required")
	}
	if r.OutChannelSPID != "" && !external {
		return fmt.Errorf("out_channel_sp_id requires external merchant locator")
	}
	return nil
}

// QueryApplyRequest 查询入驻申请单请求（query_apply）。
// 通过申请单号或业务申请编号查询，两种方式二选一，结果相同。
type QueryApplyRequest struct {
	// ApplymentID 抖音支付申请单号，由抖音支付生成并下发。
	ApplymentID string `json:"applyment_id,omitempty"`
	// BusinessCode 业务申请编号，提交时自定义的唯一编号。
	BusinessCode string `json:"business_code,omitempty"`
}

// CancelRequest 撤销入驻申请单请求（cancel_apply）。
// 仅当申请单处于待提交和审核中时可撤销，通过申请单号或业务申请编号二选一。
type CancelRequest struct {
	// ApplymentID 抖音支付申请单号，由抖音支付生成并下发。
	ApplymentID string `json:"applyment_id,omitempty"`
	// BusinessCode 业务申请编号，提交时自定义的唯一编号。
	BusinessCode string `json:"business_code,omitempty"`
}

// ModifyRequest 修改商户资料请求（modify）。
// 目前仅支持修改商户简称和客服电话，通过商户定位信息指定目标商户。
type ModifyRequest struct {
	// Locator 商户定位信息，外部定位或抖音支付间连商户号二选一。
	Locator
	// MerchantShortName 商户简称，在支付完成页向买家展示。
	MerchantShortName string `json:"merchant_short_name,omitempty"`
	// ServicePhone 客服电话，将在交易记录中向买家展示。
	ServicePhone string `json:"service_phone,omitempty"`
}

// BindAppRequest 绑定应用 APPID 请求（bind_app_id）。
// 为间连子商户配置 APPID 绑定关系，仅支持新增配置，通过商户定位信息指定目标商户。
type BindAppRequest struct {
	// Locator 商户定位信息，外部定位或抖音支付间连商户号二选一。
	Locator
	// AppID 需要给间连商户绑定的应用 APPID。
	AppID string `json:"app_id,omitempty"`
}

// QueryAppRequest 查询应用 APPID 请求（query_app_id）。
// 查询间连子商户配置的 APPID 信息，通过商户定位信息指定目标商户，可按 AppID 过滤。
type QueryAppRequest struct {
	// Locator 商户定位信息，外部定位或抖音支付间连商户号二选一。
	Locator
	// AppID 应用 APPID，用于按应用过滤查询结果。
	AppID string `json:"app_id,omitempty"`
}

// ApplyResponse 入驻申请响应（create_apply）。
type ApplyResponse struct {
	// ApplymentID 抖音支付分配的申请单号。
	ApplymentID string `json:"applyment_id,omitempty"`
}

// QueryApplyResponse 查询入驻申请单响应（query_apply）。
type QueryApplyResponse struct {
	// BusinessCode 业务申请编号。
	BusinessCode string `json:"business_code,omitempty"`
	// ApplymentID 抖音支付申请单号。
	ApplymentID string `json:"applyment_id,omitempty"`
	// IndirectMerchantID 抖音支付间连商户号，申请单状态为 APPLYMENT_STATE_PASSED 时返回。
	IndirectMerchantID string `json:"indirect_merchant_id,omitempty"`
	// ApplymentState 申请单状态，枚举包括 APPLYMENT_STATE_WAITTING_FOR_AUDIT（审核中）、
	// APPLYMENT_STATE_REJECTED（已驳回）、APPLYMENT_STATE_PASSED（已完成）、APPLYMENT_STATE_CANCELED（已作废）。
	ApplymentState string `json:"applyment_state,omitempty"`
	// AuditDetail 驳回原因详情，申请单状态为 APPLYMENT_STATE_REJECTED 时返回。
	AuditDetail []AuditDetail `json:"audit_detail,omitempty"`
}

// AuditDetail 驳回原因详情。
type AuditDetail struct {
	// Field 提交申请单的资料项字段名。
	Field string `json:"field,omitempty"`
	// RejectReason 提交资料项被驳回的原因。
	RejectReason string `json:"reject_reason,omitempty"`
}

// QueryMerchantResponse 查询间连商户响应（query_merchant）。
type QueryMerchantResponse struct {
	// OutIndirectMerchantID 外部间连商户号。
	OutIndirectMerchantID string `json:"out_indirect_merchant_id,omitempty"`
	// OutPaymentInstID 外部从业机构商户号。
	OutPaymentInstID string `json:"out_payment_inst_id,omitempty"`
	// OutChannelSPID 外部渠道商户号。
	OutChannelSPID string `json:"out_channel_sp_id,omitempty"`
	// BusinessCode 业务申请编号，对应该间连商户入驻成功的业务申请编号。
	BusinessCode string `json:"business_code,omitempty"`
	// ApplymentID 申请单号，对应该间连商户入驻成功的申请单号。
	ApplymentID string `json:"applyment_id,omitempty"`
	// IndirectMerchantID 抖音支付间连商户号。
	IndirectMerchantID string `json:"indirect_merchant_id,omitempty"`
	// SubjectInfo 主体资料。
	SubjectInfo *SubjectInfo `json:"subject_info,omitempty"`
	// IdentificationInfo 经营者/法人身份证件信息，其中证件持有人姓名、证件号码由抖音支付加密下发，SDK 自动解密。
	IdentificationInfo *IdentificationInfo `json:"identification_info,omitempty"`
	// BusinessInfo 经营信息。
	BusinessInfo *BusinessInfo `json:"business_info,omitempty"`
}

// AppInfo 应用信息。
type AppInfo struct {
	// AppID 应用 APPID。
	AppID string `json:"app_id,omitempty"`
	// AppType 应用类型（中文），枚举包括"小程序应用"、"移动应用"、"网站应用"。
	AppType string `json:"app_type,omitempty"`
	// BindStatus 绑定状态，枚举包括 BIND（已绑定）、WAIT_CONFIRM（待应用开发者确认）。
	BindStatus string `json:"bind_status,omitempty"`
}

// BindAppResponse 绑定应用 APPID 响应（bind_app_id）。
type BindAppResponse struct {
	// AppInfo 应用信息，绑定成功或待授权情况下返回。
	AppInfo *AppInfo `json:"app_info,omitempty"`
}

// QueryAppResponse 查询应用 APPID 响应（query_app_id）。
type QueryAppResponse struct {
	// AppInfoList 应用信息列表。
	AppInfoList []AppInfo `json:"app_info_list,omitempty"`
}

// UploadMeta 图片上传的媒体文件元信息（meta）。
type UploadMeta struct {
	// Filename 商户上传的媒体图片名称，须以 JPG、BMP、PNG 为后缀。
	Filename string `json:"filename"`
	// SHA256 对图片文件二进制内容进行 SHA256 计算得到的文件摘要。
	SHA256 string `json:"sha256"`
}

// UploadResponse 图片上传响应（/v1/merchant/upload）。
type UploadResponse struct {
	// MediaID 返回的媒体文件标识 ID。
	MediaID string `json:"media_id,omitempty"`
}

// ApplyStatusNotify 间连商户申请状态通知（event_type=INDIRECT_MERCHANT_APPLYMENT_STATUS）
// 回调密文解密后的业务数据。
type ApplyStatusNotify struct {
	// SPMchID 从业机构户号，由抖音支付生成并下发。
	SPMchID string `json:"sp_mchid,omitempty"`
	// ApplymentID 抖音支付生成的申请单号。
	ApplymentID string `json:"applyment_id,omitempty"`
	// BusinessCode 商户自提交的申请单号，站点页面提交情况下为空。
	BusinessCode string `json:"business_code,omitempty"`
	// IndirectMerchantID 抖音支付间连商户号，申请单状态为 APPLYMENT_STATE_PASSED 时返回。
	IndirectMerchantID string `json:"indirect_merchant_id,omitempty"`
	// ApplymentState 申请单状态，枚举包括 APPLYMENT_STATE_REJECTED（已驳回）、
	// APPLYMENT_STATE_PASSED（已完成）、APPLYMENT_STATE_CANCELED（已作废）。
	ApplymentState string `json:"applyment_state,omitempty"`
	// FailReason 失败原因详情，驳回时返回驳回原因，撤销时返回申请单已撤销。
	FailReason string `json:"fail_reason,omitempty"`
}
