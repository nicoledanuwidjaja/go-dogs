package main

import (
	"strings"
	"testing"
)

func ReverseString(s string) string {
	if s == "" {
		return ""
	}

	var newString []string
	for i := len(s) - 1; i >= 0; i-- {
		newString = append(newString, string(s[i]))
	}
	return strings.Join(newString, "")
}

// Usage: go test -v
func TestReverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// test cases
		{
			name: "nicole",
			args: args{
				s: "nicole",
			},
			want: "elocin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
