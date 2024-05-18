package formatgo

import (
	"bytes"
	"go/format"
	"os"
	"strings"

	"github.com/yyle88/erero"
	"github.com/yyle88/formatgo/internal/utils"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
	"golang.org/x/tools/imports"
)

// FormatBytes 格式化golang的源代码，当出错时依然返回中间某个阶段的代码，这样出错时返回值也还是可以用的
func FormatBytes(code []byte) ([]byte, error) {
	options := &Options{
		ImportsOptions: GetImportsOptions(),
	}
	return Format(code, options)
}

type Options struct {
	ImportsOptions *imports.Options
}

func Format(code []byte, options *Options) ([]byte, error) {
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

func FormatCode(code string) (string, error) {
	return FormatSource(code)
}

func FormatSource(code string) (string, error) {
	newSrc, err := FormatBytes([]byte(code))
	if err != nil {
		return string(newSrc), erero.WithMessage(err, "wrong")
	}
	return string(newSrc), nil
}

func FormatFile(path string) error {
	code, err := os.ReadFile(path)
	if err != nil {
		return erero.WithMessage(err, "wrong")
	}
	newSrc, err := FormatBytes(code)
	if err != nil {
		return erero.WithMessage(err, "wrong")
	}
	if bytes.Equal(code, newSrc) {
		return nil
	}
	return utils.WriteBytes(path, newSrc)
}

func FormatRoot(root string) error {
	return FormatRootChoose(root, chooseFunc)
}

func FormatRootChoose(root string, choose func(name string, path string) bool) error {
	var suffixes = []string{".go", ".GO"}
	if err := utils.FilepathWalkOnFilesWithSuffixes(root, suffixes, func(path string, info os.FileInfo) error {
		name := info.Name()
		if choose(name, path) {
			if err := FormatFile(path); err != nil {
				return erero.WithMessage(err, "wrong")
			}
		}
		return nil
	}); err != nil {
		return erero.WithMessage(err, "wrong")
	}
	return nil
}

func chooseFunc(name string, path string) bool {
	zaplog.LOG.Debug("format", zap.String("path", path), zap.String("name", name))
	return true
}

// FormatProject 除了根目录的隐藏文件，格式化其余全部代码
func FormatProject(projectPath string) error {
	return FormatProjectChoose(projectPath, chooseFunc)
}

func FormatProjectChoose(projectPath string, choose func(name string, path string) bool) error {
	mapNamePath, err := utils.LsMapNamePath(projectPath)
	if err != nil {
		return erero.WithMessage(err, "wrong")
	}
	for name, path := range mapNamePath {
		if strings.HasPrefix(name, ".") {
			continue //跳过不可见的目录，比如.git目录和.idea目录
		}
		if utils.IsRootExist(path) {
			if err := FormatRootChoose(path, choose); err != nil {
				return erero.WithMessage(err, "wrong")
			}
			continue
		} else if utils.IsFileExist(path) && strings.HasSuffix(name, ".go") {
			if err := FormatFile(path); err != nil {
				return erero.WithMessage(err, "wrong")
			}
			continue
		}
	}
	return nil
}
