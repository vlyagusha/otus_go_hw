package internalhttp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMiddleware(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		require.True(t, true)
	})
}
