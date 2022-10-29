/*
Problem definition: https://leetcode.com/problems/container-with-most-water/
*/
package mostwatercontainer

func getLargestContainer(height []int) int {

	maxVolume := 0

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {

			var vol int
			if height[i] < height[j] {
				vol = (j - i) * height[i]
			} else {
				vol = (j - i) * height[j]
			}

			if vol > maxVolume {
				maxVolume = vol
			}
		}
	}

	return maxVolume
}
