package schema_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/geyser/schema"
)

func createMethod(name string, version int) (schema.MethodKey, *schema.Method) {
	key := schema.MethodKey{Name: name, Version: version}
	sm := &schema.Method{Name: name, Version: version, HTTPMethod: http.MethodGet}
	return key, sm
}

func createMethodValid() *schema.Method {
	return &schema.Method{Name: "Method", Version: 1, HTTPMethod: http.MethodGet}
}

func createMethodInvalidName() *schema.Method {
	return &schema.Method{Name: "9Method", Version: 1, HTTPMethod: http.MethodGet}
}

func createMethodInvalidVersion() *schema.Method {
	return &schema.Method{Name: "Method", Version: 0, HTTPMethod: http.MethodGet}
}

func createMethodInvalidHTTPMethod() *schema.Method {
	return &schema.Method{Name: "Method", Version: 1, HTTPMethod: "xyz"}
}

func createMethodWithParams() *schema.Method {
	return &schema.Method{
		Name:       "MyMethod",
		Version:    1,
		HTTPMethod: http.MethodGet,
		Params:     createMethodParams(),
	}
}

func createMethodWithRequiredParams() *schema.Method {
	return &schema.Method{
		Name:       "MyMethod",
		Version:    1,
		HTTPMethod: http.MethodGet,
		Params:     createMethodParamsRequired(),
	}
}

func TestMethod_Validate(t *testing.T) {
	subject := createMethodValid()

	err := subject.Validate()

	assert.NoError(t, err)

	subject = createMethodInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = createMethodInvalidVersion()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = createMethodInvalidHTTPMethod()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}
}

func TestMethod_Key(t *testing.T) {
	expectedKey := schema.MethodKey{Name: "Method", Version: 1}
	subject := createMethodValid()

	key, err := subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	key.Name = ""
	key.Version = 0
	subject = createMethodInvalidName()

	key, err = subject.Key()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Equal(t, "", key.Name)
		assert.Equal(t, 0, key.Version)
	}

	subject = createMethodInvalidVersion()

	key, err = subject.Key()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Equal(t, "", key.Name)
		assert.Equal(t, 0, key.Version)
	}

	subject = createMethodInvalidHTTPMethod()

	key, err = subject.Key()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Equal(t, "", key.Name)
		assert.Equal(t, 0, key.Version)
	}
}

func TestMethod_ValidateParams(t *testing.T) {
	missingParams := url.Values{"param": []string{""}}
	params := url.Values{"param": []string{"value"}}

	subject := createMethodValid()

	err := subject.ValidateParams(missingParams)

	assert.NoError(t, err)

	err = subject.ValidateParams(params)

	assert.NoError(t, err)

	subject = createMethodWithParams()

	err = subject.ValidateParams(missingParams)

	assert.NoError(t, err)

	err = subject.ValidateParams(params)

	assert.NoError(t, err)

	subject = createMethodWithRequiredParams()

	err = subject.ValidateParams(missingParams)

	if assert.Error(t, err) {
		_, ok := err.(*schema.RequiredParameterError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	err = subject.ValidateParams(params)

	assert.NoError(t, err)
}
