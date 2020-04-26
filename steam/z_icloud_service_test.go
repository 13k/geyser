// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewICloudService(t *testing.T) {
	client := &steam.Client{}
	iface, err := steam.NewICloudService(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	assert.Same(t, client, iface.Client)
	assert.NotNil(t, iface.Interface)
}

func TestICloudService_BeginHTTPUpload(t *testing.T) {
	var iface *steam.ICloudService
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewICloudService(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.BeginHTTPUpload()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "BeginHTTPUpload", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestICloudService_CommitHTTPUpload(t *testing.T) {
	var iface *steam.ICloudService
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewICloudService(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.CommitHTTPUpload()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "CommitHTTPUpload", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestICloudService_Delete(t *testing.T) {
	var iface *steam.ICloudService
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewICloudService(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.Delete()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "Delete", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestICloudService_EnumerateUserFiles(t *testing.T) {
	var iface *steam.ICloudService
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewICloudService(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.EnumerateUserFiles()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "EnumerateUserFiles", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestICloudService_GetFileDetails(t *testing.T) {
	var iface *steam.ICloudService
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewICloudService(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetFileDetails()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetFileDetails", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestICloudService_GetUploadServerInfo(t *testing.T) {
	var iface *steam.ICloudService
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewICloudService(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetUploadServerInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetUploadServerInfo", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}