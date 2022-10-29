/*
Problem definition: https://www.educative.io/m/find-all-palindrome-substrings
*/
package palindrome

func isPalindrome(str string) bool {

	revString := ""
	for i := len(str) - 1; i >= 0; i-- {
		revString = revString + string(str[i])
	}

	return revString == str
}

func findPalindromeSubstrings(masterString string) []string {

	resultStrings := make([]string, 0)

	for sz := 2; sz <= len(masterString); sz++ {

		// find all substrings of size=sz
		for i := 0; i < len(masterString)-(sz-1); i++ {

			str := masterString[i : i+sz]
			if isPalindrome(str) {
				resultStrings = append(resultStrings, str)
			}
		}
	}

	return resultStrings
}
