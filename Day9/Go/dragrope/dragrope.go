//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package dragrope

import (
	"fmt"
	"image"
	"math"
	"strings"
)

func DragRope(input string, length int) int {
	directions := map[rune]image.Point{'U': {0, -1}, 'R': {1, 0}, 'D': {0, 1}, 'L': {-1, 0}}
	rope := make([]image.Point, length+1)
	visited := map[image.Point]bool{}

	for _, l := range strings.Split(input, "\n") {
		var dir rune
		var steps int
		fmt.Sscanf(l, "%c %d", &dir, &steps)

		for i := 0; i < steps; i++ {
			// move head
			rope[0] = rope[0].Add(directions[dir])

			// move tail
			for i := 1; i < len(rope); i++ {
				if d := rope[i-1].Sub(rope[i]); math.Abs(float64(d.X)) > 1 || math.Abs(float64(d.Y)) > 1 {
					rope[i] = rope[i].Add(image.Point{sgn(d.X), sgn(d.Y)})
				}
			}

			// visited by the tail of the rope
			visited[rope[len(rope)-1]] = true
		}
	}
	return len(visited)
}

func sgn(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}
