package geyser_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
)

func TestRequest_SetOptions(t *testing.T) {
	req := &geyser.Request{}

	ctx := context.Background()
	headers := map[string]string{"X-Test": "test"}
	params := url.Values{"param": []string{"value"}}
	formData := url.Values{"form": []string{"data"}}
	body := map[string]interface{}{"field": "value"}
	key := "abc123xyz"
	lang := "en"

	options := geyser.RequestOptions{
		Context:  ctx,
		Headers:  headers,
		Params:   params,
		FormData: formData,
		Body:     body,
		Key:      key,
		Lang:     lang,
	}

	req.SetOptions(options)

	assert.Same(t, ctx, req.Options.Context)
	assert.Equal(t, headers, req.Options.Headers)
	assert.Equal(t, params, req.Options.Params)
	assert.Equal(t, formData, req.Options.FormData)
	assert.Equal(t, body, req.Options.Body)
	assert.Equal(t, key, req.Options.Key)
	assert.Equal(t, lang, req.Options.Lang)
}

func TestRequest_SetResult(t *testing.T) {
	req := &geyser.Request{}
	resultMap := make(map[string]interface{})

	req.SetResult(resultMap)

	assert.Equal(t, resultMap, req.Result)

	resultStruct := &struct{ Field int }{}

	req.SetResult(resultStruct)

	assert.Same(t, resultStruct, req.Result)
}

type testRequestClient struct {
	creq *geyser.ClientRequest
}

func (c *testRequestClient) Request(req geyser.ClientRequest) (*resty.Response, error) {
	c.creq = &req
	return nil, nil
}

func TestRequest_Execute(t *testing.T) {
	c := &testRequestClient{}
	si := &schema.Interface{}
	sm := &schema.Method{}
	result := &struct{ Field int }{}

	req := &geyser.Request{
		Client:    c,
		Interface: si,
		Method:    sm,
		Result:    result,
	}

	resp, err := req.Execute()

	assert.NoError(t, err)
	assert.Nil(t, resp)

	if assert.NotNil(t, c.creq) {
		assert.Same(t, si, c.creq.Interface)
		assert.Same(t, sm, c.creq.Method)
		assert.Same(t, result, c.creq.Result)
	}
}
