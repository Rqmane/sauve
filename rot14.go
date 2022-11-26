package piscine

func Rot14(str string) string {
	Nouvstr := []rune(str)
	var resultat string

	for i := 0; i < len(Nouvstr); i++ {
		if Nouvstr[i] >= 65 && Nouvstr[i] <= 76 {
			Nouvstr[i] = Nouvstr[i] + 14
		} else if Nouvstr[i] >= 77 && Nouvstr[i] <= 90 {
			Nouvstr[i] = Nouvstr[i] - 12
		} else if Nouvstr[i] >= 97 && Nouvstr[i] <= 108 {
			Nouvstr[i] = Nouvstr[i] + 14
		} else if Nouvstr[i] >= 109 && Nouvstr[i] <= 122 {
			Nouvstr[i] = Nouvstr[i] - 12
		}
		resultat += string(Nouvstr[i])
	}
	return resultat
}
