/*
Problem definition: https://leetcode.com/problems/first-missing-positive/
*/
package firstmissingpositive

import "log"

func findFirstMissingPositive(input []int) int {

	// for every number in the array,
	// try to put the number in its index position in the array
	for i := 0; i < len(input); i++ {

		if input[i] <= 0 || input[i] > len(input) {
			continue
		}

		if input[i] == i+1 {
			continue
		}

		// swap
		t := input[i]
		input[i] = input[t-1]
		input[t-1] = t

		// testing : print array after swaps
		log.Printf("array after swapping=%#v", input)
	}

	for i := 0; i < len(input); i++ {
		if input[i] != i+1 {
			return i + 1
		}
	}

	return len(input)
}
