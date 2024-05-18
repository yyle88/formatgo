package formatgo

// FormatBytes 格式化golang的源代码
func FormatBytes(code []byte) ([]byte, error) {
	return FormatBytesWithOptions(code, NewOptions())
}

// FormatCode 格式化源代码字符串
func FormatCode(code string) (string, error) {
	return FormatCodeWithOptions(code, NewOptions())
}

// FormatSource 格式化源代码字符串
func FormatSource(code string) (string, error) {
	return FormatSourceWithOptions(code, NewOptions())
}

// FormatFile 格式化源代码文件
func FormatFile(path string) error {
	return FormatFileWithOptions(path, NewOptions())
}

// FormatRoot 格式化整个目录以及其子目录下的所有go文件
func FormatRoot(root string) error {
	return FormatRootWithOptions(root, NewRootOptions())
}

// FormatProject 格式化整个项目里所有的go文件
func FormatProject(projectRoot string) error {
	return FormatProjectWithOptions(projectRoot, NewRootOptions())
}
