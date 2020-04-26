package schema_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/geyser/schema"
)

func TestSchema_Validate(t *testing.T) {
	subject := &schema.Schema{}

	err := subject.Validate()

	assert.NoError(t, err)

	subject.Interfaces = schemaInterfacesValid()

	err = subject.Validate()

	assert.NoError(t, err)

	subject.Interfaces = schemaInterfacesInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidInterfaceNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject.Interfaces = schemaInterfacesMethods()

	err = subject.Validate()

	assert.NoError(t, err)

	subject.Interfaces = schemaInterfacesMethodsInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject.Interfaces = schemaInterfacesMethodsInvalidVersion()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject.Interfaces = schemaInterfacesMethodsInvalidHTTPMethod()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}
}
