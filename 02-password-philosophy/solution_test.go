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
			"first case",
			args{"1-3 a: abcde"},
			PasswordPolicy{upper: 3, lower: 1, letter: "a", password: "abcde"},
		},
		{
			"second case",
			args{"1-3 b: cdefg"},
			PasswordPolicy{upper: 1, lower: 3, letter: "b", password: "cdefg"},
		},
		{
			"first case",
			args{"2-9 c: ccccccccc"},
			PasswordPolicy{upper: 2, lower: 9, letter: "c", password: "ccccccccc"},
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
