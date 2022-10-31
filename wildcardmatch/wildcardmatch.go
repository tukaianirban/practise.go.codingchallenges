/*
Problem definition: https://leetcode.com/problems/wildcard-matching/
*/
package wildcardmatch

import "log"

func matchPattern(s string, p string) bool {

	poss := 0
	posp := 0
	for ; posp < len(p); posp++ {

		if poss == len(s) {
			return false
		}

		if p[posp] == s[poss] {
			poss++

			continue
		}

		if string(p[posp]) == "*" {
			for ; poss < len(s) && ((int(s[poss]) >= 97 && int(s[poss]) <= 123) || int(s[poss]) == 32); poss++ {
				continue
			}

			log.Printf("after moving poss for *, poss=%d", poss)

			continue
		}

		if string(p[posp]) == "?" {
			sascii := int(s[poss])

			if sascii < 97 || sascii > 123 {
				return false
			}

			poss++
			continue
		}

		return false
	}

	return poss == len(s)
}
