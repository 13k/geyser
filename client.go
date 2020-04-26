package geyser

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/13k/geyser/schema"
)

const (
	pathTemplate = "/{interface}/{method}/{version}"
)

// Logger is the logger interface type accepted by `Client`.
type Logger interface {
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

// ClientOption functions set options in `Client`.
type ClientOption func(*Client) error

// WithBaseURL sets the base URL for all outgoing requests.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		c.client.SetHostURL(baseURL)
		return nil
	}
}

// WithLogger sets the logger.
func WithLogger(logger Logger) ClientOption {
	return func(c *Client) error {
		c.client.SetLogger(logger)
		return nil
	}
}

// WithDebug enables debug logging.
func WithDebug() ClientOption {
	return func(c *Client) error {
		c.client.SetDebug(true)
		return nil
	}
}

// WithProxy sets the proxy.
func WithProxy(proxy string) ClientOption {
	return func(c *Client) error {
		c.client.SetProxy(proxy)
		return nil
	}
}

// WithUserAgent sets the "User-Agent" HTTP header for all outgoing requests.
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.client.SetHeader("User-Agent", userAgent)
		return nil
	}
}

// WithRetryCount enables request retrying and allows to set the number of
// tries, using a backoff mechanism.
//
// Setting to 0 disables retrying.
func WithRetryCount(retryCount int) ClientOption {
	return func(c *Client) error {
		c.client.SetRetryCount(retryCount)
		return nil
	}
}

// WithRetryMaxWaitTime sets the max wait time to sleep before retrying request.
//
// Default is 2s.
func WithRetryMaxWaitTime(retryMaxWaitTime time.Duration) ClientOption {
	return func(c *Client) error {
		c.client.SetRetryMaxWaitTime(retryMaxWaitTime)
		return nil
	}
}

// WithRetryWaitTime sets the default wait time to sleep before retrying request.
//
// Default is 100ms.
func WithRetryWaitTime(retryWaitTime time.Duration) ClientOption {
	return func(c *Client) error {
		c.client.SetRetryWaitTime(retryWaitTime)
		return nil
	}
}

// WithTimeout sets the timeout duration for each request.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) error {
		c.client.SetTimeout(timeout)
		return nil
	}
}

// WithTransport sets the HTTP transport.
func WithTransport(transport http.RoundTripper) ClientOption {
	return func(c *Client) error {
		c.client.SetTransport(transport)
		return nil
	}
}

// WithKey sets the Steam API key ("key" parameter) for all requests.
func WithKey(key string) ClientOption {
	return func(c *Client) error {
		c.client.SetQueryParam("key", key)
		return nil
	}
}

// WithLanguage sets the language ("language" parameter) for all requests.
func WithLanguage(lang string) ClientOption {
	return func(c *Client) error {
		c.client.SetQueryParam("language", lang)
		return nil
	}
}

// Client is the API entry-point.
type Client struct {
	client *resty.Client
}

// New creates a new client.
func New(options ...ClientOption) (*Client, error) {
	rclient := resty.New().
		SetQueryParam("format", "json")

	client := &Client{client: rclient}

	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// ClientRequest wraps request arguments.
type ClientRequest struct {
	Interface *schema.Interface
	Method    *schema.Method
	Options   RequestOptions
	Result    interface{}
}

// Request performs an API request to the given interface and method with given
// options and stores the result in the given result.
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
		"version":   creq.Method.FormatVersion(),
	}

	req := c.client.R().
		SetPathParams(pathParams).
		SetContext(creq.Options.Context).
		SetHeaders(creq.Options.Headers).
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
