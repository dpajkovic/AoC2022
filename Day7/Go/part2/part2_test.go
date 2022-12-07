//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	testInput, err := os.ReadFile("../testinput.txt")
	if err != nil {
		panic(err)
	}

	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Provided input",
			args: args{
				input: string(testInput),
			},
			want: 24933642,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Part2(tt.args.input))
		})
	}
}

func BenchmarkPart2(b *testing.B) {
	testInput, err := os.ReadFile("../testinput.txt")
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part2(string(testInput))
	}
}
