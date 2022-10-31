package enhancedmore

import (
	"testing"
	"time"
)

func TestFindMatchingLinesInFolder(t *testing.T) {

	path := "/Users/anirbanmukherjee/programming/go/src/github.com/tukaianirban/practise.go.codingchallenges/stringsegment"
	bufferSize := 5

	substring := "func"
	go findMatchingLinesInFolder(path, substring, bufferSize)

	time.Sleep(1 * time.Second)

	t.Log("first buffer dump :::")
	lines := getBufferDump()
	for _, line := range lines {
		t.Logf("%s", line)
	}

	setNextPage()
	time.Sleep(1 * time.Second)

	t.Log("second buffer dump :::")
	lines = getBufferDump()
	for _, line := range lines {
		t.Logf("%s", line)
	}

	time.Sleep(10 * time.Second)
}
