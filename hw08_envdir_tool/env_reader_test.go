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
}
