//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part1

import "fmt"

func Part1(input []string) int {
	count := 0
	for _, l := range input {
		var r1, r2, r3, r4 int
		fmt.Sscanf(l, "%d-%d,%d-%d", &r1, &r2, &r3, &r4)
		if (r1-r3)*(r2-r4) <= 0 {
			count++
		}
	}
	return count
}
