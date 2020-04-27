// Code generated by geyser. DO NOT EDIT.

package dota2_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/dota2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIDOTA2Fantasy(t *testing.T) {
	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err := dota2.NewIDOTA2Fantasy(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	assert.Same(t, client, ci.Client)
	assert.NotNil(t, ci.Interface)
}

func TestIDOTA2Fantasy_GetFantasyPlayerRawStats(t *testing.T) {
	var ci *dota2.IDOTA2Fantasy
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2Fantasy(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetFantasyPlayerRawStats()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetFantasyPlayerRawStats", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Fantasy_GetLeaderboards(t *testing.T) {
	var ci *dota2.IDOTA2Fantasy
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2Fantasy(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetLeaderboards()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLeaderboards", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Fantasy_GetPlayerInfo(t *testing.T) {
	var ci *dota2.IDOTA2Fantasy
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2Fantasy(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetPlayerInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerInfo", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Fantasy_GetProPlayerInfo(t *testing.T) {
	var ci *dota2.IDOTA2Fantasy
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2Fantasy(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetProPlayerInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetProPlayerInfo", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
