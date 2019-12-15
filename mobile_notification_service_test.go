// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package geyser_test

import (
	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewMobileNotificationService(t *testing.T) {
	client := &geyser.Client{}
	i, err := geyser.NewMobileNotificationService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	assert.Same(t, client, i.Client)
	assert.NotNil(t, i.Interface)
}

func TestMobileNotificationService_GetUserNotificationCounts(t *testing.T) {
	var i *geyser.MobileNotificationService
	var err error
	var req *geyser.Request
	var result *geyser.MobileNotificationServiceGetUserNotificationCounts
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewMobileNotificationService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetUserNotificationCounts()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetUserNotificationCounts", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.MobileNotificationServiceGetUserNotificationCounts)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestMobileNotificationService_SwitchSessionToPush(t *testing.T) {
	var i *geyser.MobileNotificationService
	var err error
	var req *geyser.Request
	var result *geyser.MobileNotificationServiceSwitchSessionToPush
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewMobileNotificationService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.SwitchSessionToPush()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "SwitchSessionToPush", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.MobileNotificationServiceSwitchSessionToPush)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}
