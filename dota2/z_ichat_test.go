// Code generated by geyser. DO NOT EDIT.

package dota2_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/dota2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIChat(t *testing.T) {
	client := &dota2.Client{}
	iface, err := dota2.NewIChat(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	assert.Same(t, client, iface.Client)
	assert.NotNil(t, iface.Interface)
}

func TestIChat_GetChannelMembers(t *testing.T) {
	var iface *dota2.IChat
	var err error
	var req *geyser.Request

	client := &dota2.Client{}

	iface, err = dota2.NewIChat(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetChannelMembers()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetChannelMembers", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}