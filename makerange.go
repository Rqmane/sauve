package piscine

func MakeRange(min, max int) []int {
	var Slival []int

	if min >= max {
		return Slival
	} else {
		n := max - min
		Slival = make([]int, n)
		for i := 0; i < n; i++ {
			Slival[i] = i + min
		}
		return Slival
	}
}
