//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part1

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	file, err := os.Open("../testinput.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var testInput []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := scanner.Text()
		testInput = append(testInput, l)
	}

	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Provided input",
			args: args{
				input: testInput,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Part1(tt.args.input))
	}
}

func BenchmarkPart1(b *testing.B) {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var testInput []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := scanner.Text()
		testInput = append(testInput, l)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part1(testInput)
	}
}
