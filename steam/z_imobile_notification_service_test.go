// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIMobileNotificationService(t *testing.T) {
	client := &steam.Client{}
	iface, err := steam.NewIMobileNotificationService(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	assert.Same(t, client, iface.Client)
	assert.NotNil(t, iface.Interface)
}

func TestIMobileNotificationService_GetUserNotificationCounts(t *testing.T) {
	var iface *steam.IMobileNotificationService
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIMobileNotificationService(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetUserNotificationCounts()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetUserNotificationCounts", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIMobileNotificationService_SwitchSessionToPush(t *testing.T) {
	var iface *steam.IMobileNotificationService
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIMobileNotificationService(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.SwitchSessionToPush()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "SwitchSessionToPush", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
