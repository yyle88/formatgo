package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/runpath"
)

func TestIsRootExists(t *testing.T) {
	exists, err := IsRootExists(runpath.PARENT.Path())
	require.NoError(t, err)
	require.True(t, exists)
}

func TestIsFileExists(t *testing.T) {
	exists, err := IsFileExists(runpath.Path())
	require.NoError(t, err)
	require.True(t, exists)
}
