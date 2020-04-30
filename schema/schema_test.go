package schema_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/geyser/v2/schema"
)

func TestSchema_Validate(t *testing.T) {
	subject := &schema.Schema{}

	err := subject.Validate()

	assert.NoError(t, err)

	subject.Interfaces = createInterfacesValid()

	err = subject.Validate()

	assert.NoError(t, err)

	subject.Interfaces = createInterfacesInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidInterfaceNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject.Interfaces = createInterfacesMethods()

	err = subject.Validate()

	assert.NoError(t, err)

	subject.Interfaces = createInterfacesMethodsInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject.Interfaces = createInterfacesMethodsInvalidVersion()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject.Interfaces = createInterfacesMethodsInvalidHTTPMethod()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}
}
