package atoi

import (
	"math"
	"strings"
)

func getAtoI(s string) int {

	factor := 1
	idx := 0
	number := 0

	s = strings.TrimLeft(s, " ")
	if len(s) < 1 {
		return 0
	}

	if string(s[0]) == "-" {
		factor = -1
		idx = 1

		if len(s) > 1 && string(s[1]) == "+" {
			return 0
		}
	}

	if string(s[0]) == "+" {
		idx = 1

		if len(s) > 1 && string(s[1]) == "-" {
			return 0
		}
	}

	for ; idx < len(s); idx++ {
		asciival := int(s[idx])

		// if string(s[idx]) == "-" {
		// 	factor = -1
		// 	continue
		// }

		// if string(s[idx]) == "+" {
		// 	factor = 1
		// 	continue
		// }

		if asciival >= int('0') && asciival <= int('9') {

			if (number*10 + factor*(asciival-48)) > math.MaxInt32 {
				return math.MaxInt32
			}

			if (number*10 + factor*(asciival-48)) < math.MinInt32 {
				return math.MinInt32
			}

			number = number*10 + factor*(asciival-48)
			continue
		}

		break
	}

	return number
}
