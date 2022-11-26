package piscine

func IsAlpha(str string) bool {
	for _, ia := range []rune(str) {
		if (ia < 'A' || ia > 'Z') && (ia < 'a' || ia > 'z') && (ia < '0' || ia > '9') {
			return false
		}
	}
	return true
}
