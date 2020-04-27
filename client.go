package geyser

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	pathTemplate = "/{interface}/{method}/{version}"
)

var (
	headerKeyUserAgent = http.CanonicalHeaderKey("User-Agent")
	headerKeyAccept    = http.CanonicalHeaderKey("Accept")

	queryParamKeyAPIKey   = "key"
	queryParamKeyLanguage = "language"
	queryParamKeyFormat   = "format"
	queryParamValueFormat = "json"
)

// Client is the API entry-point.
type Client interface {
	// Execute performs a Request.
	//
	// Returns `ErrRequestNoInterface` if `Request.Interface` is nil.
	// Returns `ErrRequestNoMethod` if `Request.Method` is nil.
	//
	// Before execution, it sets `Request.Time` to the current time.
	//
	// After execution, regardless of errors, it sets `Request.Raw` with the http.Request used to
	// perform the HTTP request.
	Execute(req *Request) (*Response, error)
	// SetCookies appends multiple HTTP cookies.
	// These cookies will be used with all requests sent by this Client.
	SetCookies(cookies []*http.Cookie) Client
	// SetCookieJar sets a custom http.CookieJar in Client.
	SetCookieJar(jar http.CookieJar) Client
	// SetDebug enables the debug mode on Client.
	// Debug mode logs details of every request and response.
	SetDebug(debug bool) Client
	// SetHeader sets a single header field and its value in Client.
	// This header field will be applied to all requests sent by this Client.
	SetHeader(key string, value string) Client
	// SetHeaders sets multiple header fields and their values in Client.
	// These header fields will be applied to all requests sent by this Client.
	SetHeaders(header http.Header) Client
	// HostURL returns the Client's base host URL.
	HostURL() string
	// SetHostURL sets Host URL in Client.
	// It will be used with requests sent by this client with relative URL.
	SetHostURL(host string) Client
	// SetLogger sets the logger in Client that will log request and response details.
	SetLogger(logger Logger) Client
	// SetProxy sets the Proxy URL and Port for Client.
	SetProxy(proxy string) Client
	// SetQueryParam sets a single parameter and its value in Client.
	// This parameter will be applied to all requests sent by this Client.
	SetQueryParam(key string, value string) Client
	// SetQueryParams sets multiple parameters and their values in Client.
	// These parameters will be applied to all requests sent by this Client.
	SetQueryParams(query url.Values) Client
	// SetRetryCount enables request retrying and allows to set the number of tries, using a backoff
	// algorithm.
	// Setting to 0 disables retrying.
	SetRetryCount(count int) Client
	// SetRetryMaxWaitTime sets max time to wait before retrying a request (default: 2s).
	SetRetryMaxWaitTime(maxWait time.Duration) Client
	// SetRetryWaitTime sets default time to wait before retrying a request (default: 100ms).
	SetRetryWaitTime(wait time.Duration) Client
	// SetTimeout sets the timeout for requests sent by this Client.
	SetTimeout(timeout time.Duration) Client
	// SetTransport sets a custom `*http.Transport` or any `http.RoundTripper` compatible interface
	// implementation to be used as the transport by this Client.
	SetTransport(transport http.RoundTripper) Client
}

type client struct {
	c *resty.Client
}

var _ Client = (*client)(nil)

// New creates a new Client with given options.
func New(options ...ClientOption) (Client, error) {
	rclient := resty.New().
		SetHeader(headerKeyAccept, mimeJSON).
		SetQueryParam(queryParamKeyFormat, queryParamValueFormat).
		SetContentLength(true)

	client := &client{c: rclient}

	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

func (c *client) HostURL() string {
	return c.c.HostURL
}

func (c *client) SetCookies(cookies []*http.Cookie) Client {
	c.c.SetCookies(cookies)
	return c
}

func (c *client) SetCookieJar(jar http.CookieJar) Client {
	c.c.SetCookieJar(jar)
	return c
}

func (c *client) SetDebug(b bool) Client {
	c.c.SetDebug(b)
	return c
}

func (c *client) SetHeader(k, v string) Client {
	c.c.SetHeader(k, v)
	return c
}

func (c *client) SetHeaders(h http.Header) Client {
	if c.c.Header == nil {
		c.c.Header = make(http.Header)
	}

	for field, values := range h {
		c.c.Header[field] = values
	}

	return c
}

func (c *client) SetHostURL(host string) Client {
	c.c.SetHostURL(host)
	return c
}

func (c *client) SetLogger(l Logger) Client {
	c.c.SetLogger(l)
	return c
}

func (c *client) SetProxy(p string) Client {
	c.c.SetProxy(p)
	return c
}

func (c *client) SetQueryParam(k, v string) Client {
	c.c.SetQueryParam(k, v)
	return c
}

func (c *client) SetQueryParams(q url.Values) Client {
	if c.c.QueryParam == nil {
		c.c.QueryParam = make(url.Values)
	}

	for key, values := range q {
		c.c.QueryParam[key] = values
	}

	return c
}

func (c *client) SetRetryCount(i int) Client {
	c.c.SetRetryCount(i)
	return c
}

func (c *client) SetRetryMaxWaitTime(d time.Duration) Client {
	c.c.SetRetryMaxWaitTime(d)
	return c
}

func (c *client) SetRetryWaitTime(d time.Duration) Client {
	c.c.SetRetryWaitTime(d)
	return c
}

func (c *client) SetTimeout(d time.Duration) Client {
	c.c.SetTimeout(d)
	return c
}

func (c *client) SetTransport(t http.RoundTripper) Client {
	c.c.SetTransport(t)
	return c
}

func (c *client) Execute(req *Request) (*Response, error) {
	if req.Interface == nil {
		return nil, ErrRequestNoInterface
	}

	if req.Method == nil {
		return nil, ErrRequestNoMethod
	}

	httpMethod := req.Method.HTTPMethod

	if httpMethod == "" {
		httpMethod = http.MethodGet
	}

	var params url.Values

	switch httpMethod {
	case http.MethodGet:
		params = req.Query
	case http.MethodPost, http.MethodPut:
		params = req.Form
	}

	if params != nil {
		if err := req.Method.ValidateParams(params); err != nil {
			return nil, err
		}
	}

	pathParams := map[string]string{
		"interface": req.Interface.Name,
		"method":    req.Method.Name,
		"version":   req.Method.FormatVersion(),
	}

	doNotParseResponse := req.Result == nil
	rreq := c.c.R().
		SetPathParams(pathParams).
		SetContext(req.Context).
		SetQueryParamsFromValues(req.Query).
		SetFormDataFromValues(req.Form).
		SetBody(req.Body).
		SetDoNotParseResponse(doNotParseResponse)

	for key, value := range req.Header {
		rreq.Header[key] = value
	}

	if !doNotParseResponse {
		rreq.SetResult(req.Result)
	}

	rres, err := rreq.Execute(httpMethod, pathTemplate)

	req.Raw = rreq.RawRequest
	req.Time = rreq.Time

	res := &Response{
		Request: req,
		Raw:     rres.RawResponse,
		Time:    rres.ReceivedAt(),
	}

	if err == nil {
		if doNotParseResponse {
			res.Body = res.Raw.Body
		} else {
			res.BodyData = rres.Body()
		}
	}

	return res, err
}
