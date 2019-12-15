package geyser_test

import (
	"fmt"
	"testing"

	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
)

func schemaInterface(name string, appID uint32) (geyser.SchemaInterfaceKey, *geyser.SchemaInterface) {
	key := geyser.SchemaInterfaceKey{Name: name, AppID: appID}

	if appID != 0 {
		name = fmt.Sprintf("%s_%d", name, appID)
	}

	si := &geyser.SchemaInterface{Name: name}

	return key, si
}

func schemaInterfaceValid() *geyser.SchemaInterface {
	return &geyser.SchemaInterface{Name: "IFace"}
}

func schemaInterfaceAppID() *geyser.SchemaInterface {
	return &geyser.SchemaInterface{Name: "IFace_123"}
}

func schemaInterfaceInvalidName() *geyser.SchemaInterface {
	return &geyser.SchemaInterface{Name: "iface"}
}

func schemaInterfaceMethods() *geyser.SchemaInterface {
	return &geyser.SchemaInterface{Name: "IFace", Methods: schemaMethods()}
}

func schemaInterfaceMethodsInvalidName() *geyser.SchemaInterface {
	return &geyser.SchemaInterface{Name: "IFace", Methods: schemaMethodsInvalidName()}
}

func schemaInterfaceMethodsInvalidVersion() *geyser.SchemaInterface {
	return &geyser.SchemaInterface{Name: "IFace", Methods: schemaMethodsInvalidVersion()}
}

func schemaInterfaceMethodsInvalidHTTPMethod() *geyser.SchemaInterface {
	return &geyser.SchemaInterface{Name: "IFace", Methods: schemaMethodsInvalidHTTPMethod()}
}

func TestSchemaInterface_Validate(t *testing.T) {
	subject := schemaInterfaceValid()

	err := subject.Validate()

	assert.NoError(t, err)

	subject = schemaInterfaceInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidInterfaceNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = schemaInterfaceMethods()

	err = subject.Validate()

	assert.NoError(t, err)

	subject = schemaInterfaceMethodsInvalidName()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodNameError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = schemaInterfaceMethodsInvalidVersion()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodVersionError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}

	subject = schemaInterfaceMethodsInvalidHTTPMethod()

	err = subject.Validate()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidMethodHTTPMethodError)
		assert.Truef(t, ok, "invalid error type: %T", err)
	}
}

func TestSchemaInterface_Key(t *testing.T) {
	expectedKey := geyser.SchemaInterfaceKey{Name: "IFace", AppID: 0}
	subject := schemaInterfaceValid()

	key, err := subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	expectedKey = geyser.SchemaInterfaceKey{Name: "IFace", AppID: 123}
	subject = schemaInterfaceAppID()

	key, err = subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	expectedKey = geyser.SchemaInterfaceKey{Name: "IFace", AppID: 0}
	subject = schemaInterfaceMethods()

	key, err = subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	expectedKey = geyser.SchemaInterfaceKey{Name: "IFace", AppID: 0}
	subject = schemaInterfaceMethodsInvalidName()

	key, err = subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	expectedKey = geyser.SchemaInterfaceKey{Name: "IFace", AppID: 0}
	subject = schemaInterfaceMethodsInvalidVersion()

	key, err = subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	expectedKey = geyser.SchemaInterfaceKey{Name: "IFace", AppID: 0}
	subject = schemaInterfaceMethodsInvalidHTTPMethod()

	key, err = subject.Key()

	assert.NoError(t, err)
	assert.Equal(t, expectedKey, key)

	key.Name = ""
	key.AppID = 0
	subject = schemaInterfaceInvalidName()

	key, err = subject.Key()

	if assert.Error(t, err) {
		_, ok := err.(*geyser.InvalidInterfaceNameError)

		assert.Truef(t, ok, "invalid error type: %T", err)
		assert.Equal(t, "", key.Name)
		assert.Equal(t, uint32(0), key.AppID)
	}
}
