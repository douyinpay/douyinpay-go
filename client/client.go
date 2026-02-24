package client

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/douyinpay/douyinpay-go/tools/auth"
	"github.com/douyinpay/douyinpay-go/tools/consts"
)

var (
	regJSONTypeCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	regXMLTypeCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

type ClientOption interface {
	Apply(settings *DialSettings) error
}

// ErrorOption 错误初始化参数，用于返回错误
type ErrorOption struct{ Error error }

// Apply 返回初始化错误
func (w ErrorOption) Apply(*DialSettings) error {
	return w.Error
}

type DialSettings struct {
	AgentName  string
	HTTPClient *http.Client  // HTTPClient 实例
	Signer     auth.Signer   // 加签
	Verifier   auth.Verifier // 验签
}

// Validate 校验请求配置是否有效
func (ds *DialSettings) Validate() error {
	if ds.Signer == nil {
		return fmt.Errorf("Signer is required for Client")
	}
	if ds.Verifier == nil {
		return fmt.Errorf("Verifier is required for Client")
	}
	return nil
}

// http请求参数
type APIResult struct {
	Request  *http.Request
	Response *http.Response
}

type Client struct {
	AgentName  string
	httpClient *http.Client
	signer     auth.Signer
	verifier   auth.Verifier
}

func NewClient(ctx context.Context, opts ...ClientOption) (*Client, error) {
	settings, err := initSettings(opts)
	if err != nil {
		return nil, fmt.Errorf("init client setting err:%v", err)
	}

	client := &Client{
		AgentName:  settings.AgentName,
		signer:     settings.Signer,
		httpClient: settings.HTTPClient,
		verifier:   settings.Verifier,
	}

	if client.httpClient == nil {
		client.httpClient = &http.Client{
			Timeout: consts.DefaultTimeout,
		}
	}
	return client, nil
}

func NewClientWithDialSettings(ctx context.Context, settings *DialSettings) (*Client, error) {
	if err := settings.Validate(); err != nil {
		return nil, err
	}

	client := initClientWithSettings(ctx, settings)
	return client, nil
}

func NewClientWithVerifier(client *Client, verifier auth.Verifier) *Client {
	return &Client{
		signer:     client.signer,
		httpClient: client.httpClient,
		verifier:   verifier,
	}
}

func initClientWithSettings(_ context.Context, settings *DialSettings) *Client {
	client := &Client{
		AgentName:  settings.AgentName,
		signer:     settings.Signer,
		httpClient: settings.HTTPClient,
		verifier:   settings.Verifier,
	}
	if client.httpClient == nil {
		client.httpClient = &http.Client{
			Timeout: consts.DefaultTimeout,
		}
	}
	return client
}

func initSettings(opts []ClientOption) (*DialSettings, error) {
	var (
		o   DialSettings
		err error
	)
	for _, opt := range opts {
		if err = opt.Apply(&o); err != nil {
			return nil, err
		}
	}
	if err := o.Validate(); err != nil {
		return nil, err
	}
	return &o, nil
}

// 发送http请求
type RequestEntity struct {
	Method      string
	RequestPath string
	Header      http.Header
	QueryParams url.Values
	PostBody    interface{}
	ContentType string
}

func (client *Client) Request(ctx context.Context, r *RequestEntity) (*APIResult, error) {

	pathWithQueryParam, err := SetPathWithParams(r.RequestPath, r.QueryParams)
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	var signBody string
	if r.PostBody != nil {
		if r.ContentType == "" {
			r.ContentType = consts.ApplicationJSON
		}

		body, err = setBody(r.PostBody, r.ContentType)
		if err != nil {
			return nil, err
		}
		signBody = body.String()
	}

	// Construct Request
	var request *http.Request
	if request, err = http.NewRequestWithContext(ctx, r.Method, pathWithQueryParam, body); err != nil {
		return nil, err
	}

	// Add Request Header
	for key, values := range r.Header {
		for _, v := range values {
			request.Header.Add(key, v)
		}
	}
	request.Header.Set(consts.ContentType, r.ContentType)

	// Set Authentication
	authorization, err := client.GenerateAuthorizationHeader(ctx, r.Method, request.URL.RequestURI(), signBody)
	if err != nil {
		return nil, fmt.Errorf("generate authorization err:%s", err.Error())
	}
	request.Header.Set(consts.Authorization, authorization)

	// indicate Douyinpay-Serial that client can verify
	//if serial, err := client.verifier.GetSerial(ctx); err == nil {
	//	request.Header.Set(consts.DouyinpaySerial, serial)
	//}

	if client.AgentName != "" {
		request.Header.Set(consts.DouyinpaySdkAgent, client.AgentName)
	}

	// Send HTTP Request
	result := &APIResult{
		Request: request,
	}
	result.Response, err = client.httpClient.Do(request)
	if err != nil {
		return result, err
	}
	// Check Response
	if err = CheckResponse(result.Response); err != nil {
		return result, err
	}
	// Validate Signature
	if err = client.Validate(ctx, result.Response); err != nil {
		return result, err
	}
	return result, nil
}

// CheckResponse 校验请求是否成功
// 当http回包的状态码的范围不是200-299之间的时候，会返回相应的错误信息，主要包括http状态码、回包错误码、回包错误信息提示
func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}
	slurp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("invalid response, read body error: %w", err)
	}
	_ = resp.Body.Close()

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(slurp))
	apiError := &APIError{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(slurp),
	}
	// 忽略 JSON 解析错误，均返回 apiError
	_ = json.Unmarshal(slurp, apiError)
	return apiError
}

// UnMarshalResponse 将回包组织成结构化数据
func UnMarshalResponse(httpResp *http.Response, resp interface{}) error {
	body, err := ioutil.ReadAll(httpResp.Body)
	_ = httpResp.Body.Close()

	if err != nil {
		return err
	}

	httpResp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}
	return nil
}

// setBody Set Request body from an interface
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	bodyBuf = &bytes.Buffer{}

	switch b := body.(type) {
	case string:
		_, err = bodyBuf.WriteString(b)
	case *string:
		_, err = bodyBuf.WriteString(*b)
	case []byte:
		_, err = bodyBuf.Write(b)
	case **os.File:
		_, err = bodyBuf.ReadFrom(*b)
	case io.Reader:
		_, err = bodyBuf.ReadFrom(b)
	default:
		if regJSONTypeCheck.MatchString(contentType) {
			err = json.NewEncoder(bodyBuf).Encode(body)
		} else if regXMLTypeCheck.MatchString(contentType) {
			err = xml.NewEncoder(bodyBuf).Encode(body)
		}
	}
	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

func SetPathWithParams(path string, param url.Values) (string, error) {
	varPath, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	// Adding Query Param
	query := varPath.Query()
	for k, values := range param {
		for _, v := range values {
			query.Add(k, v)
		}
	}

	// Encode the parameters.
	varPath.RawQuery = query.Encode()
	return varPath.String(), nil
}
