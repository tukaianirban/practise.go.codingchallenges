package threesum

func find3Sums(numbers []int) [][]int {

	results := make([][]int, 0)

	for i := 0; i < len(numbers)-2; i++ {

		for j := i + 1; j < len(numbers)-1; j++ {

			for k := j + 1; k < len(numbers); k++ {

				if numbers[i]+numbers[j]+numbers[k] == 0 {
					result := [3]int{numbers[i], numbers[j], numbers[k]}
					results = append(results, result[:])
				}
			}
		}
	}

	return results
}
