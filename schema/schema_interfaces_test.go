package schema_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/13k/geyser/schema"
)

func schemaInterfacesValid() schema.SchemaInterfaces {
	return schema.SchemaInterfaces{schemaInterfaceValid()}
}

func schemaInterfacesAppID() schema.SchemaInterfaces {
	return schema.SchemaInterfaces{schemaInterfaceAppID()}
}

func schemaInterfacesMethods() schema.SchemaInterfaces {
	return schema.SchemaInterfaces{schemaInterfaceMethods()}
}

func schemaInterfacesInvalidName() schema.SchemaInterfaces {
	return schema.SchemaInterfaces{schemaInterfaceInvalidName()}
}

func schemaInterfacesMethodsInvalidName() schema.SchemaInterfaces {
	return schema.SchemaInterfaces{schemaInterfaceMethodsInvalidName()}
}

func schemaInterfacesMethodsInvalidVersion() schema.SchemaInterfaces {
	return schema.SchemaInterfaces{schemaInterfaceMethodsInvalidVersion()}
}

func schemaInterfacesMethodsInvalidHTTPMethod() schema.SchemaInterfaces {
	return schema.SchemaInterfaces{schemaInterfaceMethodsInvalidHTTPMethod()}
}

func TestNewSchemaInterfaces(t *testing.T) {
	subject, err := schema.NewSchemaInterfaces()

	assert.NoError(t, err)
	assert.Len(t, subject, 0)

	subject, err = schema.NewSchemaInterfaces(schemaInterfaceValid())

	assert.NoError(t, err)
	assert.Len(t, subject, 1)

	subject, err = schema.NewSchemaInterfaces(schemaInterfaceMethods())

	assert.NoError(t, err)
	assert.Len(t, subject, 1)

	subject, err = schema.NewSchemaInterfaces(schemaInterfaceInvalidName())

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidInterfaceNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}

	subject, err = schema.NewSchemaInterfaces(schemaInterfaceMethodsInvalidName())

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}

	subject, err = schema.NewSchemaInterfaces(schemaInterfaceMethodsInvalidVersion())

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}

	subject, err = schema.NewSchemaInterfaces(schemaInterfaceMethodsInvalidHTTPMethod())

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}
}

func TestMustNewSchemaInterfaces(t *testing.T) {
	var subject schema.SchemaInterfaces

	require.NotPanics(t, func() {
		subject = schema.MustNewSchemaInterfaces()
	})

	assert.Len(t, subject, 0)

	require.NotPanics(t, func() {
		subject = schema.MustNewSchemaInterfaces(schemaInterfaceValid())
	})

	assert.Len(t, subject, 1)

	require.NotPanics(t, func() {
		subject = schema.MustNewSchemaInterfaces(schemaInterfaceMethods())
	})

	assert.Len(t, subject, 1)

	require.Panics(t, func() {
		subject = schema.MustNewSchemaInterfaces(schemaInterfaceInvalidName())
	})

	require.Panics(t, func() {
		subject = schema.MustNewSchemaInterfaces(schemaInterfaceMethodsInvalidName())
	})

	require.Panics(t, func() {
		subject = schema.MustNewSchemaInterfaces(schemaInterfaceMethodsInvalidVersion())
	})

	require.Panics(t, func() {
		subject = schema.MustNewSchemaInterfaces(schemaInterfaceMethodsInvalidHTTPMethod())
	})
}

func TestSchemaInterfaces_Validate(t *testing.T) {
	subject := schemaInterfacesValid()

	err := subject.Validate()

	assert.NoError(t, err)

	subject = schemaInterfacesInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidInterfaceNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = schemaInterfacesMethods()

	err = subject.Validate()

	assert.NoError(t, err)

	subject = schemaInterfacesMethodsInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = schemaInterfacesMethodsInvalidVersion()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = schemaInterfacesMethodsInvalidHTTPMethod()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}
}

func TestSchemaInterfaces_Get(t *testing.T) {
	key := schema.SchemaInterfaceKey{Name: "IFace", AppID: 0}
	subject := schemaInterfacesInvalidName()
	si, err := subject.Get(key)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidInterfaceNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, si)
	}

	key = schema.SchemaInterfaceKey{Name: "IFace", AppID: 0}
	subject = schemaInterfacesValid()
	si, err = subject.Get(key)

	if assert.NoError(t, err) {
		require.NotNil(t, si)
		assert.Equal(t, "IFace", si.Name)
	}

	key = schema.SchemaInterfaceKey{Name: "IFace", AppID: 123}
	subject = schemaInterfacesAppID()
	si, err = subject.Get(key)

	if assert.NoError(t, err) {
		require.NotNil(t, si)
		assert.Equal(t, "IFace_123", si.Name)
	}

	key = schema.SchemaInterfaceKey{Name: "IFace", AppID: 0}
	subject = schemaInterfacesMethodsInvalidName()
	si, err = subject.Get(key)

	if assert.NoError(t, err) {
		require.NotNil(t, si)
		assert.Equal(t, "IFace", si.Name)
	}

	key = schema.SchemaInterfaceKey{Name: "IFace", AppID: 0}
	subject = schemaInterfacesMethodsInvalidVersion()
	si, err = subject.Get(key)

	if assert.NoError(t, err) {
		require.NotNil(t, si)
		assert.Equal(t, "IFace", si.Name)
	}

	key = schema.SchemaInterfaceKey{Name: "IFace", AppID: 0}
	subject = schemaInterfacesMethodsInvalidHTTPMethod()
	si, err = subject.Get(key)

	if assert.NoError(t, err) {
		require.NotNil(t, si)
		assert.Equal(t, "IFace", si.Name)
	}

	key1, si1 := schemaInterface("IFace1", 111)
	key2, si2 := schemaInterface("IFace1", 222)
	key3, si3 := schemaInterface("IFace2", 0)
	key4, si4 := schemaInterface("IFace3", 0)
	key5, si5 := schemaInterface("IFace4", 444)
	key6 := schema.SchemaInterfaceKey{Name: "IFace1", AppID: 0}
	key7 := schema.SchemaInterfaceKey{Name: "IFace1", AppID: 333}
	key8 := schema.SchemaInterfaceKey{Name: "IFace2", AppID: 222}
	key9 := schema.SchemaInterfaceKey{Name: "IFace3", AppID: 333}
	key10 := schema.SchemaInterfaceKey{Name: "IFace4", AppID: 0}

	require.NotPanics(t, func() {
		subject = schema.MustNewSchemaInterfaces(si1, si2, si3, si4, si5)
	})

	si, err = subject.Get(key1)

	require.NoError(t, err)
	require.NotNil(t, si)
	assert.Same(t, si1, si)

	si, err = subject.Get(key2)

	require.NoError(t, err)
	require.NotNil(t, si)
	assert.Same(t, si2, si)

	si, err = subject.Get(key3)

	require.NoError(t, err)
	require.NotNil(t, si)
	assert.Same(t, si3, si)

	si, err = subject.Get(key4)

	require.NoError(t, err)
	require.NotNil(t, si)
	assert.Same(t, si4, si)

	si, err = subject.Get(key5)

	require.NoError(t, err)
	require.NotNil(t, si)
	assert.Same(t, si5, si)

	si, err = subject.Get(key6)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InterfaceNotFoundError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, si)
	}

	si, err = subject.Get(key7)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InterfaceNotFoundError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, si)
	}

	si, err = subject.Get(key8)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InterfaceNotFoundError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, si)
	}

	si, err = subject.Get(key9)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InterfaceNotFoundError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, si)
	}

	si, err = subject.Get(key10)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InterfaceNotFoundError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, si)
	}
}

func TestSchemaInterfaces_GroupByBaseName(t *testing.T) {
	subject := schemaInterfacesInvalidName()
	groups, err := subject.GroupByBaseName()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidInterfaceNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, groups)
	}

	require.NotPanics(t, func() {
		subject = schema.MustNewSchemaInterfaces(
			&schema.SchemaInterface{Name: "IFace1_111"},
			&schema.SchemaInterface{Name: "IFace1_222"},
			&schema.SchemaInterface{Name: "IFace2"},
			&schema.SchemaInterface{Name: "IFace3_A"},
			&schema.SchemaInterface{Name: "IFace4_444"},
		)
	})

	groups, err = subject.GroupByBaseName()

	require.NoError(t, err)
	assert.Len(t, groups, 4)

	group, ok := groups[""]

	assert.Falsef(t, ok, "group %q should not exist", "")
	assert.Len(t, group, 0)

	group, ok = groups["IFace1"]

	if assert.Truef(t, ok, "group %q not found", "IFace1") {
		assert.NotNil(t, group)
		assert.Len(t, group, 2)
	}

	group, ok = groups["IFace2"]

	if assert.Truef(t, ok, "group %q not found", "IFace2") {
		assert.NotNil(t, group)
		assert.Len(t, group, 1)
	}

	group, ok = groups["IFace3_A"]

	if assert.Truef(t, ok, "group %q not found", "IFace3_A") {
		assert.NotNil(t, group)
		assert.Len(t, group, 1)
	}

	group, ok = groups["IFace4"]

	if assert.Truef(t, ok, "group %q not found", "IFace4") {
		assert.NotNil(t, group)
		assert.Len(t, group, 1)
	}
}

func TestSchemaInterfacesGroup_Name(t *testing.T) {
	interfaces := schema.SchemaInterfaces{
		&schema.SchemaInterface{Name: "Interface_1"},
		&schema.SchemaInterface{Name: "Interface_2"},
		&schema.SchemaInterface{Name: "Interface_3"},
	}

	subject := schema.SchemaInterfacesGroup{}

	for _, si := range interfaces {
		key, err := si.Key()

		require.NoError(t, err)

		subject[key] = si
	}

	name := subject.Name()

	assert.Equal(t, "Interface", name)
}

func TestSchemaInterfacesGroup_AppIDs(t *testing.T) {
	interfaces := schema.SchemaInterfaces{
		&schema.SchemaInterface{Name: "Interface_1"},
		&schema.SchemaInterface{Name: "Interface_2"},
		&schema.SchemaInterface{Name: "Interface_3"},
	}

	subject := schema.SchemaInterfacesGroup{}

	for _, si := range interfaces {
		key, err := si.Key()

		require.NoError(t, err)

		subject[key] = si
	}

	appIDs := subject.AppIDs()

	assert.Len(t, appIDs, 3)
	assert.NotContains(t, appIDs, uint32(0))
	assert.Contains(t, appIDs, uint32(1))
	assert.Contains(t, appIDs, uint32(2))
	assert.Contains(t, appIDs, uint32(3))
	assert.NotContains(t, appIDs, uint32(4))
}

func TestSchemaInterfacesGroup_GroupMethods(t *testing.T) {
	var interfaces schema.SchemaInterfaces

	require.NotPanics(t, func() {
		interfaces = schema.SchemaInterfaces{
			&schema.SchemaInterface{
				Name: "Interface_1",
				Methods: schema.MustNewSchemaMethods(
					&schema.SchemaMethod{Name: "Method1", Version: 1, HTTPMethod: http.MethodGet},
					&schema.SchemaMethod{Name: "Method2", Version: 1, HTTPMethod: http.MethodGet},
					&schema.SchemaMethod{Name: "Method3", Version: 1, HTTPMethod: http.MethodGet},
				),
			},
			&schema.SchemaInterface{
				Name: "Interface_2",
				Methods: schema.MustNewSchemaMethods(
					&schema.SchemaMethod{Name: "Method1", Version: 1, HTTPMethod: http.MethodGet},
					&schema.SchemaMethod{Name: "Method1", Version: 2, HTTPMethod: http.MethodGet},
					&schema.SchemaMethod{Name: "Method1", Version: 3, HTTPMethod: http.MethodGet},
				),
			},
			&schema.SchemaInterface{
				Name: "Interface_3",
				Methods: schema.MustNewSchemaMethods(
					&schema.SchemaMethod{Name: "Method1", Version: 1, HTTPMethod: http.MethodGet},
					&schema.SchemaMethod{Name: "Method2", Version: 2, HTTPMethod: http.MethodGet},
					&schema.SchemaMethod{Name: "Method3", Version: 3, HTTPMethod: http.MethodGet},
				),
			},
		}
	})

	subject := schema.SchemaInterfacesGroup{}

	for _, si := range interfaces {
		key, err := si.Key()

		require.NoError(t, err)

		subject[key] = si
	}

	groups, err := subject.GroupMethods()

	require.NoError(t, err)
	require.Len(t, groups, 3)

	group, ok := groups[""]

	require.Falsef(t, ok, "method group %q should not exist", "")
	require.Len(t, group, 0)

	group, ok = groups["Method1"]

	require.Truef(t, ok, "method group %q not found", "Method1")
	require.Len(t, group, 3)

	key := schema.SchemaMethodKey{Name: "Method1", Version: 1}
	sm, ok := group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method1")
	require.Equal(t, "Method1", sm.Name)
	require.Equal(t, 1, sm.Version)

	key = schema.SchemaMethodKey{Name: "Method1", Version: 2}
	sm, ok = group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method1")
	require.Equal(t, "Method1", sm.Name)
	require.Equal(t, 2, sm.Version)

	key = schema.SchemaMethodKey{Name: "Method1", Version: 3}
	sm, ok = group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method1")
	require.Equal(t, "Method1", sm.Name)
	require.Equal(t, 3, sm.Version)

	group, ok = groups["Method2"]

	require.Truef(t, ok, "method group %q not found", "Method2")
	require.Len(t, group, 2)

	key = schema.SchemaMethodKey{Name: "Method2", Version: 1}
	sm, ok = group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method2")
	require.Equal(t, "Method2", sm.Name)
	require.Equal(t, 1, sm.Version)

	key = schema.SchemaMethodKey{Name: "Method2", Version: 2}
	sm, ok = group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method2")
	require.Equal(t, "Method2", sm.Name)
	require.Equal(t, 2, sm.Version)

	group, ok = groups["Method3"]

	require.Truef(t, ok, "method group %q not found", "Method3")
	require.Len(t, group, 2)

	key = schema.SchemaMethodKey{Name: "Method3", Version: 1}
	sm, ok = group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method3")
	require.Equal(t, "Method3", sm.Name)
	require.Equal(t, 1, sm.Version)

	key = schema.SchemaMethodKey{Name: "Method3", Version: 3}
	sm, ok = group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method3")
	require.Equal(t, "Method3", sm.Name)
	require.Equal(t, 3, sm.Version)
}
