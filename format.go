package formatgo

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"strings"

	"github.com/yyle88/erero"
	"github.com/yyle88/formatgo/internal/utils"
	"golang.org/x/tools/imports"
)

// FormatBytesWithOptions 格式化golang的源代码，当出错时依然返回中间某个阶段的代码，这样出错时返回值也还是可以用的
func FormatBytesWithOptions(code []byte, options *Options) ([]byte, error) {
	newSrc, err := format.Source(code)
	if err != nil {
		return code, erero.WithMessage(err, "wrong")
	}
	code = newSrc // 存储前一阶段的正确结果

	// 接下来顺带把 imports 整理整理
	newSrc, err = imports.Process("", code, options.ImportsOptions)
	if err != nil {
		return code, erero.WithMessage(err, "wrong")
	}
	return newSrc, nil
}

// FormatCodeWithOptions 格式化源代码字符串，当出错时依然返回中间某个阶段的代码，这样出错时返回值也还是可以用的
func FormatCodeWithOptions(code string, options *Options) (string, error) {
	newSrc, err := FormatBytesWithOptions([]byte(code), options)
	if err != nil {
		return string(newSrc), erero.WithMessage(err, "wrong")
	}
	return string(newSrc), nil
}

// FormatSourceWithOptions 跟 FormatCodeWithOptions 完全是一样的，只是函数名称不同，看外部调用者喜欢哪个名称吧
func FormatSourceWithOptions(code string, options *Options) (string, error) {
	return FormatCodeWithOptions(code, options)
}

// FormatFileWithOptions 格式化源代码文件
func FormatFileWithOptions(path string, options *Options) error {
	source, err := os.ReadFile(path)
	if err != nil {
		return erero.WithMessage(err, "wrong")
	}
	newSrc, err := FormatBytesWithOptions(source, options)
	if err != nil {
		return erero.WithMessage(err, "wrong")
	}
	if bytes.Equal(source, newSrc) {
		return nil
	}
	return utils.WriteFile(path, newSrc)
}

// FormatRootWithOptions 格式化整个目录以及其子目录下的所有go文件
func FormatRootWithOptions(root string, options *RootOptions) error {
	return formatRootWithOptions(root, 0, options)
}

// FormatProjectWithOptions 格式化整个项目里所有的go文件
func FormatProjectWithOptions(projectRoot string, options *RootOptions) error {
	return formatRootWithOptions(projectRoot, 0, options)
}

func formatRootWithOptions(root string, depth int, options *RootOptions) error {
	mapNamePath, err := utils.LsMapNamePath(root)
	if err != nil {
		return erero.WithMessage(err, "wrong")
	}
	for name, path := range mapNamePath {
		if strings.HasPrefix(name, ".") {
			if depth < options.MinSkipHiddenDepth { //在若干层以内跳过隐藏目录
				continue //跳过不可见的目录，比如.git目录和.idea目录
			}
		}

		if utils.IsRootExist(path) {
			if options.FilterRootFunction(depth, path, name) {
				if err := formatRootWithOptions(path, depth+1, options); err != nil {
					return erero.WithMessage(err, "wrong")
				}
			}
		} else if utils.IsFileExist(path) {
			if filepath.Ext(name) == ".go" || utils.IsStringHasAnySuffix(name, options.FileNameSuffixes) {
				if options.FilterFileFunction(depth, path, name) {
					if err := FormatFileWithOptions(path, options.FileOptions); err != nil {
						return erero.WithMessage(err, "wrong")
					}
				}
			}
		}
	}
	return nil
}
