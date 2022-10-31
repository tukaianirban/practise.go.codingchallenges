package bracketpatterns

import "testing"

func TestFindBracketPatterns(t *testing.T) {

	input := "{{{{}}}}"
	findBracketPatterns(input, "")
}
