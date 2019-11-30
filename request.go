package geyser

import (
	"context"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type RequestOptions struct {
	Context  context.Context
	Headers  map[string]string
	Params   url.Values
	FormData url.Values
	Body     interface{}
	Key      string
	Lang     string
}

type Request struct {
	Client    *Client
	Interface *SchemaInterface
	Method    *SchemaMethod
	Options   RequestOptions
	Result    interface{}
}

func (req *Request) SetOptions(options RequestOptions) *Request {
	req.Options = options
	return req
}

func (req *Request) SetResult(result interface{}) *Request {
	req.Result = result
	return req
}

func (req *Request) Execute() (*resty.Response, error) {
	return req.Client.Request(ClientRequest{
		Interface: req.Interface,
		Method:    req.Method,
		Options:   req.Options,
		Result:    req.Result,
	})
}
