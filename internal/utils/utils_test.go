package utils

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFileModePerm(t *testing.T) {
	_, path, _, ok := runtime.Caller(0)
	require.True(t, ok)
	t.Log(path)

	t.Log(GetFileModePerm(path))
}
