package schema_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/geyser/schema"
)

func schemaMethodParams() schema.SchemaMethodParams {
	return schema.SchemaMethodParams{schemaMethodParam()}
}

func schemaMethodParamsRequired() schema.SchemaMethodParams {
	return schema.SchemaMethodParams{schemaMethodParamRequired()}
}

func TestNewSchemaMethodParams(t *testing.T) {
	params := schema.NewSchemaMethodParams()

	assert.Len(t, params, 0)

	params = schema.NewSchemaMethodParams(schemaMethodParam())

	assert.Len(t, params, 1)

	params = schema.NewSchemaMethodParams(schemaMethodParam(), schemaMethodParamRequired())

	assert.Len(t, params, 2)
}

func TestSchemaMethodParams_ValidateParams(t *testing.T) {
	missingParams := url.Values{"param": []string{""}}
	params := url.Values{"param": []string{"value"}}

	var subject schema.SchemaMethodParams

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
		_, ok := err.(*schema.RequiredParameterError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	err = subject.ValidateParams(params)

	assert.NoError(t, err)
}
