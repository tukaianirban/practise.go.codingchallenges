/*
Problem definition: https://www.geeksforgeeks.org/maximum-size-sub-matrix-with-all-1s-in-a-binary-matrix/?ref=lbp
*/
package maxsizebinarysubmatrix

var matrix = [6][5]int{
	{0, 1, 1, 0, 1},
	{1, 1, 0, 1, 0},
	{0, 1, 1, 1, 0},
	{1, 1, 1, 1, 0},
	{1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0},
}

func findLargestSubMatrix() (int, int, int) {

	maxCount := 0
	maxCountX := -1
	maxCountY := -1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			// if this is the starting of a new sub-matrix
			// happens only if the item before is not a '1', and this item = '1'
			if (j == 0 || matrix[i][j-1] == 0) && matrix[i][j] == 1 {

				rowstart := i
				colstart := j

				dim := 2
				for {

					flag := true
					for k := rowstart; flag && k < rowstart+dim && k < len(matrix); k++ {
						for l := colstart; flag && l < colstart+dim && l < len(matrix[i]); l++ {

							if matrix[k][l] == 0 {
								flag = false
							}
						}
					}

					if flag {
						dim++
					} else {
						break
					}
				}

				// log.Printf("Max size at x=%d,y=%d is %d", i, j, ((dim - 1) * (dim - 1)))

				if (dim-1)*(dim-1) > maxCount {
					maxCountX = i
					maxCountY = j
					maxCount = (dim - 1) * (dim - 1)
				}
			}
		}
	}

	return maxCountX + 1, maxCountY + 1, maxCount
}
