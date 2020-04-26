// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"github.com/13k/geyser/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIDOTA2Match(t *testing.T) {
	client := &steam.Client{}
	appIDs := []uint32{570, 205790}

	for _, appID := range appIDs {
		iface, err := steam.NewIDOTA2Match(client, appID)

		require.NoError(t, err)
		require.NotNil(t, iface)

		assert.Same(t, client, iface.Client)
		assert.NotNil(t, iface.Interface)
	}
}

func TestIDOTA2Match_GetLeagueListing(t *testing.T) {
	var iface *steam.IDOTA2Match
	var err error
	var req *geyser.Request
	var ok bool

	client := &steam.Client{}

	iface, err = steam.NewIDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetLeagueListing()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetLeagueListing()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLeagueListing", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Match_GetLiveLeagueGames(t *testing.T) {
	var iface *steam.IDOTA2Match
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetLiveLeagueGames()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLiveLeagueGames", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetLiveLeagueGames()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLiveLeagueGames", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Match_GetMatchDetails(t *testing.T) {
	var iface *steam.IDOTA2Match
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetMatchDetails()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchDetails", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetMatchDetails()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchDetails", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Match_GetMatchHistory(t *testing.T) {
	var iface *steam.IDOTA2Match
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetMatchHistory()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchHistory", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetMatchHistory()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchHistory", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Match_GetMatchHistoryBySequenceNum(t *testing.T) {
	var iface *steam.IDOTA2Match
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetMatchHistoryBySequenceNum()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchHistoryBySequenceNum", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetMatchHistoryBySequenceNum()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMatchHistoryBySequenceNum", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Match_GetTeamInfoByTeamID(t *testing.T) {
	var iface *steam.IDOTA2Match
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTeamInfoByTeamID()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTeamInfoByTeamID", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTeamInfoByTeamID()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTeamInfoByTeamID", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Match_GetTopLiveEventGame(t *testing.T) {
	var iface *steam.IDOTA2Match
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTopLiveEventGame()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTopLiveEventGame", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTopLiveEventGame()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTopLiveEventGame", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Match_GetTopLiveGame(t *testing.T) {
	var iface *steam.IDOTA2Match
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTopLiveGame()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTopLiveGame", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTopLiveGame()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTopLiveGame", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Match_GetTopWeekendTourneyGames(t *testing.T) {
	var iface *steam.IDOTA2Match
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTopWeekendTourneyGames()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTopWeekendTourneyGames", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTopWeekendTourneyGames()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTopWeekendTourneyGames", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2Match_GetTournamentPlayerStats(t *testing.T) {
	var iface *steam.IDOTA2Match
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTournamentPlayerStats(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTournamentPlayerStats", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIDOTA2Match(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTournamentPlayerStats(2)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTournamentPlayerStats", req.Method.Name)
		assert.Equal(t, 2, req.Method.Version)
	}

	iface, err = steam.NewIDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTournamentPlayerStats(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTournamentPlayerStats", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIDOTA2Match(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTournamentPlayerStats(2)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTournamentPlayerStats", req.Method.Name)
		assert.Equal(t, 2, req.Method.Version)
	}
}