/*
Problem definition: given a list of strings, find all combinations of all strings in the list
*/
package combinations

import (
	"math"
)

func findCombinations(words []string) []string {

	resultStrings := make([]string, 0)
	totalLength := math.Pow(2, float64(len(words)))

	for i := 0; i < int(totalLength); i++ {

		resultString := ""
		// break down i into it's binary form and track the strings on that indexes
		val := i
		pos := 0

		for val > 0 {
			if val%2 == 1 {
				resultString = resultString + words[pos]
			}

			pos++
			val = int(val / 2)
		}

		resultStrings = append(resultStrings, resultString)
	}

	return resultStrings
}
