package partnerbill

import (
	"fmt"
	neturl "net/url"

	"github.com/douyinpay/douyinpay-go/tools/consts"
)

// Bill 表示服务商账单申请接口的响应体。
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

// ApplyTradeBillRequest 表示服务商申请交易账单的请求参数。
type ApplyTradeBillRequest struct {
	// 字段含义：商户号。
	// 必填规则：必填。
	// 格式规则：string，[1,32]。
	// 业务规则：由抖音支付生成并下发，支持服务商和平台商户传入。
	// 示例：6020230301343998
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：子商户号。
	// 必填规则：选填。
	// 格式规则：string，[1,32]。
	// 业务规则：不填则默认返回服务商下的交易或退款数据；如需下载某个子商户下的交易或退款数据，则传入该字段；平台商户不支持该字段。
	// 示例：6020230307605084
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：账单日期。
	// 必填规则：必填。
	// 格式规则：yyyy-MM-dd，[1,10]。
	// 业务规则：仅支持三个月内的账单下载申请。
	// 示例：2024-10-10
	BillDate string `json:"bill_date,omitempty"`
	// 字段含义：压缩类型。
	// 必填规则：必填。
	// 格式规则：string，[1,32]。
	// 业务规则：GZIP 表示返回 .gzip 格式的压缩包账单。
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
	// 字段含义：商户号。
	// 必填规则：必填。
	// 格式规则：string，[1,32]。
	// 业务规则：由抖音支付生成并下发，支持服务商和平台商户传入。
	// 示例：6020230301343998
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：子商户号。
	// 必填规则：选填。
	// 格式规则：string，[1,32]。
	// 业务规则：当前特约商户资金账单接口文档未定义该字段，通常无需传入；如平台能力扩展支持，以最新接口文档为准。
	// 示例：6020230307605084
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：账单日期。
	// 必填规则：必填。
	// 格式规则：yyyy-MM-dd，[1,10]。
	// 业务规则：仅支持三个月内的账单下载申请。
	// 示例：2024-10-10
	BillDate string `json:"bill_date,omitempty"`
	// 字段含义：账户类型。
	// 必填规则：选填。
	// 格式规则：string，[1,32]。
	// 业务规则：可选值包括 BaseAccount（基本账户）、OperationAccount（运营账户）和 FeeAccount（手续费账户）；不填默认值为 BaseAccount。
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
	// 字段含义：商户号。
	// 必填规则：必填。
	// 格式规则：string，[1,32]。
	// 业务规则：由抖音支付生成并下发，支持服务商和平台商户传入。
	// 示例：6020230301343998
	SpMchid string `json:"sp_mchid,omitempty"`
	// 字段含义：子商户号。
	// 必填规则：选填。
	// 格式规则：string，[1,32]。
	// 业务规则：不填则默认返回服务商下的所有分账账单；如需下载某个子商户下的分账账单，则传入指定子商户号。
	// 示例：6020230307605084
	SubMchid string `json:"sub_mchid,omitempty"`
	// 字段含义：账单日期。
	// 必填规则：必填。
	// 格式规则：yyyy-MM-dd，[1,10]。
	// 业务规则：仅支持三个月内的账单下载申请。
	// 示例：2024-10-10
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
