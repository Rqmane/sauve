package piscine

func AlphaCount(str string) int {
	nombre := 0
	for _, lettres := range []rune(str) {
		if (lettres >= 'A' && lettres <= 'Z') || (lettres >= 'a' && lettres <= 'z') {
			nombre = nombre + 1
		}
	}
	return nombre
}
