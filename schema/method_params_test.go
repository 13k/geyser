package schema_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/geyser/v2/schema"
)

func createMethodParams() schema.MethodParams {
	return schema.MethodParams{createMethodParam()}
}

func createMethodParamsRequired() schema.MethodParams {
	return schema.MethodParams{createMethodParamRequired()}
}

func TestNewMethodParams(t *testing.T) {
	params := schema.NewMethodParams()

	assert.Len(t, params, 0)

	params = schema.NewMethodParams(createMethodParam())

	assert.Len(t, params, 1)

	params = schema.NewMethodParams(createMethodParam(), createMethodParamRequired())

	assert.Len(t, params, 2)
}

func TestMethodParams_ValidateParams(t *testing.T) {
	missingParams := url.Values{"param": []string{""}}
	params := url.Values{"param": []string{"value"}}

	var subject schema.MethodParams

	err := subject.ValidateParams(missingParams)

	assert.NoError(t, err)

	subject = createMethodParams()

	err = subject.ValidateParams(missingParams)

	assert.NoError(t, err)

	err = subject.ValidateParams(params)

	assert.NoError(t, err)

	subject = createMethodParamsRequired()

	err = subject.ValidateParams(missingParams)

	if assert.Error(t, err) {
		_, ok := err.(*schema.RequiredParameterError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	err = subject.ValidateParams(params)

	assert.NoError(t, err)
}
