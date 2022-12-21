//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import (
	"container/ring"
	"strconv"
	"strings"
)

func Part2(input string) int {
	split := strings.Fields(input)

	r := ring.New(len(split))
	idx := map[int]*ring.Ring{}
	z := (*ring.Ring)(nil)

	for i, s := range split {
		v, _ := strconv.Atoi(s)
		if v == 0 {
			z = r
		}

		decryptionKey := 811589153
		r.Value = v * decryptionKey
		idx[i] = r
		r = r.Next()
	}
	repetitions := 10
	for i := 0; i < repetitions; i++ {
		for i := 0; i < len(idx); i++ {
			r = idx[i].Prev()
			c := r.Unlink(1)
			r.Move(c.Value.(int) % (len(idx) - 1)).Link(c)
		}
	}
	return z.Move(1000).Value.(int) + z.Move(2000).Value.(int) + z.Move(3000).Value.(int)
}
