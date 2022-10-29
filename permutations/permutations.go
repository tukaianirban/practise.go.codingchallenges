/**
Problem statement: Given an array of strings, print all permutations of the list
**/
package permutations

import (
	"strings"
)

func FindPermutations(input []string) []string {

	chStringPermutes := make(chan string)
	go _recursePermutations(input, chStringPermutes, 0)

	resultStrings := make([]string, 0)
	for res := range chStringPermutes {
		if res == "" {
			break
		}

		resultStrings = append(resultStrings, res)
	}

	return resultStrings
}

func _recursePermutations(input []string, results chan string, index int) {

	// find permutations by swaps with other indexes
	for i := index; i < len(input); i++ {
		newResult := make([]string, len(input))

		copy(newResult, input)

		// swap between pos: index and i
		if i != index {
			str := newResult[index]
			newResult[index] = newResult[i]
			newResult[i] = str
		}

		// log.Printf("index=%d swap at i=%d calling recursion for new string=%#v", index, i, newResult)

		if index >= len(input)-2 {
			results <- strings.Join(newResult, "")
		} else {
			_recursePermutations(newResult, results, index+1)
		}
	}

	if index == 0 {
		results <- ""
	}
}
