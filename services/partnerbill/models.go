package partnerbill

import (
	"fmt"
	neturl "net/url"

	"github.com/douyinpay/douyinpay-go/tools/consts"
)

type Bill struct {
	HashType    string `json:"hash_type,omitempty"`    // hash类型
	HashValue   string `json:"hash_value,omitempty"`   // hash值
	DownloadUrl string `json:"download_url,omitempty"` // 下载地址
}

// ApplyTradeBillRequest 商户下载交易或结算账单
type ApplyTradeBillRequest struct {
	SpMchid  string `json:"sp_mchid,omitempty"`  // 商户号
	SubMchid string `json:"sub_mchid,omitempty"` // 子商户号
	BillDate string `json:"bill_date,omitempty"` // 账期
	TarType  string `json:"tar_type,omitempty"`  // 压缩类型
}

func (r ApplyTradeBillRequest) Validate() error {
	if r.SpMchid == "" || r.BillDate == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}

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

func (r ApplyTradeBillRequest) GetPath() string {
	return consts.ApplyTradeBillPath
}

// ApplyFundFlowBillRequest 商户下载资金账单
type ApplyFundFlowBillRequest struct {
	SpMchid     string `json:"sp_mchid,omitempty"`     // 商户号
	SubMchid    string `json:"sub_mchid,omitempty"`    // 子商户号
	BillDate    string `json:"bill_date,omitempty"`    // 账期
	AccountType string `json:"account_type,omitempty"` // 账户类型
	TarType     string `json:"tar_type,omitempty"`     // 压缩类型
}

func (r ApplyFundFlowBillRequest) Validate() error {
	if r.SpMchid == "" || r.BillDate == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}

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

func (r ApplyFundFlowBillRequest) GetPath() string {
	return consts.ApplyFundFlowBillPath
}

// ApplySplitBillRequest 商户下载分账账单
type ApplySplitBillRequest struct {
	SpMchid  string `json:"sp_mchid,omitempty"`  // 商户号
	SubMchid string `json:"sub_mchid,omitempty"` // 子商户号
	BillDate string `json:"bill_date,omitempty"` // 账期
	TarType  string `json:"tar_type,omitempty"`  // 压缩类型
}

func (r ApplySplitBillRequest) Validate() error {
	if r.SpMchid == "" || r.BillDate == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}

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

func (r ApplySplitBillRequest) GetPath() string {
	return consts.ApplySplitBillPath
}
