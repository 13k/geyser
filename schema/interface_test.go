package schema_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/geyser/schema"
)

func createInterface(name string, appID uint32) (schema.InterfaceKey, *schema.Interface) {
	key := schema.InterfaceKey{Name: name, AppID: appID}

	if appID != 0 {
		name = fmt.Sprintf("%s_%d", name, appID)
	}

	si := &schema.Interface{Name: name}

	return key, si
}

func createInterfaceValid() *schema.Interface {
	return &schema.Interface{Name: "IFace"}
}

func createInterfaceAppID() *schema.Interface {
	return &schema.Interface{Name: "IFace_123"}
}

func createInterfaceInvalidName() *schema.Interface {
	return &schema.Interface{Name: "iface"}
}

func createInterfaceMethods() *schema.Interface {
	return &schema.Interface{Name: "IFace", Methods: createMethods()}
}

func createInterfaceMethodsInvalidName() *schema.Interface {
	return &schema.Interface{Name: "IFace", Methods: createMethodsInvalidName()}
}

func createInterfaceMethodsInvalidVersion() *schema.Interface {
	return &schema.Interface{Name: "IFace", Methods: createMethodsInvalidVersion()}
}

func createInterfaceMethodsInvalidHTTPMethod() *schema.Interface {
	return &schema.Interface{Name: "IFace", Methods: createMethodsInvalidHTTPMethod()}
}

func TestInterface_Validate(t *testing.T) {
	subject := createInterfaceValid()

	err := subject.Validate()

	assert.NoError(t, err)

	subject = createInterfaceInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidInterfaceNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = createInterfaceMethods()

	err = subject.Validate()

	assert.NoError(t, err)

	subject = createInterfaceMethodsInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = createInterfaceMethodsInvalidVersion()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodVersionError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = createInterfaceMethodsInvalidHTTPMethod()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidMethodHTTPMethodError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}
}

func TestInterface_Key(t *testing.T) {
	expectedKey := schema.InterfaceKey{Name: "IFace", AppID: 0}
	subject := createInterfaceValid()

	key, err := subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	expectedKey = schema.InterfaceKey{Name: "IFace", AppID: 123}
	subject = createInterfaceAppID()

	key, err = subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	expectedKey = schema.InterfaceKey{Name: "IFace", AppID: 0}
	subject = createInterfaceMethods()

	key, err = subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	expectedKey = schema.InterfaceKey{Name: "IFace", AppID: 0}
	subject = createInterfaceMethodsInvalidName()

	key, err = subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	expectedKey = schema.InterfaceKey{Name: "IFace", AppID: 0}
	subject = createInterfaceMethodsInvalidVersion()

	key, err = subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	expectedKey = schema.InterfaceKey{Name: "IFace", AppID: 0}
	subject = createInterfaceMethodsInvalidHTTPMethod()

	key, err = subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	key.Name = ""
	key.AppID = 0
	subject = createInterfaceInvalidName()

	key, err = subject.Key()

	if assert.Error(t, err) {
		_, ok := err.(*schema.InvalidInterfaceNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Equal(t, "", key.Name)
		assert.Equal(t, uint32(0), key.AppID)
	}
}
