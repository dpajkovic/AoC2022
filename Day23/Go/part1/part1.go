//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part1

import (
	"image"
	"strings"
)

func Part1(input string) int {
	grid := make(map[image.Point]bool)
	for y, s := range strings.Fields(input) {
		for x, r := range s {
			if r == '#' {
				grid[image.Point{x, y}] = true
			}
		}
	}
	sides := [][]image.Point{
		{{0, -1}, {1, -1}, {-1, -1}}, // S
		{{0, 1}, {1, 1}, {-1, 1}},    // N
		{{-1, 0}, {-1, -1}, {-1, 1}}, // W
		{{1, 0}, {1, -1}, {1, 1}},    // E
	}
	for i := 0; ; i++ {
		proposed := map[image.Point]image.Point{}
		count := map[image.Point]int{}

		for p := range grid {
			neighbours := make([]int, len(sides))
			for i := range sides {
				for _, q := range sides[i] {
					if grid[p.Add(q)] {
						neighbours[i]++
					}
				}
			}

			if neighbours[0]+neighbours[1]+neighbours[2]+neighbours[3] == 0 {
				continue
			}

			for d := 0; d < 4; d++ {
				if direction := (i + d) % 4; neighbours[direction] == 0 {
					proposed[p] = p.Add(sides[direction][0])
					count[proposed[p]]++
					break
				}
			}
		}

		newGrid := make(map[image.Point]bool)
		for p := range grid {
			if _, ok := proposed[p]; ok && count[proposed[p]] == 1 {
				newGrid[proposed[p]] = true
			} else {
				newGrid[p] = true
			}
		}

		if i == 9 {
			var r image.Rectangle
			for p := range newGrid {
				r = r.Union(image.Rectangle{p, p.Add(image.Point{1, 1})})
			}
			return r.Dx()*r.Dy() - len(newGrid)
		}
		grid = newGrid
	}
}
