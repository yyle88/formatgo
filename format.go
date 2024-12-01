package formatgo

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"strings"

	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/formatgo/internal/utils"
	"golang.org/x/tools/imports"
)

// FormatBytesWithOptions formats Go source code from a byte slice.
// Even if an error occurs during formatting, it returns the intermediate code.
// FormatBytesWithOptions 格式化 Go 的源代码（以字节切片形式提供）。
// 即使在格式化期间发生错误，也会返回中间结果代码。
func FormatBytesWithOptions(code []byte, options *Options) ([]byte, error) {
	// Step 1: Format source code
	// 步骤 1：格式化源代码
	if newSrc, err := format.Source(code); err != nil {
		return code, erero.Wro(err) // Return the original code if an error occurs // 如果发生错误，返回原始代码
	} else {
		code = newSrc // Save the successful intermediate result // 保存格式化后的中间结果
	}

	// Step 2: Condense import statements if enabled
	// 步骤 2：如果启用，合并导入语句
	if options.CondenseImport {
		if newSrc, err := CleanCodeImportNewlines(code); err != nil {
			return code, erero.Wro(err)
		} else {
			code = newSrc // Save the successful intermediate result // 保存格式化后的导入语句中间结果
		}
	}

	// Step 3: Format imports if enabled
	// 步骤 3：如果启用，格式化导入语句
	if options.IsFormatImport {
		if newSrc, err := imports.Process("", code, options.ImportsOptions); err != nil {
			return code, erero.Wro(err)
		} else {
			code = newSrc // Save the successful intermediate result // 保存格式化后的导入语句中间结果
		}
	}
	return code, nil
}

// FormatCodeWithOptions formats Go source code from a string.
// Even if an error occurs during formatting, it returns the intermediate code as a string.
// FormatCodeWithOptions 格式化 Go 的源代码（以字符串形式提供）。
// 即使在格式化期间发生错误，也会返回中间结果代码（字符串形式）。
func FormatCodeWithOptions(code string, options *Options) (string, error) {
	newSrc, err := FormatBytesWithOptions([]byte(code), options)
	if err != nil {
		return string(newSrc), erero.Wro(err) // Return intermediate result even on error // 即使发生错误，仍然返回中间结果
	}
	return string(newSrc), nil
}

// FormatFileWithOptions formats a Go source code file.
// It reads the file, formats it, and writes back the result if changes are made.
// FormatFileWithOptions 格式化一个 Go 源代码文件。
// 它会读取文件内容，进行格式化，如果内容有变化则写回结果。
func FormatFileWithOptions(path string, options *Options) error {
	source, err := os.ReadFile(path)
	if err != nil {
		return erero.Wro(err)
	}
	newSrc, err := FormatBytesWithOptions(source, options)
	if err != nil {
		return erero.Wro(err)
	}
	// Skip writing if no changes are detected
	// 如果没有变化，跳过写入
	if bytes.Equal(source, newSrc) {
		return nil
	}
	return utils.WriteFileKeepMode(path, newSrc) // Write the formatted content
	// 写入格式化后的内容
}

// FormatRootWithOptions formats all Go files in a directory and its subdirectories.
// It recursively processes each directory and applies the provided options.
// FormatRootWithOptions 格式化指定目录及其子目录下的所有 Go 文件。
// 它递归处理每个目录，并根据提供的选项进行格式化。
func FormatRootWithOptions(root string, options *RootOptions) error {
	return formatRootRecursive(root, 0, options)
}

// formatRootRecursive is a helper function for FormatRootWithOptions.
// It recursively processes directories and files based on the given options.
// formatRootRecursive 是 FormatRootWithOptions 的辅助函数。
// 它根据给定的选项递归处理目录和文件。
func formatRootRecursive(root string, depth int, options *RootOptions) error {
	mapNamePath, err := utils.LsAsMap(root)
	if err != nil {
		return erero.Wro(err)
	}

	for name, path := range mapNamePath {
		// Skip hidden directories/files based on depth
		// 根据深度跳过隐藏的目录/文件
		if strings.HasPrefix(name, ".") {
			if depth < options.SkipHiddenDepth {
				continue // Skip hidden directories like .git or .idea // 跳过隐藏的目录，如 .git 或 .idea
			}
		}

		// Process subdirectories
		// 处理子目录
		if done.VBE(utils.IsRootExists(path)).Done() {
			if options.FilterRoot(depth, path, name) {
				if err := formatRootRecursive(path, depth+1, options); err != nil {
					return erero.Wro(err)
				}
			}
		} else if done.VBE(utils.IsFileExists(path)).Done() {
			// Check if the file is a Go file or matches the specified suffixes
			// 检查文件是否为 Go 文件，或是否与指定的后缀匹配
			if filepath.Ext(name) == ".go" || utils.HasAnySuffix(name, options.FileHasSuffixes) {
				if options.FilterFile(depth, path, name) {
					if err := FormatFileWithOptions(path, options.FileOptions); err != nil {
						return erero.Wro(err)
					}
				}
			}
		}
	}
	return nil
}
