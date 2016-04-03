package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Newからの戻り値がnilだよ！")
	} else {
		msg := "トレースパッケージ「来ちゃったっ…//」"
		tracer.Trace(msg)
		msg += "\n"
		if buf.String() != msg {
			t.Errorf("ピピーっ！テスト警察ですっ！'%s'は誤った文字列ですよっ！！", buf.String())
		}
	}
}
