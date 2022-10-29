/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
*/
package phnumcombos

var digitcharsMap = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

var combosStrings []string
var resultStrings []string

func findPhoneNumberCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	combosStrings = make([]string, len(digits))
	for idx, ch := range digits {
		combosStrings[idx] = digitcharsMap[int(ch)-48]
	}

	resultStrings = make([]string, 0)

	recurseComboStrings(0, "")
	return resultStrings
}

func recurseComboStrings(idx int, str string) {

	isLastIndex := (idx == len(combosStrings)-1)

	for _, ch := range combosStrings[idx] {

		if !isLastIndex {
			recurseComboStrings(idx+1, str+string(ch))
		} else {
			resultStrings = append(resultStrings, str+string(ch))
		}
	}
}
