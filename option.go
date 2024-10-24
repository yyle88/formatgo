package formatgo

import (
	"math"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
	"golang.org/x/tools/imports"
)

type Options struct {
	ImportsOptions        *imports.Options
	IsCleanImportNewlines bool
	IsProcessFormatImport bool
}

func NewOptions() *Options {
	return &Options{
		ImportsOptions:        GetImportsOptions(),
		IsCleanImportNewlines: true,
		IsProcessFormatImport: true,
	}
}

type RootOptions struct {
	FileOptions        *Options
	FilterRootFunction func(depth int, path string, name string) bool //只格式化目录名/路径符合的
	FilterFileFunction func(depth int, path string, name string) bool //只格式化文件名/路径符合的，当然前提是目录已经符合
	FileNameSuffixes   []string                                       //要格式化的文件名称的后缀范围，由于是go代码的format，即使是为空也必然包含".go"这个后缀
	MinSkipHiddenDepth int                                            //在多少层深度内跳过隐藏目录，比如第一层的 ".git" ".idea" 目录等
}

func NewRootOptions() *RootOptions {
	return &RootOptions{
		FileOptions: NewOptions(),
		FilterRootFunction: func(depth int, path string, name string) bool {
			zaplog.LOG.Debug("format_root", zap.Int("depth", depth), zap.String("path", path), zap.String("name", name))
			return true
		},
		FilterFileFunction: func(depth int, path string, name string) bool {
			zaplog.LOG.Debug("format_file", zap.Int("depth", depth), zap.String("path", path), zap.String("name", name))
			return true
		},
		FileNameSuffixes: []string{ //默认值设置为空也行，在逻辑里有只要是 ".go" 就算通过，剩下的才通过这个后缀筛选过滤
			".go",
			".GO",
		},
		MinSkipHiddenDepth: math.MaxInt, //直接设置为最大值，就是任何层的隐藏文件都不进行格式化，认为这才是符合99%场景的，当然实际上默认值设置为1也行，毕竟99%的项目除了第0层以外就没有别的隐藏文件啦
	}
}
