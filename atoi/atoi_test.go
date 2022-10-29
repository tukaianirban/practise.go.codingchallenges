package atoi

import "testing"

func TestGetAtoI(t *testing.T) {

	vals := []string{"-123abc", "4193 with words", "   -42", "-4294967296", "-91283472332", "   +42", ""}
	for _, val := range vals {
		t.Logf("%s converted to number is %d", val, getAtoI(val))
	}
}
