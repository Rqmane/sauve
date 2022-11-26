package piscine

func Split(str, charset string) []string {
	lStr := 0
	leCharset := 0
	wordf := 0
	for range str {
		lStr++
	}
	for range charset {
		leCharset++
	}
	for t := 0; t < lStr-leCharset; t++ {
		if str[t:t+leCharset] == charset {
			wordf++
		}
	}
	newStr := make([]string, wordf+1)
	r := -1
	k := 0
	for i := 0; i < lStr-leCharset; i++ {
		if str[i:i+leCharset] == charset {
			r++
			newStr[r] = str[i-k : i]
			k = 0
			i = i + leCharset - 1
		} else {
			k++
		}
	}

	q := 0
	for _, v := range newStr {
		for range v {
			q++
		}
		q += leCharset
	}
	q -= leCharset
	r++
	newStr[r] = str[q:]

	return newStr
}
