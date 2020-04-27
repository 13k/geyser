package geyser_test

import (
	"net/http"
	"net/url"
	"time"

	"github.com/13k/geyser"
)

type NoopClient struct {
	req *geyser.Request
	res *geyser.Response
	err error

	cookies      []*http.Cookie
	jar          http.CookieJar
	debug        bool
	header       http.Header
	hostURL      string
	logger       geyser.Logger
	proxy        string
	query        url.Values
	retryCount   int
	retryMaxWait time.Duration
	retryWait    time.Duration
	timeout      time.Duration
	transport    http.RoundTripper
}

var _ geyser.Client = (*NoopClient)(nil)

func (c *NoopClient) Execute(req *geyser.Request) (*geyser.Response, error) {
	c.req = req
	return c.res, c.err
}

func (c *NoopClient) HostURL() string {
	return c.hostURL
}

func (c *NoopClient) SetCookies(cookies []*http.Cookie) geyser.Client {
	c.cookies = cookies
	return c
}

func (c *NoopClient) SetCookieJar(jar http.CookieJar) geyser.Client {
	c.jar = jar
	return c
}

func (c *NoopClient) SetDebug(debug bool) geyser.Client {
	c.debug = debug
	return c
}

func (c *NoopClient) SetHeader(key string, value string) geyser.Client {
	if c.header == nil {
		c.header = make(http.Header)
	}

	c.header.Set(key, value)

	return c
}

func (c *NoopClient) SetHeaders(header http.Header) geyser.Client {
	if c.header == nil {
		c.header = make(http.Header)
	}

	for k, v := range header {
		c.header[k] = v
	}

	return c
}

func (c *NoopClient) SetHostURL(host string) geyser.Client {
	c.hostURL = host
	return c
}

func (c *NoopClient) SetLogger(logger geyser.Logger) geyser.Client {
	c.logger = logger
	return c
}

func (c *NoopClient) SetProxy(proxy string) geyser.Client {
	c.proxy = proxy
	return c
}

func (c *NoopClient) SetQueryParam(key string, value string) geyser.Client {
	if c.query == nil {
		c.query = make(url.Values)
	}

	c.query.Set(key, value)

	return c
}

func (c *NoopClient) SetQueryParams(query url.Values) geyser.Client {
	if c.query == nil {
		c.query = make(url.Values)
	}

	for k, v := range query {
		c.query[k] = v
	}

	return c
}

func (c *NoopClient) SetRetryCount(count int) geyser.Client {
	c.retryCount = count
	return c
}

func (c *NoopClient) SetRetryMaxWaitTime(maxWait time.Duration) geyser.Client {
	c.retryMaxWait = maxWait
	return c
}

func (c *NoopClient) SetRetryWaitTime(wait time.Duration) geyser.Client {
	c.retryWait = wait
	return c
}

func (c *NoopClient) SetTimeout(timeout time.Duration) geyser.Client {
	c.timeout = timeout
	return c
}

func (c *NoopClient) SetTransport(transport http.RoundTripper) geyser.Client {
	c.transport = transport
	return c
}
