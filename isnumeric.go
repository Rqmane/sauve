package piscine

func IsNumeric(str string) bool {
	for _, in := range []rune(str) {
		if in < '0' || in > '9' {
			return false
		}
	}
	return true
}
