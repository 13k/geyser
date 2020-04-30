// Code generated by geyser. DO NOT EDIT.

package dota2_test

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/dota2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIWorkshop(t *testing.T) {
	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err := dota2.NewIWorkshop(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	assert.Same(t, client, ci.Client)
	assert.NotNil(t, ci.Interface)
}

func TestIWorkshop_GetChangeLog(t *testing.T) {
	var ci *dota2.IWorkshop
	var err error
	var req *geyser.Request

	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = dota2.NewIWorkshop(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetChangeLog()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetChangeLog", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
