package dota2_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/13k/geyser/dota2"
)

func TestNew(t *testing.T) {
	client, err := dota2.New()

	require.NoError(t, err)
	require.NotNil(t, client)
	require.Equal(t, dota2.HostURL, client.HostURL())
}
