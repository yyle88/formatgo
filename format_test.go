package formatgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatBytesWithOptions(t *testing.T) {
	const code = `
		package formatgo
		import (
			"github.com/kr/pretty"
		
			"testing"
		)
		
		func TestGetImportsOptions(t *testing.T) {
			_ ,_ = pretty.Println(GetImportsOptions())
		}
		`

	newSrc, err := FormatBytesWithOptions([]byte(code), NewOptions())
	require.NoError(t, err)

	t.Log(string(newSrc))

	const want = `package formatgo

import (
	"testing"

	"github.com/kr/pretty"
)

func TestGetImportsOptions(t *testing.T) {
	_, _ = pretty.Println(GetImportsOptions())
}
`
	require.Equal(t, want, string(newSrc))
}
