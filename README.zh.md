# formatgo

`formatgo` 是一个 Go 包，用于格式化 Go 源代码，无论是字节切片、字符串还是文件，甚至是包含 Go 文件的整个目录。

## 英文文档

[English README](README.md)

## 安装

你可以通过以下命令安装 `formatgo` 包：

```bash
go get github.com/yyle88/formatgo
```

## 使用方法

该包提供了多个函数来格式化 Go 代码，下面是主要的函数及其用法：

### `FormatBytes`

从字节切片格式化 Go 源代码。

```go
formattedCode, err := formatgo.FormatBytes(code []byte)
```

- `code`: Go 源代码（字节切片）。
- 返回值：格式化后的代码（字节切片）或者格式化出错时的错误。

### `FormatCode`

从字符串格式化 Go 源代码。

```go
formattedCode, err := formatgo.FormatCode(code string)
```

- `code`: Go 源代码（字符串）。
- 返回值：格式化后的代码（字符串）或者格式化出错时的错误。

### `FormatFile`

格式化指定路径下的 Go 源代码文件。

```go
err := formatgo.FormatFile(path string)
```

- `path`: Go 源代码文件的路径。
- 返回值：格式化失败时的错误。

### `FormatRoot`

格式化指定根目录及其子目录下的所有 Go 源代码文件。

```go
err := formatgo.FormatRoot(root string)
```

- `root`: 要开始格式化的根目录路径。
- 返回值：格式化过程中发生错误时的错误。

## 示例

以下是一个简单的例子，演示如何从字符串格式化 Go 代码：

```go
package main

import (
	"fmt"
	"github.com/yyle88/formatgo"
)

func main() {
	code := `package main

import "fmt"

func main() {fmt.Println("Hello, world!")}`
	
	formattedCode, err := formatgo.FormatCode(code)
	if err != nil {
		fmt.Println("格式化代码时出错:", err)
		return
	}
	
	fmt.Println("格式化后的代码:", formattedCode)
}
```

---

## 许可证类型

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

---

## 贡献新代码

非常欢迎贡献代码！贡献流程：

1. 在 GitHub 上 Fork 仓库 （通过网页界面操作）。
2. 克隆Forked项目 (`git clone https://github.com/yourname/repo-name.git`)。
3. 在克隆的项目里 (`cd repo-name`)
4. 创建功能分支（`git checkout -b feature/xxx`）。
5. 添加代码 (`git add .`)。
6. 提交更改（`git commit -m "添加功能 xxx"`）。
7. 推送分支（`git push origin feature/xxx`）。
8. 发起 Pull Request （通过网页界面操作）。

请确保测试通过并更新相关文档。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
