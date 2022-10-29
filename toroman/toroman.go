package toroman

func convertToRoman(number int) string {

	result := ""

	integers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var rem int
	for idx, integer := range integers {

		rem = number / integer
		for count := 1; count <= rem; count++ {
			result = result + romans[idx]
		}

		number = number % integers[idx]
	}

	return result
}
