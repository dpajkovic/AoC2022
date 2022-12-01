//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import (
	"sort"
	"strconv"
)

type max3 struct {
	m1  int
	m2  int
	m3  int
	min int
}

func (m *max3) add(i int) {
	var a sort.IntSlice = sort.IntSlice{m.m1, m.m2, m.m3, i}
	a.Sort()
	m.m1, m.m2, m.m3 = a[1], a[2], a[3]
	m.min = m.m1
}

func (m max3) sum() int {
	return m.m1 + m.m2 + m.m3
}

func Part2(input []string) int {
	var max max3
	sum := 0
	for _, s := range input {
		if len(s) == 0 {
			if sum > max.min {
				max.add(sum)
			}
			sum = 0
			continue
		}
		i, _ := strconv.Atoi(s)
		sum += i
	}
	if sum > max.min {
		max.add(sum)
	}
	return max.sum()
}
