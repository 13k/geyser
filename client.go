package geyser

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	baseURL      = "https://api.steampowered.com"
	pathTemplate = "/{interface}/{method}/{version}"
)

type Logger interface {
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

type ClientOption func(*Client) error

func WithLogger(logger Logger) ClientOption {
	return func(c *Client) error {
		c.client.SetLogger(logger)
		return nil
	}
}

func WithDebug() ClientOption {
	return func(c *Client) error {
		c.client.SetDebug(true)
		return nil
	}
}

func WithProxy(proxy string) ClientOption {
	return func(c *Client) error {
		c.client.SetProxy(proxy)
		return nil
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.client.SetHeader("User-Agent", userAgent)
		return nil
	}
}

func WithRetryCount(retryCount int) ClientOption {
	return func(c *Client) error {
		c.client.SetRetryCount(retryCount)
		return nil
	}
}

func WithRetryMaxWaitTime(retryMaxWaitTime time.Duration) ClientOption {
	return func(c *Client) error {
		c.client.SetRetryMaxWaitTime(retryMaxWaitTime)
		return nil
	}
}

func WithRetryWaitTime(retryWaitTime time.Duration) ClientOption {
	return func(c *Client) error {
		c.client.SetRetryWaitTime(retryWaitTime)
		return nil
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) error {
		c.client.SetTimeout(timeout)
		return nil
	}
}

func WithTransport(transport http.RoundTripper) ClientOption {
	return func(c *Client) error {
		c.client.SetTransport(transport)
		return nil
	}
}

func WithKey(key string) ClientOption {
	return func(c *Client) error {
		c.client.SetQueryParam("key", key)
		return nil
	}
}

func WithLanguage(lang string) ClientOption {
	return func(c *Client) error {
		c.client.SetQueryParam("language", lang)
		return nil
	}
}

type Client struct {
	client *resty.Client
}

func New(options ...ClientOption) (*Client, error) {
	rclient := resty.New().
		SetHostURL(baseURL).
		SetQueryParam("format", "json")

	client := &Client{client: rclient}

	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

type ClientRequest struct {
	Interface *SchemaInterface
	Method    *SchemaMethod
	Options   RequestOptions
	Result    interface{}
}

func (c *Client) Request(creq ClientRequest) (*resty.Response, error) {
	var params url.Values

	switch creq.Method.HTTPMethod {
	case http.MethodGet:
		params = creq.Options.Params
	case http.MethodPost, http.MethodPut:
		params = creq.Options.FormData
	}

	if params != nil {
		if err := creq.Method.ValidateParams(params); err != nil {
			return nil, err
		}
	}

	pathParams := map[string]string{
		"interface": creq.Interface.Name,
		"method":    creq.Method.Name,
		"version":   creq.Method.getVersionPathParam(),
	}

	req := c.client.R().
		SetContext(creq.Options.Context).
		SetPathParams(pathParams).
		SetQueryParamsFromValues(creq.Options.Params).
		SetFormDataFromValues(creq.Options.FormData).
		SetBody(creq.Options.Body)

	if creq.Result != nil {
		req.SetResult(creq.Result)
	}

	if creq.Options.Key != "" {
		req.SetQueryParam("key", creq.Options.Key)
	}

	if creq.Options.Lang != "" {
		req.SetQueryParam("language", creq.Options.Lang)
	}

	return req.Execute(creq.Method.HTTPMethod, pathTemplate)
}
