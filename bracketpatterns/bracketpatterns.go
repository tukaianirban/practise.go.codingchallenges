/*
Problem definition: Given a pattern of brackets ( ex: {{{}}} ), find out all the permutations of closed bracket patterns possible
*/
package bracketpatterns

import "log"

func findBracketPatterns(remainingBrackets string, designed string) {

	if len(remainingBrackets) <= 0 {
		log.Printf(designed + remainingBrackets)
		return
	}

	// extract one set and recurse over the rest
	newRemainingBrackets := remainingBrackets
	count := 1
	for len(newRemainingBrackets) > 0 {
		newDesigned := string(remainingBrackets[:count]) + string(remainingBrackets[len(remainingBrackets)-count:])
		newRemainingBrackets = remainingBrackets[count : len(remainingBrackets)-count]

		// log.Printf("input ::: designed=%s  remaining=%s", designed, remainingBrackets)
		// log.Printf("recurse ::: newdesigned=%s newremaining=%s", newDesigned, newRemainingBrackets)

		findBracketPatterns(newRemainingBrackets, newDesigned+designed)

		count++
	}
}
