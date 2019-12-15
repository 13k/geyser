package geyser_test

import (
	"net/url"
	"testing"

	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
)

func schemaMethodParams() geyser.SchemaMethodParams {
	return geyser.SchemaMethodParams{schemaMethodParam()}
}

func schemaMethodParamsRequired() geyser.SchemaMethodParams {
	return geyser.SchemaMethodParams{schemaMethodParamRequired()}
}

func TestNewSchemaMethodParams(t *testing.T) {
	params := geyser.NewSchemaMethodParams()

	assert.Len(t, params, 0)

	params = geyser.NewSchemaMethodParams(schemaMethodParam())

	assert.Len(t, params, 1)

	params = geyser.NewSchemaMethodParams(schemaMethodParam(), schemaMethodParamRequired())

	assert.Len(t, params, 2)
}

func TestSchemaMethodParams_ValidateParams(t *testing.T) {
	missingParams := url.Values{"param": []string{""}}
	params := url.Values{"param": []string{"value"}}

	var subject geyser.SchemaMethodParams

	err := subject.ValidateParams(missingParams)

	assert.NoError(t, err)

	subject = schemaMethodParams()

	err = subject.ValidateParams(missingParams)

	assert.NoError(t, err)

	err = subject.ValidateParams(params)

	assert.NoError(t, err)

	subject = schemaMethodParamsRequired()

	err = subject.ValidateParams(missingParams)

	if assert.Error(t, err) {
		_, ok := err.(*geyser.RequiredParameterError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	err = subject.ValidateParams(params)

	assert.NoError(t, err)
}
