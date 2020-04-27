package geyser

import (
	"io"
	"net/http"
	"time"
)

// Response represents an HTTP response.
type Response struct {
	// Originating Request.
	Request *Request
	// Underlying http.Response.
	Raw *http.Response
	// Body is the raw response body reader.
	// It's nil if `Request.Result` is non-nil (body already read).
	// If non-nil, must be manually consumed and closed.
	Body io.ReadCloser
	// BodyData contains the response body data that was read to be deserialized into `Request.Result`.
	// It's nil if `Request.Result` is nil.
	BodyData []byte
	// Time at which the response was received.
	Time time.Time

	duration time.Duration
}

// Header returns the response header.
//
// Shortcut for `Raw.Header`. Returns nil if `Raw` is nil.
func (r *Response) Header() http.Header {
	if r.Raw == nil {
		return nil
	}

	return r.Raw.Header
}

// Cookies returns the response cookies.
//
// Shortcut for `Raw.Cookies()`. Returns nil if `Raw` is nil.
func (r *Response) Cookies() []*http.Cookie {
	if r.Raw == nil {
		return nil
	}

	return r.Raw.Cookies()
}

// Status returns the HTTP status string for the executed request.
//
// Shortcut for `Raw.Status`. Returns empty string if `Raw` is nil.
func (r *Response) Status() string {
	if r.Raw == nil {
		return ""
	}

	return r.Raw.Status
}

// StatusCode returns the HTTP status code for the executed request.
//
// Shortcut for `Raw.StatusCode`. Returns 0 if `Raw` is nil.
func (r *Response) StatusCode() int {
	if r.Raw == nil {
		return 0
	}

	return r.Raw.StatusCode
}

// IsError returns true if HTTP status `code >= 400` otherwise false.
func (r *Response) IsError() bool {
	return r.StatusCode() >= 400
}

// IsSuccess returns true if HTTP status `code >= 200 and < 300` otherwise false.
func (r *Response) IsSuccess() bool {
	return r.StatusCode() >= 200 && r.StatusCode() < 300
}

// Duration returns the duration since the request was made until receiving the response.
//
// It's the result of `Time.Sub(Request.Time)`.
//
// Returns zero if any of the following is true: `Time` is the zero Time, `Request` is nil,
// `Request.Time` is the zero Time.
func (r *Response) Duration() time.Duration {
	if r.duration == 0 {
		if !r.Time.IsZero() && r.Request != nil && !r.Request.Time.IsZero() {
			r.duration = r.Time.Sub(r.Request.Time)
		}
	}

	return r.duration
}

// Result returns the parsed result.
//
// Shortcut for `Request.Result`.
func (r *Response) Result() interface{} {
	if r.Request == nil {
		return nil
	}

	return r.Request.Result
}
