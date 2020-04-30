// Code generated by geyser. DO NOT EDIT.

package dota2_test

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/dota2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIDOTA2Match(t *testing.T) {
	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err := dota2.NewIDOTA2Match(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	assert.Same(t, client, ci.Client)
	assert.NotNil(t, ci.Interface)
}

func TestIDOTA2Match_GetMatchMVPVotes(t *testing.T) {
	var ci *dota2.IDOTA2Match
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2Match(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetMatchMVPVotes()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchMVPVotes", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Match_GetTopLiveEventGame(t *testing.T) {
	var ci *dota2.IDOTA2Match
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2Match(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetTopLiveEventGame()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTopLiveEventGame", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Match_GetTopLiveGame(t *testing.T) {
	var ci *dota2.IDOTA2Match
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2Match(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetTopLiveGame()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTopLiveGame", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Match_GetTopWeekendTourneyGames(t *testing.T) {
	var ci *dota2.IDOTA2Match
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIDOTA2Match(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetTopWeekendTourneyGames()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTopWeekendTourneyGames", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
