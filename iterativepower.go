package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	resultat := 1
	if power > 0 {
		for index := 1; index <= power; index++ {
			resultat = resultat * nb
		}
	}
	return resultat
}
