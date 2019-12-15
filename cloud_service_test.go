// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package geyser_test

import (
	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCloudService(t *testing.T) {
	client := &geyser.Client{}
	i, err := geyser.NewCloudService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	assert.Same(t, client, i.Client)
	assert.NotNil(t, i.Interface)
}

func TestCloudService_BeginHTTPUpload(t *testing.T) {
	var i *geyser.CloudService
	var err error
	var req *geyser.Request
	var result *geyser.CloudServiceBeginHTTPUpload
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewCloudService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.BeginHTTPUpload()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "BeginHTTPUpload", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.CloudServiceBeginHTTPUpload)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestCloudService_CommitHTTPUpload(t *testing.T) {
	var i *geyser.CloudService
	var err error
	var req *geyser.Request
	var result *geyser.CloudServiceCommitHTTPUpload
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewCloudService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.CommitHTTPUpload()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "CommitHTTPUpload", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.CloudServiceCommitHTTPUpload)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestCloudService_Delete(t *testing.T) {
	var i *geyser.CloudService
	var err error
	var req *geyser.Request
	var result *geyser.CloudServiceDelete
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewCloudService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.Delete()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "Delete", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.CloudServiceDelete)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestCloudService_EnumerateUserFiles(t *testing.T) {
	var i *geyser.CloudService
	var err error
	var req *geyser.Request
	var result *geyser.CloudServiceEnumerateUserFiles
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewCloudService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.EnumerateUserFiles()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "EnumerateUserFiles", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.CloudServiceEnumerateUserFiles)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestCloudService_GetFileDetails(t *testing.T) {
	var i *geyser.CloudService
	var err error
	var req *geyser.Request
	var result *geyser.CloudServiceGetFileDetails
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewCloudService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetFileDetails()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetFileDetails", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.CloudServiceGetFileDetails)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestCloudService_GetUploadServerInfo(t *testing.T) {
	var i *geyser.CloudService
	var err error
	var req *geyser.Request
	var result *geyser.CloudServiceGetUploadServerInfo
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewCloudService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetUploadServerInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetUploadServerInfo", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.CloudServiceGetUploadServerInfo)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}
