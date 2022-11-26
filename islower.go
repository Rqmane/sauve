package piscine

func IsLower(str string) bool {
	for _, l := range []rune(str) {
		if l < 'a' || l > 'z' {
			return false
		}
	}
	return true
}
