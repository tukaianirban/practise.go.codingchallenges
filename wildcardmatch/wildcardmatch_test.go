package wildcardmatch

import "testing"

func TestMatchPattern(t *testing.T) {

	s := "aa"
	p := "*9"

	t.Logf("pattern matching result=%t", matchPattern(s, p))
}
