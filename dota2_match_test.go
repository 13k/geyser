// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package geyser_test

import (
	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDOTA2Match(t *testing.T) {
	client := &geyser.Client{}
	appIDs := []uint32{570, 205790}

	for _, appID := range appIDs {
		i, err := geyser.NewDOTA2Match(client, appID)

		require.NoError(t, err)
		require.NotNil(t, i)

		assert.Same(t, client, i.Client)
		assert.NotNil(t, i.Interface)
	}
}

func TestDOTA2Match_GetLeagueListing(t *testing.T) {
	var i *geyser.DOTA2Match
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2MatchGetLeagueListing
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetLeagueListing()

	require.Error(t, err)

	_, ok = err.(*geyser.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	i, err = geyser.NewDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetLeagueListing()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLeagueListing", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetLeagueListing)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Match_GetLiveLeagueGames(t *testing.T) {
	var i *geyser.DOTA2Match
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2MatchGetLiveLeagueGames
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetLiveLeagueGames()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLiveLeagueGames", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetLiveLeagueGames)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetLiveLeagueGames()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLiveLeagueGames", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetLiveLeagueGames)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Match_GetMatchDetails(t *testing.T) {
	var i *geyser.DOTA2Match
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2MatchGetMatchDetails
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetMatchDetails()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchDetails", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetMatchDetails)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetMatchDetails()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchDetails", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetMatchDetails)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Match_GetMatchHistory(t *testing.T) {
	var i *geyser.DOTA2Match
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2MatchGetMatchHistory
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetMatchHistory()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchHistory", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetMatchHistory)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetMatchHistory()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchHistory", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetMatchHistory)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Match_GetMatchHistoryBySequenceNum(t *testing.T) {
	var i *geyser.DOTA2Match
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2MatchGetMatchHistoryBySequenceNum
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetMatchHistoryBySequenceNum()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchHistoryBySequenceNum", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetMatchHistoryBySequenceNum)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetMatchHistoryBySequenceNum()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchHistoryBySequenceNum", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetMatchHistoryBySequenceNum)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Match_GetTeamInfoByTeamID(t *testing.T) {
	var i *geyser.DOTA2Match
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2MatchGetTeamInfoByTeamID
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetTeamInfoByTeamID()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTeamInfoByTeamID", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetTeamInfoByTeamID)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetTeamInfoByTeamID()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTeamInfoByTeamID", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetTeamInfoByTeamID)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Match_GetTopLiveEventGame(t *testing.T) {
	var i *geyser.DOTA2Match
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2MatchGetTopLiveEventGame
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Match(client, 570)

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

	result, ok = req.Result.(*geyser.DOTA2MatchGetTopLiveEventGame)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Match(client, 205790)

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

	result, ok = req.Result.(*geyser.DOTA2MatchGetTopLiveEventGame)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Match_GetTopLiveGame(t *testing.T) {
	var i *geyser.DOTA2Match
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2MatchGetTopLiveGame
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Match(client, 570)

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

	result, ok = req.Result.(*geyser.DOTA2MatchGetTopLiveGame)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Match(client, 205790)

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

	result, ok = req.Result.(*geyser.DOTA2MatchGetTopLiveGame)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Match_GetTopWeekendTourneyGames(t *testing.T) {
	var i *geyser.DOTA2Match
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2MatchGetTopWeekendTourneyGames
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Match(client, 570)

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

	result, ok = req.Result.(*geyser.DOTA2MatchGetTopWeekendTourneyGames)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Match(client, 205790)

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

	result, ok = req.Result.(*geyser.DOTA2MatchGetTopWeekendTourneyGames)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Match_GetTournamentPlayerStats(t *testing.T) {
	var i *geyser.DOTA2Match
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2MatchGetTournamentPlayerStats
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetTournamentPlayerStats(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTournamentPlayerStats", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetTournamentPlayerStats)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetTournamentPlayerStats(2)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTournamentPlayerStats", req.Method.Name)
		assert.Equal(t, 2, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetTournamentPlayerStats)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetTournamentPlayerStats(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTournamentPlayerStats", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetTournamentPlayerStats)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetTournamentPlayerStats(2)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTournamentPlayerStats", req.Method.Name)
		assert.Equal(t, 2, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2MatchGetTournamentPlayerStats)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}
