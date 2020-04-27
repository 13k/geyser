// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIGameNotificationsService(t *testing.T) {
	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err := steam.NewIGameNotificationsService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	assert.Same(t, client, ci.Client)
	assert.NotNil(t, ci.Interface)
}

func TestIGameNotificationsService_UserCreateSession(t *testing.T) {
	var ci *steam.IGameNotificationsService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIGameNotificationsService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.UserCreateSession()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "UserCreateSession", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIGameNotificationsService_UserDeleteSession(t *testing.T) {
	var ci *steam.IGameNotificationsService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIGameNotificationsService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.UserDeleteSession()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "UserDeleteSession", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIGameNotificationsService_UserUpdateSession(t *testing.T) {
	var ci *steam.IGameNotificationsService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIGameNotificationsService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.UserUpdateSession()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "UserUpdateSession", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
