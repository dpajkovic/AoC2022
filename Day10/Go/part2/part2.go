//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import (
	"fmt"
	"strings"
)

func Part2(input string) string {
	width := 40
	clock := 0
	accX := 1

	result := ""

	for _, l := range strings.Split(input, "\n") {
		var ins string
		var v int
		fmt.Sscanf(l, "%s %d", &ins, &v)
		// advance once for noop
		if clock%width >= accX-1 && clock%width <= accX+1 {
			result += "#"
		} else {
			result += "."
		}

		if clock%width == width-1 {
			result += "\n"
		}
		clock++
		if ins == "addx" {
			// advance second time for add, add to accumulator
			if clock%width >= accX-1 && clock%width <= accX+1 {
				result += "#"
			} else {
				result += "."
			}

			if clock%width == width-1 {
				result += "\n"
			}
			clock++
			accX += v
		}
	}
	return strings.TrimSpace(result)
}
