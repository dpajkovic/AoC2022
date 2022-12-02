//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

var lut2 map[string]int = map[string]int{
	"A X": 0 + 3, "A Y": 3 + 1, "A Z": 6 + 2,
	"B X": 0 + 1, "B Y": 3 + 2, "B Z": 6 + 3,
	"C X": 0 + 2, "C Y": 3 + 3, "C Z": 6 + 1,
}

func Part2(input []string) int {
	sum := 0
	for _, r := range input {
		sum += lut2[r]
	}
	return sum
}
