package geyser_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
)

func TestRequest_GettersSetters(t *testing.T) {
	assert := assert.New(t)

	si := &schema.Interface{}
	sm := &schema.Method{}
	ctx := context.Background()
	form := url.Values{"form": {"data"}}
	body := map[string]interface{}{"field": "value"}
	result := map[string]interface{}{}

	req := geyser.NewRequest(si, sm).
		SetContext(ctx).
		SetHeader("X-Keep", "keep").
		SetHeader("X-Override", "initial").
		SetHeaders(http.Header{
			"X-Append":   {"append"},
			"X-Override": {"overridden"},
		}).
		SetQueryParam("keep", "keep").
		SetQueryParam("override", "initial").
		SetQueryParams(url.Values{
			"append":   {"append"},
			"override": {"overridden"},
		}).
		SetAPIKey("abc123xyz").
		SetLanguage("de").
		SetForm(form).
		SetBody(body).
		SetResult(result)

	assert.Same(si, req.Interface)
	assert.Same(sm, req.Method)
	assert.Same(ctx, req.Context)
	assert.Exactly(form, req.Form)
	assert.Exactly(body, req.Body)
	assert.Exactly(result, req.Result)

	expectedHeader := http.Header{
		"X-Keep":     {"keep"},
		"X-Override": {"overridden"},
		"X-Append":   {"append"},
	}

	expectedQuery := url.Values{
		"keep":     {"keep"},
		"override": {"overridden"},
		"append":   {"append"},
		"key":      {"abc123xyz"},
		"language": {"de"},
	}

	assert.Exactly(expectedHeader, req.Header)
	assert.Exactly(expectedQuery, req.Query)
}
