package norepeatsubstring

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {

	longestSubstring := ""

	for length := len(s); length > 0; length-- {

		for i := 0; i <= len(s)-length; i++ {

			substr := s[i : i+length]

			// find any repeating chars in substring
			hasRepeat := false
			for j := 0; j < len(substr)-1; j++ {

				if strings.Count(substr, string(substr[j])) > 1 {
					hasRepeat = true
					break
				}
			}

			if hasRepeat {
				continue
			}

			// finding the max of the substrs found
			if len(substr) > len(longestSubstring) {
				longestSubstring = substr
			}
		}
	}

	return len(longestSubstring)
}
