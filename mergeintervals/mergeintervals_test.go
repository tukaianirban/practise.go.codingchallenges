package mergeintervals

import "testing"

func TestGetMergedIntervals(t *testing.T) {

	var input = [8]interval{
		{10, 20},
		{18, 30},
		{31, 40},
		{40, 50},
		{51, 60},
		{59, 80},
		{59, 100},
		{150, 200},
	}

	results := getMergedIntervals(input[:])
	t.Logf("merged intervals = %#v", results)
}
