package geyser_test

import (
	"net/http"
	"testing"

	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func schemaMethods() geyser.SchemaMethods {
	return geyser.SchemaMethods{schemaMethodValid()}
}

func schemaMethodsInvalidName() geyser.SchemaMethods {
	return geyser.SchemaMethods{schemaMethodInvalidName()}
}

func schemaMethodsInvalidVersion() geyser.SchemaMethods {
	return geyser.SchemaMethods{schemaMethodInvalidVersion()}
}

func schemaMethodsInvalidHTTPMethod() geyser.SchemaMethods {
	return geyser.SchemaMethods{schemaMethodInvalidHTTPMethod()}
}

func TestNewSchemaMethods(t *testing.T) {
	subject, err := geyser.NewSchemaMethods()

	assert.NoError(t, err)
	assert.Len(t, subject, 0)

	subject, err = geyser.NewSchemaMethods(schemaMethodValid())

	assert.NoError(t, err)
	assert.Len(t, subject, 1)

	subject, err = geyser.NewSchemaMethods(schemaMethodInvalidName())

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}

	subject, err = geyser.NewSchemaMethods(schemaMethodInvalidVersion())

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodVersionError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}

	subject, err = geyser.NewSchemaMethods(schemaMethodInvalidHTTPMethod())

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodHTTPMethodError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}
}

func TestMustNewSchemaMethods(t *testing.T) {
	var subject geyser.SchemaMethods

	require.NotPanics(t, func() {
		subject = geyser.MustNewSchemaMethods()
	})

	assert.Len(t, subject, 0)

	require.NotPanics(t, func() {
		subject = geyser.MustNewSchemaMethods(schemaMethodValid())
	})

	assert.Len(t, subject, 1)

	require.Panics(t, func() {
		subject = geyser.MustNewSchemaMethods(schemaMethodInvalidName())
	})

	require.Panics(t, func() {
		subject = geyser.MustNewSchemaMethods(schemaMethodInvalidVersion())
	})

	require.Panics(t, func() {
		subject = geyser.MustNewSchemaMethods(schemaMethodInvalidHTTPMethod())
	})
}

func TestSchemaMethods_Validate(t *testing.T) {
	subject := schemaMethods()

	err := subject.Validate()

	assert.NoError(t, err)

	subject = schemaMethodsInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = schemaMethodsInvalidVersion()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodVersionError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = schemaMethodsInvalidHTTPMethod()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodHTTPMethodError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}
}

func TestSchemaMethods_Get(t *testing.T) {
	key := geyser.SchemaMethodKey{Name: "Method", Version: 1}
	subject := schemaMethodsInvalidName()

	sm, err := subject.Get(key)

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, sm)
	}

	subject = schemaMethodsInvalidVersion()

	sm, err = subject.Get(key)

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodVersionError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, sm)
	}

	subject = schemaMethodsInvalidHTTPMethod()

	sm, err = subject.Get(key)

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodHTTPMethodError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, sm)
	}

	key1, sm1 := schemaMethod("Method1", 1)
	key2, sm2 := schemaMethod("Method1", 2)
	key3, sm3 := schemaMethod("Method1", 3)
	key4, sm4 := schemaMethod("Method2", 2)
	key5, sm5 := schemaMethod("Method2", 3)
	key6, sm6 := schemaMethod("Method3", 3)
	key7 := geyser.SchemaMethodKey{Name: "Method1", Version: 0}
	key8 := geyser.SchemaMethodKey{Name: "Method2", Version: 1}
	key9 := geyser.SchemaMethodKey{Name: "Method3", Version: 2}

	require.NotPanics(t, func() {
		subject = geyser.MustNewSchemaMethods(sm1, sm2, sm3, sm4, sm5, sm6)
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
		_, ok := err.(*geyser.InterfaceMethodNotFoundError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, sm)
	}

	sm, err = subject.Get(key8)

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InterfaceMethodNotFoundError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, sm)
	}

	sm, err = subject.Get(key9)

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InterfaceMethodNotFoundError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, sm)
	}
}

func TestSchemaMethods_GroupByName(t *testing.T) {
	subject := schemaMethodsInvalidName()
	groups, err := subject.GroupByName()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, groups)
	}

	subject = schemaMethodsInvalidVersion()
	groups, err = subject.GroupByName()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodVersionError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, groups)
	}

	subject = schemaMethodsInvalidHTTPMethod()
	groups, err = subject.GroupByName()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodHTTPMethodError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, groups)
	}

	subject = []*geyser.SchemaMethod{
		&geyser.SchemaMethod{Name: "Method", Version: 1, HTTPMethod: http.MethodGet},
		&geyser.SchemaMethod{Name: "Method", Version: 2, HTTPMethod: http.MethodGet},
		&geyser.SchemaMethod{Name: "Method", Version: 3, HTTPMethod: http.MethodGet},
		&geyser.SchemaMethod{Name: "Method2", Version: 1, HTTPMethod: http.MethodGet},
		&geyser.SchemaMethod{Name: "Method3", Version: 1, HTTPMethod: http.MethodGet},
		&geyser.SchemaMethod{Name: "Method3", Version: 2, HTTPMethod: http.MethodGet},
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

func TestSchemaMethodsGroup_Versions(t *testing.T) {
	methods := geyser.SchemaMethods{
		&geyser.SchemaMethod{Name: "Method", Version: 1, HTTPMethod: http.MethodGet},
		&geyser.SchemaMethod{Name: "Method", Version: 2, HTTPMethod: http.MethodGet},
		&geyser.SchemaMethod{Name: "Method", Version: 3, HTTPMethod: http.MethodGet},
	}

	subject := geyser.SchemaMethodsGroup{}

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
