package stringsegment

import (
	"strings"
)

/**
Logic: Traverse through the dictionary one by one in linear order.
For every word, extract it out as many times as possible from the master, leaving the remainingMaster string.
proceed the traversal through the dictionary with the remainingMaster.
Maintain a parallel int array which tracks the number of times the string at that index (in dictionary) occurs
in the masterword

**/

func findSegments2(masterWord string, dictionary []string) []string {

	occurences := make([]int, len(dictionary))
	remainingMasterWord := masterWord

	for idx, dictWord := range dictionary {

		// extract the dictionary word at idx as many times as possible
		rem := remainingMasterWord
		count := 0
		flag := true

		for flag {
			pos := strings.Index(rem, dictWord)
			if pos < 0 {
				flag = false
				break
			}

			count++

			fhalf := rem[:pos]
			shalf := rem[pos+len(dictWord):]
			rem = fhalf + shalf
		}

		remainingMasterWord = rem
		occurences[idx] = count

		if remainingMasterWord == "" {
			break
		}
	}

	if len(remainingMasterWord) > 0 {
		return nil
	}

	segments := make([]string, 0)
	for idx, count := range occurences {
		if count == 0 {
			continue
		}

		segments = append(segments, dictionary[idx])
	}

	return segments
}
