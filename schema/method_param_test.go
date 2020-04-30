package schema_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/geyser/v2/schema"
)

func createMethodParam() *schema.MethodParam {
	return &schema.MethodParam{
		Name:        "param",
		Type:        "type",
		Optional:    true,
		Description: "description",
	}
}

func createMethodParamRequired() *schema.MethodParam {
	return &schema.MethodParam{
		Name:        "param",
		Type:        "type",
		Optional:    false,
		Description: "description",
	}
}

func TestMethodParam_ValidateValue(t *testing.T) {
	subject := createMethodParam()

	err := subject.ValidateValue("")

	assert.NoError(t, err)

	err = subject.ValidateValue("value")

	assert.NoError(t, err)

	subject = createMethodParamRequired()

	err = subject.ValidateValue("")

	if assert.Error(t, err) {
		_, ok := err.(*schema.RequiredParameterError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	err = subject.ValidateValue("value")

	assert.NoError(t, err)
}
