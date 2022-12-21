//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part1

import (
	"fmt"
	"strings"
)

func Part1(input string) int {
	width := 40
	clock := 0
	accX := 1

	result := 0

	for _, l := range strings.Split(input, "\n") {
		var ins string
		var v int
		fmt.Sscanf(l, "%s %d", &ins, &v)
		// advance once for noop
		if clock++; (clock+width/2)%width == 0 {
			result += clock * accX
		}
		if ins == "addx" {
			// advance second time for add, add to accumulator
			if clock++; (clock+width/2)%width == 0 {
				result += clock * accX
			}
			accX += v
		}
	}
	return result
}
