package romantointeger

import "testing"

func TestConvertRomanToInteger(t *testing.T) {

	romans := []string{"MMI", "III", "LVIII", "MCMXCIV"}
	for _, roman := range romans {
		t.Logf("Roman=%s Integer=%d", roman, convertRomanToInteger(roman))
	}
}
