// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package geyser_test

import (
	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDOTA2Ticket(t *testing.T) {
	client := &geyser.Client{}
	appIDs := []uint32{570, 205790}

	for _, appID := range appIDs {
		i, err := geyser.NewDOTA2Ticket(client, appID)

		require.NoError(t, err)
		require.NotNil(t, i)

		assert.Same(t, client, i.Client)
		assert.NotNil(t, i.Interface)
	}
}

func TestDOTA2Ticket_ClaimBadgeReward(t *testing.T) {
	var i *geyser.DOTA2Ticket
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2TicketClaimBadgeReward
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Ticket(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.ClaimBadgeReward()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "ClaimBadgeReward", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2TicketClaimBadgeReward)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Ticket(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.ClaimBadgeReward()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "ClaimBadgeReward", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2TicketClaimBadgeReward)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Ticket_GetSteamIDForBadgeID(t *testing.T) {
	var i *geyser.DOTA2Ticket
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2TicketGetSteamIDForBadgeID
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Ticket(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetSteamIDForBadgeID()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSteamIDForBadgeID", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2TicketGetSteamIDForBadgeID)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Ticket(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetSteamIDForBadgeID()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSteamIDForBadgeID", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2TicketGetSteamIDForBadgeID)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Ticket_SetSteamAccountPurchased(t *testing.T) {
	var i *geyser.DOTA2Ticket
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2TicketSetSteamAccountPurchased
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Ticket(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.SetSteamAccountPurchased()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "SetSteamAccountPurchased", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2TicketSetSteamAccountPurchased)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Ticket(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.SetSteamAccountPurchased()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "SetSteamAccountPurchased", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2TicketSetSteamAccountPurchased)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestDOTA2Ticket_SteamAccountValidForBadgeType(t *testing.T) {
	var i *geyser.DOTA2Ticket
	var err error
	var req *geyser.Request
	var result *geyser.DOTA2TicketSteamAccountValidForBadgeType
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewDOTA2Ticket(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.SteamAccountValidForBadgeType()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "SteamAccountValidForBadgeType", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2TicketSteamAccountValidForBadgeType)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewDOTA2Ticket(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.SteamAccountValidForBadgeType()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "SteamAccountValidForBadgeType", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.DOTA2TicketSteamAccountValidForBadgeType)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}
