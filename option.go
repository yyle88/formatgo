package formatgo

import (
	"math"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
	"golang.org/x/tools/imports"
)

// Options holds configuration options for imports formatting.
// Options 结构体保存了导入格式化的配置选项。
type Options struct {
	ImportsOptions *imports.Options // Options for formatting imports  // 导入格式化的选项
	CondenseImport bool             // Whether to condense imports by removing empty lines // 是否压缩导入部分，去除空行
	IsFormatImport bool             // Whether to format import statements // 是否格式化导入语句
}

// NewOptions creates and returns a new Options instance with default values.
// NewOptions 创建并返回一个带有默认值的 Options 实例。
func NewOptions() *Options {
	return &Options{
		ImportsOptions: NewImportsOptions(), // 使用默认的导入选项创建
		CondenseImport: true,                // 默认开启导入部分压缩
		IsFormatImport: true,                // 默认启用导入格式化
	}
}

// RootOptions holds configuration options for the root directory formatting.
// RootOptions 结构体保存了根目录格式化的配置选项。
type RootOptions struct {
	FileOptions     *Options                                       // File formatting options // 文件格式化选项
	FilterRoot      func(depth int, path string, name string) bool // Filter to format only directories that match certain criteria // 用于过滤目录名/路径的条件，只有符合条件的目录才会被格式化
	FilterFile      func(depth int, path string, name string) bool // Filter to format only files that match certain criteria // 用于过滤文件名/路径的条件，只有符合条件的文件才会被格式化
	FileHasSuffixes []string                                       // Suffixes of files to be formatted // 要格式化的文件后缀列表
	SkipHiddenDepth int                                            // Depth level to skip hidden directories (e.g., .git, .idea) // 在多少层深度内跳过隐藏目录（例如 .git，.idea）
}

// NewRootOptions creates and returns a new RootOptions instance with default values.
// NewRootOptions 创建并返回一个带有默认值的 RootOptions 实例。
func NewRootOptions() *RootOptions {
	return &RootOptions{
		FileOptions: NewOptions(), // 默认使用 NewOptions() 创建 FileOptions
		FilterRoot: func(depth int, path string, name string) bool {
			// Debug log for root filtering
			// 输出根目录过滤的调试日志
			zaplog.LOG.Debug("format_root", zap.Int("depth", depth), zap.String("path", path), zap.String("name", name))
			return true // 默认允许所有根目录
		},
		FilterFile: func(depth int, path string, name string) bool {
			// Debug log for file filtering
			// 输出文件过滤的调试日志
			zaplog.LOG.Debug("format_file", zap.Int("depth", depth), zap.String("path", path), zap.String("name", name))
			return true // 默认允许所有文件
		},
		FileHasSuffixes: []string{ // Default file suffixes to format
			".go", // Go source files
			".GO", // Go source files with uppercase extension
		},
		SkipHiddenDepth: math.MaxInt, // Skip hidden directories at all depths (default behavior)
		// 跳过所有层级的隐藏目录（默认行为）
	}
}
