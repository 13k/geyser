// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewITFSystem(t *testing.T) {
	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	appIDs := []uint32{440}

	for _, appID := range appIDs {
		ci, err := steam.NewITFSystem(client, appID)

		require.NoError(t, err)
		require.NotNil(t, ci)

		assert.Same(t, client, ci.Client)
		assert.NotNil(t, ci.Interface)
	}
}

func TestITFSystem_GetWorldStatus(t *testing.T) {
	var ci *steam.ITFSystem
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewITFSystem(client, 440)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetWorldStatus()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetWorldStatus", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
