package utils

import (
	"testing"

	"github.com/yyle88/runpath"
)

func TestGetFileMode(t *testing.T) {
	t.Log(GetFileMode(runpath.Path()))
}
