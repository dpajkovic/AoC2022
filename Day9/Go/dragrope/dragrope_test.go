//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package dragrope

import (
	"os"
	"testing"
)

func Test_sgn(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "positive",
			args: args{
				x: 456,
			},
			want: 1,
		},
		{
			name: "negative",
			args: args{
				x: -654,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sgn(tt.args.x); got != tt.want {
				t.Errorf("sgn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDragRope(t *testing.T) {
	type args struct {
		input  string
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Provided input part 1",
			args: args{
				input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
				length: 1,
			},
			want: 13,
		},
		{
			name: "Provided input part 2",
			args: args{
				input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
				length: 9,
			},
			want: 1,
		},
		{
			name: "Provided input part 2, alternate",
			args: args{
				input: `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`,
				length: 36,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DragRope(tt.args.input, tt.args.length); got != tt.want {
				t.Errorf("DragRope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDragRopePart1(b *testing.B) {
	testInput, err := os.ReadFile("../input.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DragRope(string(testInput), 1)
	}
}

func BenchmarkDragRopePart2(b *testing.B) {
	testInput, err := os.ReadFile("../input.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DragRope(string(testInput), 9)
	}
}
