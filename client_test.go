package geyser_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
)

type ClientTestSuite struct {
	suite.Suite

	sv    *httptest.Server
	svURL *url.URL
}

func TestClient(t *testing.T) {
	suite.Run(t, &ClientTestSuite{})
}

type RequestInfoResult struct {
	Method    string         `json:"method"`
	Scheme    string         `json:"scheme"`
	Host      string         `json:"host"`
	URL       *url.URL       `json:"url"`
	StringURL string         `json:"string_url"`
	FullURL   string         `json:"full_url"`
	Header    http.Header    `json:"header"`
	Cookies   []*http.Cookie `json:"cookies"`
}

func respondWithJSON(handler func(req *http.Request) (int, interface{})) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		status, body := handler(req)

		if body != nil {
			res.Header().Set("Content-Type", "application/json")
		}

		res.WriteHeader(status)

		if body != nil {
			enc := json.NewEncoder(res)
			_ = enc.Encode(body)
		}
	})
}

func (s *ClientTestSuite) SetupTest() {
	s.sv = httptest.NewServer(respondWithJSON(func(req *http.Request) (int, interface{}) {
		scheme := "http"

		if req.TLS != nil {
			scheme = "https"
		}

		fullURL := fmt.Sprintf("%s://%s%s", scheme, req.Host, req.URL)

		s.T().Logf("request: Method=%q Scheme=%q Host=%q URL=%q", req.Method, scheme, req.Host, req.URL)

		reqInfo := map[string]interface{}{
			"method":     req.Method,
			"scheme":     scheme,
			"host":       req.Host,
			"url":        req.URL,
			"string_url": req.URL.String(),
			"full_url":   fullURL,
			"header":     req.Header,
			"cookies":    req.Cookies(),
		}

		switch req.URL.Path {
		case
			"/Client/SetCookies/v1",
			"/Client/SetHeaders/v1",
			"/Client/SetHostURL/v1",
			"/Client/SetQueryParams/v1",
			"/Client/ExecuteRaw/v1",
			"/Client/ExecuteResult/v1":
			return http.StatusOK, reqInfo
		default:
			return http.StatusNotFound, nil
		}
	}))

	s.svURL, _ = url.Parse(s.sv.URL)

	s.T().Logf("server URL: %q", s.svURL)
}

func (s *ClientTestSuite) TearDownTest() {
	s.sv.Close()
}

func (s *ClientTestSuite) TestSetCookies() {
	c, err := geyser.New(geyser.WithHostURL(s.sv.URL))

	if !s.NoError(err) {
		return
	}

	cookies := []*http.Cookie{
		{Name: "test", Value: "test"},
	}

	c.SetCookies(cookies)

	result := &RequestInfoResult{}
	si := &schema.Interface{Name: "Client"}
	sm := &schema.Method{Name: "SetCookies", Version: 1, HTTPMethod: http.MethodGet}
	req := geyser.NewRequest(si, sm).SetResult(result)
	res, err := c.Execute(req)

	s.Require().NoError(err)
	s.Require().NotNil(req.Raw)
	s.NotZero(req.Time)
	s.Require().NotNil(res.Request)
	s.Require().NotNil(res.Raw)
	s.NotZero(res.Time)
	s.Require().Equal(http.StatusOK, res.StatusCode())
	s.Require().Nil(res.Body)
	s.Require().NotNil(res.BodyData)
	s.Require().Same(result, res.Result())

	s.Equal(cookies, result.Cookies)
}

func (s *ClientTestSuite) TestSetHeaders() {
	c, err := geyser.New(geyser.WithHostURL(s.sv.URL))

	if !s.NoError(err) {
		return
	}

	c.SetHeader("X-Keep", "keep")
	c.SetHeader("X-Override", "initial")
	c.SetHeaders(http.Header{
		"X-Append":   {"append"},
		"X-Override": {"override"},
	})

	expected := http.Header{
		"X-Keep":     {"keep"},
		"X-Append":   {"append"},
		"X-Override": {"override"},
	}

	result := &RequestInfoResult{}
	si := &schema.Interface{Name: "Client"}
	sm := &schema.Method{Name: "SetHeaders", Version: 1, HTTPMethod: http.MethodGet}
	req := geyser.NewRequest(si, sm).SetResult(result)
	res, err := c.Execute(req)

	s.Require().NoError(err)
	s.Require().NotNil(req.Raw)
	s.NotZero(req.Time)
	s.Require().NotNil(res.Request)
	s.Require().NotNil(res.Raw)
	s.NotZero(res.Time)
	s.Require().Equal(http.StatusOK, res.StatusCode())
	s.Require().Nil(res.Body)
	s.Require().NotNil(res.BodyData)
	s.Require().Same(result, res.Result())

	for k, v := range expected {
		s.Contains(result.Header, k)
		s.Equal(v, result.Header[k])
	}
}

func (s *ClientTestSuite) TestSetHostURL() {
	c, err := geyser.New()

	if !s.NoError(err) {
		return
	}

	c.SetHostURL(s.sv.URL)

	if !s.Equal(s.sv.URL, c.HostURL()) {
		return
	}

	result := &RequestInfoResult{}
	si := &schema.Interface{Name: "Client"}
	sm := &schema.Method{Name: "SetHostURL", Version: 1, HTTPMethod: http.MethodGet}
	req := geyser.NewRequest(si, sm).SetResult(result)
	res, err := c.Execute(req)

	s.Require().NoError(err)
	s.Require().NotNil(req.Raw)
	s.NotZero(req.Time)
	s.Require().NotNil(res.Request)
	s.Require().NotNil(res.Raw)
	s.NotZero(res.Time)
	s.Require().Equal(http.StatusOK, res.StatusCode())
	s.Require().Nil(res.Body)
	s.Require().NotNil(res.BodyData)
	s.Require().Same(result, res.Result())

	s.Equal(s.svURL.Host, result.Host)
}

func (s *ClientTestSuite) TestSetQueryParams() {
	c, err := geyser.New(geyser.WithHostURL(s.sv.URL))

	if !s.NoError(err) {
		return
	}

	c.SetQueryParam("keep", "keep")
	c.SetQueryParam("override", "initial")
	c.SetQueryParams(url.Values{
		"append":   {"append"},
		"override": {"override"},
	})

	result := &RequestInfoResult{}
	si := &schema.Interface{Name: "Client"}
	sm := &schema.Method{Name: "SetQueryParams", Version: 1, HTTPMethod: http.MethodGet}
	req := geyser.NewRequest(si, sm).SetResult(result)
	res, err := c.Execute(req)

	s.Require().NoError(err)
	s.Require().NotNil(req.Raw)
	s.NotZero(req.Time)
	s.Require().NotNil(res.Request)
	s.Require().NotNil(res.Raw)
	s.NotZero(res.Time)
	s.Require().Equal(http.StatusOK, res.StatusCode())
	s.Require().Nil(res.Body)
	s.Require().NotNil(res.BodyData)
	s.Require().Same(result, res.Result())

	expected := url.Values{
		"keep":     {"keep"},
		"append":   {"append"},
		"override": {"override"},
	}

	actual := result.URL.Query()

	for k, v := range expected {
		s.Contains(actual, k)
		s.Equal(v, actual[k])
	}
}

func (s *ClientTestSuite) TestExecuteResult() {
	c, err := geyser.New(
		geyser.WithHostURL(s.sv.URL),
		geyser.WithHeaders(http.Header{"X-Client": {"client"}}),
		geyser.WithQueryParams(url.Values{"client": {"client"}}),
	)

	s.Require().NoError(err)

	result := &RequestInfoResult{}
	si := &schema.Interface{Name: "Client"}
	sm := &schema.Method{Name: "ExecuteResult", Version: 1, HTTPMethod: http.MethodGet}

	req := geyser.NewRequest(si, sm).
		SetResult(result).
		SetHeader("X-Request", "request").
		SetQueryParam("request", "request")

	res, err := c.Execute(req)

	s.Require().NoError(err)
	s.Require().NotNil(req.Raw)
	s.NotZero(req.Time)
	s.Require().NotNil(res.Request)
	s.Require().NotNil(res.Raw)
	s.NotZero(res.Time)
	s.Require().Equal(http.StatusOK, res.StatusCode())
	s.Require().Nil(res.Body)
	s.Require().NotNil(res.BodyData)
	s.Require().Same(result, res.Result())

	expectedHeader := http.Header{
		"Accept":    {"application/json"},
		"X-Client":  {"client"},
		"X-Request": {"request"},
	}

	expectedQuery := url.Values{
		"format":  {"json"},
		"client":  {"client"},
		"request": {"request"},
	}

	actualHeader := result.Header
	actualQuery := result.URL.Query()

	for k, v := range expectedHeader {
		s.Contains(actualHeader, k)
		s.Equal(v, actualHeader[k])
	}

	for k, v := range expectedQuery {
		s.Contains(actualQuery, k)
		s.Equal(v, actualQuery[k])
	}
}

func (s *ClientTestSuite) TestExecuteRaw() {
	c, err := geyser.New(
		geyser.WithHostURL(s.sv.URL),
		geyser.WithHeaders(http.Header{"X-Client": {"client"}}),
		geyser.WithQueryParams(url.Values{"client": {"client"}}),
	)

	s.Require().NoError(err)

	si := &schema.Interface{Name: "Client"}
	sm := &schema.Method{Name: "ExecuteRaw", Version: 1, HTTPMethod: http.MethodGet}

	req := geyser.NewRequest(si, sm).
		SetHeader("X-Request", "request").
		SetQueryParam("request", "request")

	res, err := c.Execute(req)

	s.Require().NoError(err)
	s.Require().NotNil(req.Raw)
	s.NotZero(req.Time)
	s.Require().NotNil(res.Request)
	s.Require().NotNil(res.Raw)
	s.NotZero(res.Time)
	s.Require().Equal(http.StatusOK, res.StatusCode())
	s.Require().NotNil(res.Body)
	s.Require().Nil(res.BodyData)
	s.Require().Nil(res.Result())

	expectedHeader := http.Header{
		"Accept":    {"application/json"},
		"X-Client":  {"client"},
		"X-Request": {"request"},
	}

	expectedQuery := url.Values{
		"format":  {"json"},
		"client":  {"client"},
		"request": {"request"},
	}

	result := &RequestInfoResult{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(result)

	if !s.NoError(err) {
		return
	}

	actualHeader := result.Header
	actualQuery := result.URL.Query()

	for k, v := range expectedHeader {
		s.Contains(actualHeader, k)
		s.Equal(v, actualHeader[k])
	}

	for k, v := range expectedQuery {
		s.Contains(actualQuery, k)
		s.Equal(v, actualQuery[k])
	}
}
