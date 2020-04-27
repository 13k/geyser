package geyser

import (
	"net/http"
	"net/url"
	"time"
)

// ClientOption functions set options in `Client`.
type ClientOption func(Client) error

// WithCookieJar sets the cookie jar used with all requests.
func WithCookieJar(jar http.CookieJar) ClientOption {
	return func(c Client) error {
		c.SetCookieJar(jar)
		return nil
	}
}

// WithDebug enables debug logging.
func WithDebug() ClientOption {
	return func(c Client) error {
		c.SetDebug(true)
		return nil
	}
}

// WithHeaders sets the headers sent with all requests.
func WithHeaders(headers http.Header) ClientOption {
	return func(c Client) error {
		c.SetHeaders(headers)
		return nil
	}
}

// WithHostURL sets the host URL for all requests.
func WithHostURL(baseURL string) ClientOption {
	return func(c Client) error {
		c.SetHostURL(baseURL)
		return nil
	}
}

// WithKey sets the Steam API key ("key" parameter) for all requests.
func WithKey(key string) ClientOption {
	return func(c Client) error {
		c.SetQueryParam(queryParamKeyAPIKey, key)
		return nil
	}
}

// WithLanguage sets the language ("language" parameter) for all requests.
func WithLanguage(lang string) ClientOption {
	return func(c Client) error {
		c.SetQueryParam(queryParamKeyLanguage, lang)
		return nil
	}
}

// WithLogger sets the logger.
func WithLogger(logger Logger) ClientOption {
	return func(c Client) error {
		c.SetLogger(logger)
		return nil
	}
}

// WithProxy sets the proxy.
func WithProxy(proxy string) ClientOption {
	return func(c Client) error {
		c.SetProxy(proxy)
		return nil
	}
}

// WithQueryParams sets the query parameters sent with all requests.
func WithQueryParams(query url.Values) ClientOption {
	return func(c Client) error {
		c.SetQueryParams(query)
		return nil
	}
}

// WithRetryCount enables request retrying and allows to set the number of tries, using a backoff
// mechanism.
//
// Setting to 0 disables retrying.
func WithRetryCount(retryCount int) ClientOption {
	return func(c Client) error {
		c.SetRetryCount(retryCount)
		return nil
	}
}

// WithRetryMaxWaitTime sets the max time to wait before retrying a request.
//
// Default is 2s.
func WithRetryMaxWaitTime(retryMaxWaitTime time.Duration) ClientOption {
	return func(c Client) error {
		c.SetRetryMaxWaitTime(retryMaxWaitTime)
		return nil
	}
}

// WithRetryWaitTime sets the default time to wait before retrying a request.
//
// Default is 100ms.
func WithRetryWaitTime(retryWaitTime time.Duration) ClientOption {
	return func(c Client) error {
		c.SetRetryWaitTime(retryWaitTime)
		return nil
	}
}

// WithTimeout sets the timeout duration for each request.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c Client) error {
		c.SetTimeout(timeout)
		return nil
	}
}

// WithTransport sets the underlying HTTP transport used to perform requests.
func WithTransport(transport http.RoundTripper) ClientOption {
	return func(c Client) error {
		c.SetTransport(transport)
		return nil
	}
}

// WithUserAgent sets the "User-Agent" HTTP header for all requests.
func WithUserAgent(userAgent string) ClientOption {
	return func(c Client) error {
		c.SetHeader(headerKeyUserAgent, userAgent)
		return nil
	}
}
