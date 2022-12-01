//	Copyright (c) Miloš Rackov 2022
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_max3_add(t *testing.T) {

	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields max3
		args   args
		want   max3
	}{
		{
			name: "nothing to add",
			fields: max3{
				m1:  4,
				m2:  5,
				m3:  6,
				min: 4,
			},
			args: args{
				i: 1,
			},
			want: max3{
				m1:  4,
				m2:  5,
				m3:  6,
				min: 4,
			},
		},
		{
			name: "something to add",
			fields: max3{
				m1:  4,
				m2:  5,
				m3:  6,
				min: 4,
			},
			args: args{
				i: 8,
			},
			want: max3{
				m1:  5,
				m2:  6,
				m3:  8,
				min: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := max3{
				m1: tt.fields.m1,
				m2: tt.fields.m2,
				m3: tt.fields.m3,
			}
			m.add(tt.args.i)
			assert.Equal(t, tt.want, m)
		})
	}
}

func Test_max3_sum(t *testing.T) {

	tests := []struct {
		name   string
		fields max3
		want   int
	}{
		{
			name: "0",
			fields: max3{
				m1:  0,
				m2:  0,
				m3:  0,
				min: 0,
			},
			want: 0,
		},
		{
			name: "more",
			fields: max3{
				m1:  1,
				m2:  2,
				m3:  3,
				min: 1,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := max3{
				m1:  tt.fields.m1,
				m2:  tt.fields.m2,
				m3:  tt.fields.m3,
				min: tt.fields.min,
			}
			assert.Equal(t, tt.want, m.sum())
		})
	}
}

func TestPart2(t *testing.T) {

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
			name: "Provided test input",
			args: args{
				input: testInput,
			},
			want: 45000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Part2(tt.args.input))
		})
	}
}
