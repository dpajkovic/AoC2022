//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package main

import (
	"AoC2022D13/comparator"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	inputDay13, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	pkts, part1 := []any{}, 0
	for i, s := range strings.Split(strings.TrimSpace(string(inputDay13)), "\n\n") {
		s := strings.Split(s, "\n")
		var a, b any
		json.Unmarshal([]byte(s[0]), &a)
		json.Unmarshal([]byte(s[1]), &b)
		pkts = append(pkts, a, b)

		if comparator.Cmp(a, b) <= 0 {
			part1 += i + 1
		}
	}
	fmt.Println(part1)

	pkts = append(pkts, []any{[]any{2.0}}, []any{[]any{6.0}})
	sort.Slice(pkts, func(i, j int) bool { return comparator.Cmp(pkts[i], pkts[j]) < 0 })

	part2 := 1
	for i, p := range pkts {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			part2 *= i + 1
		}
	}
	fmt.Println(part2)
	// p1 := part1.Part1(string(inputDay13))
	// fmt.Printf("Result for part 1: %d\n", p1)
	// p2 := part2.Part2(string(inputDay13))
	// fmt.Printf("Result for part 2: %d\n", p2)

}
