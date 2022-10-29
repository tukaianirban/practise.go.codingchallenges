/*
Problem definition: https://www.geeksforgeeks.org/given-n-x-n-square-matrix-find-sum-sub-squares-size-k-x-k/?ref=lbp
*/
package maxsumksquare

var matrix = [5][5]int{{1, 1, 1, 1, 1},
	{2, 2, 2, 2, 2},
	{3, 3, 3, 3, 3},
	{4, 4, 4, 4, 4},
	{5, 5, 5, 5, 5},
}

// return sum of subsquare starting at top left at x,y; dim=k
func findSum(x, y, k int) int {

	sum := 0
	for i := x; i < x+k; i++ {
		for j := y; j < y+k; j++ {
			sum += matrix[i][j]
		}
	}

	return sum
}

// find sums of k-sized subsquares
func traverseMatrix(k int) [][]int {

	result := make([][]int, len(matrix)-(k-1))
	for i := 0; i < len(matrix)-(k-1); i++ {
		result[i] = make([]int, len(matrix[i])-(k-1))
		for j := 0; j < len(matrix)-(k-1); j++ {

			result[i][j] = findSum(i, j, k)
		}
	}

	return result
}
