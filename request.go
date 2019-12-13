package geyser

import (
	"context"
	"net/url"

	"github.com/go-resty/resty/v2"
)

// Requester is the minimal interface for a `Request` to `Execute`.
type Requester interface {
	Request(ClientRequest) (*resty.Response, error)
}

// RequestOptions wrap options to pass to a request.
type RequestOptions struct {
	Context  context.Context
	Headers  map[string]string
	Params   url.Values
	FormData url.Values
	Body     interface{}
	Key      string
	Lang     string
}

// Request represents a "bound" request with client, interface and method.
type Request struct {
	Client    Requester
	Interface *SchemaInterface
	Method    *SchemaMethod
	Options   RequestOptions
	Result    interface{}
}

// SetOptions sets request options.
//
// The options provided here override the "global" client options.
func (req *Request) SetOptions(options RequestOptions) *Request {
	req.Options = options
	return req
}

// SetResult sets the result object where the response will be deserialized
// into.
//
// It expects a JSON "unmarshable" object.
func (req *Request) SetResult(result interface{}) *Request {
	req.Result = result
	return req
}

// Execute executes the request.
func (req *Request) Execute() (*resty.Response, error) {
	return req.Client.Request(ClientRequest{
		Interface: req.Interface,
		Method:    req.Method,
		Options:   req.Options,
		Result:    req.Result,
	})
}
