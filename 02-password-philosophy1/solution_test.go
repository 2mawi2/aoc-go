package main

import (
	"reflect"
	"testing"
)

func Test_fromLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want PasswordPolicy
	}{
		{
			"first sample",
			args{"1-3 a: abcde"},
			PasswordPolicy{upper: 3, lower: 1, letter: "a", password: "abcde"},
		},
		{
			"second sample",
			args{"1-3 b: cdefg"},
			PasswordPolicy{upper: 3, lower: 1, letter: "b", password: "cdefg"},
		},
		{
			"third sample",
			args{"2-9 c: ccccccccc"},
			PasswordPolicy{upper: 9, lower: 2, letter: "c", password: "ccccccccc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fromLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
