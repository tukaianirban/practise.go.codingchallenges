package toroman

import "testing"

func TestConvertToRoman(t *testing.T) {

	testnumbers := []int{3, 58, 1994, 0}
	for _, testnum := range testnumbers {
		t.Logf("Integer=%d  Roman=%s", testnum, convertToRoman(testnum))
	}
}
