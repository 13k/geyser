package steam_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/13k/geyser/steam"
)

func TestNew(t *testing.T) {
	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)
	require.Equal(t, steam.HostURL, client.HostURL())
}
