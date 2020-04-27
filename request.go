package geyser

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/13k/geyser/schema"
)

// Request represents an HTTP request to an Interface Method.
type Request struct {
	// Schema Interface to request.
	Interface *schema.Interface
	// Schema Method to request.
	Method *schema.Method
	// Context in which the request is to be executed.
	Context context.Context
	// Header fields to be sent with the request.
	Header http.Header
	// Cookies to be sent with the request.
	Cookies []*http.Cookie
	// Query parameters to be sent in the request URL.
	Query url.Values
	// Values to be form-encoded in request body.
	Form url.Values
	// Raw request body.
	Body interface{}
	// Result object where the response will be deserialized into.
	Result interface{}
	// Underlying http.Request.
	// It's nil if the Request was not executed.
	Raw *http.Request
	// Time at which the request execution started.
	// It's zero Time if the Request was not executed.
	Time time.Time
}

// NewRequest creates a new Request for the given Interface and Method.
func NewRequest(si *schema.Interface, sm *schema.Method) *Request {
	return &Request{
		Interface: si,
		Method:    sm,
	}
}

// SetInterface sets the Interface for the request.
func (r *Request) SetInterface(i *schema.Interface) *Request {
	r.Interface = i
	return r
}

// SetMethod sets the Interface Method for the request.
func (r *Request) SetMethod(m *schema.Method) *Request {
	r.Method = m
	return r
}

// SetContext sets the request context.
func (r *Request) SetContext(ctx context.Context) *Request {
	r.Context = ctx
	return r
}

// SetHeader sets a single header field.
func (r *Request) SetHeader(key, value string) *Request {
	if r.Header == nil {
		r.Header = make(http.Header)
	}

	r.Header.Set(key, value)

	return r
}

// SetHeaders sets multiple header fields.
func (r *Request) SetHeaders(header http.Header) *Request {
	if r.Header == nil {
		r.Header = make(http.Header)
	}

	for k, v := range header {
		r.Header[k] = v
	}

	return r
}

// SetCookies sets the HTTP cookies.
func (r *Request) SetCookies(cookies []*http.Cookie) *Request {
	r.Cookies = cookies
	return r
}

// SetQueryParam sets a single query parameter.
func (r *Request) SetQueryParam(key, value string) *Request {
	if r.Query == nil {
		r.Query = make(url.Values)
	}

	r.Query.Set(key, value)

	return r
}

// SetQueryParams sets multiple query parameters.
func (r *Request) SetQueryParams(query url.Values) *Request {
	if r.Query == nil {
		r.Query = make(url.Values)
	}

	for k, v := range query {
		r.Query[k] = v
	}

	return r
}

// SetForm sets the values to be form-encoded in request body.
func (r *Request) SetForm(f url.Values) *Request {
	r.Form = f
	return r
}

// SetBody sets the raw request body.
func (r *Request) SetBody(b interface{}) *Request {
	r.Body = b
	return r
}

// SetKey sets the "key" request parameter.
func (r *Request) SetAPIKey(v string) *Request {
	return r.SetQueryParam(queryParamKeyAPIKey, v)
}

// SetLang sets the "language" request parameter.
func (r *Request) SetLanguage(v string) *Request {
	return r.SetQueryParam(queryParamKeyLanguage, v)
}

// SetResult sets the result object where the response will be deserialized into.
func (r *Request) SetResult(result interface{}) *Request {
	r.Result = result
	return r
}
