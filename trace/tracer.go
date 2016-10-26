package trace

import (
	"fmt"
	"io"
)

// Tracer はコード内での出来事を記録する
type Tracer struct {
	out io.Writer
}

// Trace は任意のオブジェクト任意個標準出力する
func (t Tracer) Trace(a ...interface{}) {
	if t.out == nil {
		return
	}
	fmt.Fprintln(t.out, a...)
}

// New は新しいTracerのインスタンスを生成する
func New(w io.Writer) Tracer {
	return Tracer{out: w}
}
