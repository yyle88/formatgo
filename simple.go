package formatgo

// FormatBytes formats Go source code from a byte slice.
// FormatBytes 格式化golang的源代码
func FormatBytes(code []byte) ([]byte, error) {
	return FormatBytesWithOptions(code, NewOptions())
}

// FormatCode formats Go source code from a string.
// FormatCode 格式化源代码字符串
func FormatCode(code string) (string, error) {
	return FormatCodeWithOptions(code, NewOptions())
}

// FormatFile formats a Go source code file at the given path.
// FormatFile 格式化源代码文件
func FormatFile(path string) error {
	return FormatFileWithOptions(path, NewOptions())
}

// FormatRoot formats all Go source files in the specified root directory and its subdirectories.
// FormatRoot 格式化整个目录以及其子目录下的所有go文件
func FormatRoot(root string) error {
	return FormatRootWithOptions(root, NewRootOptions())
}
