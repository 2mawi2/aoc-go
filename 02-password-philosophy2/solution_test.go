package main

import (
	"testing"
)

func Test_isValidPasswordPolicy(t *testing.T) {
	type args struct {
		policy PasswordPolicy
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"lower letter contained",
			args{PasswordPolicy{
				upper:    3,
				lower:    1,
				letter:   "a",
				password: "abcde",
			}},
			true,
		},
		{
			"upper letter contained",
			args{PasswordPolicy{
				upper:    3,
				lower:    1,
				letter:   "c",
				password: "abcde",
			}},
			true,
		},
		{
			"none of the letters contained",
			args{PasswordPolicy{
				upper:    3,
				lower:    1,
				letter:   "x",
				password: "abcde",
			}},
			false,
		},
		{
			"both letters contained",
			args{PasswordPolicy{
				upper:    3,
				lower:    1,
				letter:   "a",
				password: "abade",
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.policy.isValid(); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
