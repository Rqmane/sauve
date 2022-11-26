package piscine

func Sqrt(nb int) int {
	var Sroot int
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			Sroot = i
			return Sroot
		}
	}
	return 0
}
