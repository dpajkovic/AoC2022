//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import (
	"fmt"
	"path"
	"strings"
)

func Part2(input string) int {
	files, currentDir := map[string]int{}, ""
	for _, s := range strings.Split(input, "\n") {
		if strings.HasPrefix(s, "$ cd") {
			currentDir = path.Join(currentDir, strings.Fields(s)[2])
		} else {
			var size int
			var name string
			fmt.Sscanf(s, "%d %s", &size, &name)
			files[path.Join(currentDir, name)] = size
		}
	}

	dirs := map[string]int{}
	for f, s := range files {
		for d := path.Dir(f); d != "/"; d = path.Dir(d) {
			dirs[d] += s
		}
		dirs["/"] += s
	}

	result := dirs["/"]
	for _, s := range dirs {
		if s+70000000-dirs["/"] >= 30000000 && s < result {
			result = s
		}
	}
	return result
}
