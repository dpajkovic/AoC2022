//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import (
	"image"
	"strconv"
	"strings"
)

func Part2(input string) int {

	woods := map[image.Point]int{}
	ySize := 0
	xSize := 0

	for y, l := range strings.Fields(input) {
		xSize = 0
		ySize++
		for x, r := range l {
			woods[image.Point{x, y}], _ = strconv.Atoi(string(r))
			xSize++
		}
	}

	woodSize := image.Rect(0, 0, xSize, ySize)

	result := 0
	for tree := range woods {
		score := 1

		for _, direction := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			i := 0
			for currentPoint := tree.Add(direction); ; currentPoint = tree.Add(direction.Mul(i + 1)) {

				if !currentPoint.In(woodSize) {
					score = score * i
					break
				}

				if woods[currentPoint] >= woods[tree] {
					score *= i + 1
					break
				}
				i++
			}
		}

		if score > result {
			result = score
		}
	}
	return result
}
