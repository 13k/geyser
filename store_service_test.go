// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package geyser_test

import (
	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewStoreService(t *testing.T) {
	client := &geyser.Client{}
	i, err := geyser.NewStoreService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	assert.Same(t, client, i.Client)
	assert.NotNil(t, i.Interface)
}

func TestStoreService_GetAppList(t *testing.T) {
	var i *geyser.StoreService
	var err error
	var req *geyser.Request
	var result *geyser.StoreServiceGetAppList
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewStoreService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetAppList()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetAppList", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.StoreServiceGetAppList)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}