/*
Problem definition: https://leetcode.com/problems/zigzag-conversion/
*/
package zigzagpattern

func createZigZag(s string, numRows int) string {

	matrix := make([]string, numRows)

	go getChars(s)

	r := 0

	flag := true
	for flag {

		r = 0
		for ; r < numRows; r++ {
			ch := <-chChars
			if ch == "" {
				flag = false
				break
			}

			matrix[r] = matrix[r] + ch
		}

		if !flag {
			break
		}

		r = r - 2
		for ; r > 0; r-- {
			ch := <-chChars
			if ch == "" {
				flag = false
				break
			}

			matrix[r] = matrix[r] + ch
		}
	}

	result := ""
	for _, str := range matrix {
		result = result + str
	}
	return result
}

var chChars = make(chan string)

func getChars(s string) {

	for _, ch := range s {
		chChars <- string(ch)
	}

	chChars <- ""
}
