package main

import "testing"

func Test_getAtRepeatedPosition(t *testing.T) {
	type args struct {
		line     string
		position int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"gets string at center of first repetition",
			args{
				line:     "..M#.......",
				position: 2,
			},
			"M",
		},
		{
			"gets string the beginning of first repetition",
			args{
				line:     "S.##......E",
				position: 0,
			},
			"S",
		},
		{
			"gets string the end of first repetition",
			args{
				line:     "S.##......E",
				position: 10,
			},
			"E",
		},
		{
			"gets string at center of second repetition",
			args{
				line:     "..M#.......",
				position: 13,
			},
			"M",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAtRepeatedPosition(tt.args.line, tt.args.position); got != tt.want {
				t.Errorf("getAtRepeatedPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_traverseMap(t *testing.T) {
	type args struct {
		lines  []string
		deltaX int
		deltaY int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"counts trees in simple map with 1 / 3 slope",
			args{
				lines: []string{
					".......",
					"...#...",
					"......#",
				},
				deltaX: 3,
				deltaY: 1,
			},
			2,
		},
		{
			"counts trees in simple map with 1 / 5 slope",
			args{
				lines: []string{
					"...........",
					".....#.....",
					"......#...#",
				},
				deltaX: 5,
				deltaY: 1,
			},
			2,
		},
		{
			"counts trees in simple map with 2 / 2 slope",
			args{
				lines: []string{
					"...",
					"..#",
					"..#",
				},
				deltaX: 5,
				deltaY: 1,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := traverseMap(tt.args.lines, tt.args.deltaX, tt.args.deltaY); got != tt.want {
				t.Errorf("traverseMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
