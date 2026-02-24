package bill

import (
	"fmt"
	neturl "net/url"

	"github.com/douyinpay/douyinpay-go/tools/consts"
)

// BillApplyRequest 商户下载交易或结算账单
type BillApplyRequest struct {
	// 商户号
	Mchid string `json:"mchid,omitempty"`
	// 应用id
	BillDate string `json:"bill_date,omitempty"`
	// 子商户的商户号，由抖音支付生成并下发。服务商模式下必须传递此参数
	BillType string `json:"bill_type,omitempty"`
	// 原支付交易对应的抖音订单号
	TarType string `json:"tar_type,omitempty"`
}

func (r BillApplyRequest) Validate() error {
	if r.Mchid == "" {
		return fmt.Errorf("field `Mchid` is required and must be specified in BillApplyRequest")
	}
	if r.BillDate == "" {
		return fmt.Errorf("field `BillDate` is required and must be specified in BillApplyRequest")
	}
	return nil
}

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

func (r BillApplyRequest) GetPath() string {
	return consts.BillApplyPath
}

type Bill struct {
	// 商户号
	HashType string `json:"hash_type,omitempty"`
	// 应用id
	HashValue string `json:"hash_value,omitempty"`
	// 子商户的商户号，由抖音支付生成并下发。服务商模式下必须传递此参数
	DownloadUrl string `json:"download_url,omitempty"`
}

// ApplyFundFlowBillRequest 商户下载资金账单
type ApplyFundFlowBillRequest struct {
	Mchid       string `json:"mchid,omitempty"`        // 商户号
	BillDate    string `json:"bill_date,omitempty"`    // 账期
	AccountType string `json:"account_type,omitempty"` // 账户类型
	TarType     string `json:"tar_type,omitempty"`     // 压缩类型
}

func (r ApplyFundFlowBillRequest) Validate() error {
	if r.Mchid == "" || r.BillDate == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}

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

func (r ApplyFundFlowBillRequest) GetPath() string {
	return consts.ApplyFundFlowBillPath
}

// ApplySplitBillRequest 商户下载分账账单
type ApplySplitBillRequest struct {
	Mchid    string `json:"mchid,omitempty"`     // 商户号
	BillDate string `json:"bill_date,omitempty"` // 账期
	TarType  string `json:"tar_type,omitempty"`  // 压缩类型
}

func (r ApplySplitBillRequest) Validate() error {
	if r.Mchid == "" || r.BillDate == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}

func (r ApplySplitBillRequest) GetQueryParams() neturl.Values {
	queryParams := neturl.Values{}
	queryParams.Add("mchid", r.Mchid)
	queryParams.Add("bill_date", r.BillDate)
	queryParams.Add("tar_type", r.TarType)
	return queryParams
}

func (r ApplySplitBillRequest) GetPath() string {
	return consts.ApplySplitBillPath
}
