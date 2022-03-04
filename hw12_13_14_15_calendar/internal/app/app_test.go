package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApp(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		require.True(t, true)
	})
}
