//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import (
	"fmt"
	"strings"
	"unicode"
)

func Part2(input string) string {
	split := strings.Split(input, "\n\n")
	crates := strings.Split(split[0], "\n")
	keys := crates[len(crates)-1]

	stack := map[rune]string{}
	for _, s := range crates {
		for i, r := range s {
			if unicode.IsLetter(r) {
				stack[[]rune(keys)[i]] += string(r)
			}
		}
	}

	for _, s := range strings.Split(strings.TrimSpace(split[1]), "\n") {
		var qty int
		var from, to rune
		fmt.Sscanf(s, "move %d from %c to %c", &qty, &from, &to)

		stack[to] = stack[from][:qty] + stack[to]
		stack[from] = stack[from][qty:]
	}

	result := ""
	for _, r := range strings.ReplaceAll(keys, " ", "") {
		result += stack[r][:1]
	}

	return result
}
