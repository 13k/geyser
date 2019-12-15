// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package dota2_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/dota2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDOTA2Match(t *testing.T) {
	client := &dota2.Client{}
	i, err := dota2.NewDOTA2Match(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	assert.Same(t, client, i.Client)
	assert.NotNil(t, i.Interface)
}

func TestDOTA2Match_GetMatchMVPVotes(t *testing.T) {
	var i *dota2.DOTA2Match
	var err error
	var req *geyser.Request
	var result *dota2.DOTA2MatchGetMatchMVPVotes
	var ok bool

	client := &dota2.Client{}

	i, err = dota2.NewDOTA2Match(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetMatchMVPVotes()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchMVPVotes", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*dota2.DOTA2MatchGetMatchMVPVotes)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Match_GetTopLiveEventGame(t *testing.T) {
	var i *dota2.DOTA2Match
	var err error
	var req *geyser.Request
	var result *dota2.DOTA2MatchGetTopLiveEventGame
	var ok bool

	client := &dota2.Client{}

	i, err = dota2.NewDOTA2Match(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetTopLiveEventGame()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTopLiveEventGame", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*dota2.DOTA2MatchGetTopLiveEventGame)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Match_GetTopLiveGame(t *testing.T) {
	var i *dota2.DOTA2Match
	var err error
	var req *geyser.Request
	var result *dota2.DOTA2MatchGetTopLiveGame
	var ok bool

	client := &dota2.Client{}

	i, err = dota2.NewDOTA2Match(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetTopLiveGame()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTopLiveGame", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*dota2.DOTA2MatchGetTopLiveGame)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Match_GetTopWeekendTourneyGames(t *testing.T) {
	var i *dota2.DOTA2Match
	var err error
	var req *geyser.Request
	var result *dota2.DOTA2MatchGetTopWeekendTourneyGames
	var ok bool

	client := &dota2.Client{}

	i, err = dota2.NewDOTA2Match(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetTopWeekendTourneyGames()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTopWeekendTourneyGames", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*dota2.DOTA2MatchGetTopWeekendTourneyGames)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}
