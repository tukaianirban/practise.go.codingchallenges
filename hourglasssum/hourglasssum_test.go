package hourglasssum

import "testing"

func TestHourGlassSum(t *testing.T) {

	v := findMaxSum()

	t.Logf("max sum=%d", v)
}
