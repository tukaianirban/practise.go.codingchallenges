/*
Problem definition: https://leetcode.com/problems/trapping-rain-water/
*/
package traprainwater

import "log"

func getTrappedRainWater(height []int) int {

	totalVolume := 0

	for i := 0; i < len(height); i++ {

		// for any given item, find the next one that is >= its height
		j := i + 1
		interimheights := height[i]
		for ; j < len(height) && height[j] < height[i]; j++ {
			interimheights += height[j]
		}

		if j == len(height) {
			// no container end found at this height
			continue
		}

		vol := (j-i)*height[i] - interimheights
		log.Printf("vol between %d and %d is %d; interimheight=%d", i, j, vol, interimheights)

		totalVolume += vol

		i = j - 1
	}

	return totalVolume
}
