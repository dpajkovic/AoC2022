//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import "fmt"

func Part2(input []string) int {
	count := 0
	for _, l := range input {
		var r1, r2, r3, r4 int
		fmt.Sscanf(l, "%d-%d,%d-%d", &r1, &r2, &r3, &r4)
		if (r2 >= r3) && (r4 >= r1) {
			count++
		}
	}
	return count
}
