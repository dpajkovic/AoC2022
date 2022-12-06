//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package duplicate

import (
	"strings"
)

func HasDuplicate(s string) bool {
	found := false
	i := 0

	for !found && i < len(s) {
		if strings.Count(s, s[i:i+1]) > 1 {
			found = true
		}
		i++
	}
	return found
}
