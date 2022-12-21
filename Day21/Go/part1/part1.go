//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part1

import (
	"strconv"
	"strings"
)

var monkeys = map[string]string{}

func Part1(input string) int {
	for _, s := range strings.Split(input, "\n") {
		s := strings.Split(s, ": ")
		monkeys[s[0]] = s[1]
	}
	return solve("root")
}

func solve(expr string) int {
	if v, err := strconv.Atoi(monkeys[expr]); err == nil {
		return v
	}

	s := strings.Fields(monkeys[expr])
	return map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}[s[1]](solve(s[0]), solve(s[2]))
}
