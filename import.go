package formatgo

import (
	"os"
	"strings"

	"github.com/yyle88/erero"
	"github.com/yyle88/formatgo/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"golang.org/x/tools/imports"
)

func GetImportsOptions() *imports.Options {
	return &imports.Options{
		TabWidth:  4,
		TabIndent: true,
		Comments:  true,
		Fragment:  true,
	}
}

func CleanImportNewlinesInFile(path string) error {
	source, err := os.ReadFile(path)
	if err != nil {
		return erero.Wro(err)
	}
	newSrc, err := CleanImportNewlines(source)
	if err != nil {
		return erero.Wro(err)
	}
	err = utils.WriteFileKeepFileMode(path, newSrc)
	if err != nil {
		return erero.Wro(err)
	}
	return nil
}

// CleanImportNewlines processes the provided source code in byte form,
// focusing on the import section of a Go file. It removes consecutive
// empty lines within the import statements, ensuring that there is at
// most one newline between them.
func CleanImportNewlines(source []byte) ([]byte, error) {
	astFile, err := syntaxgo_ast.NewAstFromSource(source)
	if err != nil {
		return nil, erero.Wro(err)
	}
	if len(astFile.Imports) == 0 {
		return source, nil
	}

	node := syntaxgo_ast.NewNode(
		astFile.Imports[0].Pos(),
		astFile.Imports[len(astFile.Imports)-1].End(),
	)
	oldImports := node.GetCode(source)
	if oldImports == "" {
		return source, nil
	}
	newImports := condenseNewlines(oldImports)
	if oldImports == newImports {
		return source, nil
	}

	return syntaxgo_ast.ChangeNodeBytes(source, node, []byte(newImports)), nil
}

// condenseNewlines takes a string input and removes consecutive empty lines,
// condensing multiple newlines into a single newline.
// The function iteratively replaces instances of double newlines with a single
// newline until no more consecutive newlines remain,
// ensuring that the output string contains at most one newline between lines。
func condenseNewlines(source string) string {
	for origin := ""; origin != source; source = strings.ReplaceAll(source, "\n\n", "\n") {
		origin = source
	}
	return source
}
