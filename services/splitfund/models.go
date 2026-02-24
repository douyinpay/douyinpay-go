package splitfund

// 分账接口
type ReceiverInfoDTO struct {
	Type        string `json:"type"`           // 可能是个人openid或者商户号，本期只支持商户号：MERCHANT_ID
	Account     string `json:"account"`        // 分账接收方账号类型是MERCHANT_ID时，是商户号
	Name        string `json:"name,omitempty"` // 分账接收方账号类型是个人时选填，为商户时无需填写
	Amount      int64  `json:"amount"`         //  分账金额（分）
	Description string `json:"description"`    // 分账的原因描述，分账账单中需要体现
}

type SplitFundReq struct {
	AppId           string             `json:"appid"`            // 商户应用号
	MerchantId      string             `json:"mchid"`            // 直连商户号，区分与追光内部的SmchId。追光内部商户SmchId为机构商户号，外部商户为特约商户号
	TradeNo         string             `json:"transaction_id"`   // 交易订单号
	OutOrderNo      string             `json:"out_order_no"`     // 外部商户单号，必填，非空
	Receivers       []*ReceiverInfoDTO `json:"receivers"`        // 分账方信息，必填
	UnfreezeUnsplit bool               `json:"unfreeze_unsplit"` // 能力等同是否完结分账
}

type SplitFundResp struct {
	TradeNo    string `json:"transaction_id"` // 交易单号
	OutOrderNo string `json:"out_order_no"`   // 外部商户单号
	OrderId    string `json:"order_id"`       // 抖音支付分账单号
}

// 分账查询接口
type QuerySplitFundReq struct {
	MerchantId string `json:"mchid"`                    // 直连商户号，区分与追光内部的SmchId。追光内部商户SmchId为机构商户号，外部商户为特约商户号
	TradeNo    string `json:"transaction_id"`           // 交易订单号
	OutOrderNo string `            json:"out_order_no"` // 外部商户单号
	OrderId    string `            json:"order_id"`     // 抖音支付分账单号
}

type QuerySplitFundResp struct {
	TradeNo                  string                    `json:"transaction_id,omitempty"` // 正向支付交易单号，必填，非空
	OutTradeNo               string                    `json:"out_order_no,omitempty"`   // 结算中心总单id，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OrderId                  string                    `json:"order_id,omitempty"`       // 结算中心请求id
	State                    string                    `json:"state,omitempty"`          // 分账单状态（每个接收方的分账结果请查看receivers中的result字段），枚举值：PROCESSING：处理中；FINISHED：分账完成
	ReceiversSplitResultList []*ReceiverSplitResultDTO `json:"receivers,omitempty"`      //分账接收方列表
}

type ReceiverSplitResultDTO struct {
	Amount      int64  `json:"amount"`                // 分账金额，单位为分，只能为整数，不能超过原订单支付金额及最大分账比例金额
	Description string `json:"description"`           // 分账的原因描述，分账账单中需要体现
	Type        string `json:"type"`                  // 1、MERCHANT_ID：商户号  2、PERSONAL_OPENID：个人openid（待确认）
	Account     string `json:"account"`               // 分账接收方类型为MERCHANT_ID时，分账接收方账号为商户号
	Result      string `json:"result"`                // 枚举值：PENDING：待分账  2、SUCCESS：分账成功  3、CLOSED：已关闭
	FailReason  string `json:"fail_reason"`           // 分账失败原因，当分账结果result为CLOSED（已关闭）时，返回该字段
	CreateTime  string `json:"create_time"`           // 分账创建时间，例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒
	FinishTime  string `json:"finish_time,omitempty"` // 分账完成时间
	DetailId    string `json:"detail_id"`             // 抖音支付分账明细单号，每笔分账业务执行的明细单号，可与资金账单对账使用
}
