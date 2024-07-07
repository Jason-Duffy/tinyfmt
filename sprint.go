package tinyfmt

import (
	"github.com/Jason-Duffy/tinystrconv"
)

// Sprint concatenates the string representations of the provided arguments.
func Sprint(arguments ...interface{}) string {
	var result string
	for _, argument := range arguments {
		switch value := argument.(type) {
		case string:
			result += value
		case int:
			str, _ := tinystrconv.IntToString(value, 10)
			result += str
		case bool:
			result += tinystrconv.BoolToString(value)
		case float64:
			str, _ := tinystrconv.FloatToString(value, -1) // Use -1 for full precision
			result += str
		default:
			result += "<unsupported>"
		}
	}
	return result
}

// Sprintf formats the provided arguments according to the format specifier.
func Sprintf(format string, arguments ...interface{}) (string, error) {
	return Format(format, arguments...)
}
