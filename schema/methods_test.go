package schema_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/13k/geyser/v2/schema"
)

func createMethods() schema.Methods {
	return schema.Methods{createMethodValid()}
}

func createMethodsInvalidName() schema.Methods {
	return schema.Methods{createMethodInvalidName()}
}

func createMethodsInvalidVersion() schema.Methods {
	return schema.Methods{createMethodInvalidVersion()}
}

func createMethodsInvalidHTTPMethod() schema.Methods {
	return schema.Methods{createMethodInvalidHTTPMethod()}
}

func TestNewMethods(t *testing.T) {
	subject, err := schema.NewMethods()

	assert.NoError(t, err)
	assert.Len(t, subject, 0)

	subject, err = schema.NewMethods(createMethodValid())

	assert.NoError(t, err)
	assert.Len(t, subject, 1)

	subject, err = schema.NewMethods(createMethodInvalidName())

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}

	subject, err = schema.NewMethods(createMethodInvalidVersion())

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}

	subject, err = schema.NewMethods(createMethodInvalidHTTPMethod())

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}
}

func TestMustNewMethods(t *testing.T) {
	var subject schema.Methods

	require.NotPanics(t, func() {
		subject = schema.MustNewMethods()
	})

	assert.Len(t, subject, 0)

	require.NotPanics(t, func() {
		subject = schema.MustNewMethods(createMethodValid())
	})

	assert.Len(t, subject, 1)

	require.Panics(t, func() {
		subject = schema.MustNewMethods(createMethodInvalidName())
	})

	require.Panics(t, func() {
		subject = schema.MustNewMethods(createMethodInvalidVersion())
	})

	require.Panics(t, func() {
		subject = schema.MustNewMethods(createMethodInvalidHTTPMethod())
	})
}

func TestMethods_Validate(t *testing.T) {
	subject := createMethods()

	err := subject.Validate()

	assert.NoError(t, err)

	subject = createMethodsInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = createMethodsInvalidVersion()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = createMethodsInvalidHTTPMethod()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}
}

func TestMethods_Get(t *testing.T) {
	key := schema.MethodKey{Name: "Method", Version: 1}
	subject := createMethodsInvalidName()

	sm, err := subject.Get(key)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, sm)
	}

	subject = createMethodsInvalidVersion()

	sm, err = subject.Get(key)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, sm)
	}

	subject = createMethodsInvalidHTTPMethod()

	sm, err = subject.Get(key)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, sm)
	}

	key1, sm1 := createMethod("Method1", 1)
	key2, sm2 := createMethod("Method1", 2)
	key3, sm3 := createMethod("Method1", 3)
	key4, sm4 := createMethod("Method2", 2)
	key5, sm5 := createMethod("Method2", 3)
	key6, sm6 := createMethod("Method3", 3)
	key7 := schema.MethodKey{Name: "Method1", Version: 0}
	key8 := schema.MethodKey{Name: "Method2", Version: 1}
	key9 := schema.MethodKey{Name: "Method3", Version: 2}

	require.NotPanics(t, func() {
		subject = schema.MustNewMethods(sm1, sm2, sm3, sm4, sm5, sm6)
	})

	sm, err = subject.Get(key1)

	require.NoError(t, err)
	require.NotNil(t, sm)
	assert.Same(t, sm1, sm)

	sm, err = subject.Get(key2)

	require.NoError(t, err)
	require.NotNil(t, sm)
	assert.Same(t, sm2, sm)

	sm, err = subject.Get(key3)

	require.NoError(t, err)
	require.NotNil(t, sm)
	assert.Same(t, sm3, sm)

	sm, err = subject.Get(key4)

	require.NoError(t, err)
	require.NotNil(t, sm)
	assert.Same(t, sm4, sm)

	sm, err = subject.Get(key5)

	require.NoError(t, err)
	require.NotNil(t, sm)
	assert.Same(t, sm5, sm)

	sm, err = subject.Get(key6)

	require.NoError(t, err)
	require.NotNil(t, sm)
	assert.Same(t, sm6, sm)

	sm, err = subject.Get(key7)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InterfaceMethodNotFoundError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, sm)
	}

	sm, err = subject.Get(key8)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InterfaceMethodNotFoundError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, sm)
	}

	sm, err = subject.Get(key9)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InterfaceMethodNotFoundError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, sm)
	}
}

func TestMethods_GroupByName(t *testing.T) {
	subject := createMethodsInvalidName()
	groups, err := subject.GroupByName()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, groups)
	}

	subject = createMethodsInvalidVersion()
	groups, err = subject.GroupByName()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, groups)
	}

	subject = createMethodsInvalidHTTPMethod()
	groups, err = subject.GroupByName()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, groups)
	}

	subject = []*schema.Method{
		{Name: "Method", Version: 1, HTTPMethod: http.MethodGet},
		{Name: "Method", Version: 2, HTTPMethod: http.MethodGet},
		{Name: "Method", Version: 3, HTTPMethod: http.MethodGet},
		{Name: "Method2", Version: 1, HTTPMethod: http.MethodGet},
		{Name: "Method3", Version: 1, HTTPMethod: http.MethodGet},
		{Name: "Method3", Version: 2, HTTPMethod: http.MethodGet},
	}

	groups, err = subject.GroupByName()

	require.NoError(t, err)
	require.Len(t, groups, 3)

	group, ok := groups[""]

	require.Falsef(t, ok, "method group %q should not exist", "")
	require.Len(t, group, 0)

	group, ok = groups["Method"]

	require.Truef(t, ok, "method group %q not found", "Method")
	require.Len(t, group, 3)

	group, ok = groups["Method2"]

	require.Truef(t, ok, "method group %q not found", "Method2")
	require.Len(t, group, 1)

	group, ok = groups["Method3"]

	require.Truef(t, ok, "method group %q not found", "Method3")
	require.Len(t, group, 2)
}

func TestMethodsGroup_Versions(t *testing.T) {
	methods := schema.Methods{
		&schema.Method{Name: "Method", Version: 1, HTTPMethod: http.MethodGet},
		&schema.Method{Name: "Method", Version: 2, HTTPMethod: http.MethodGet},
		&schema.Method{Name: "Method", Version: 3, HTTPMethod: http.MethodGet},
	}

	subject := schema.MethodsGroup{}

	for _, method := range methods {
		key, err := method.Key()

		require.NoError(t, err)

		subject[key] = method
	}

	versions := subject.Versions()

	assert.Len(t, versions, 3)
	assert.NotContains(t, versions, 0)
	assert.Contains(t, versions, 1)
	assert.Contains(t, versions, 2)
	assert.Contains(t, versions, 3)
	assert.NotContains(t, versions, 4)
}
