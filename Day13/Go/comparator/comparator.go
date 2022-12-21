//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package comparator

func Cmp(a, b any) int {
	var al, bl []any
	aok, bok := false, false

	switch a.(type) {
	case float64:
		al, aok = []any{a}, true
	case []any, []float64:
		al = a.([]any)
	}

	switch b.(type) {
	case float64:
		bl, bok = []any{b}, true
	case []any, []float64:
		bl = b.([]any)
	}

	if aok && bok {
		return int(al[0].(float64) - bl[0].(float64))
	}
	for i := 0; i < len(al) && i < len(bl); i++ {
		if c := Cmp(al[i], bl[i]); c != 0 {
			return c
		}
	}
	return len(al) - len(bl)
}
