// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIPortal2Leaderboards(t *testing.T) {
	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	appIDs := []uint32{620}

	for _, appID := range appIDs {
		ci, err := steam.NewIPortal2Leaderboards(client, appID)

		require.NoError(t, err)
		require.NotNil(t, ci)

		assert.Same(t, client, ci.Client)
		assert.NotNil(t, ci.Interface)
	}
}

func TestIPortal2Leaderboards_GetBucketizedData(t *testing.T) {
	var ci *steam.IPortal2Leaderboards
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIPortal2Leaderboards(client, 620)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetBucketizedData()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetBucketizedData", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
