package main

import (
	"crypto/sha256"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("test ErrUnsupportedFile", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "go_hw07")
		require.NoError(t, err)

		err = Copy("/dev/urandom", tempFile.Name(), 0, 0)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUnsupportedFile)
	})

	t.Run("test ErrOffsetExceedsFileSize", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "go_hw07")
		require.NoError(t, err)

		err = Copy("testdata/input.txt", tempFile.Name(), 8000, 0)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("test empty file", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "go_hw07")
		require.NoError(t, err)

		err = Copy("testdata/input_empty.txt", tempFile.Name(), 0, 0)
		require.NoError(t, err)
		stat, err := tempFile.Stat()
		require.NoError(t, err)
		require.Equal(t, int64(0), stat.Size())
	})

	t.Run("test offset 0 limit 0", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "go_hw07")
		require.NoError(t, err)

		err = Copy("testdata/input.txt", tempFile.Name(), 0, 0)
		require.NoError(t, err)

		inputFile, err := os.OpenFile("testdata/out_offset0_limit0.txt", os.O_RDONLY, 0o644)
		require.NoError(t, err)

		inputFileHash := sha256.New()
		_, err = io.Copy(inputFileHash, inputFile)
		require.NoError(t, err)

		outFileHash := sha256.New()
		_, err = io.Copy(outFileHash, tempFile)
		require.NoError(t, err)
		require.Equal(t, inputFileHash.Sum(nil), outFileHash.Sum(nil))
	})

	t.Run("test offset 0 limit 10", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "go_hw07")
		require.NoError(t, err)

		err = Copy("testdata/input.txt", tempFile.Name(), 0, 10)
		require.NoError(t, err)

		inputFile, err := os.OpenFile("testdata/out_offset0_limit10.txt", os.O_RDONLY, 0o644)
		require.NoError(t, err)

		inputFileHash := sha256.New()
		_, err = io.Copy(inputFileHash, inputFile)
		require.NoError(t, err)

		outFileHash := sha256.New()
		_, err = io.Copy(outFileHash, tempFile)
		require.NoError(t, err)
		require.Equal(t, inputFileHash.Sum(nil), outFileHash.Sum(nil))
	})

	t.Run("test offset 0 limit 1000", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "go_hw07")
		require.NoError(t, err)

		err = Copy("testdata/input.txt", tempFile.Name(), 0, 1000)
		require.NoError(t, err)

		inputFile, err := os.OpenFile("testdata/out_offset0_limit1000.txt", os.O_RDONLY, 0o644)
		require.NoError(t, err)

		inputFileHash := sha256.New()
		_, err = io.Copy(inputFileHash, inputFile)
		require.NoError(t, err)

		outFileHash := sha256.New()
		_, err = io.Copy(outFileHash, tempFile)
		require.NoError(t, err)
		require.Equal(t, inputFileHash.Sum(nil), outFileHash.Sum(nil))
	})

	t.Run("test offset 0 limit 10000", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "go_hw07")
		require.NoError(t, err)

		err = Copy("testdata/input.txt", tempFile.Name(), 0, 10000)
		require.NoError(t, err)

		inputFile, err := os.OpenFile("testdata/out_offset0_limit10000.txt", os.O_RDONLY, 0o644)
		require.NoError(t, err)

		inputFileHash := sha256.New()
		_, err = io.Copy(inputFileHash, inputFile)
		require.NoError(t, err)

		outFileHash := sha256.New()
		_, err = io.Copy(outFileHash, tempFile)
		require.NoError(t, err)
		require.Equal(t, inputFileHash.Sum(nil), outFileHash.Sum(nil))
	})

	t.Run("test offset 100 limit 1000", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "go_hw07")
		require.NoError(t, err)

		err = Copy("testdata/input.txt", tempFile.Name(), 100, 1000)
		require.NoError(t, err)

		inputFile, err := os.OpenFile("testdata/out_offset100_limit1000.txt", os.O_RDONLY, 0o644)
		require.NoError(t, err)

		inputFileHash := sha256.New()
		_, err = io.Copy(inputFileHash, inputFile)
		require.NoError(t, err)

		outFileHash := sha256.New()
		_, err = io.Copy(outFileHash, tempFile)
		require.NoError(t, err)
		require.Equal(t, inputFileHash.Sum(nil), outFileHash.Sum(nil))
	})

	t.Run("test offset 6000 limit 1000", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "go_hw07")
		require.NoError(t, err)

		err = Copy("testdata/input.txt", tempFile.Name(), 6000, 1000)
		require.NoError(t, err)

		inputFile, err := os.OpenFile("testdata/out_offset6000_limit1000.txt", os.O_RDONLY, 0o644)
		require.NoError(t, err)

		inputFileHash := sha256.New()
		_, err = io.Copy(inputFileHash, inputFile)
		require.NoError(t, err)

		outFileHash := sha256.New()
		_, err = io.Copy(outFileHash, tempFile)
		require.NoError(t, err)
		require.Equal(t, inputFileHash.Sum(nil), outFileHash.Sum(nil))
	})
}
