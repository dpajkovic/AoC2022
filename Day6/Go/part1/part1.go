//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part1

import (
	"AoC2022D06/duplicate"
)

func Part1(input string) int {
	i := 3
	found := false

	for !found && i < len(input) {
		found = !duplicate.HasDuplicate(input[i-3 : i+1])
		i++
	}
	return i
}
