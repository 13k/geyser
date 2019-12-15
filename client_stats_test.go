// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package geyser_test

import (
	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewClientStats(t *testing.T) {
	client := &geyser.Client{}
	appIDs := []uint32{1046930}

	for _, appID := range appIDs {
		i, err := geyser.NewClientStats(client, appID)

		require.NoError(t, err)
		require.NotNil(t, i)

		assert.Same(t, client, i.Client)
		assert.NotNil(t, i.Interface)
	}
}

func TestClientStats_ReportEvent(t *testing.T) {
	var i *geyser.ClientStats
	var err error
	var req *geyser.Request
	var result *geyser.ClientStatsReportEvent
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewClientStats(client, 1046930)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.ReportEvent()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "ReportEvent", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.ClientStatsReportEvent)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}
