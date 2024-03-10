package mystrings

// Reverse reverses a string left to right
func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

