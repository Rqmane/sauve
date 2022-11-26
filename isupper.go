package piscine

func IsUpper(str string) bool {
	for _, a := range []rune(str) {
		if a < 'A' || a > 'Z' {
			return false
		}
	}
	return true
}
