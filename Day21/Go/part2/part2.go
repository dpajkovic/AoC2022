//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import (
	"sort"
	"strconv"
	"strings"
)

var monkeys = map[string]string{}

func Part2(input string) int {
	for _, s := range strings.Split(input, "\n") {
		s := strings.Split(s, ": ")
		monkeys[s[0]] = s[1]
	}

	monkeys["humn"] = "0"
	s := strings.Fields(monkeys["root"])
	if solve(s[0]) < solve(s[2]) {
		s[0], s[2] = s[2], s[0]
	}

	i, _ := sort.Find(1e16, func(v int) int {
		monkeys["humn"] = strconv.Itoa(v)
		return solve(s[0]) - solve(s[2])
	})
	return i
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
