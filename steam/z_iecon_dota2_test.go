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

func TestNewIEconDOTA2(t *testing.T) {
	client := &steam.Client{}
	appIDs := []uint32{570, 205790}

	for _, appID := range appIDs {
		iface, err := steam.NewIEconDOTA2(client, appID)

		require.NoError(t, err)
		require.NotNil(t, iface)

		assert.Same(t, client, iface.Client)
		assert.NotNil(t, iface.Interface)
	}
}

func TestIEconDOTA2_GetEventStatsForAccount(t *testing.T) {
	var iface *steam.IEconDOTA2
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIEconDOTA2(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetEventStatsForAccount()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetEventStatsForAccount", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconDOTA2(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetEventStatsForAccount()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetEventStatsForAccount", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIEconDOTA2_GetGameItems(t *testing.T) {
	var iface *steam.IEconDOTA2
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIEconDOTA2(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetGameItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetGameItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconDOTA2(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetGameItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetGameItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIEconDOTA2_GetHeroes(t *testing.T) {
	var iface *steam.IEconDOTA2
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIEconDOTA2(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetHeroes()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetHeroes", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconDOTA2(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetHeroes()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetHeroes", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIEconDOTA2_GetItemCreators(t *testing.T) {
	var iface *steam.IEconDOTA2
	var err error
	var req *geyser.Request
	var ok bool

	client := &steam.Client{}

	iface, err = steam.NewIEconDOTA2(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetItemCreators()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetItemCreators", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconDOTA2(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetItemCreators()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)
}

func TestIEconDOTA2_GetItemIconPath(t *testing.T) {
	var iface *steam.IEconDOTA2
	var err error
	var req *geyser.Request
	var ok bool

	client := &steam.Client{}

	iface, err = steam.NewIEconDOTA2(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetItemIconPath()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconDOTA2(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetItemIconPath()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetItemIconPath", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIEconDOTA2_GetRarities(t *testing.T) {
	var iface *steam.IEconDOTA2
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIEconDOTA2(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetRarities()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetRarities", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconDOTA2(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetRarities()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetRarities", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIEconDOTA2_GetTournamentPrizePool(t *testing.T) {
	var iface *steam.IEconDOTA2
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewIEconDOTA2(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTournamentPrizePool()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTournamentPrizePool", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconDOTA2(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTournamentPrizePool()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTournamentPrizePool", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}