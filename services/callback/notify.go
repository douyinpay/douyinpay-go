package callback

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/tools/auth"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"
)

// Handler 回调通知 Handler
type Handler struct {
	encryptKey string
	verifier   auth.Verifier
}

// NewNotifyHandler 创建通知处理器
func NewNotifyHandler(key string, verifier auth.Verifier) *Handler {
	return &Handler{
		encryptKey: key,
		verifier:   verifier,
	}
}

func (h *Handler) ParseCallback(ctx context.Context, request *http.Request, content interface{}) (*Request, error) {
	//获取body
	body, header, err := GetRequestBodyAndHeader(request)
	if err != nil {
		return nil, err
	}
	//验签
	if err = h.validateHTTPMessage(ctx, header, body); err != nil {
		return nil, fmt.Errorf("invalid notify request: %v", err)
	}
	//反序列化
	ret := new(Request)
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, fmt.Errorf("parse request body error: %v", err)
	}
	//解密
	rsc := ret.Resource
	var plaintext string
	if rsc.Algorithm == consts.EncryptTypeAES256GCM {
		plaintext, err = utils.DecryptAES256GCM(h.encryptKey, rsc.AssociatedData, rsc.Nonce, rsc.Ciphertext)
		if err != nil {
			return ret, fmt.Errorf("decrypt request error: %v", err)
		}
	} else {
		plaintext, err = utils.DecryptSM4(h.encryptKey, rsc.AssociatedData, rsc.Nonce, rsc.Ciphertext)
		if err != nil {
			return ret, fmt.Errorf("decrypt request error: %v", err)
		}
	}
	ret.Resource.Plaintext = plaintext

	if err = json.Unmarshal([]byte(plaintext), &content); err != nil {
		return ret, fmt.Errorf("unmarshal plaintext to content failed: %v", err)
	}

	return ret, nil
}

func GetRequestBodyAndHeader(request *http.Request) ([]byte, *client.ResponseHeader, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("read request body err: %v", err)
	}

	_ = request.Body.Close()
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	header, err := client.GetHeader(request.Header)
	if err != nil {
		return nil, nil, fmt.Errorf("read request header err: %v", err)
	}

	return body, &header, nil
}

func (h *Handler) validateHTTPMessage(ctx context.Context, header *client.ResponseHeader, body []byte) error {
	if h.verifier == nil {
		return fmt.Errorf("please init client with Verifier")
	}

	if header == nil {
		return fmt.Errorf("header is nil")
	}

	//if math.Abs(float64(time.Now().Unix()-header.Timestamp)) >= consts.TwentyFourHourAndFourMinute {
	//	return fmt.Errorf("timestamp=[%d] expires, request-id=[%s]", header.Timestamp, header.RequestID)
	//}

	message := fmt.Sprintf(consts.VerifyMessageFormat, header.Timestamp, header.Nonce, string(body))

	if err := h.verifier.Verify(ctx, header.Serial, message, header.Signature); err != nil {
		return fmt.Errorf(
			"validate verify fail serial=[%s] request-id=[%s] err=%w",
			header.Serial, header.RequestID, err,
		)
	}
	return nil
}
