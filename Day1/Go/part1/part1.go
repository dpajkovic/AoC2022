//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part1

import "strconv"

func Part1(input []string) int {
	max := 0
	sum := 0
	for _, s := range input {
		if len(s) == 0 {
			if sum > max {
				max = sum
			}
			sum = 0
			continue
		}
		i, _ := strconv.Atoi(s)
		sum += i
	}
	return max
}
