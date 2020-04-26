package schema_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/geyser/schema"
)

func schemaMethodParam() *schema.SchemaMethodParam {
	return &schema.SchemaMethodParam{
		Name:        "param",
		Type:        "type",
		Optional:    true,
		Description: "description",
	}
}

func schemaMethodParamRequired() *schema.SchemaMethodParam {
	return &schema.SchemaMethodParam{
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
		_, ok := err.(*schema.RequiredParameterError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	err = subject.ValidateValue("value")

	assert.NoError(t, err)
}
