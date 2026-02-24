package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/douyinpay/douyinpay-go/tools/consts"
)

type ResponseHeader struct {
	RequestID string
	Serial    string
	Signature string
	Nonce     string
	Timestamp int64
}

func (v *Client) Validate(ctx context.Context, response *http.Response) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("read response body err:[%s]", err.Error())
	}
	response.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return v.validateHTTPMessage(ctx, response.Header, body)
}

func (c *Client) validateHTTPMessage(ctx context.Context, header http.Header, body []byte) error {
	if c.verifier == nil {
		return fmt.Errorf("please init client with Verifier")
	}

	headerArgs, err := GetHeader(header)
	if err != nil {
		return err
	}

	if math.Abs(float64(time.Now().Unix()-headerArgs.Timestamp)) >= consts.FiveMinute {
		return fmt.Errorf("timestamp=[%d] expires, request-id=[%s]", headerArgs.Timestamp, headerArgs.RequestID)
	}

	message := fmt.Sprintf("%d\n%s\n%s\n", headerArgs.Timestamp, headerArgs.Nonce, string(body))

	if err := c.verifier.Verify(ctx, headerArgs.Serial, message, headerArgs.Signature); err != nil {
		return fmt.Errorf(
			"validate verify fail serial=[%s] request-id=[%s] err=%w",
			headerArgs.Serial, headerArgs.RequestID, err,
		)
	}
	return nil
}

func GetHeader(header http.Header) (ResponseHeader, error) {

	requestID := strings.TrimSpace(header.Get(consts.RequestID))

	getHeaderString := func(key string) (string, error) {
		val := strings.TrimSpace(header.Get(key))
		if val == "" {
			return "", fmt.Errorf("key `%s` is empty in header, request-id=[%s]", key, requestID)
		}
		return val, nil
	}

	getHeaderInt64 := func(key string) (int64, error) {
		val, err := getHeaderString(key)
		if err != nil {
			return 0, nil
		}
		ret, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid `%s` in header, request-id=[%s], err:%w", key, requestID, err)
		}
		return ret, nil
	}

	ret := ResponseHeader{
		RequestID: requestID,
	}
	var err error

	if ret.Serial, err = getHeaderString(consts.DouyinPaySerial); err != nil {
		return ret, err
	}

	if ret.Signature, err = getHeaderString(consts.DouyinPaySignature); err != nil {
		return ret, err
	}

	if ret.Timestamp, err = getHeaderInt64(consts.DouyinPayTimestamp); err != nil {
		return ret, err
	}

	if ret.Nonce, err = getHeaderString(consts.DouyinPayNonce); err != nil {
		return ret, err
	}

	return ret, nil
}
