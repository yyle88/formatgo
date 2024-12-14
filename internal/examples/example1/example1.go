package example1

import (
	"fmt"
	"os"
	"time"

	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
)

// ShowExampleFileCode 显示Example文件的代码
func ShowExampleFileCode() {
	since := time.Now()

	sourceCode := GetExampleFileCode()

	fmt.Println(time.Since(since))

	fmt.Println(string(sourceCode))
}

// GetExampleFileCode 获取Example文件的代码
func GetExampleFileCode() []byte {
	return rese.V1(os.ReadFile(GetExampleFilePath()))
}

// GetExampleFilePath 获取Example文件的路径
func GetExampleFilePath() string {
	return runpath.Path()
}
