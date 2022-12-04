//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package main

import (
	"AoC2022D04/part1"
	"AoC2022D04/part2"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var inputDay3 []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := scanner.Text()
		inputDay3 = append(inputDay3, l)
	}

	p1 := part1.Part1(inputDay3)
	fmt.Printf("Result for part 1: %d\n", p1)
	p2 := part2.Part2(inputDay3)
	fmt.Printf("Result for part 2: %d\n", p2)
}
