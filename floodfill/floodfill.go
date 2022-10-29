/*
Simulate the floodfill algorithm used in painting programs
Provided a matrix where each pixel has a color.
You will be given a starting point, and you must fill the starting point
as well any adjacent pixels that are of the same color with the new color provided.

Solution: Find paint the pixel you are on, then find all adjacent pixels that have the same color
Use recursion to trigger painting all adjacents found
*/
package floodfill

import (
	"fmt"
)

var screen = [][]int{{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 0, 0},
	{1, 0, 0, 1, 1, 0, 2, 2},
	{1, 2, 2, 2, 2, 0, 2, 0},
	{1, 1, 1, 2, 2, 0, 2, 0},
	{1, 1, 1, 2, 2, 2, 2, 0},
	{1, 1, 1, 1, 1, 2, 1, 1},
	{1, 1, 1, 1, 1, 2, 2, 1},
}

func TriggerFloodFill(x int, y int, newColor int) {
	oldColor := screen[x][y]

	Floodfill(x, y, oldColor, newColor)
}

func Floodfill(x int, y int, oldColor int, newColor int) {

	// if out of bounds, return
	if x < 0 || x >= len(screen) || y < 0 || y >= len(screen[0]) {
		return
	}

	// ensure you are at a point that matches the pixel color to bepainted
	if screen[x][y] != oldColor {
		return
	}

	screen[x][y] = newColor

	var imin, imax, jmin, jmax int
	if x == 0 {
		imin = 0
	} else {
		imin = x - 1
	}

	if x >= len(screen)-1 {
		imax = len(screen) - 1
	} else {
		imax = x + 1
	}

	if y == 0 {
		jmin = 0
	} else {
		jmin = y - 1
	}

	if y >= len(screen[0])-1 {
		jmax = len(screen[0]) - 1
	} else {
		jmax = y + 1
	}

	for i := imin; i <= imax; i++ {
		for j := jmin; j <= jmax; j++ {
			if i == x && j == y {
				continue
			}

			if screen[i][j] == oldColor {
				Floodfill(i, j, oldColor, newColor)
			}
		}
	}
}

func PrintScreen() {

	for i := 0; i < len(screen); i++ {

		for j := 0; j < len(screen[i]); j++ {
			fmt.Printf("%d ", screen[i][j])
		}

		fmt.Println("")
	}
}
