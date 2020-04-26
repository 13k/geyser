// Code generated by geyser. DO NOT EDIT.

package dota2_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/dota2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIDOTA2AutomatedTourney(t *testing.T) {
	client := &dota2.Client{}
	iface, err := dota2.NewIDOTA2AutomatedTourney(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	assert.Same(t, client, iface.Client)
	assert.NotNil(t, iface.Interface)
}

func TestIDOTA2AutomatedTourney_GetParticipationDetails(t *testing.T) {
	var iface *dota2.IDOTA2AutomatedTourney
	var err error
	var req *geyser.Request

	client := &dota2.Client{}

	iface, err = dota2.NewIDOTA2AutomatedTourney(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetParticipationDetails()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetParticipationDetails", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2AutomatedTourney_GetPlayerHistory(t *testing.T) {
	var iface *dota2.IDOTA2AutomatedTourney
	var err error
	var req *geyser.Request

	client := &dota2.Client{}

	iface, err = dota2.NewIDOTA2AutomatedTourney(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetPlayerHistory()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerHistory", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIDOTA2AutomatedTourney_GetTournamentDetails(t *testing.T) {
	var iface *dota2.IDOTA2AutomatedTourney
	var err error
	var req *geyser.Request

	client := &dota2.Client{}

	iface, err = dota2.NewIDOTA2AutomatedTourney(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTournamentDetails()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTournamentDetails", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}