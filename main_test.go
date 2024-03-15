package main

import (
	"testing"
)

func Test_isCircularPalindrome(t *testing.T) {
	tests := []struct {
		name       string
		args       string
		wantResult bool
	}{
		{
			name:       "racecar",
			args:       "racecar",
			wantResult: true,
		},
		{
			name:       "hello world",
			args:       "hello world",
			wantResult: false,
		},
		{
			name:       "mAlAyAlaM",
			args:       "mAlAyAlaM",
			wantResult: true,
		},
		{
			name:       "abcabca",
			args:       "abcabca",
			wantResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := isCircularPalindrome(tt.args); gotResult != tt.wantResult {
				t.Errorf("isCircularPalindrome() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
