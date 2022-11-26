package piscine

func BasicAtoi2(s string) int {
	strVal := []rune(s)

	n := len(s)

	intVal := 0

	for i := 0; i < n; i++ {
		if strVal[i] < '0' || strVal[i] > '9' {
			return 0
		} else {

			intVal *= 10
			intVal += int(strVal[i]) - '0'

		}
	}
	return intVal
}
