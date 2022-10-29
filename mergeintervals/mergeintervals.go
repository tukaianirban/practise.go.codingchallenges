package mergeintervals

type interval struct {
	start int
	end   int
}

func getMergedIntervals(intervals []interval) []interval {

	results := make([]interval, 0)

	results = append(results, intervals[0])

	for i := 1; i < len(intervals); i++ {

		if intervals[i].start < results[len(results)-1].end {
			// merge
			// lastinterval.end = intervals[i].end
			results[len(results)-1].end = intervals[i].end
		} else {
			// add new interval to results
			results = append(results, intervals[i])
		}
	}

	return results
}
