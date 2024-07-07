package tinyfmt

import (
	"testing"
)

func TestSprint(t *testing.T) {
	testCases := []struct {
		arguments []interface{}
		want      string
	}{
		{[]interface{}{"Hello, ", "world!"}, "Hello, world!"},                                     // Test concatenating strings
		{[]interface{}{"Value: ", 42}, "Value: 42"},                                               // Test concatenating string and integer
		{[]interface{}{"Bool: ", true}, "Bool: true"},                                             // Test concatenating string and boolean
		{[]interface{}{"Float: ", 3.14159}, "Float: 3.1415900000000001"},                          // Test concatenating string and float
		{[]interface{}{"Mixed: ", "string", ", ", 123, ", ", false}, "Mixed: string, 123, false"}, // Test concatenating mixed types
	}

	for _, testCase := range testCases {
		got := Sprint(testCase.arguments...)
		if got != testCase.want {
			t.Errorf("Sprint(%v) = %q, want %q", testCase.arguments, got, testCase.want)
		}
	}
}

func TestSprintf(t *testing.T) {
	testCases := []struct {
		format    string
		arguments []interface{}
		want      string
		shouldErr bool
	}{
		{"Hello, %s!", []interface{}{"world"}, "Hello, world!", false},   // Test formatting string
		{"Value: %d", []interface{}{42}, "Value: 42", false},             // Test formatting integer
		{"Hex: %x", []interface{}{255}, "Hex: 0xff", false},              // Test formatting hexadecimal
		{"Binary: %b", []interface{}{7}, "Binary: 0b111", false},         // Test formatting binary
		{"Octal: %o", []interface{}{64}, "Octal: 0o100", false},          // Test formatting octal
		{"Float: %.2f", []interface{}{3.14159}, "Float: 3.14", false},    // Test formatting float with precision
		{"Float: %.5f", []interface{}{3.14159}, "Float: 3.14159", false}, // Test formatting float with higher precision
		{"Bool: %v", []interface{}{true}, "Bool: true", false},           // Test formatting boolean true
		{"Bool: %v", []interface{}{false}, "Bool: false", false},         // Test formatting boolean false
		{"Invalid: %q", []interface{}{42}, "", true},                     // Test unsupported format specifier
		{"Missing arg: %d %d", []interface{}{42}, "", true},              // Test missing argument
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
