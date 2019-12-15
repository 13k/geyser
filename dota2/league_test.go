// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package dota2_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/dota2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewLeague(t *testing.T) {
	client := &dota2.Client{}
	i, err := dota2.NewLeague(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	assert.Same(t, client, i.Client)
	assert.NotNil(t, i.Interface)
}

func TestLeague_GetLeagueInfo(t *testing.T) {
	var i *dota2.League
	var err error
	var req *geyser.Request
	var result *dota2.LeagueGetLeagueInfo
	var ok bool

	client := &dota2.Client{}

	i, err = dota2.NewLeague(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetLeagueInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLeagueInfo", req.Method.Name)
		assert.Equal(t, 2, req.Method.Version)
	}

	result, ok = req.Result.(*dota2.LeagueGetLeagueInfo)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}
