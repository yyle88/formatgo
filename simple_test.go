package formatgo

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatCode(t *testing.T) {
	const code = `
		package main

		import (
            "fmt"
            "strconv"
        )
		import "time"

		//这是main函数的注释
		func main() {
			fmt.Println("abc") //随便打印个字符串
			fmt.Println(time.Now()) //随便打印当前时间
			fmt.Println(strconv.Itoa(1))
		}
	`

	t.Log(code)

	newCode, err := FormatCode(code)
	require.NoError(t, err)
	t.Log(newCode)
}

func TestFormatFile(t *testing.T) {
	_, path, _, ok := runtime.Caller(0)
	require.True(t, ok)
	t.Log(path)
	require.NoError(t, FormatFile(path))
}

func TestFormatProject(t *testing.T) {
	_, path, _, ok := runtime.Caller(0)
	require.True(t, ok)
	t.Log(path)
	root := filepath.Dir(path)
	t.Log(root)
	require.True(t, strings.HasSuffix(root, "/formatgo")) //这样99.99%能够确保目录路径是正确的
	//既然已经开发了项目格式化的代码，就把本项目自己格式化下吧
	require.NoError(t, FormatProject(root))
}
