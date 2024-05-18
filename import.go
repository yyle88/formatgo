package formatgo

import "golang.org/x/tools/imports"

func GetImportsOptions() *imports.Options {
	return &imports.Options{
		TabWidth:  4,
		TabIndent: true,
		Comments:  true,
		Fragment:  true,
	}
}
