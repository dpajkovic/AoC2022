//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part1

import (
	"AoC2022D03/lettersum"
	"strings"
)

func Part1(input []string) int {
	letters := ""
	for _, r := range input {
		var l rune
		i := 0
		for l == 0 {
			if strings.ContainsRune(r[len(r)/2:], rune(r[i])) {
				l = rune(r[i])
			}
			i++
		}
		letters += string(l)
	}
	return lettersum.Sum(letters)
}
