//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package duplicate

import "testing"

func TestHasDuplicate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no dups",
			args: args{
				s: "ABVGD",
			},
			want: false,
		},
		{
			name: "yes dups",
			args: args{
				s: "ABVGDDJ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasDuplicate(tt.args.s); got != tt.want {
				t.Errorf("HasDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
