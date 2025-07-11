package formatgo

import (
	"os"
	"strings"

	"github.com/yyle88/erero"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
	"golang.org/x/tools/imports"
)

// NewImportsOptions returns a new imports.Options with default settings.
// NewImportsOptions 返回一个具有默认设置的 imports.Options。
func NewImportsOptions() *imports.Options {
	return &imports.Options{
		TabWidth:  4,    // 设置制表符宽度为 4
		TabIndent: true, // 启用制表符缩进
		Comments:  true, // 保留注释
		Fragment:  true, // 启用代码片段
	}
}

// CleanFileImportNewlines reads a Go source file, condenses the import section
// by removing consecutive empty lines, and writes the result back to the file.
// CleanFileImportNewlines 读取 Go 源代码文件，压缩导入部分
// 通过去除连续的空行并将结果写回文件。
func CleanFileImportNewlines(path string) error {
	source, err := os.ReadFile(path)
	if err != nil {
		return erero.Wro(err) // 返回读取文件时的错误
	}

	// Condense the import lines to remove empty lines
	// 压缩导入语句以去除空行
	newSrc, err := CleanCodeImportNewlines(source)
	if err != nil {
		return erero.Wro(err) // 返回压缩导入行时的错误
	}

	// Write the new source code back to the file
	// 将新源代码写回文件
	// (when file exist) WriteFile truncates it before writing, without changing permissions.
	err = os.WriteFile(path, newSrc, 0644)
	if err != nil {
		return erero.Wro(err) // 返回写入文件时的错误
	}

	return nil
}

// CleanCodeImportNewlines processes the source code and condenses consecutive
// empty lines in the import section to a single newline.
// CleanCodeImportNewlines 处理源代码，并将导入部分中的连续空行压缩为单一的换行符。
func CleanCodeImportNewlines(source []byte) ([]byte, error) {
	astBundle, err := syntaxgo_ast.NewAstBundleV1(source)
	if err != nil {
		return nil, erero.Wro(err) // 返回解析 AST 时的错误
	}

	astFile, _ := astBundle.GetBundle()

	// If no imports exist, return the source as it is
	// 如果没有导入语句，则直接返回原始源代码
	if len(astFile.Imports) == 0 {
		return source, nil
	}

	// Create a new node that covers the range of import statements
	// 创建一个新的节点，覆盖所有导入语句的范围
	node := syntaxgo_astnode.NewNode(
		astFile.Imports[0].Pos(),
		astFile.Imports[len(astFile.Imports)-1].End(),
	)

	// Get the text of the import section
	// 获取导入部分的文本
	oldImports := node.GetText(source)
	if oldImports == "" {
		return source, nil
	}

	// Condense consecutive newlines into a single newline in the import section
	// 在导入部分将连续的换行符压缩为一个换行符
	newImports := condenseLines(oldImports)
	if oldImports == newImports {
		return source, nil // No changes made, return the source as it is // 如果没有变化，则返回原始源代码
	}

	// Change the import section to the condensed version
	// 将导入部分替换为压缩后的版本
	return syntaxgo_astnode.ChangeNodeCode(source, node, []byte(newImports)), nil
}

// condenseLines replaces multiple consecutive newlines with a single newline.
// It iteratively replaces instances of double newlines until only single newlines remain.
// condenseLines 将多个连续的换行符替换为单个换行符。
// 它通过迭代地替换双换行符，直到只剩下单个换行符。
func condenseLines(source string) string {
	for origin := ""; origin != source; source = strings.ReplaceAll(source, "\n\n", "\n") {
		origin = source
	}
	return source
}
