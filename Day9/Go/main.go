//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package main

import (
	"AoC2022D09/dragrope"
	"fmt"
	"os"
)

func main() {
	inputDay9, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	p1 := dragrope.DragRope(string(inputDay9), 1)
	fmt.Printf("Result for part 1: %d\n", p1)
	p2 := dragrope.DragRope(string(inputDay9), 9)
	fmt.Printf("Result for part 2: %d\n", p2)
}
