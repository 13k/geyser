// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIDOTA2MatchStats(t *testing.T) {
	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	appIDs := []uint32{570, 205790}

	for _, appID := range appIDs {
		ci, err := steam.NewIDOTA2MatchStats(client, appID)

		require.NoError(t, err)
		require.NotNil(t, ci)

		assert.Same(t, client, ci.Client)
		assert.NotNil(t, ci.Interface)
	}
}

func TestIDOTA2MatchStats_GetRealtimeStats(t *testing.T) {
	var ci *steam.IDOTA2MatchStats
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIDOTA2MatchStats(client, 570)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetRealtimeStats()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetRealtimeStats", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	ci, err = steam.NewIDOTA2MatchStats(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetRealtimeStats()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetRealtimeStats", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
