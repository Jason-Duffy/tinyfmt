// =============================================================================
// Project: tinyfmt
// File: sprint_test.go
// Description: Test suite for sprint functions in tinyfmt package.
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
	"math"
	"testing"
)

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

func TestSprint(t *testing.T) {
	type ExampleStruct struct {
		Name  string
		Value int
	}

	testCases := []struct {
		arguments []interface{}
		want      string
	}{
		{[]interface{}{"Hello, ", "world!"}, "Hello, world!"},                                     // Test concatenating strings
		{[]interface{}{"Value: ", 42}, "Value: 42"},                                               // Test concatenating string and integer
		{[]interface{}{"Bool: ", true}, "Bool: true"},                                             // Test concatenating string and boolean
		{[]interface{}{"Float: ", 3.14159}, "Float: 3.141590000000000"},                           // Test concatenating string and float
		{[]interface{}{"Mixed: ", "string", ", ", 123, ", ", false}, "Mixed: string, 123, false"}, // Test concatenating mixed types
		{[]interface{}{"Empty: ", ""}, "Empty: "},                                                 // Test concatenating with empty string
		{[]interface{}{nil}, "<unsupported>"},                                                     // Test unsupported type (nil)
		{[]interface{}{1, 2, 3, 4, 5}, "12345"},                                                   // Test concatenating multiple integers
		{[]interface{}{"Multiple: ", true, ", ", 1, ", ", "text"}, "Multiple: true, 1, text"},     // Test multiple different types
		{[]interface{}{math.MaxInt64}, "9223372036854775807"},                                     // Test edge case for maximum integer value
		{[]interface{}{[]int{1, 2, 3}}, "[1 2 3]"},                                                // Test formatting slice
		{[]interface{}{map[string]int{"key": 1}}, "{key:1}"},                                      // Test formatting map
		{[]interface{}{ExampleStruct{"example", 123}}, "{Name:example Value:123}"},                // Test formatting struct
	}

	for _, testCase := range testCases {
		got := Sprint(testCase.arguments...)
		if got != testCase.want {
			t.Errorf("Sprint(%v) = %q, want %q", testCase.arguments, got, testCase.want)
		}
	}
}

func TestSprintf(t *testing.T) {
	type ExampleStruct struct {
		Name  string
		Value int
	}

	testCases := []struct {
		format    string
		arguments []interface{}
		want      string
		shouldErr bool
	}{
		{"Hello, %s!", []interface{}{"world"}, "Hello, world!", false},                                          // Test formatting string
		{"Value: %d", []interface{}{42}, "Value: 42", false},                                                    // Test formatting integer
		{"Hex: %x", []interface{}{255}, "Hex: 0xff", false},                                                     // Test formatting hexadecimal
		{"Binary: %b", []interface{}{7}, "Binary: 0b111", false},                                                // Test formatting binary
		{"Octal: %o", []interface{}{64}, "Octal: 0o100", false},                                                 // Test formatting octal
		{"Float: %.2f", []interface{}{3.14159}, "Float: 3.14", false},                                           // Test formatting float with precision
		{"Float: %.5f", []interface{}{3.14159}, "Float: 3.14159", false},                                        // Test formatting float with higher precision
		{"Bool: %v", []interface{}{true}, "Bool: true", false},                                                  // Test formatting boolean true
		{"Bool: %v", []interface{}{false}, "Bool: false", false},                                                // Test formatting boolean false
		{"Multiple: %d, %s, %v", []interface{}{42, "test", true}, "Multiple: 42, test, true", false},            // Test multiple format specifiers
		{"Precision: %.0f", []interface{}{123.456}, "Precision: 123", false},                                    // Test float with precision 0
		{"Invalid: %q", []interface{}{42}, "", true},                                                            // Test unsupported format specifier
		{"Missing arg: %d %d", []interface{}{42}, "", true},                                                     // Test missing argument
		{"Edge case: %d", []interface{}{math.MaxInt64}, "Edge case: 9223372036854775807", false},                // Test edge case for large integer
		{"Negative: %d", []interface{}{-123}, "Negative: -123", false},                                          // Test negative integer
		{"Escape: %%", []interface{}{}, "Escape: %", false},                                                     // Test escape percentage sign
		{"Slice: %v", []interface{}{[]int{1, 2, 3}}, "Slice: [1 2 3]", false},                                   // Test formatting slice
		{"Map: %v", []interface{}{map[string]int{"key": 1}}, "Map: {key:1}", false},                             // Test formatting map
		{"Struct: %v", []interface{}{ExampleStruct{"example", 123}}, "Struct: {Name:example Value:123}", false}, // Test formatting struct
	}

	for _, testCase := range testCases {
		got, err := Sprintf(testCase.format, testCase.arguments...)
		if (err != nil) != testCase.shouldErr {
			t.Errorf("Sprintf(%q, %v) error = %v, wantErr %v", testCase.format, testCase.arguments, err, testCase.shouldErr)
			continue
		}
		if got != testCase.want {
			t.Errorf("Sprintf(%q, %v) = %q, want %q", testCase.format, testCase.arguments, got, testCase.want)
		}
	}
}
