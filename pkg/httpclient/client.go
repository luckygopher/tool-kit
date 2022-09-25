package httpclient

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type Client struct {
	config      Config
	logger      *zap.Logger
	restyClient *resty.Client
}

var HTTPClient *Client

func InitHTTPClient(conf Config, logger *zap.Logger) {
	HTTPClient = NewClient(conf, logger)
}

func NewClient(config Config, logger *zap.Logger) *Client {
	restyClient = GetRestyClient(config)
	if config.Verbose {
		restyClient.EnableTrace()
	}
	restyClient.SetDebug(config.Debug)

	return &Client{
		config:      config,
		logger:      logger,
		restyClient: restyClient,
	}
}

func (c Client) Post(ctx context.Context, url string, headers map[string]string, body interface{}, result interface{}) ([]byte, error) {
	res := make([]byte, 0)
	headers["Content-Type"] = "application/json;charset=UTF-8"
	r, err := c.restyClient.R().SetContext(ctx).SetBody(body).SetHeaders(headers).Post(url)
	if err != nil {
		c.logger.Error("send request failed", zap.Error(err))
		return res, err
	}

	if r.StatusCode() != http.StatusCreated && r.StatusCode() != http.StatusOK {
		c.logger.Error("rsp status_code not equal 200 or 201",
			zap.String("rsp.Body", string(r.Body())),
			zap.Int("rsp.StatusCode", r.StatusCode()))
		return res, errors.New(string(r.Body()))
	}

	if err := json.Unmarshal(r.Body(), &result); err != nil {
		c.logger.Error(" json unmarshal failed", zap.Error(err), zap.String("response.body", string(r.Body())))
		return res, err
	}
	return r.Body(), nil
}

func (c Client) Get(ctx context.Context, url string, headers map[string]string, pathParams, queryParams map[string]string, result interface{}) ([]byte, error) {
	res := make([]byte, 0)
	r, err := c.restyClient.R().SetContext(ctx).SetPathParams(pathParams).SetQueryParams(queryParams).SetHeaders(headers).Get(url)
	if err != nil {
		c.logger.Error("send request failed", zap.Error(err))
		return res, err
	}

	if r.StatusCode() != http.StatusCreated && r.StatusCode() != http.StatusOK {
		c.logger.Error("rsp status_code not equal 200 or 201",
			zap.String("rsp.Body", string(r.Body())),
			zap.Int("rsp.StatusCode", r.StatusCode()))
		return res, errors.New(string(r.Body()))
	}

	if err := json.Unmarshal(r.Body(), &result); err != nil {
		c.logger.Error(" json unmarshal failed", zap.Error(err), zap.String("response.body", string(r.Body())))
		return res, err
	}
	return r.Body(), nil
}

func (c Client) PostForm(ctx context.Context, url string, param map[string]string, result interface{}) ([]byte, error) {
	res := make([]byte, 0)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	r, err := c.restyClient.R().SetContext(ctx).SetFormData(param).SetHeaders(headers).Post(url)
	if err != nil {
		c.logger.Error("send request failed", zap.Error(err))
		return res, err
	}
	if r.StatusCode() != http.StatusCreated && r.StatusCode() != http.StatusOK {
		c.logger.Error("rsp status_code not equal 200 or 201",
			zap.String("rsp.Body", string(r.Body())),
			zap.Int("rsp.StatusCode", r.StatusCode()))
		return res, errors.New(string(r.Body()))
	}
	if err := json.Unmarshal(r.Body(), &result); err != nil {
		c.logger.Error(" json unmarshal failed", zap.Error(err), zap.String("response.body", string(r.Body())))
		return res, err
	}
	return r.Body(), nil
}

func (c Client) SetProxy(proxyUrl string) {
	c.restyClient.SetProxy(proxyUrl)
}

func (c Client) EncodeWithSha256(rawStr string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(rawStr))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// JoinStringsInASCII 排序
// data 排序内容
// sep 连接符
// includeEmpty 是否包含空值，true则包含空值，否则不包含，注意此参数不影响参数名的存在
// exceptKeys 被排除的参数名，不参与排序及拼接
func (c Client) JoinStringsInASCII(data map[string]interface{}, sep string, includeEmpty bool, exceptKeys ...string) string {
	var list []string
	m := make(map[string]int)
	if len(exceptKeys) > 0 {
		for _, except := range exceptKeys {
			m[except] = 1
		}
	}
	for k := range data {
		if _, ok := m[k]; ok {
			continue
		}
		value := data[k]
		if !includeEmpty && value == "" {
			continue
		}
		list = append(list, fmt.Sprintf("%s=%v", k, value))
	}
	sort.Strings(list)
	return strings.Join(list, sep)
}

// ToMap 转map
func (c Client) ToMap(context interface{}) (map[string]interface{}, error) {
	var (
		res map[string]interface{}
		r   []byte
		err error
	)
	if r, err = json.Marshal(context); err != nil {
		return nil, err
	}
	d := json.NewDecoder(bytes.NewReader(r))
	d.UseNumber()
	if err := d.Decode(&res); err != nil {
		return nil, err
	}
	for k, v := range res {
		res[k] = v
	}
	return res, nil
}
