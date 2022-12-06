//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package main

import (
	"AoC2022D05/part1"
	"AoC2022D05/part2"
	"fmt"
	"os"
)

func main() {
	inputDay6, _ := os.ReadFile("input.txt")

	p1 := part1.Part1(string(inputDay6))
	fmt.Printf("Result for part 1: %s\n", p1)
	p2 := part2.Part2(string(inputDay6))
	fmt.Printf("Result for part 2: %s\n", p2)

}
