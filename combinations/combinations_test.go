package combinations

import "testing"

func TestFindCombinations(t *testing.T) {

	words := []string{"A", "B", "C"}

	result := findCombinations(words)

	t.Logf("list of all combinations :::")
	for _, str := range result {
		t.Logf("%s", str)
	}
}
