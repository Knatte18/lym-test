package main

import (
	"fmt"
	"strings"
)

// Dummy subpath fixture for weft relpath-mirroring tests.
func main() {
	fmt.Println(FormatGreeting(""))
}

// FormatGreeting returns a friendly greeting for name, falling back to the
// default name "friend" when name is empty or whitespace-only.
func FormatGreeting(name string) string {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		trimmed = "friend"
	}
	return "Hello, " + trimmed + "!"
}
