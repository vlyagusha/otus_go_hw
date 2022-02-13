package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		environment, err := ReadDir("./testdata/env")

		require.NoError(t, err)
		require.Equal(t, 5, len(environment))
		require.Equal(t, EnvValue{
			Value:      "bar",
			NeedRemove: false,
		}, environment["BAR"])
		require.Equal(t, EnvValue{
			Value:      "",
			NeedRemove: true,
		}, environment["EMPTY"])
		require.Equal(t, EnvValue{
			Value:      "   foo\nwith new line",
			NeedRemove: false,
		}, environment["FOO"])
		require.Equal(t, EnvValue{
			Value:      "\"hello\"",
			NeedRemove: false,
		}, environment["HELLO"])
		require.Equal(t, EnvValue{
			Value:      "",
			NeedRemove: true,
		}, environment["UNSET"])
	})

	t.Run("error: illegal char in filename", func(t *testing.T) {
		_, err := ReadDir("./testdata/env2")

		require.Error(t, err)
		require.ErrorIs(t, err, ErrIllegalCharInFileName)
	})

	t.Run("error: unsupported file", func(t *testing.T) {
		_, err := ReadDir("./testdata/env3")

		require.Error(t, err)
		require.ErrorIs(t, err, ErrUnsupportedFile)
	})

	t.Run("error: dir doesnt exist", func(t *testing.T) {
		_, err := ReadDir("./testdata/env_doesnt_exist")

		require.Error(t, err)
	})
}
