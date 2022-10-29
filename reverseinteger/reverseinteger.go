package reverseinteger

import (
	"math"
)

func findReverse(number int) int {

	maxInt := math.MaxInt32
	minInt := math.MinInt32

	resultNumber := 0
	for number != 0 {
		if resultNumber > 0 && number < 0 {
			resultNumber = -1 * resultNumber
		}

		if (resultNumber*10+number%10) > maxInt || (resultNumber*10+number%10) < minInt {
			return 0
		}

		resultNumber = resultNumber*10 + number%10

		number = number / 10
	}

	return resultNumber
}
