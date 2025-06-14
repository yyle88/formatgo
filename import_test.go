package formatgo

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/kr/pretty"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
)

func TestGetImportsOptions(t *testing.T) {
	_, _ = pretty.Println(NewImportsOptions())
}

func TestNewImportsOptions(t *testing.T) {
	spew.Dump(NewImportsOptions())
}

func TestSeeImportsOptions(t *testing.T) {
	data := rese.A1(json.Marshal(NewImportsOptions()))
	// 这里先这样使用，将来可以用 neatjson 包，因为 neatjson 包引用该包，因此该项目不引用neatjson包
	var result bytes.Buffer
	must.Done(json.Indent(&result, data, "", "\t"))
	t.Log(result.String())
}
