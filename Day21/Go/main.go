//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package main

import (
	"AoC2022D21/part1"
	"AoC2022D21/part2"
	"fmt"
	"os"
)

func main() {
	inputDay21, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	p1 := part1.Part1(string(inputDay21))
	fmt.Printf("Result for part 1: %d\n", p1)
	p2 := part2.Part2(string(inputDay21))
	fmt.Printf("Result for part 2: %d\n", p2)
}
