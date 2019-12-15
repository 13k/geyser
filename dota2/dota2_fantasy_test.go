// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package dota2_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/dota2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDOTA2Fantasy(t *testing.T) {
	client := &dota2.Client{}
	i, err := dota2.NewDOTA2Fantasy(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	assert.Same(t, client, i.Client)
	assert.NotNil(t, i.Interface)
}

func TestDOTA2Fantasy_GetFantasyPlayerRawStats(t *testing.T) {
	var i *dota2.DOTA2Fantasy
	var err error
	var req *geyser.Request
	var result *dota2.DOTA2FantasyGetFantasyPlayerRawStats
	var ok bool

	client := &dota2.Client{}

	i, err = dota2.NewDOTA2Fantasy(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetFantasyPlayerRawStats()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetFantasyPlayerRawStats", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*dota2.DOTA2FantasyGetFantasyPlayerRawStats)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Fantasy_GetLeaderboards(t *testing.T) {
	var i *dota2.DOTA2Fantasy
	var err error
	var req *geyser.Request
	var result *dota2.DOTA2FantasyGetLeaderboards
	var ok bool

	client := &dota2.Client{}

	i, err = dota2.NewDOTA2Fantasy(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetLeaderboards()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLeaderboards", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*dota2.DOTA2FantasyGetLeaderboards)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Fantasy_GetPlayerInfo(t *testing.T) {
	var i *dota2.DOTA2Fantasy
	var err error
	var req *geyser.Request
	var result *dota2.DOTA2FantasyGetPlayerInfo
	var ok bool

	client := &dota2.Client{}

	i, err = dota2.NewDOTA2Fantasy(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetPlayerInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerInfo", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*dota2.DOTA2FantasyGetPlayerInfo)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Fantasy_GetProPlayerInfo(t *testing.T) {
	var i *dota2.DOTA2Fantasy
	var err error
	var req *geyser.Request
	var result *dota2.DOTA2FantasyGetProPlayerInfo
	var ok bool

	client := &dota2.Client{}

	i, err = dota2.NewDOTA2Fantasy(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetProPlayerInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetProPlayerInfo", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*dota2.DOTA2FantasyGetProPlayerInfo)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}
