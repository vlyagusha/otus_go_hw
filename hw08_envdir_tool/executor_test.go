package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		environment, err := ReadDir("./testdata/env")

		require.NoError(t, err)
		require.Equal(t, 5, len(environment))

		_, exists := os.LookupEnv("bar")
		require.False(t, exists)
		_, exists = os.LookupEnv("foo")
		require.False(t, exists)
		_, exists = os.LookupEnv("hello")
		require.False(t, exists)

		code := RunCmd([]string{"echo", "hello world"}, environment)
		require.Equal(t, 0, code)

		_, exists = os.LookupEnv("bar")
		require.False(t, exists)
		_, exists = os.LookupEnv("foo")
		require.False(t, exists)
		_, exists = os.LookupEnv("hello")
		require.False(t, exists)
	})
}
