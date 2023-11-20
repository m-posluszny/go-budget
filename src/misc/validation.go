package misc

func ValidateLength(s string, min int, max int) bool {
	l := len(s)
	return min <= l && l <= max
}
