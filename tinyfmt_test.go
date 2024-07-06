package tinyfmt

import (
	"testing"
)

// TestSprint tests the Sprint function for various inputs
func TestSprint(t *testing.T) {
	// Define a slice of test cases, each with input arguments and the expected output
	testCases := []struct {
		inputArgs []interface{}
		want      string
	}{
		{[]interface{}{"Hello, world!"}, "Hello, world!"}, // Test case for a single string
		{[]interface{}{"Number: ", 42}, "Number: 42"},     // Test case for a string and an integer
		{[]interface{}{"Value: ", true}, "Value: true"},   // Test case for a string and a boolean
		{[]interface{}{"Value: ", false}, "Value: false"}, // Test case for a string and a boolean
	}

	// Iterate over each test case
	for _, testCase := range testCases {
		// Call the Sprint function with the test input arguments
		got := Sprint(testCase.inputArgs...)
		// Check if the output matches the expected value
		if got != testCase.want {
			// If not, report an error
			t.Errorf("Sprint(%v) = %q, want %q", testCase.inputArgs, got, testCase.want)
		}
	}
}
