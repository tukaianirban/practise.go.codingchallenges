package zigzagpattern

import "testing"

func TestCreateZigZag(t *testing.T) {

	str := "PAYPALISHIRING"
	t.Logf(" transform of %s  with numRows=%d is: %#v", str, 3, createZigZag(str, 3))
	t.Logf(" transform of %s  with numRows=%d is: %#v", str, 4, createZigZag(str, 4))
	t.Logf(" transform of %s  with numRows=%d is: %#v", "AB", 1, createZigZag("AB", 1))
}
