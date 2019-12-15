package geyser_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
)

func schemaMethod(name string, version int) (geyser.SchemaMethodKey, *geyser.SchemaMethod) {
	key := geyser.SchemaMethodKey{Name: name, Version: version}
	sm := &geyser.SchemaMethod{Name: name, Version: version, HTTPMethod: http.MethodGet}
	return key, sm
}

func schemaMethodValid() *geyser.SchemaMethod {
	return &geyser.SchemaMethod{Name: "Method", Version: 1, HTTPMethod: http.MethodGet}
}

func schemaMethodInvalidName() *geyser.SchemaMethod {
	return &geyser.SchemaMethod{Name: "9Method", Version: 1, HTTPMethod: http.MethodGet}
}

func schemaMethodInvalidVersion() *geyser.SchemaMethod {
	return &geyser.SchemaMethod{Name: "Method", Version: 0, HTTPMethod: http.MethodGet}
}

func schemaMethodInvalidHTTPMethod() *geyser.SchemaMethod {
	return &geyser.SchemaMethod{Name: "Method", Version: 1, HTTPMethod: "xyz"}
}

func schemaMethodWithParams() *geyser.SchemaMethod {
	return &geyser.SchemaMethod{
		Name:       "MyMethod",
		Version:    1,
		HTTPMethod: http.MethodGet,
		Params:     schemaMethodParams(),
	}
}

func schemaMethodWithRequiredParams() *geyser.SchemaMethod {
	return &geyser.SchemaMethod{
		Name:       "MyMethod",
		Version:    1,
		HTTPMethod: http.MethodGet,
		Params:     schemaMethodParamsRequired(),
	}
}

func TestSchemaMethod_Validate(t *testing.T) {
	subject := schemaMethodValid()

	err := subject.Validate()

	assert.NoError(t, err)

	subject = schemaMethodInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = schemaMethodInvalidVersion()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodVersionError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = schemaMethodInvalidHTTPMethod()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodHTTPMethodError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}
}

func TestSchemaMethod_Key(t *testing.T) {
	expectedKey := geyser.SchemaMethodKey{Name: "Method", Version: 1}
	subject := schemaMethodValid()

	key, err := subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	key.Name = ""
	key.Version = 0
	subject = schemaMethodInvalidName()

	key, err = subject.Key()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Equal(t, "", key.Name)
		assert.Equal(t, 0, key.Version)
	}

	subject = schemaMethodInvalidVersion()

	key, err = subject.Key()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodVersionError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Equal(t, "", key.Name)
		assert.Equal(t, 0, key.Version)
	}

	subject = schemaMethodInvalidHTTPMethod()

	key, err = subject.Key()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodHTTPMethodError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Equal(t, "", key.Name)
		assert.Equal(t, 0, key.Version)
	}
}

func TestSchemaMethod_ValidateParams(t *testing.T) {
	missingParams := url.Values{"param": []string{""}}
	params := url.Values{"param": []string{"value"}}

	subject := schemaMethodValid()

	err := subject.ValidateParams(missingParams)

	assert.NoError(t, err)

	err = subject.ValidateParams(params)

	assert.NoError(t, err)

	subject = schemaMethodWithParams()

	err = subject.ValidateParams(missingParams)

	assert.NoError(t, err)

	err = subject.ValidateParams(params)

	assert.NoError(t, err)

	subject = schemaMethodWithRequiredParams()

	err = subject.ValidateParams(missingParams)

	if assert.Error(t, err) {
		_, ok := err.(*geyser.RequiredParameterError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	err = subject.ValidateParams(params)

	assert.NoError(t, err)
}
