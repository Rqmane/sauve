package piscine

func FindNextPrime(nb int) int {
	for i := nb; i >= nb; i++ {
		if IsPrime1(i) == true {
			return i
		}
	}
	return 0
}

func IsPrime1(ind int) bool {
	if ind >= 2 {
		if ind%2 == 0 && ind != 2 {
			return false
		}
		for i := 3; i*i <= ind; i = i + 2 {
			if ind%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}
