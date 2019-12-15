// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package geyser_test

import (
	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewContentServerDirectoryService(t *testing.T) {
	client := &geyser.Client{}
	i, err := geyser.NewContentServerDirectoryService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	assert.Same(t, client, i.Client)
	assert.NotNil(t, i.Interface)
}

func TestContentServerDirectoryService_GetDepotPatchInfo(t *testing.T) {
	var i *geyser.ContentServerDirectoryService
	var err error
	var req *geyser.Request
	var result *geyser.ContentServerDirectoryServiceGetDepotPatchInfo
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewContentServerDirectoryService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetDepotPatchInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetDepotPatchInfo", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.ContentServerDirectoryServiceGetDepotPatchInfo)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestContentServerDirectoryService_GetServersForSteamPipe(t *testing.T) {
	var i *geyser.ContentServerDirectoryService
	var err error
	var req *geyser.Request
	var result *geyser.ContentServerDirectoryServiceGetServersForSteamPipe
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewContentServerDirectoryService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetServersForSteamPipe()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetServersForSteamPipe", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.ContentServerDirectoryServiceGetServersForSteamPipe)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}