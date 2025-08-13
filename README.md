[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/formatgo/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/formatgo/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/formatgo)](https://pkg.go.dev/github.com/yyle88/formatgo)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/formatgo/master.svg)](https://coveralls.io/github/yyle88/formatgo?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/formatgo.svg)](https://github.com/yyle88/formatgo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/formatgo)](https://goreportcard.com/report/github.com/yyle88/formatgo)

# formatgo

`formatgo` is a Go package that provides utilities for formatting Go source code, whether it's in a byte slice, string, or a file, and even for entire directories containing Go files.

## CHINESE README

[中文说明](README.zh.md)

## Installation

To install the `formatgo` package, you can run the following command:

```bash
go get github.com/yyle88/formatgo
```

## Usage

The package provides several functions for formatting Go code. Below are the main functions that you can use:

### `FormatBytes`

Formats Go source code from a byte slice.

```go
formattedCode, err := formatgo.FormatBytes(code []byte)
```

- `code`: The source code as a byte slice.
- Returns the formatted code as a byte slice or an error if something goes wrong.

### `FormatCode`

Formats Go source code from a string.

```go
formattedCode, err := formatgo.FormatCode(code string)
```

- `code`: The source code as a string.
- Returns the formatted code as a string or an error if something goes wrong.

### `FormatFile`

Formats a Go source code file at the given path.

```go
err := formatgo.FormatFile(path string)
```

- `path`: The path to the Go source code file.
- Returns an error if the formatting fails.

### `FormatRoot`

Formats all Go source files in the specified root directory and its subdirectories.

```go
err := formatgo.FormatRoot(root string)
```

- `root`: The root directory to start formatting files from.
- Returns an error if something goes wrong during the formatting process.

## Example

Here’s a simple example of how to format Go code from a string:

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
		fmt.Println("Error formatting code:", err)
		return
	}
	
	fmt.Println("Formatted Code:", formattedCode)
}
```

---

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repo on GitHub (using the webpage interface).
2. Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. Navigate to the cloned project (`cd repo-name`)
4. Create a feature branch (`git checkout -b feature/xxx`).
5. Stage changes (`git add .`)
6. Commit changes (`git commit -m "Add feature xxx"`).
7. Push to the branch (`git push origin feature/xxx`).
8. Open a pull request on GitHub (on the GitHub webpage).

Please ensure tests pass and include relevant documentation updates.

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with this package!** 🎉

Give me stars. Thank you!!!

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/formatgo.svg?variant=adaptive)](https://starchart.cc/yyle88/formatgo)
