package atomic

import (
	"testing"
)

func TestInt32SetAndGet(t *testing.T) {
	var v Int32
	v.Set(10)

	got := v.Get()

	if got != int32(10) {
		t.Errorf("failed, got=<%d> expect=<%d>",
			got,
			10)
	}
}
