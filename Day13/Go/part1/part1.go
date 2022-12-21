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

	var start, end image.Point
	height := map[image.Point]rune{}
	for x, s := range strings.Fields(input) {
		for y, r := range s {
			if r == 'S' {
				start, r = image.Point{x, y}, 'a'
			} else if r == 'E' {
				end, r = image.Point{x, y}, 'z'
			}
			height[image.Point{x, y}] = r
		}
	}

	distance := map[image.Point]int{end: 0}
	queue := []image.Point{end}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, direction := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			next := current.Add(direction)
			_, visited := distance[next]
			_, valid := height[next]

			if !visited && valid && height[current] <= height[next]+1 {
				distance[next] = distance[current] + 1
				queue = append(queue, next)
			}
		}
	}

	return distance[start]
}
