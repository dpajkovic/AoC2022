//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package comparator

import "testing"

func TestCmp(t *testing.T) {
	type args struct {
		a any
		b any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "equal numbers",
			args: args{
				a: 1.0,
				b: 1.0,
			},
			want: 0,
		},
		{
			name: "unequal numbers",
			args: args{
				a: 3.0,
				b: 5.0,
			},
			want: -2,
		},
		{
			name: "unequal lists",
			args: args{
				a: []float64{1, 1, 3, 1, 1},
				b: []float64{1, 1, 5, 1, 1},
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cmp(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Cmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
