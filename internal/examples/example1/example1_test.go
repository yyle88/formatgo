package example1

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/rese"
)

func TestShowExampleFileCode(t *testing.T) {
	ShowExampleFileCode()
}

func TestFormatSrcAndCompare(t *testing.T) {
	const code = `
package example1

import (
	"fmt"
	
"github.com/yyle88/rese"

		"github.com/yyle88/runpath"

	"os"
	"time"
)

// 		ShowExampleFileCode 显示Example文件的代码
func ShowExampleFileCode() {
	since := time.Now()

sourceCode := GetExampleFileCode()

		fmt.Println(time.Since(since))

fmt.Println(string(sourceCode))
}

// 		GetExampleFileCode 获取Example文件的代码
func GetExampleFileCode() []byte {
	return rese.V1(os.ReadFile(GetExampleFilePath()))
}

// 		GetExampleFilePath 获取Example文件的路径
func GetExampleFilePath() string {
	return runpath.Path()
}

`
	newCode := rese.C1(formatgo.FormatCode(code))

	t.Log(newCode)

	expectedCode := string(GetExampleFileCode())

	t.Log(expectedCode)

	require.Equal(t, expectedCode, newCode)
}
