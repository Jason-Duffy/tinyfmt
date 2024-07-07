// =============================================================================
// Project: tinyfmt
// File: error_test.go
// Description: Test suite for error functions in tinyfmt package.
// Datasheet/Docs:
//
// Author: Jason Duffy
// Created on: 07/07/2024
//
// Copyright: (C) 2024, Jason Duffy
// License: See LICENSE file in the project root for full license information.
// Disclaimer: See DISCLAIMER file in the project root for full disclaimer.
// =============================================================================

// -------------------------------------------------------------------------- //
//                               Import Statement                             //
// -------------------------------------------------------------------------- //

package tinyfmt

import (
	"errors"
	"testing"
)

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

func TestErrorf(t *testing.T) {
	testCases := []struct {
		format    string
		arguments []interface{}
		want      string
	}{
		{"Error: %s", []interface{}{"something went wrong"}, "Error: something went wrong"}, // Test formatting a string error message
		{"Code: %d", []interface{}{404}, "Code: 404"},                                       // Test formatting an integer error code
		{"Invalid: %q", []interface{}{42}, "unsupported format specifier"},                  // Test with unsupported format specifier
		{"Missing arg: %d %d", []interface{}{42}, "missing argument for %d"},                // Test with missing argument
		{"", []interface{}{}, ""}, // Test with empty format string
		{"Nil arg: %v", []interface{}{nil}, "argument for %v is not a bool"},                               // Test with nil argument
		{"Multiple: %d, %s, %v", []interface{}{42, "test", true}, "Multiple: 42, test, true"},              // Test with multiple format specifiers
		{"Unsupported type: %v", []interface{}{map[string]int{"key": 1}}, "argument for %v is not a bool"}, // Test with unsupported type
		{"Percent sign: %%", []interface{}{}, "Percent sign: %"},                                           // Test with escaped percent sign
	}

	for _, testCase := range testCases {
		err := Errorf(testCase.format, testCase.arguments...)
		if err.Error() != testCase.want {
			t.Errorf("Errorf(%q, %v) = %q, want %q", testCase.format, testCase.arguments, err.Error(), testCase.want)
		}
	}
}

// Additional test for compatibility with errors.New()
func TestErrorfWithErrorsNew(t *testing.T) {
	err1 := Errorf("This is a custom error with value: %d", 123)
	err2 := errors.New("This is a standard error")

	if err1.Error() == err2.Error() {
		t.Errorf("Errorf generated error should not match errors.New() error: got %q, want different", err1.Error())
	}

	expected := "This is a custom error with value: 123"
	if err1.Error() != expected {
		t.Errorf("Errorf generated error = %q, want %q", err1.Error(), expected)
	}
}
