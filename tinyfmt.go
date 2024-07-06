package tinyfmt

// Sprint formats the provided arguments and returns the resulting string
func Sprint(arguments ...interface{}) string {
	// Initialize an empty string to store the result
	var result string
	// Iterate over each argument
	for _, argument := range arguments {
		// Perform a type switch to handle different types of arguments
		switch value := argument.(type) {
		case string:
			// If the argument is a string, append it to the result
			result += value
		case int:
			// If the argument is an int, convert it to a string and append it to the result
			result += itoa(value)
		case bool:
			if value {
				result += "true"
			} else {
				result += "false"
			}
		default:
			// For unsupported types, append a placeholder
			result += "<unsupported>"
		}
	}
	// Return the concatenated result
	return result
}
