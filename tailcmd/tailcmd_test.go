package tailcmd

import (
	"runtime"
	"testing"
)

func TestGetNLines(t *testing.T) {

	_, err := GetNLines("./samplelogs.txt", 5)
	if err != nil {
		t.Fatalf("error getting lines from file=%s", err.Error())
	}

	// t.Logf("%#v", lines)
}

func Test2GetNLines(t *testing.T) {

	var m1, m2 runtime.MemStats

	runtime.GC()
	runtime.ReadMemStats(&m1)
	_, err := GetNLines("./samplelogs.txt", 5)
	runtime.ReadMemStats(&m2)

	if err != nil {
		t.Errorf("error getting lines from file=%s", err.Error())
	}

	t.Logf("total: %d", m2.TotalAlloc-m1.TotalAlloc)
	t.Logf("mallocs: %d", m2.Mallocs-m1.Mallocs)
}

func BenchmarkGetNLines(b *testing.B) {

	lines, err := GetNLines("./samplelogs.txt", 5)
	if err != nil {
		b.Fatalf("error getting lines from file=%s", err.Error())
	}

	b.Logf("%#v", lines)
}
