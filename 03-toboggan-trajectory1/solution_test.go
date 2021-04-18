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
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"counts trees in simple map",
			args{lines: []string{
				".......",
				"...#...",
				"......#",
			}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := traverseMap(tt.args.lines); got != tt.want {
				t.Errorf("traverseMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
