// =============================================================================
// Project: tinyfmt
// File: format_test.go
// Description: Test suite for the Format function in the tinyfmt package.
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

import "testing"

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

func TestFormat(t *testing.T) {
	testCases := []struct {
		format    string
		args      []interface{}
		want      string
		shouldErr bool
	}{
		{"Hello, %s!", []interface{}{"world"}, "Hello, world!", false},                                                                        // Test formatting string
		{"Value: %d", []interface{}{42}, "Value: 42", false},                                                                                  // Test formatting integer
		{"Hex: %x", []interface{}{255}, "Hex: 0xff", false},                                                                                   // Test formatting hexadecimal
		{"Binary: %b", []interface{}{7}, "Binary: 0b111", false},                                                                              // Test formatting binary
		{"Octal: %o", []interface{}{64}, "Octal: 0o100", false},                                                                               // Test formatting octal
		{"Float: %.2f", []interface{}{3.14159}, "Float: 3.14", false},                                                                         // Test formatting float with precision
		{"Float: %.5f", []interface{}{3.14159}, "Float: 3.14159", false},                                                                      // Test formatting float with higher precision
		{"Bool: %v", []interface{}{true}, "Bool: true", false},                                                                                // Test formatting boolean true
		{"Bool: %v", []interface{}{false}, "Bool: false", false},                                                                              // Test formatting boolean false
		{"Invalid: %q", []interface{}{42}, "", true},                                                                                          // Test unsupported format specifier
		{"Missing arg: %d %d", []interface{}{42}, "", true},                                                                                   // Test missing argument
		{"Percent: %%", []interface{}{}, "Percent: %", false},                                                                                 // Test escaping percent sign
		{"Negative: %d", []interface{}{-42}, "Negative: -42", false},                                                                          // Test formatting negative integer
		{"Large int: %d", []interface{}{9223372036854775807}, "Large int: 9223372036854775807", false},                                        // Test formatting large integer
		{"Small float: %.10f", []interface{}{0.0000001234}, "Small float: 0.0000001234", false},                                               // Test formatting small float with high precision
		{"String: %s %s", []interface{}{"Hello", "world"}, "String: Hello world", false},                                                      // Test formatting multiple strings
		{"Mix: %s %d %x %b %o %.2f %v", []interface{}{"Mix", 42, 255, 7, 64, 3.14159, true}, "Mix: Mix 42 0xff 0b111 0o100 3.14 true", false}, // Test mixed format specifiers
		{"Null string: %s", []interface{}{""}, "Null string: ", false},                                                                        // Test formatting null string
		{"Zero int: %d", []interface{}{0}, "Zero int: 0", false},                                                                              // Test formatting zero integer
		{"Precision: %.0f", []interface{}{123.456}, "Precision: 123", false},                                                                  // Test formatting float with zero precision
		{"Bool as int: %d", []interface{}{true}, "", true},                                                                                    // Test formatting boolean as integer should error
	}

	for _, testCase := range testCases {
		got, err := Format(testCase.format, testCase.args...)
		if (err != nil) != testCase.shouldErr {
			t.Errorf("Format(%q, %v) error = %v, wantErr %v", testCase.format, testCase.args, err, testCase.shouldErr)
			continue
		}
		if got != testCase.want {
			t.Errorf("Format(%q, %v) = %q, want %q", testCase.format, testCase.args, got, testCase.want)
		}
	}
}
