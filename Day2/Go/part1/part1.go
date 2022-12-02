//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part1

var lut1 map[string]int = map[string]int{
	"A X": 3 + 1, "A Y": 6 + 2, "A Z": 0 + 3,
	"B X": 0 + 1, "B Y": 3 + 2, "B Z": 6 + 3,
	"C X": 6 + 1, "C Y": 0 + 2, "C Z": 3 + 3,
}

func Part1(input []string) int {
	sum := 0
	for _, r := range input {
		sum += lut1[r]
	}
	return sum
}
