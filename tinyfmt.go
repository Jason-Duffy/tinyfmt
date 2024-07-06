package tinyfmt

// Sprint formats the provided arguments and returns the resulting string
// Takes a variadic parameter (a list of arguments of any type) and returns a string.
func Sprint(a ...interface{}) string {
	// If only one argument was passed in
	if len(a) == 1 {
		// Check if the first argument is of type string
		if str, ok := a[0].(string); ok {
			// If it is, return the string
			return str
		}
	}
	// If the condition above wasn't met, return an empty string.
	return ""
}
