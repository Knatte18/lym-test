package main

import "testing"

func TestGreeting(t *testing.T) {
	tests := []struct {
		description string
		name        string
		want        string
	}{
		{description: "named", name: "Ada", want: "Hello, Ada!"},
		{description: "empty", name: "", want: "Hello, world!"},
		{description: "whitespace-only", name: " \t\n", want: "Hello, world!"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if got := Greeting(test.name); got != test.want {
				t.Errorf("Greeting(%q) = %q, want %q", test.name, got, test.want)
			}
		})
	}
}

func TestShoutGreeting(t *testing.T) {
	tests := []struct {
		description string
		name        string
		want        string
	}{
		{description: "named", name: "Ada", want: "HELLO, ADA!"},
		{description: "empty", name: "", want: "HELLO, WORLD!"},
		{description: "whitespace-only", name: " \t\n", want: "HELLO, WORLD!"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if got := ShoutGreeting(test.name); got != test.want {
				t.Errorf("ShoutGreeting(%q) = %q, want %q", test.name, got, test.want)
			}
		})
	}
}
