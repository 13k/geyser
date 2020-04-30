// Code generated by geyser. DO NOT EDIT.

package dota2_test

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/dota2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIDOTA2League(t *testing.T) {
	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err := dota2.NewIDOTA2League(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	assert.Same(t, client, ci.Client)
	assert.NotNil(t, ci.Interface)
}

func TestIDOTA2League_GetLeagueData(t *testing.T) {
	var ci *dota2.IDOTA2League
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2League(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetLeagueData()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLeagueData", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2League_GetLeagueInfoList(t *testing.T) {
	var ci *dota2.IDOTA2League
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2League(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetLeagueInfoList()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLeagueInfoList", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2League_GetLeagueNodeResults(t *testing.T) {
	var ci *dota2.IDOTA2League
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2League(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetLeagueNodeResults()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLeagueNodeResults", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2League_GetLiveGames(t *testing.T) {
	var ci *dota2.IDOTA2League
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2League(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetLiveGames()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLiveGames", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2League_GetPredictionResults(t *testing.T) {
	var ci *dota2.IDOTA2League
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2League(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetPredictionResults()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPredictionResults", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2League_GetPredictions(t *testing.T) {
	var ci *dota2.IDOTA2League
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2League(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetPredictions()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPredictions", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2League_GetPrizePool(t *testing.T) {
	var ci *dota2.IDOTA2League
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2League(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetPrizePool()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPrizePool", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
