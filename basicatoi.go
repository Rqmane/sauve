package piscine

func BasicAtoi(s string) int {
	strVal := []rune(s)

	n := len(s)

	intVal := 0

	for i := 0; i < n; i++ {
		if strVal[i] < '0' || strVal[i] > '9' {
			return intVal
		} else {

			intVal *= 10
			intVal += int(strVal[i]) - '0'

		}
	}
	return intVal
}
