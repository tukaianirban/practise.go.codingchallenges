/*
Problem: https://www.educative.io/m/string-segmentation
Given a dictionary of words, and a master word; find if the master word can be completely
segmented into words in the dictionary
*/
package stringsegment

import (
	"log"
	"strings"
)

// solve for 2 part segmentation only
// assumptions: masterword is segmented in 2 parts only
// assumptions: each part begins at index=0 or otherwise.
func findStringSegment(masterWord string, dictionary []string) (string, string) {

	complements := make(map[string]bool)

	for _, dictword := range dictionary {

		// if the word is contained, then proceed or else cotninue
		idx := strings.Index(masterWord, dictword)
		if idx < 0 {
			continue
		}

		var comp string
		if idx == 0 {
			comp = masterWord[len(dictword):]
		} else {
			comp = masterWord[:idx]
		}

		log.Printf("word=%s complement=%s", dictword, comp)

		if _, ok := complements[comp]; ok {
			return dictword, comp
		}

		complements[dictword] = true
		log.Printf("complements map=%#v", complements)
	}

	return "", ""
}

func findStringSegments(findWord string, dictionary []string) []string {

	masterWord := findWord

	for len(findWord) > 0 {

	}

	// // traverse thorough the whole array
	// for _, dictword := range dictionary {

	// 	var comp string
	// 	idx := strings.Index(masterWord, dictword)
	// 	if idx == 0 {
	// 		comp = masterWord[len(dictword):]
	// 	} else {
	// 		comp = masterWord[:idx]
	// 	}

	// 	isComplementPresent := isContainedInStringArray(dictionary, comp, []string{dictword})
	// 	if isComplementPresent {
	// 		return []string{dictword, comp}
	// 	}

	// }
}

func isContainedInStringArray(strArray []string, word string, excepts []string) bool {

	for _, val := range strArray {
		if word == val {
			isFound := true

			// check if this val is present in the excepts list or not
			for _, ex := range excepts {
				if val == ex {
					isFound = false
				}
			}

			// if isFound=true, then not matched in excepts
			if !isFound {
				return true
			}
		}
	}

	return false
}

/*
returns the complement part of the masterword
and adds the found part into the excepts.
Returns: complement of matched word, and copy of excepts with matched word added in
*/
func findNextInDictionary(dictionary []string, masterWord string, excepts []string) ([]string, string) {

	for _, word := range dictionary {

		idx := strings.Index(masterWord, word)
		if idx < 0 {
			continue
		}

		var comp string
		if idx == 0 {
			comp = masterWord[len(word):]
		} else {
			comp = masterWord[:idx]
		}

		// check if the complement is present in the dictionary
		if isContainedInStringArray(dictionary, comp, excepts) {
			return append(excepts, comp), ""
		}

		return append(excepts, word), comp
	}

	return excepts, masterWord
}
