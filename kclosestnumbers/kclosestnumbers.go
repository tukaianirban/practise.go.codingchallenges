/*
Problem definition: Find ‘k’ closest numbers

Problem Statement: Given a sorted number array and two integers ‘K’ and ‘X’, find ‘K’ closest numbers to ‘X’ in the array.
Return the numbers in the sorted order. ‘X’ is not necessarily present in the array.
*/
package kclosestnumbers

import (
	"fmt"
	"log"
)

/*
Returns:
- immediately lower number index (or index of X)
- immediate higher number index (or index of X)
- error: non-nil if something goes wrong
*/
func findIndexes(numbersList []int, start int, end int, X int) (int, int, error) {

	log.Printf("called with begin=%d end=%d", start, end)

	length := end - start + 1

	if numbersList[start] > X {
		return -1, -1, fmt.Errorf("not found")
	}

	if numbersList[end] < X {
		return -1, -1, fmt.Errorf("not found")
	}

	idx := (end - start) / 2

	if length > 3 {
		if b, e, err := findIndexes(numbersList, start, start+idx+1, X); err == nil && b != -1 && e != -1 {
			return b, e, nil
		}

		if b, e, err := findIndexes(numbersList, start+idx+1, end, X); err == nil && b != -1 && e != -1 {
			return b, e, nil
		}
	}

	if length == 3 {
		for i := start; i <= end; i++ {
			if numbersList[i] == X {
				return i, i, nil
			}
		}

		if numbersList[start] < X && numbersList[start+1] > X {
			return start, start + 1, nil
		}

		if numbersList[start+1] < X && numbersList[end] > X {
			return start + 1, end, nil
		}
	}

	if length == 2 {

		if numbersList[start] == X {
			return start, start, nil
		}

		if numbersList[end] == X {
			return end, end, nil
		}

		return start, end, nil
	}

	if numbersList[start] != X {
		return -1, -1, fmt.Errorf("not found")
	}

	return start, end, nil
}

func findClosest(numbersList []int, K int, X int) []int {

	if X < numbersList[0] {
		return numbersList[:K]
	}

	if X > numbersList[len(numbersList)-1] {
		return numbersList[(len(numbersList) - K):]
	}

	// find location of X in the array
	lower, higher, err := findIndexes(numbersList, 0, len(numbersList)-1, X)
	if err != nil {
		return nil
	}

	if lower > higher {
		return nil
	}

	log.Printf("lower=%d, higher=%d", lower, higher)

	// if higher == lower, then X is present
	if lower == higher {

		if lower == 0 {
			return numbersList[:K]
		}

		if higher == len(numbersList)-1 {
			return numbersList[(len(numbersList) - K):]
		}

		if (numbersList[lower] - numbersList[lower-1]) < (numbersList[lower+1] - numbersList[lower]) {
			var l, h int
			if lower-K+1 < 0 {
				l = 0
			} else {
				l = lower - K + 1
			}

			if lower+1 > len(numbersList) {
				h = len(numbersList)
			} else {
				h = lower + 1
			}

			return numbersList[l:h]
		} else {

			var h int
			if higher+K > len(numbersList) {
				h = len(numbersList)
			} else {
				h = higher + K
			}

			return numbersList[higher:h]
		}
	}

	// if lower < higher, then X is not present, but this is the range to consider
	if (X - numbersList[lower]) < (numbersList[higher] - X) {
		var l, h int
		if lower-K+1 < 0 {
			l = 0
		} else {
			l = lower - K + 1
		}

		if lower+1 > len(numbersList) {
			h = len(numbersList)
		} else {
			h = lower + 1
		}

		return numbersList[l:h]
	} else {
		var h int
		if higher+K > len(numbersList) {
			h = len(numbersList)
		} else {
			h = higher + K
		}

		return numbersList[higher:h]
	}
}
