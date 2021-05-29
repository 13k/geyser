package web

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	// HostURL is the base domain URL that is used by Client.
	HostURL = "https://steamcommunity.com"

	userAgent = "geyser-web"

	pathGetRSAKey     = "/login/getrsakey"
	pathRenderCaptcha = "/login/rendercaptcha"
	pathSubmitLogin   = "/login/dologin"
)

var (
	hostURL = mustParseURL(HostURL)
)

// Client is an HTTP client for web requests.
type Client struct {
	c       *resty.Client
	jar     http.CookieJar
	domain  *url.URL
	mirrors []*url.URL
}

// NewClient creates a new web Client with HostURL as base domain.
func NewClient() *Client {
	jar, _ := cookiejar.New(nil) //nolint: errcheck

	c := resty.New().
		SetDebug(true).
		SetHostURL(HostURL).
		SetCookieJar(jar).
		SetHeader("User-Agent", userAgent)

	return &Client{
		c:      c,
		jar:    jar,
		domain: hostURL,
	}
}

// CookieJar returns the cookie jar used to store and send cookies.
func (c *Client) CookieJar() http.CookieJar {
	return c.jar
}

// domains returns all domains against which the Client can perform requests.
//
// The first domain returned is the base domain.
func (c *Client) domains() []*url.URL {
	return append([]*url.URL{c.domain}, c.mirrors...)
}

// cookies returns all stored cookies, keyed by domain.
func (c *Client) cookies() map[*url.URL][]*http.Cookie {
	domains := c.domains()
	m := make(map[*url.URL][]*http.Cookie, len(domains))

	for _, domain := range domains {
		cookies := c.jar.Cookies(domain)
		m[domain] = cookies
	}

	return m
}

// addMirrors adds domain mirrors.
func (c *Client) addMirrors(domains []*url.URL) {
	c.mirrors = append(c.mirrors, domains...)
}

// transferCookies copies the cookies from base domain to all domains.
func (c *Client) transferCookies() {
	cookies := c.jar.Cookies(c.domain)

	for _, mirror := range c.mirrors {
		c.jar.SetCookies(mirror, cookies)
	}
}

// setCookies sets cookies on all domains.
func (c *Client) setCookies(cookies []*http.Cookie) {
	c.jar.SetCookies(c.domain, cookies)

	for _, mirror := range c.mirrors {
		c.jar.SetCookies(mirror, cookies)
	}
}

// URL returns an URL with given path relative to the Client's base domain.
func (c *Client) URL(path string, query url.Values) *url.URL {
	url := &url.URL{
		Path:     path,
		RawQuery: query.Encode(),
	}

	return c.domain.ResolveReference(url)
}

// CaptchaURL returns the URL to a Steam captcha challenge page.
func (c *Client) CaptchaURL(gid int64) *url.URL {
	return c.URL(pathRenderCaptcha, url.Values{
		"gid": []string{strconv.FormatInt(gid, 10)},
	})
}

func (c *Client) post(
	path string,
	form requestForm,
	result interface{},
	invalidateCache bool,
) (*resty.Response, error) { //nolint: unparam
	values, err := form.FormValues()

	if err != nil {
		return nil, fmt.Errorf("geyser/web/steam: error encoding form: %w", err)
	}

	if invalidateCache {
		values["donotcache"] = []string{strconv.FormatInt(time.Now().UnixNano(), 10)}
	}

	req := c.c.R().
		SetFormDataFromValues(values).
		SetResult(result)

	res, err := req.Post(path)

	if err != nil {
		return nil, fmt.Errorf("geyser/steam/web: error performing request: %w", err)
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("geyser/steam/web: HTTP error: %s", res.Status())
	}

	return res, nil
}

// RSAKey fetches the RSA key for an user.
func (c *Client) RSAKey(username string) (*RSAKeyResult, error) {
	form := &rsaKeyForm{
		Username: username,
	}

	result := &RSAKeyResult{}

	if _, err := c.post(pathGetRSAKey, form, result, true); err != nil {
		return nil, err
	}

	return result, nil
}

// Login sends a logon request.
func (c *Client) Login(credentials *LoginCredentials) (*LoginResult, error) {
	form, err := credentials.form()

	if err != nil {
		return nil, fmt.Errorf("geyser/steam/web: error processing credentials: %w", err)
	}

	result := &LoginResult{}

	if _, err := c.post(pathSubmitLogin, form, result, true); err != nil {
		return nil, err
	}

	return result, nil
}
