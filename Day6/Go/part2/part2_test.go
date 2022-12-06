//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import (
	"os"
	"testing"
)

func TestPart2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Provided input 1",
			args: args{
				input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			want: 19,
		},
		{
			name: "Provided input 2",
			args: args{
				input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			want: 23,
		},
		{
			name: "Provided input 3",
			args: args{
				input: "nppdvjthqldpwncqszvftbrmjlhg",
			},
			want: 23,
		},
		{
			name: "Provided input 4",
			args: args{
				input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			want: 29,
		},
		{
			name: "Provided input 5",
			args: args{
				input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.input); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPart2(b *testing.B) {
	testInput, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part2(string(testInput))
	}
}
