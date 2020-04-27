package geyser_test

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/13k/geyser"
)

func TestClientOption_WithCookieJar(t *testing.T) {
	jar, err := cookiejar.New(nil)

	require.NoError(t, err)

	c := &NoopClient{}
	opt := geyser.WithCookieJar(jar)
	err = opt(c)

	require.NoError(t, err)
	assert.Same(t, jar, c.jar)
}

func TestClientOption_WithDebug(t *testing.T) {
	c := &NoopClient{}
	opt := geyser.WithDebug()
	err := opt(c)

	require.NoError(t, err)
	assert.True(t, c.debug)
}

func TestClientOption_WithHeaders(t *testing.T) {
	headers := http.Header{"X-Test": {"test"}}
	c := &NoopClient{}
	opt := geyser.WithHeaders(headers)
	err := opt(c)

	require.NoError(t, err)
	require.NotNil(t, c.header)
	assert.Exactly(t, headers, c.header)
}

func TestClientOption_WithHostURL(t *testing.T) {
	baseURL := "https://geyser.test"
	c := &NoopClient{}
	opt := geyser.WithHostURL(baseURL)
	err := opt(c)

	require.NoError(t, err)
	assert.Equal(t, baseURL, c.hostURL)
}

func TestClientOption_WithKey(t *testing.T) {
	key := "secret"
	c := &NoopClient{}
	opt := geyser.WithKey(key)
	err := opt(c)

	require.NoError(t, err)
	require.NotNil(t, c.query)
	assert.Equal(t, key, c.query.Get("key"))
}

func TestClientOption_WithLanguage(t *testing.T) {
	lang := "de"
	c := &NoopClient{}
	opt := geyser.WithLanguage(lang)
	err := opt(c)

	require.NoError(t, err)
	require.NotNil(t, c.query)
	assert.Equal(t, lang, c.query.Get("language"))
}

func TestClientOption_WithLogger(t *testing.T) {
	logger := &NoopLogger{}
	c := &NoopClient{}
	opt := geyser.WithLogger(logger)
	err := opt(c)

	require.NoError(t, err)
	assert.Same(t, logger, c.logger)
}

func TestClientOption_WithProxy(t *testing.T) {
	proxy := "https://geyser.proxy"
	c := &NoopClient{}
	opt := geyser.WithProxy(proxy)
	err := opt(c)

	require.NoError(t, err)
	assert.Equal(t, proxy, c.proxy)
}

func TestClientOption_WithQueryParams(t *testing.T) {
	query := url.Values{}
	c := &NoopClient{}
	opt := geyser.WithQueryParams(query)
	err := opt(c)

	require.NoError(t, err)
	require.NotNil(t, c.query)
	assert.Exactly(t, query, c.query)
}

func TestClientOption_WithRetryCount(t *testing.T) {
	count := 13
	c := &NoopClient{}
	opt := geyser.WithRetryCount(count)
	err := opt(c)

	require.NoError(t, err)
	assert.Equal(t, count, c.retryCount)
}

func TestClientOption_WithRetryMaxWaitTime(t *testing.T) {
	retryMaxWaitTime := 13 * time.Second
	c := &NoopClient{}
	opt := geyser.WithRetryMaxWaitTime(retryMaxWaitTime)
	err := opt(c)

	require.NoError(t, err)
	assert.Equal(t, retryMaxWaitTime, c.retryMaxWait)
}

func TestClientOption_WithRetryWaitTime(t *testing.T) {
	retryWaitTime := 2 * time.Second
	c := &NoopClient{}
	opt := geyser.WithRetryWaitTime(retryWaitTime)
	err := opt(c)

	require.NoError(t, err)
	assert.Equal(t, retryWaitTime, c.retryWait)
}

func TestClientOption_WithTimeout(t *testing.T) {
	timeout := 30 * time.Second
	c := &NoopClient{}
	opt := geyser.WithTimeout(timeout)
	err := opt(c)

	require.NoError(t, err)
	assert.Equal(t, timeout, c.timeout)
}

func TestClientOption_WithTransport(t *testing.T) {
	transport := http.DefaultTransport
	c := &NoopClient{}
	opt := geyser.WithTransport(transport)
	err := opt(c)

	require.NoError(t, err)
	assert.Same(t, transport, c.transport)
}

func TestClientOption_WithUserAgent(t *testing.T) {
	userAgent := "agent007"
	c := &NoopClient{}
	opt := geyser.WithUserAgent(userAgent)
	err := opt(c)

	require.NoError(t, err)
	require.NotNil(t, c.header)
	assert.Equal(t, userAgent, c.header.Get("User-Agent"))
}
