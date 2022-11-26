package piscine

func IsPrintable(str string) bool {
	for _, ip := range []rune(str) {
		if !(ip >= 32 && ip <= 127) {
			return false
		}
	}
	return true
}
