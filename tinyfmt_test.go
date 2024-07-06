package tinyfmt

import (
	"testing"
)

// TestSprint tests the Sprint function for a single string
func TestSprint(t *testing.T) {
	// The "got" variable holds the result of calling the Sprint function with the argument "Hello, world!".
	got := Sprint("Hello, world!")
	// The "want" variable represents the expected result of the Sprint function.
	want := "Hello, world!"

	// Compare the actual result ("got") with the expected result ("want"), and raise an error if they don't match.
	if got != want {
		t.Errorf("Sprint() = %q, want %q", got, want)
	}
}
