/*
Problem description: https://www.geeksforgeeks.org/count-number-islands-every-island-separated-line/?ref=lbp
*/
package islandcounter

import (
	"log"
	"strconv"
)

var matrix = [6][3]string{
	{"O", "O", "O"},
	{"X", "X", "O"},
	{"X", "X", "O"},
	{"O", "O", "X"},
	{"O", "O", "X"},
	{"X", "X", "O"},
}

func painter() int {

	count := 0
	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix[i]); j++ {

			if matrix[i][j] == "X" {

				if i > 0 {
					_, err := strconv.Atoi(matrix[i-1][j])
					if err == nil {
						// above cell is a number; paint it here too
						matrix[i][j] = matrix[i-1][j]
						continue
					}
				}

				if j > 0 {
					_, err := strconv.Atoi(matrix[i][j-1])
					if err == nil {
						matrix[i][j] = matrix[i][j-1]
						continue
					}
				}

				count++
				matrix[i][j] = strconv.Itoa(count)

			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		log.Printf("%#v", matrix[i])
	}

	return count
}
