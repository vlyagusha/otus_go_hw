package logger

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogger(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		require.True(t, true)
	})
}
