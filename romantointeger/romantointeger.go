package romantointeger

func convertRomanToInteger(s string) int {

	result := 0
	integers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for idx, roman := range romans {

		for len(s) >= len(roman) && string(s[0:len(roman)]) == roman {
			result = result + integers[idx]
			s = s[len(roman):]
		}
	}

	return result
}
