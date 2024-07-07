// =============================================================================
// Project: tinyfmt
// File: print_test.go
// Description: Test suite for print functions in tinyfmt package.
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
	"bytes"
	"io"
	"os"
	"testing"
)

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

func TestPrintToIo(t *testing.T) {
	testCases := []struct {
		format    string
		arguments []interface{}
		want      string
		shouldErr bool
	}{
		{"Hello, %s!", []interface{}{"world"}, "Hello, world!", false},
		{"Value: %d", []interface{}{42}, "Value: 42", false},
		{"Hex: %x", []interface{}{255}, "Hex: 0xff", false},
		{"Binary: %b", []interface{}{7}, "Binary: 0b111", false},
		{"Octal: %o", []interface{}{64}, "Octal: 0o100", false},
		{"Float: %.2f", []interface{}{3.14159}, "Float: 3.14", false},
		{"Float: %.5f", []interface{}{3.14159}, "Float: 3.14159", false},
		{"Bool: %v", []interface{}{true}, "Bool: true", false},
		{"Bool: %v", []interface{}{false}, "Bool: false", false},
		{"Invalid: %q", []interface{}{42}, "", true},
		{"Missing arg: %d %d", []interface{}{42}, "", true},
	}

	for _, testCase := range testCases {
		var buf bytes.Buffer
		err := PrintToIo(&buf, testCase.format, testCase.arguments...)
		got := buf.String()

		if (err != nil) != testCase.shouldErr {
			t.Errorf("PrintToIo(%q, %v) error = %v, wantErr %v", testCase.format, testCase.arguments, err, testCase.shouldErr)
			continue
		}
		if got != testCase.want {
			t.Errorf("PrintToIo(%q, %v) = %q, want %q", testCase.format, testCase.arguments, got, testCase.want)
		}
	}
}

func TestPrintf(t *testing.T) {
	testCases := []struct {
		format    string
		arguments []interface{}
		want      string
		shouldErr bool
	}{
		{"Hello, %s!", []interface{}{"world"}, "Hello, world!", false},
		{"Value: %d", []interface{}{42}, "Value: 42", false},
		{"Hex: %x", []interface{}{255}, "Hex: 0xff", false},
		{"Binary: %b", []interface{}{7}, "Binary: 0b111", false},
		{"Octal: %o", []interface{}{64}, "Octal: 0o100", false},
		{"Float: %.2f", []interface{}{3.14159}, "Float: 3.14", false},
		{"Float: %.5f", []interface{}{3.14159}, "Float: 3.14159", false},
		{"Bool: %v", []interface{}{true}, "Bool: true", false},
		{"Bool: %v", []interface{}{false}, "Bool: false", false},
		{"Invalid: %q", []interface{}{42}, "", true},
		{"Missing arg: %d %d", []interface{}{42}, "", true},
	}

	for _, testCase := range testCases {
		// Create a pipe to capture stdout
		r, w, _ := os.Pipe()
		origStdout := os.Stdout
		os.Stdout = w

		err := Printf(testCase.format, testCase.arguments...)

		w.Close()
		os.Stdout = origStdout

		var buf bytes.Buffer
		io.Copy(&buf, r)
		got := buf.String()

		if (err != nil) != testCase.shouldErr {
			t.Errorf("Printf(%q, %v) error = %v, wantErr %v", testCase.format, testCase.arguments, err, testCase.shouldErr)
			continue
		}
		if got != testCase.want {
			t.Errorf("Printf(%q, %v) = %q, want %q", testCase.format, testCase.arguments, got, testCase.want)
		}
	}
}
