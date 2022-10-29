package permutations

import "testing"

func TestFindPermutations(t *testing.T) {

	var input = []string{"ab", "cd", "ef"}

	results := FindPermutations(input)

	t.Log("Permutation results :::")
	for _, res := range results {
		t.Logf("%s", res)
	}

}
