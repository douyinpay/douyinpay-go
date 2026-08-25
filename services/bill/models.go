package bill

import (
	"fmt"
	neturl "net/url"

	"github.com/douyinpay/douyinpay-go/tools/consts"
)

// BillApplyRequest 表示直连商户申请交易账单的请求参数。
type BillApplyRequest struct {
	// 字段含义：直连商户号。
	// 必填规则：必填。
	// 格式规则：string，[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：账单日期。
	// 必填规则：必填。
	// 格式规则：yyyy-MM-dd，[1,10]。
	// 业务规则：仅支持三个月内的账单下载申请。
	// 示例：2023-02-25
	BillDate string `json:"bill_date,omitempty"`
	// 字段含义：账单类型。
	// 必填规则：必填。
	// 格式规则：string，[1,32]。
	// 业务规则：TRADE 表示返回当日所有交易订单信息；其他取值以开放平台最新文档为准。
	// 示例：TRADE
	BillType string `json:"bill_type,omitempty"`
	// 字段含义：压缩类型。
	// 必填规则：必填。
	// 格式规则：string，[1,32]。
	// 业务规则：GZIP 表示返回 .gzip 格式的压缩包账单。
	// 示例：GZIP
	TarType string `json:"tar_type,omitempty"`
}

// Validate 校验交易账单请求的必填字段。
func (r BillApplyRequest) Validate() error {
	if r.Mchid == "" {
		return fmt.Errorf("field `Mchid` is required and must be specified in BillApplyRequest")
	}
	if r.BillDate == "" {
		return fmt.Errorf("field `BillDate` is required and must be specified in BillApplyRequest")
	}
	return nil
}

// GetQueryParams 返回交易账单请求对应的 query 参数。
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

// GetPath 返回交易账单接口路径。
func (r BillApplyRequest) GetPath() string {
	return consts.BillApplyPath
}

// Bill 表示账单申请接口的响应体。
type Bill struct {
	// 字段含义：哈希类型。
	// 格式规则：string，[1,32]。
	// 业务规则：SHA1 表示账单文件摘要使用 SHA1 算法。
	// 示例：SHA1
	HashType string `json:"hash_type,omitempty"`
	// 字段含义：哈希值。
	// 格式规则：string，[1,1024]。
	// 业务规则：原始账单（gzip 需要解压缩）的摘要值，用于校验文件完整性。
	// 示例：b3d51ec31534a5e027b49e78a61e20770973f70f
	HashValue string `json:"hash_value,omitempty"`
	// 字段含义：账单下载地址。
	// 格式规则：URL，[1,2048]。
	// 业务规则：供下一步请求账单文件的下载地址，该地址 5 分钟内有效。
	// 示例：https://download.douyinpay.com/v1/billdownload/file?token=shgvbeh1BWB84eXGz8rptvm5Po2uTKOnkqwc8W2DS721jY9rGL6ETWHEpARlSswz
	DownloadUrl string `json:"download_url,omitempty"`
}

// ApplyFundFlowBillRequest 表示直连商户申请资金账单的请求参数。
type ApplyFundFlowBillRequest struct {
	// 字段含义：直连商户号。
	// 必填规则：必填。
	// 格式规则：string，[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：账单日期。
	// 必填规则：必填。
	// 格式规则：yyyy-MM-dd，[1,10]。
	// 业务规则：仅支持三个月内的账单下载申请。
	// 示例：2023-02-25
	BillDate string `json:"bill_date,omitempty"`
	// 字段含义：账户类型。
	// 必填规则：必填。
	// 格式规则：string，[1,32]。
	// 业务规则：可选值包括 BaseAccount（基本账户）和 OperationAccount（运营账户），默认值为 BaseAccount。
	// 示例：BaseAccount
	AccountType string `json:"account_type,omitempty"`
	// 字段含义：压缩类型。
	// 必填规则：必填。
	// 格式规则：string，[1,32]。
	// 业务规则：GZIP 表示返回 .gzip 格式的压缩包账单。
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
	// 格式规则：string，[1,32]。
	// 业务规则：直连商户的商户号，由抖音支付生成并下发。
	// 示例：6020230307605084
	Mchid string `json:"mchid,omitempty"`
	// 字段含义：账单日期。
	// 必填规则：必填。
	// 格式规则：yyyy-MM-dd，[1,10]。
	// 业务规则：仅支持三个月内的账单下载申请。
	// 示例：2023-02-25
	BillDate string `json:"bill_date,omitempty"`
	// 字段含义：压缩类型。
	// 必填规则：必填。
	// 格式规则：string，[1,32]。
	// 业务规则：GZIP 表示返回 .gzip 格式的压缩包账单。
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
