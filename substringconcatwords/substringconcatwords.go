/*
Problem definition: https://leetcode.com/problems/substring-with-concatenation-of-all-words/
*/
package substringconcatwords

import (
	"practice_scripting/permutations"
	"strings"
)

func findMatchingSubstrings(s string, words []string) []int {

	stringPermutes := permutations.FindPermutations(words)

	resultPositions := make([]int, 0)

	for _, str := range stringPermutes {

		idx := strings.Index(s, str)
		if idx >= 0 {
			resultPositions = append(resultPositions, idx)
		}
	}

	return resultPositions
}
