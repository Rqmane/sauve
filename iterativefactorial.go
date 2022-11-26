package piscine

func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 || nb > 20 {
		return 0
	} else {
		for i := nb; i > 0; i-- {
			result *= int(i)
		}
	}
	return result
}
