package main

import "testing"

func TestGreeting(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "Ada", want: "Hello, Ada!"},
		{name: "", want: "Hello, world!"},
		{name: " \t\n", want: "Hello, world!"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Greeting(test.name); got != test.want {
				t.Errorf("Greeting(%q) = %q, want %q", test.name, got, test.want)
			}
		})
	}
}
