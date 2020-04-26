package schema_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/13k/geyser/schema"
)

func createInterfacesValid() schema.Interfaces {
	return schema.Interfaces{createInterfaceValid()}
}

func createInterfacesAppID() schema.Interfaces {
	return schema.Interfaces{createInterfaceAppID()}
}

func createInterfacesMethods() schema.Interfaces {
	return schema.Interfaces{createInterfaceMethods()}
}

func createInterfacesInvalidName() schema.Interfaces {
	return schema.Interfaces{createInterfaceInvalidName()}
}

func createInterfacesMethodsInvalidName() schema.Interfaces {
	return schema.Interfaces{createInterfaceMethodsInvalidName()}
}

func createInterfacesMethodsInvalidVersion() schema.Interfaces {
	return schema.Interfaces{createInterfaceMethodsInvalidVersion()}
}

func createInterfacesMethodsInvalidHTTPMethod() schema.Interfaces {
	return schema.Interfaces{createInterfaceMethodsInvalidHTTPMethod()}
}

func TestNewInterfaces(t *testing.T) {
	subject, err := schema.NewInterfaces()

	assert.NoError(t, err)
	assert.Len(t, subject, 0)

	subject, err = schema.NewInterfaces(createInterfaceValid())

	assert.NoError(t, err)
	assert.Len(t, subject, 1)

	subject, err = schema.NewInterfaces(createInterfaceMethods())

	assert.NoError(t, err)
	assert.Len(t, subject, 1)

	subject, err = schema.NewInterfaces(createInterfaceInvalidName())

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidInterfaceNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}

	subject, err = schema.NewInterfaces(createInterfaceMethodsInvalidName())

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}

	subject, err = schema.NewInterfaces(createInterfaceMethodsInvalidVersion())

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}

	subject, err = schema.NewInterfaces(createInterfaceMethodsInvalidHTTPMethod())

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, subject)
	}
}

func TestMustNewInterfaces(t *testing.T) {
	var subject schema.Interfaces

	require.NotPanics(t, func() {
		subject = schema.MustNewInterfaces()
	})

	assert.Len(t, subject, 0)

	require.NotPanics(t, func() {
		subject = schema.MustNewInterfaces(createInterfaceValid())
	})

	assert.Len(t, subject, 1)

	require.NotPanics(t, func() {
		subject = schema.MustNewInterfaces(createInterfaceMethods())
	})

	assert.Len(t, subject, 1)

	require.Panics(t, func() {
		subject = schema.MustNewInterfaces(createInterfaceInvalidName())
	})

	require.Panics(t, func() {
		subject = schema.MustNewInterfaces(createInterfaceMethodsInvalidName())
	})

	require.Panics(t, func() {
		subject = schema.MustNewInterfaces(createInterfaceMethodsInvalidVersion())
	})

	require.Panics(t, func() {
		subject = schema.MustNewInterfaces(createInterfaceMethodsInvalidHTTPMethod())
	})
}

func TestInterfaces_Validate(t *testing.T) {
	subject := createInterfacesValid()

	err := subject.Validate()

	assert.NoError(t, err)

	subject = createInterfacesInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidInterfaceNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = createInterfacesMethods()

	err = subject.Validate()

	assert.NoError(t, err)

	subject = createInterfacesMethodsInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = createInterfacesMethodsInvalidVersion()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = createInterfacesMethodsInvalidHTTPMethod()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}
}

func TestInterfaces_Get(t *testing.T) {
	key := schema.InterfaceKey{Name: "IFace", AppID: 0}
	subject := createInterfacesInvalidName()
	si, err := subject.Get(key)

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidInterfaceNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, si)
	}

	key = schema.InterfaceKey{Name: "IFace", AppID: 0}
	subject = createInterfacesValid()
	si, err = subject.Get(key)

	if assert.NoError(t, err) {
		require.NotNil(t, si)
		assert.Equal(t, "IFace", si.Name)
	}

	key = schema.InterfaceKey{Name: "IFace", AppID: 123}
	subject = createInterfacesAppID()
	si, err = subject.Get(key)

	if assert.NoError(t, err) {
		require.NotNil(t, si)
		assert.Equal(t, "IFace_123", si.Name)
	}

	key = schema.InterfaceKey{Name: "IFace", AppID: 0}
	subject = createInterfacesMethodsInvalidName()
	si, err = subject.Get(key)

	if assert.NoError(t, err) {
		require.NotNil(t, si)
		assert.Equal(t, "IFace", si.Name)
	}

	key = schema.InterfaceKey{Name: "IFace", AppID: 0}
	subject = createInterfacesMethodsInvalidVersion()
	si, err = subject.Get(key)

	if assert.NoError(t, err) {
		require.NotNil(t, si)
		assert.Equal(t, "IFace", si.Name)
	}

	key = schema.InterfaceKey{Name: "IFace", AppID: 0}
	subject = createInterfacesMethodsInvalidHTTPMethod()
	si, err = subject.Get(key)

	if assert.NoError(t, err) {
		require.NotNil(t, si)
		assert.Equal(t, "IFace", si.Name)
	}

	key1, si1 := createInterface("IFace1", 111)
	key2, si2 := createInterface("IFace1", 222)
	key3, si3 := createInterface("IFace2", 0)
	key4, si4 := createInterface("IFace3", 0)
	key5, si5 := createInterface("IFace4", 444)
	key6 := schema.InterfaceKey{Name: "IFace1", AppID: 0}
	key7 := schema.InterfaceKey{Name: "IFace1", AppID: 333}
	key8 := schema.InterfaceKey{Name: "IFace2", AppID: 222}
	key9 := schema.InterfaceKey{Name: "IFace3", AppID: 333}
	key10 := schema.InterfaceKey{Name: "IFace4", AppID: 0}

	require.NotPanics(t, func() {
		subject = schema.MustNewInterfaces(si1, si2, si3, si4, si5)
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

func TestInterfaces_GroupByBaseName(t *testing.T) {
	subject := createInterfacesInvalidName()
	groups, err := subject.GroupByBaseName()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidInterfaceNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Nil(t, groups)
	}

	require.NotPanics(t, func() {
		subject = schema.MustNewInterfaces(
			&schema.Interface{Name: "IFace1_111"},
			&schema.Interface{Name: "IFace1_222"},
			&schema.Interface{Name: "IFace2"},
			&schema.Interface{Name: "IFace3_A"},
			&schema.Interface{Name: "IFace4_444"},
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

func TestInterfacesGroup_Name(t *testing.T) {
	interfaces := schema.Interfaces{
		&schema.Interface{Name: "Interface_1"},
		&schema.Interface{Name: "Interface_2"},
		&schema.Interface{Name: "Interface_3"},
	}

	subject := schema.InterfacesGroup{}

	for _, si := range interfaces {
		key, err := si.Key()

		require.NoError(t, err)

		subject[key] = si
	}

	name := subject.Name()

	assert.Equal(t, "Interface", name)
}

func TestInterfacesGroup_AppIDs(t *testing.T) {
	interfaces := schema.Interfaces{
		&schema.Interface{Name: "Interface_1"},
		&schema.Interface{Name: "Interface_2"},
		&schema.Interface{Name: "Interface_3"},
	}

	subject := schema.InterfacesGroup{}

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

func TestInterfacesGroup_GroupMethods(t *testing.T) {
	var interfaces schema.Interfaces

	require.NotPanics(t, func() {
		interfaces = schema.Interfaces{
			&schema.Interface{
				Name: "Interface_1",
				Methods: schema.MustNewMethods(
					&schema.Method{Name: "Method1", Version: 1, HTTPMethod: http.MethodGet},
					&schema.Method{Name: "Method2", Version: 1, HTTPMethod: http.MethodGet},
					&schema.Method{Name: "Method3", Version: 1, HTTPMethod: http.MethodGet},
				),
			},
			&schema.Interface{
				Name: "Interface_2",
				Methods: schema.MustNewMethods(
					&schema.Method{Name: "Method1", Version: 1, HTTPMethod: http.MethodGet},
					&schema.Method{Name: "Method1", Version: 2, HTTPMethod: http.MethodGet},
					&schema.Method{Name: "Method1", Version: 3, HTTPMethod: http.MethodGet},
				),
			},
			&schema.Interface{
				Name: "Interface_3",
				Methods: schema.MustNewMethods(
					&schema.Method{Name: "Method1", Version: 1, HTTPMethod: http.MethodGet},
					&schema.Method{Name: "Method2", Version: 2, HTTPMethod: http.MethodGet},
					&schema.Method{Name: "Method3", Version: 3, HTTPMethod: http.MethodGet},
				),
			},
		}
	})

	subject := schema.InterfacesGroup{}

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

	key := schema.MethodKey{Name: "Method1", Version: 1}
	sm, ok := group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method1")
	require.Equal(t, "Method1", sm.Name)
	require.Equal(t, 1, sm.Version)

	key = schema.MethodKey{Name: "Method1", Version: 2}
	sm, ok = group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method1")
	require.Equal(t, "Method1", sm.Name)
	require.Equal(t, 2, sm.Version)

	key = schema.MethodKey{Name: "Method1", Version: 3}
	sm, ok = group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method1")
	require.Equal(t, "Method1", sm.Name)
	require.Equal(t, 3, sm.Version)

	group, ok = groups["Method2"]

	require.Truef(t, ok, "method group %q not found", "Method2")
	require.Len(t, group, 2)

	key = schema.MethodKey{Name: "Method2", Version: 1}
	sm, ok = group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method2")
	require.Equal(t, "Method2", sm.Name)
	require.Equal(t, 1, sm.Version)

	key = schema.MethodKey{Name: "Method2", Version: 2}
	sm, ok = group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method2")
	require.Equal(t, "Method2", sm.Name)
	require.Equal(t, 2, sm.Version)

	group, ok = groups["Method3"]

	require.Truef(t, ok, "method group %q not found", "Method3")
	require.Len(t, group, 2)

	key = schema.MethodKey{Name: "Method3", Version: 1}
	sm, ok = group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method3")
	require.Equal(t, "Method3", sm.Name)
	require.Equal(t, 1, sm.Version)

	key = schema.MethodKey{Name: "Method3", Version: 3}
	sm, ok = group[key]

	require.Truef(t, ok, "method %s/%d not found in group %q", key.Name, key.Version, "Method3")
	require.Equal(t, "Method3", sm.Name)
	require.Equal(t, 3, sm.Version)
}
