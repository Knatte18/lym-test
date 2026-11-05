package main

import "strings"

// Greeting returns a greeting for name, using "world" when name is empty or whitespace.
func Greeting(name string) string {
	if strings.TrimSpace(name) == "" {
		name = "world"
	}
	return "Hello, " + name + "!"
}
