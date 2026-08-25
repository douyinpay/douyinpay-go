package bill

import (
	"fmt"
	neturl "net/url"

	"github.com/douyinpay/douyinpay-go/tools/consts"
)

// BillApplyRequest 表示直连商户申请交易账单或结算账单的请求参数。
type BillApplyRequest struct {
	// 字段含义：直连商户号。
	// 必填规则：必填。
	// 格式规则：字符串。
	// 业务规则：由抖音支付生成并下发。
	// 示例：6020230307605084
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：账单日期。
	// 必填规则：必填。
	// 格式规则：yyyy-MM-dd。
	// 业务规则：仅支持申请近三个月内且为昨日及以前的账单。
	// 示例：2026-08-24
	BillDate string `json:"bill_date,omitempty"`
	// 字段含义：账单类型。
	// 必填规则：按业务场景选填；申请交易账单或结算账单时建议显式传入。
	// 格式规则：枚举字符串。
	// 业务规则：常见取值为 TRADE（交易账单）和 SETTLEMENT（结算账单）；如开放品牌交易账单能力，以开放平台最新文档为准。
	// 示例：TRADE
	BillType string `json:"bill_type,omitempty"`
	// 字段含义：压缩类型。
	// 必填规则：选填。
	// 格式规则：枚举字符串。
	// 业务规则：常用值为 GZIP，返回 gzip 压缩包账单。
	// 示例：GZIP
	TarType string `json:"tar_type,omitempty"`
}

// Validate 校验交易账单或结算账单请求的必填字段。
func (r BillApplyRequest) Validate() error {
	if r.Mchid == "" {
		return fmt.Errorf("field `Mchid` is required and must be specified in BillApplyRequest")
	}
	if r.BillDate == "" {
		return fmt.Errorf("field `BillDate` is required and must be specified in BillApplyRequest")
	}
	return nil
}

// GetQueryParams 返回交易账单或结算账单请求对应的 query 参数。
func (r BillApplyRequest) GetQueryParams() neturl.Values {
	queryParams := neturl.Values{}
	queryParams.Add("mchid", r.Mchid)
	queryParams.Add("bill_date", r.BillDate)
	if r.BillType != "" {
		queryParams.Add("bill_type", r.BillType)
	}
	if r.TarType != "" {
		queryParams.Add("tar_type", r.TarType)
	}
	return queryParams
}

// GetPath 返回交易账单或结算账单接口路径。
func (r BillApplyRequest) GetPath() string {
	return consts.BillApplyPath
}

// Bill 表示账单申请接口的响应体。
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

// ApplyFundFlowBillRequest 表示直连商户申请资金账单的请求参数。
type ApplyFundFlowBillRequest struct {
	// 字段含义：直连商户号。
	// 必填规则：必填。
	// 格式规则：字符串。
	// 业务规则：由抖音支付生成并下发。
	// 示例：6020230307605084
	Mchid string `json:"mchid,omitempty"`
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
	if r.Mchid == "" || r.BillDate == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}

// GetQueryParams 返回资金账单请求对应的 query 参数。
func (r ApplyFundFlowBillRequest) GetQueryParams() neturl.Values {
	queryParams := neturl.Values{}
	queryParams.Add("mchid", r.Mchid)
	queryParams.Add("bill_date", r.BillDate)
	if r.AccountType != "" {
		queryParams.Add("account_type", r.AccountType)
	}
	queryParams.Add("tar_type", r.TarType)
	return queryParams
}

// GetPath 返回资金账单接口路径。
func (r ApplyFundFlowBillRequest) GetPath() string {
	return consts.ApplyFundFlowBillPath
}

// ApplySplitBillRequest 表示直连商户申请分账账单的请求参数。
type ApplySplitBillRequest struct {
	// 字段含义：直连商户号。
	// 必填规则：必填。
	// 格式规则：字符串。
	// 业务规则：由抖音支付生成并下发。
	// 示例：6020230307605084
	Mchid string `json:"mchid,omitempty"`
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
	if r.Mchid == "" || r.BillDate == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}

// GetQueryParams 返回分账账单请求对应的 query 参数。
func (r ApplySplitBillRequest) GetQueryParams() neturl.Values {
	queryParams := neturl.Values{}
	queryParams.Add("mchid", r.Mchid)
	queryParams.Add("bill_date", r.BillDate)
	queryParams.Add("tar_type", r.TarType)
	return queryParams
}

// GetPath 返回分账账单接口路径。
func (r ApplySplitBillRequest) GetPath() string {
	return consts.ApplySplitBillPath
}
