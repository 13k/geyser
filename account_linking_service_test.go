// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package geyser_test

import (
	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewAccountLinkingService(t *testing.T) {
	client := &geyser.Client{}
	i, err := geyser.NewAccountLinkingService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	assert.Same(t, client, i.Client)
	assert.NotNil(t, i.Interface)
}

func TestAccountLinkingService_GetLinkedAccountInfo(t *testing.T) {
	var i *geyser.AccountLinkingService
	var err error
	var req *geyser.Request
	var result *geyser.AccountLinkingServiceGetLinkedAccountInfo
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewAccountLinkingService(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetLinkedAccountInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetLinkedAccountInfo", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.AccountLinkingServiceGetLinkedAccountInfo)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}
