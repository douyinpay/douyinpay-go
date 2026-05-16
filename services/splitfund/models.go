package splitfund

type SplitFundRequest struct {
	AppID          string         `json:"appid,omitempty"`
	MchID          string         `json:"mchid,omitempty"`
	TransactionID  string         `json:"transaction_id,omitempty"`
	OutOrderNo     string         `json:"out_order_no,omitempty"`
	Receivers      []ReceiverInfo `json:"receivers,omitempty"`
	UnfreezeUnsplit *bool          `json:"unfreeze_unsplit,omitempty"`
	NotifyURL      string         `json:"notify_url,omitempty"`
}

type ReceiverInfo struct {
	Type        string `json:"type,omitempty"`
	Account     string `json:"account,omitempty"`
	Name        string `json:"name,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
}

type SplitFundResponse struct {
	TransactionID string `json:"transaction_id,omitempty"`
	OutOrderNo    string `json:"out_order_no,omitempty"`
	OrderID       string `json:"order_id,omitempty"`
	MchID         string `json:"mchid,omitempty"`
}

type QuerySplitFundRequest struct {
	MchID         string `json:"mchid,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	OutOrderNo    string `json:"out_order_no,omitempty"`
	OrderID       string `json:"order_id,omitempty"`
}

type QuerySplitFundResponse struct {
	MchID           string               `json:"mchid,omitempty"`
	TransactionID   string               `json:"transaction_id,omitempty"`
	OutOrderNo      string               `json:"out_order_no,omitempty"`
	OrderID         string               `json:"order_id,omitempty"`
	State           string               `json:"state,omitempty"`
	Receivers       []ReceiverSplitResult `json:"receivers,omitempty"`
	FinishAmount    int                  `json:"finish_amount,omitempty"`
	FinishDesc      string               `json:"finish_description,omitempty"`
	SplitFinishTime string               `json:"split_finish_time,omitempty"`
}

type ReceiverSplitResult struct {
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Account     string `json:"account,omitempty"`
	Result      string `json:"result,omitempty"`
	FailReason  string `json:"fail_reason,omitempty"`
	CreateTime  string `json:"create_time,omitempty"`
	FinishTime  string `json:"finish_time,omitempty"`
	DetailID    string `json:"detail_id,omitempty"`
}

type ReturnSplitFundRequest struct {
	MchID       string `json:"mchid,omitempty"`
	OrderID     string `json:"order_id,omitempty"`
	OutOrderNo  string `json:"out_order_no,omitempty"`
	OutReturnNo string `json:"out_return_no,omitempty"`
	ReturnMchID string `json:"return_mchid,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
}

type ReturnSplitFundResponse struct {
	OutOrderNo  string `json:"out_order_no,omitempty"`
	OutReturnNo string `json:"out_return_no,omitempty"`
	ReturnID    string `json:"return_id,omitempty"`
	ReturnMchID string `json:"return_mchid,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Result      string `json:"result,omitempty"`
	FailReason  string `json:"fail_reason,omitempty"`
	CreateTime  string `json:"create_time,omitempty"`
	FinishTime  string `json:"finish_time,omitempty"`
}

type QueryReturnSplitFundRequest struct {
	OutReturnNo string `json:"out_return_no,omitempty"`
	MchID       string `json:"mchid,omitempty"`
	OutOrderNo  string `json:"out_order_no,omitempty"`
}

type QueryReturnSplitFundResponse struct {
	OrderID     string `json:"order_id,omitempty"`
	OutOrderNo  string `json:"out_order_no,omitempty"`
	OutReturnNo string `json:"out_return_no,omitempty"`
	ReturnID    string `json:"return_id,omitempty"`
	ReturnMchID string `json:"return_mchid,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Result      string `json:"result,omitempty"`
	FailReason  string `json:"fail_reason,omitempty"`
	CreateTime  string `json:"create_time,omitempty"`
	FinishTime  string `json:"finish_time,omitempty"`
}

type FinishSplitFundRequest struct {
	TransactionID string `json:"transaction_id,omitempty"`
	OutOrderNo    string `json:"out_order_no,omitempty"`
	MchID         string `json:"mchid,omitempty"`
	Description   string `json:"description,omitempty"`
	NotifyURL     string `json:"notify_url,omitempty"`
}

type FinishSplitFundResponse struct {
	MchID         string `json:"mchid,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	OutOrderNo    string `json:"out_order_no,omitempty"`
	OrderID       string `json:"order_id,omitempty"`
}

type QueryUnSplitFundRequest struct {
	TransactionID string `json:"transaction_id,omitempty"`
	MchID         string `json:"mchid,omitempty"`
}

type QueryUnSplitFundResponse struct {
	MchID         string `json:"mchid,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	UnsplitAmount int    `json:"unsplit_amount,omitempty"`
}

type AddSplitReceiverRequest struct {
	MchID          string `json:"mchid,omitempty"`
	AppID          string `json:"appid,omitempty"`
	Type           string `json:"type,omitempty"`
	Account        string `json:"account,omitempty"`
	Name           string `json:"name,omitempty"`
	RelationType   string `json:"relation_type,omitempty"`
	CustomRelation string `json:"custom_relation,omitempty"`
}

type AddSplitReceiverResponse struct {
	Type           string `json:"type,omitempty"`
	Account        string `json:"account,omitempty"`
	Name           string `json:"name,omitempty"`
	RelationType   string `json:"relation_type,omitempty"`
	CustomRelation string `json:"custom_relation,omitempty"`
}

type DeleteSplitReceiverRequest struct {
	MchID   string `json:"mchid,omitempty"`
	AppID   string `json:"appid,omitempty"`
	Type    string `json:"type,omitempty"`
	Account string `json:"account,omitempty"`
}

type DeleteSplitReceiverResponse struct {
	MchID   string `json:"mchid,omitempty"`
	Type    string `json:"type,omitempty"`
	Account string `json:"account,omitempty"`
}
