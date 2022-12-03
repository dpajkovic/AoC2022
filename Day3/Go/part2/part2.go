//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import (
	"AoC2022D03/lettersum"
	"strings"
)

func Part2(input []string) int {
	letters := ""
	for i := 0; i < len(input); i += 3 {
		r1 := input[i]
		r2 := input[i+1]
		r3 := input[i+2]
		var l rune
		j := 0
		for l == 0 {
			if strings.ContainsRune(r2, rune(r1[j])) && strings.ContainsRune(r3, rune(r1[j])) {
				l = rune(r1[j])
			}
			j++
		}
		letters += string(l)
	}
	return lettersum.Sum(letters)
}
