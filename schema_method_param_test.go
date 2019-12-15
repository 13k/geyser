package geyser_test

import (
	"testing"

	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
)

func schemaMethodParam() *geyser.SchemaMethodParam {
	return &geyser.SchemaMethodParam{
		Name:        "param",
		Type:        "type",
		Optional:    true,
		Description: "description",
	}
}

func schemaMethodParamRequired() *geyser.SchemaMethodParam {
	return &geyser.SchemaMethodParam{
		Name:        "param",
		Type:        "type",
		Optional:    false,
		Description: "description",
	}
}

func TestSchemaMethodParam_ValidateValue(t *testing.T) {
	subject := schemaMethodParam()

	err := subject.ValidateValue("")

	assert.NoError(t, err)

	err = subject.ValidateValue("value")

	assert.NoError(t, err)

	subject = schemaMethodParamRequired()

	err = subject.ValidateValue("")

	if assert.Error(t, err) {
		_, ok := err.(*geyser.RequiredParameterError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	err = subject.ValidateValue("value")

	assert.NoError(t, err)
}
