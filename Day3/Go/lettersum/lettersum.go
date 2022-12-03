//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package lettersum

func Sum(input string) int {
	sum := 0
	for _, ch := range input {
		if int(ch) > 95 {
			sum += int(ch) - 96
		} else {
			sum += int(ch) - 38
		}
	}
	return sum
}
