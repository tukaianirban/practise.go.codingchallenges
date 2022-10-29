/*
Find the Kth largest element in a number stream
*/

package khighest

var initialArray = [6]int{4, 1, 3, 12, 7, 14}

// can be run after findKHighest has run atleast 1 time
func Add(newnumber int) int {

	// if new number is lower than the Kth highest number existing, then nothing changes
	if newnumber < khighestNumbers[len(khighestNumbers)-1] {
		return khighestNumbers[len(khighestNumbers)-1]
	}

	// log.Printf("new number is lesser than current K highest; no change")

	// if new number already exists in the Kth highest numbers array, then nothing changes
	exists := false
	for i := 0; i < len(khighestNumbers); i++ {
		if khighestNumbers[i] == newnumber {
			exists = true
		}
	}

	if exists {
		// log.Printf("new number already in K highest; no change")
		return khighestNumbers[len(khighestNumbers)-1]
	}

	i := len(khighestNumbers) - 2
	for ; i >= 0 && khighestNumbers[i] < newnumber; i-- {
		khighestNumbers[i+1] = khighestNumbers[i]
	}
	khighestNumbers[i+1] = newnumber

	// log.Printf("after adding newnumber, array=%#v", khighestNumbers)

	return khighestNumbers[len(khighestNumbers)-1]
}

var khighestNumbers []int

func findKHighest(K int) int {

	khighestNumbers = make([]int, 0)

	// run through the current list and set out the highest K numbers
	kthhighest := initialArray[0]
	initDone := false
	for i := 1; i <= K; i++ {

		max := initialArray[0]
		for j := 0; j < len(initialArray); j++ {

			if !initDone {
				if initialArray[j] > max {
					max = initialArray[j]
				}
			} else {
				if initialArray[j] > max && initialArray[j] < kthhighest {
					max = initialArray[j]
				}
			}
		}

		initDone = true

		khighestNumbers = append(khighestNumbers, max)
		kthhighest = max
	}

	// log.Printf("in initial array, K highest numbers are %#v", khighestNumbers)

	return khighestNumbers[len(khighestNumbers)-1]
}
