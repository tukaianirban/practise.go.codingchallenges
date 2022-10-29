package hourglasssum

// var matrix = [5][5]int{{1, 1, 1, 0, 0},
// 	{0, 1, 0, 0, 0},
// 	{1, 1, 1, 0, 0},
// 	{0, 0, 0, 0, 0},
// 	{0, 0, 0, 0, 0},
// }

var matrix = [5][5]int{{0, 3, 0, 0, 0},
	{0, 1, 0, 0, 0},
	{1, 1, 1, 0, 0},
	{0, 0, 2, 4, 4},
	{0, 0, 0, 2, 4}}

func findSum(x, y int) int {

	if x == 0 || x == len(matrix)-1 || y == 0 || y == len(matrix[0])-1 {
		return -1
	}

	sum := matrix[x-1][y-1] + matrix[x-1][y] + matrix[x-1][y+1] + matrix[x][y]
	sum = sum + matrix[x+1][y-1] + matrix[x+1][y] + matrix[x+1][y+1]
	return sum
}

func findMaxSum() int {

	maxSum := 0

	for x := 1; x < len(matrix)-1; x++ {

		for y := 1; y < len(matrix[0])-1; y++ {

			newSum := findSum(x, y)
			if newSum > maxSum {
				maxSum = newSum
			}
		}
	}

	return maxSum
}
