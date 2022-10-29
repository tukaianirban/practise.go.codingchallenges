package longestsubstrpalindrome

func findLongestPalindromeSubstring(s string) string {

	longestPalindrome := ""

	for i := len(s); i > 0; i-- {

		for j := 0; j <= len(s)-i; j++ {

			substr := s[j : j+i]

			// check if substr is palindrome and compare
			isPalindrome := true
			for k := 0; k < len(substr)/2; k++ {

				if substr[k] != substr[len(substr)-k-1] {
					isPalindrome = false
					break
				}
			}

			if !isPalindrome {
				continue
			}

			if len(substr) > len(longestPalindrome) {
				longestPalindrome = substr
			}
		}
	}

	return longestPalindrome
}

func findLongestPalindromeSubstring2(s string) string {

	maxLenStr := expandAndFind(s, 0, 1)

	for i := 1; i < len(s)-1; i++ {

		if pstr := expandAndFind(s, i-1, i+1); len(pstr) > len(maxLenStr) {
			maxLenStr = pstr
		}

		if pstr := expandAndFind(s, i-1, i+1); len(pstr) > len(maxLenStr) {
			maxLenStr = pstr
		}

	}

	return maxLenStr
}

func expandAndFind(s string, left int, right int) string {

	for {

		if s[left] != s[right] {
			return s[left+1 : right]
		}

		left--
		right++

		if left <= 0 || right >= len(s)-1 {
			return s[left+1 : right]
		}

	}
}
