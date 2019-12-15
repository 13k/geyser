package geyser_test

import (
	"testing"

	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
)

func TestSchema_Validate(t *testing.T) {
	subject := &geyser.Schema{}

	err := subject.Validate()

	assert.NoError(t, err)

	subject.Interfaces = schemaInterfacesValid()

	err = subject.Validate()

	assert.NoError(t, err)

	subject.Interfaces = schemaInterfacesInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidInterfaceNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject.Interfaces = schemaInterfacesMethods()

	err = subject.Validate()

	assert.NoError(t, err)

	subject.Interfaces = schemaInterfacesMethodsInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject.Interfaces = schemaInterfacesMethodsInvalidVersion()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodVersionError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject.Interfaces = schemaInterfacesMethodsInvalidHTTPMethod()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodHTTPMethodError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}
}
