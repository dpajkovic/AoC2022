//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package main

import (
	"AoC2022D01/part1"
	"AoC2022D01/part2"
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

	var inputDay1 []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := scanner.Text()
		inputDay1 = append(inputDay1, l)
	}

	p1 := part1.Part1(inputDay1)
	fmt.Printf("Result for part 1: %d\n", p1)
	p2 := part2.Part2(inputDay1)
	fmt.Printf("Result for part 2: %d\n", p2)

}
