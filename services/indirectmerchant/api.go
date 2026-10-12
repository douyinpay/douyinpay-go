package indirectmerchant

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/services/callback"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

type ApiService services.Service

// ApiIndirectMerchantService 是间连商户服务的描述性别名。
type ApiIndirectMerchantService = ApiService

// base 为间连商户接口的公共前缀，与其他服务一致需携带 API 主域名。
const base = consts.DouyinPayServer + "/v1/merchant/indirect_merchant/"

// uploadPath 为图片上传接口的完整请求地址。
const uploadPath = consts.DouyinPayServer + "/v1/merchant/upload"

// notifyEventType 间连商户申请状态通知的事件类型。
const notifyEventType = "INDIRECT_MERCHANT_APPLYMENT_STATUS"

func (a *ApiService) request(ctx context.Context, method, path string, q url.Values, req, resp interface{}) (*client.APIResult, error) {
	var body interface{} = req
	if req == nil {
		body = nil
	}
	result, err := (*services.Service)(a).Client.Request(ctx, &client.RequestEntity{Method: method, RequestPath: path, QueryParams: q, PostBody: body})
	if err != nil {
		return result, err
	}
	if resp != nil {
		err = client.UnMarshalResponse(result.Response, resp)
	}
	return result, err
}

// CreateApply 提交间连商户入驻申请（create_apply）。
// 证件持有人姓名、证件号码为敏感字段，发送前使用平台公钥加密；加密作用于请求副本，不修改调用方传入的对象。
func (a *ApiService) CreateApply(ctx context.Context, req ApplyRequest) (*ApplyResponse, *client.APIResult, error) {
	if err := req.Validate(); err != nil {
		return nil, nil, err
	}
	// 对敏感身份信息做副本加密，保持调用方传入的 req.IdentificationInfo 不变。
	if req.IdentificationInfo != nil {
		encrypted, err := a.encryptIdentification(*req.IdentificationInfo)
		if err != nil {
			return nil, nil, err
		}
		req.IdentificationInfo = &encrypted
	}
	out := new(ApplyResponse)
	r, e := a.request(ctx, http.MethodPost, base+"create_apply", nil, req, out)
	return out, r, e
}

// encryptIdentification 使用平台公钥加密证件持有人敏感字段（姓名、证件号码），返回加密后的副本。
// 入参按值传递，不修改调用方对象；当字段为空时跳过，当密钥未配置或加密失败时返回带上下文的错误。
func (a *ApiService) encryptIdentification(info IdentificationInfo) (IdentificationInfo, error) {
	cli := (*services.Service)(a).Client
	if info.IdentificationName != "" {
		cipher, err := cli.EncryptSensitive(info.IdentificationName)
		if err != nil {
			return info, fmt.Errorf("encrypt identification_name: %s", err.Error())
		}
		info.IdentificationName = cipher
	}
	if info.IdentificationNumber != "" {
		cipher, err := cli.EncryptSensitive(info.IdentificationNumber)
		if err != nil {
			return info, fmt.Errorf("encrypt identification_number: %s", err.Error())
		}
		info.IdentificationNumber = cipher
	}
	return info, nil
}

// decryptIdentification 使用服务商私钥原位解密下行证件持有人敏感字段（姓名、证件号码）。
// 当字段为空时跳过，当密钥未配置或密文非法时返回带上下文的错误。
func (a *ApiService) decryptIdentification(info *IdentificationInfo) error {
	cli := (*services.Service)(a).Client
	if info.IdentificationName != "" {
		plain, err := cli.DecryptSensitive(info.IdentificationName)
		if err != nil {
			return fmt.Errorf("decrypt identification_name: %s", err.Error())
		}
		info.IdentificationName = plain
	}
	if info.IdentificationNumber != "" {
		plain, err := cli.DecryptSensitive(info.IdentificationNumber)
		if err != nil {
			return fmt.Errorf("decrypt identification_number: %s", err.Error())
		}
		info.IdentificationNumber = plain
	}
	return nil
}

// QueryApply 查询入驻申请单（query_apply），申请单号或业务申请编号二选一。
func (a *ApiService) QueryApply(ctx context.Context, req QueryApplyRequest) (*QueryApplyResponse, *client.APIResult, error) {
	q := url.Values{}
	if req.ApplymentID != "" {
		q.Set("applyment_id", req.ApplymentID)
	}
	if req.BusinessCode != "" {
		q.Set("business_code", req.BusinessCode)
	}
	out := new(QueryApplyResponse)
	r, e := a.request(ctx, http.MethodGet, base+"query_apply", q, nil, out)
	return out, r, e
}

// CancelApply 撤销入驻申请单（cancel_apply），仅在待提交和审核中可撤销。
func (a *ApiService) CancelApply(ctx context.Context, req CancelRequest) (*client.APIResult, error) {
	r, e := a.request(ctx, http.MethodPost, base+"cancel_apply", nil, req, nil)
	return r, e
}

// QueryMerchant 查询间连商户信息（query_merchant）。
// 响应中的证件持有人姓名、证件号码为抖音支付加密下发字段，使用服务商私钥解密后回填。
func (a *ApiService) QueryMerchant(ctx context.Context, req Locator) (*QueryMerchantResponse, *client.APIResult, error) {
	if e := req.Validate(); e != nil {
		return nil, nil, e
	}
	q := url.Values{}
	if req.OutIndirectMerchantID != "" {
		q.Set("out_indirect_merchant_id", req.OutIndirectMerchantID)
	}
	if req.OutPaymentInstID != "" {
		q.Set("out_payment_inst_id", req.OutPaymentInstID)
	}
	if req.OutChannelSPID != "" {
		q.Set("out_channel_sp_id", req.OutChannelSPID)
	}
	if req.IndirectMerchantID != "" {
		q.Set("indirect_merchant_id", req.IndirectMerchantID)
	}
	out := new(QueryMerchantResponse)
	r, e := a.request(ctx, http.MethodGet, base+"query_merchant", q, nil, out)
	if e != nil {
		return out, r, e
	}
	if out.IdentificationInfo != nil {
		if err := a.decryptIdentification(out.IdentificationInfo); err != nil {
			return out, r, err
		}
	}
	return out, r, e
}

// Modify 修改商户资料（modify），目前仅支持修改商户简称和客服电话。
func (a *ApiService) Modify(ctx context.Context, req ModifyRequest) (*client.APIResult, error) {
	if e := req.Locator.Validate(); e != nil {
		return nil, e
	}
	r, e := a.request(ctx, http.MethodPost, base+"modify", nil, req, nil)
	return r, e
}

// BindAppID 为间连子商户绑定应用 APPID（bind_app_id），仅支持新增配置。
func (a *ApiService) BindAppID(ctx context.Context, req BindAppRequest) (*BindAppResponse, *client.APIResult, error) {
	if e := req.Locator.Validate(); e != nil {
		return nil, nil, e
	}
	out := new(BindAppResponse)
	r, e := a.request(ctx, http.MethodPost, base+"bind_app_id", nil, req, out)
	return out, r, e
}

// QueryAppID 查询间连子商户配置的应用 APPID 信息（query_app_id），可按 AppID 过滤。
func (a *ApiService) QueryAppID(ctx context.Context, req QueryAppRequest) (*QueryAppResponse, *client.APIResult, error) {
	if e := req.Locator.Validate(); e != nil {
		return nil, nil, e
	}
	q := url.Values{}
	if req.OutIndirectMerchantID != "" {
		q.Set("out_indirect_merchant_id", req.OutIndirectMerchantID)
	}
	if req.OutPaymentInstID != "" {
		q.Set("out_payment_inst_id", req.OutPaymentInstID)
	}
	if req.OutChannelSPID != "" {
		q.Set("out_channel_sp_id", req.OutChannelSPID)
	}
	if req.IndirectMerchantID != "" {
		q.Set("indirect_merchant_id", req.IndirectMerchantID)
	}
	if req.AppID != "" {
		q.Set("app_id", req.AppID)
	}
	out := new(QueryAppResponse)
	r, e := a.request(ctx, http.MethodGet, base+"query_app_id", q, nil, out)
	return out, r, e
}

// Upload 上传商户资质图片（/v1/merchant/upload），返回图片 MediaID。
// 请求为 multipart/form-data，参与签名计算的仅为 meta 的 JSON 串。
func (a *ApiService) Upload(ctx context.Context, filename, mime string, file []byte) (*UploadResponse, *client.APIResult, error) {
	if len(file) == 0 {
		return nil, nil, fmt.Errorf("file is empty")
	}
	sum := sha256.Sum256(file)
	meta := UploadMeta{Filename: filename, SHA256: hex.EncodeToString(sum[:])}
	b, err := json.Marshal(meta)
	if err != nil {
		return nil, nil, fmt.Errorf("marshal upload meta: %s", err.Error())
	}
	result, e := (*services.Service)(a).Client.RequestMultipart(ctx, &client.MultipartRequestEntity{
		Method:      http.MethodPost,
		RequestPath: uploadPath,
		SignBody:    string(b),
		Meta:        string(b),
		FileName:    filename,
		FileType:    mime,
		File:        file,
	})
	out := new(UploadResponse)
	if e == nil {
		e = client.UnMarshalResponse(result.Response, out)
	}
	return out, result, e
}

// ParseApplyStatusNotify 验签解密间连商户申请状态通知（event_type=INDIRECT_MERCHANT_APPLYMENT_STATUS）。
// 先由 callback.Handler 完成验签和 AES-256-GCM 解密，再将明文反序列化为业务数据。
func ParseApplyStatusNotify(ctx context.Context, handler *callback.Handler, request *http.Request) (*callback.Request, *ApplyStatusNotify, error) {
	if handler == nil {
		return nil, nil, fmt.Errorf("handler is required")
	}
	out := new(ApplyStatusNotify)
	n, e := handler.ParseCallback(ctx, request, out)
	if e != nil {
		return n, nil, e
	}
	if n.EventType != "" && n.EventType != notifyEventType {
		return n, nil, fmt.Errorf("unexpected event_type %s", n.EventType)
	}
	return n, out, nil
}
