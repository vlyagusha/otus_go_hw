package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("test ErrUnsupportedFile", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "*")
		require.NoError(t, err)

		err = Copy("/dev/urandom", tempFile.Name(), 0, 0)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUnsupportedFile)
	})

	t.Run("test ErrOffsetExceedsFileSize", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "*")
		require.NoError(t, err)

		err = Copy("testdata/input.txt", tempFile.Name(), 8000, 0)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)
	})
}
