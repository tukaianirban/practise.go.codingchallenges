package reverseinteger

import "testing"

func TestFindReverse(t *testing.T) {

	numbers := []int{123, -123, 120, 4294967296, 429496729}
	for _, number := range numbers {
		t.Logf("reverse of %d is %d", number, findReverse(number))
	}

}
