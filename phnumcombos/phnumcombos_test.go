package phnumcombos

import "testing"

func TestFindPhoneNUmberCombinations(t *testing.T) {

	str := "23"
	t.Logf("Combinations for the input=%s are:", str)

	results := findPhoneNumberCombinations(str)

	for _, result := range results {

		t.Logf("%s", result)
	}
}
