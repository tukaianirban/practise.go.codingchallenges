/*
Problem definition: https://leetcode.com/problems/substring-with-concatenation-of-all-words/
*/
package substringconcatwords

import (
	"strings"

	"practise.go.codingchallenges/permutations"
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
