package tinyfmt

import "testing"

// TestItoa tests the itoa function for converting integers to strings
func TestItoa(t *testing.T) {
	testCases := []struct {
		input int
		want  string
	}{
		{0, "0"},       // Test case for zero
		{123, "123"},   // Test case for a positive integer
		{-123, "-123"}, // Test case for a negative integer
	}

	for _, testCase := range testCases {
		got := itoa(testCase.input)
		if got != testCase.want {
			t.Errorf("itoa(%d) = %q, want %q", testCase.input, got, testCase.want)
		}
	}
}
