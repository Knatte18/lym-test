package main

import (
	"strings"
	"testing"
)

func TestFormatGreeting(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		contains string
		excludes string
	}{
		{"empty", "", "friend", ""},
		{"whitespace only", "   ", "friend", ""},
		{"ordinary name", "Alice", "Alice", "friend"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatGreeting(tc.input)
			if !strings.Contains(got, tc.contains) {
				t.Errorf("FormatGreeting(%q) = %q, want it to contain %q", tc.input, got, tc.contains)
			}
			if tc.excludes != "" && strings.Contains(got, tc.excludes) {
				t.Errorf("FormatGreeting(%q) = %q, want it to NOT contain %q", tc.input, got, tc.excludes)
			}
		})
	}
}
