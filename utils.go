package tinyfmt

// itoa converts an integer to its string representation.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result []byte
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = append([]byte{byte(digit) + '0'}, result...)
		n /= 10
	}

	if negative {
		result = append([]byte{'-'}, result...)
	}

	return string(result)
}
