package splitfund

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/tools/crypto"
)

type ApiService services.Service

// SplitFund 分账
func (a *ApiService) SplitFund(ctx context.Context, req SplitFundRequest) (resp *SplitFundResponse, result *client.APIResult, err error) {
	if err = a.encryptSplitFundReceiverNames(ctx, &req); err != nil {
		return nil, nil, err
	}
	headers, err := a.buildPlatformCertificateSerialHeaders(ctx)
	if err != nil {
		return nil, nil, err
	}

	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      headers,
	}
	r.RequestPath = consts.DouyinPayServer + consts.SplitFundPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	resp = new(SplitFundResponse)
	if err = client.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QuerySplitFund 分账查询
func (a *ApiService) QuerySplitFund(ctx context.Context, req QuerySplitFundRequest) (resp *QuerySplitFundResponse, result *client.APIResult, err error) {
	if req.OutOrderNo == "" {
		return nil, nil, fmt.Errorf("QuerySplitFundRequest required field `OutOrderNo` is empty")
	}
	localVarQueryParams := neturl.Values{}
	if req.MchID != "" {
		localVarQueryParams.Add("mchid", req.MchID)
	}
	if req.TransactionID != "" {
		localVarQueryParams.Add("transaction_id", req.TransactionID)
	}
	if req.OrderID != "" {
		localVarQueryParams.Add("order_id", req.OrderID)
	}

	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
		QueryParams: localVarQueryParams,
	}
	rawUrl := consts.DouyinPayServer + consts.QuerySplitFundPath
	r.RequestPath = strings.Replace(rawUrl, "{out_trade_no}", neturl.PathEscape(req.OutOrderNo), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	resp = new(QuerySplitFundResponse)
	if err = client.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ReturnSplitFund 分账回退
func (a *ApiService) ReturnSplitFund(ctx context.Context, req ReturnSplitFundRequest) (resp *ReturnSplitFundResponse, result *client.APIResult, err error) {
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.ReturnSplitFundPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	resp = new(ReturnSplitFundResponse)
	if err = client.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryReturnSplitFund 分账回退查询
func (a *ApiService) QueryReturnSplitFund(ctx context.Context, req QueryReturnSplitFundRequest) (resp *QueryReturnSplitFundResponse, result *client.APIResult, err error) {
	if req.OutReturnNo == "" {
		return nil, nil, fmt.Errorf("QueryReturnSplitFundRequest required field `OutReturnNo` is empty")
	}
	localVarQueryParams := neturl.Values{}
	if req.MchID != "" {
		localVarQueryParams.Add("mchid", req.MchID)
	}
	if req.OutOrderNo != "" {
		localVarQueryParams.Add("out_order_no", req.OutOrderNo)
	}

	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
		QueryParams: localVarQueryParams,
	}
	rawUrl := consts.DouyinPayServer + consts.QueryReturnSplitFundPath
	r.RequestPath = strings.Replace(rawUrl, "{out_return_no}", neturl.PathEscape(req.OutReturnNo), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	resp = new(QueryReturnSplitFundResponse)
	if err = client.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// FinishSplitFund 完结分账
func (a *ApiService) FinishSplitFund(ctx context.Context, req FinishSplitFundRequest) (resp *FinishSplitFundResponse, result *client.APIResult, err error) {
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.FinishSplitFundPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	resp = new(FinishSplitFundResponse)
	if err = client.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryUnSplitFund 查询剩余待分金额
func (a *ApiService) QueryUnSplitFund(ctx context.Context, req QueryUnSplitFundRequest) (resp *QueryUnSplitFundResponse, result *client.APIResult, err error) {
	if req.TransactionID == "" {
		return nil, nil, fmt.Errorf("QueryUnSplitFundRequest required field `TransactionID` is empty")
	}
	localVarQueryParams := neturl.Values{}
	if req.MchID != "" {
		localVarQueryParams.Add("mchid", req.MchID)
	}
	r := &client.RequestEntity{
		Method:      nethttp.MethodGet,
		ContentType: consts.ApplicationJSON,
		Header:      nethttp.Header{},
		QueryParams: localVarQueryParams,
	}
	rawUrl := consts.DouyinPayServer + consts.QueryUnsplitFundPath
	r.RequestPath = strings.Replace(rawUrl, "{transaction_id}", neturl.PathEscape(req.TransactionID), -1)

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	resp = new(QueryUnSplitFundResponse)
	if err = client.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// AddSplitReceiver 添加分账接收方
func (a *ApiService) AddSplitReceiver(ctx context.Context, req AddSplitReceiverRequest) (resp *AddSplitReceiverResponse, result *client.APIResult, err error) {
	if req.Name != "" {
		encryptedName, err := a.encryptSensitiveName(ctx, req.Name)
		if err != nil {
			return nil, nil, err
		}
		req.Name = encryptedName
	}
	headers, err := a.buildPlatformCertificateSerialHeaders(ctx)
	if err != nil {
		return nil, nil, err
	}

	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      headers,
	}
	r.RequestPath = consts.DouyinPayServer + consts.AddSplitReceiverPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	resp = new(AddSplitReceiverResponse)
	if err = client.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, result, err
	}
	if resp != nil && resp.Name != "" {
		decryptedName, err := a.decryptSensitiveName(ctx, resp.Name)
		if err != nil {
			return nil, result, err
		}
		resp.Name = decryptedName
	}
	return resp, result, nil
}

// DeleteSplitReceiver 删除分账接收方
func (a *ApiService) DeleteSplitReceiver(ctx context.Context, req DeleteSplitReceiverRequest) (resp *DeleteSplitReceiverResponse, result *client.APIResult, err error) {
	r := &client.RequestEntity{
		Method:      nethttp.MethodPost,
		ContentType: consts.ApplicationJSON,
		PostBody:    req,
		Header:      nethttp.Header{},
	}
	r.RequestPath = consts.DouyinPayServer + consts.DeleteSplitReceiverPath

	result, err = a.Client.Request(ctx, r)
	if err != nil {
		return nil, result, err
	}
	resp = new(DeleteSplitReceiverResponse)
	if err = client.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ApiService) encryptSplitFundReceiverNames(ctx context.Context, req *SplitFundRequest) error {
	if req == nil || len(req.Receivers) == 0 {
		return nil
	}
	for index := range req.Receivers {
		if req.Receivers[index].Name == "" {
			continue
		}
		encryptedName, err := a.encryptSensitiveName(ctx, req.Receivers[index].Name)
		if err != nil {
			return err
		}
		req.Receivers[index].Name = encryptedName
	}
	return nil
}

func (a *ApiService) encryptSensitiveName(ctx context.Context, name string) (string, error) {
	if name == "" {
		return name, nil
	}
	encryptor, err := a.getEncryptor()
	if err != nil {
		return "", err
	}
	return encryptor.Encrypt(ctx, name)
}

func (a *ApiService) decryptSensitiveName(ctx context.Context, ciphertext string) (string, error) {
	if ciphertext == "" {
		return ciphertext, nil
	}
	decryptor, err := a.getDecryptor()
	if err != nil {
		return "", err
	}
	return decryptor.Decrypt(ctx, ciphertext)
}

func (a *ApiService) buildPlatformCertificateSerialHeaders(ctx context.Context) (nethttp.Header, error) {
	encryptor, err := a.getEncryptor()
	if err != nil {
		return nil, err
	}
	serial, err := encryptor.GetPlatformSerial(ctx)
	if err != nil {
		return nil, err
	}
	if serial == "" {
		return nil, fmt.Errorf("platform certificate serial is empty")
	}
	headers := nethttp.Header{}
	headers.Set(consts.DouyinpaySerial, serial)
	return headers, nil
}

func (a *ApiService) getEncryptor() (crypto.Encryptor, error) {
	if a.Client == nil || a.Client.Encryptor() == nil {
		return nil, fmt.Errorf("client has no encryptor configured")
	}
	return a.Client.Encryptor(), nil
}

func (a *ApiService) getDecryptor() (crypto.Decryptor, error) {
	if a.Client == nil || a.Client.Decryptor() == nil {
		return nil, fmt.Errorf("client has no decryptor configured")
	}
	return a.Client.Decryptor(), nil
}
