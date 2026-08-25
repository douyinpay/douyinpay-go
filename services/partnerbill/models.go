package partnerbill

import (
	"fmt"
	neturl "net/url"

	"github.com/douyinpay/douyinpay-go/tools/consts"
)

// Bill 表示服务商账单申请接口的响应体。
type Bill struct {
	// 字段含义：哈希类型。
	// 格式规则：字符串。
	// 业务规则：用于描述账单文件校验值的摘要算法类型，当前常见返回值为 SHA1。
	// 示例：SHA1
	HashType string `json:"hash_type,omitempty"`
	// 字段含义：哈希值。
	// 格式规则：字符串。
	// 业务规则：用于校验下载后的账单文件完整性。
	// 示例：3b7e4b5c0f9c7d4a8fbc1234567890abcdef1234
	HashValue string `json:"hash_value,omitempty"`
	// 字段含义：账单下载地址。
	// 格式规则：URL。
	// 业务规则：下载地址存在时效限制，建议获取后尽快下载并完成文件校验。
	// 示例：https://download.example.com/bill.gz
	DownloadUrl string `json:"download_url,omitempty"`
}

// ApplyTradeBillRequest 表示服务商申请交易账单的请求参数。
type ApplyTradeBillRequest struct {
	// 字段含义：服务商商户号。
	// 必填规则：必填。
	// 格式规则：字符串。
	// 业务规则：由抖音支付生成并下发，服务商模式必传。
	// 示例：699000000000001
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：子商户号。
	// 必填规则：选填。
	// 格式规则：字符串。
	// 业务规则：传入后仅返回指定子商户对应的账单数据。
	// 示例：699000000000101
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：账单日期。
	// 必填规则：必填。
	// 格式规则：yyyy-MM-dd。
	// 业务规则：仅支持申请近三个月内且为昨日及以前的账单。
	// 示例：2026-08-24
	BillDate string `json:"bill_date,omitempty"`
	// 字段含义：压缩类型。
	// 必填规则：选填。
	// 格式规则：枚举字符串。
	// 业务规则：常用值为 GZIP，返回 gzip 压缩包账单。
	// 示例：GZIP
	TarType string `json:"tar_type,omitempty"`
}

// Validate 校验交易账单请求的必填字段。
func (r ApplyTradeBillRequest) Validate() error {
	if r.SpMchid == "" || r.BillDate == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}

// GetQueryParams 返回交易账单请求对应的 query 参数。
func (r ApplyTradeBillRequest) GetQueryParams() neturl.Values {
	queryParams := neturl.Values{}
	queryParams.Add("sp_mchid", r.SpMchid)
	queryParams.Add("bill_date", r.BillDate)
	queryParams.Add("tar_type", r.TarType)
	if r.SubMchid != "" {
		queryParams.Add("sub_mchid", r.SubMchid)
	}
	return queryParams
}

// GetPath 返回交易账单接口路径。
func (r ApplyTradeBillRequest) GetPath() string {
	return consts.ApplyTradeBillPath
}

// ApplyFundFlowBillRequest 表示服务商申请资金账单的请求参数。
type ApplyFundFlowBillRequest struct {
	// 字段含义：服务商商户号。
	// 必填规则：必填。
	// 格式规则：字符串。
	// 业务规则：由抖音支付生成并下发，服务商模式必传。
	// 示例：699000000000001
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：子商户号。
	// 必填规则：选填。
	// 格式规则：字符串。
	// 业务规则：传入后仅返回指定子商户对应的账单数据。
	// 示例：699000000000101
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：账单日期。
	// 必填规则：必填。
	// 格式规则：yyyy-MM-dd。
	// 业务规则：仅支持申请近三个月内且为昨日及以前的账单。
	// 示例：2026-08-24
	BillDate string `json:"bill_date,omitempty"`
	// 字段含义：账户类型。
	// 必填规则：选填。
	// 格式规则：枚举字符串。
	// 业务规则：常见取值包括 BaseAccount（基本户）和 OperationAccount（运营户），其他取值以开放平台最新文档为准。
	// 示例：BaseAccount
	AccountType string `json:"account_type,omitempty"`
	// 字段含义：压缩类型。
	// 必填规则：选填。
	// 格式规则：枚举字符串。
	// 业务规则：常用值为 GZIP，返回 gzip 压缩包账单。
	// 示例：GZIP
	TarType string `json:"tar_type,omitempty"`
}

// Validate 校验资金账单请求的必填字段。
func (r ApplyFundFlowBillRequest) Validate() error {
	if r.SpMchid == "" || r.BillDate == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}

// GetQueryParams 返回资金账单请求对应的 query 参数。
func (r ApplyFundFlowBillRequest) GetQueryParams() neturl.Values {
	queryParams := neturl.Values{}
	queryParams.Add("sp_mchid", r.SpMchid)
	queryParams.Add("bill_date", r.BillDate)
	queryParams.Add("tar_type", r.TarType)
	if r.AccountType != "" {
		queryParams.Add("account_type", r.AccountType)
	}
	if r.SubMchid != "" {
		queryParams.Add("sub_mchid", r.SubMchid)
	}
	return queryParams
}

// GetPath 返回资金账单接口路径。
func (r ApplyFundFlowBillRequest) GetPath() string {
	return consts.ApplyFundFlowBillPath
}

// ApplySplitBillRequest 表示服务商申请分账账单的请求参数。
type ApplySplitBillRequest struct {
	// 字段含义：服务商商户号。
	// 必填规则：必填。
	// 格式规则：字符串。
	// 业务规则：由抖音支付生成并下发，服务商模式必传。
	// 示例：699000000000001
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：子商户号。
	// 必填规则：选填。
	// 格式规则：字符串。
	// 业务规则：传入后仅返回指定子商户对应的账单数据。
	// 示例：699000000000101
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：账单日期。
	// 必填规则：必填。
	// 格式规则：yyyy-MM-dd。
	// 业务规则：仅支持申请近三个月内且为昨日及以前的账单。
	// 示例：2026-08-24
	BillDate string `json:"bill_date,omitempty"`
	// 字段含义：压缩类型。
	// 必填规则：选填。
	// 格式规则：枚举字符串。
	// 业务规则：常用值为 GZIP，返回 gzip 压缩包账单。
	// 示例：GZIP
	TarType string `json:"tar_type,omitempty"`
}

// Validate 校验分账账单请求的必填字段。
func (r ApplySplitBillRequest) Validate() error {
	if r.SpMchid == "" || r.BillDate == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}

// GetQueryParams 返回分账账单请求对应的 query 参数。
func (r ApplySplitBillRequest) GetQueryParams() neturl.Values {
	queryParams := neturl.Values{}
	queryParams.Add("sp_mchid", r.SpMchid)
	queryParams.Add("bill_date", r.BillDate)
	queryParams.Add("tar_type", r.TarType)
	if r.SubMchid != "" {
		queryParams.Add("sub_mchid", r.SubMchid)
	}
	return queryParams
}

// GetPath 返回分账账单接口路径。
func (r ApplySplitBillRequest) GetPath() string {
	return consts.ApplySplitBillPath
}
