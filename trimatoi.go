package piscine

func TrimAtoi(s string) int {
	numb := 0
	sign := 1
	for _, com := range s {
		if com >= '0' && com <= '9' {
			numb = numb*10 + int(com-48)
		} else if com == '-' && numb == 0 {
			sign = -1
		}
	}
	return numb * sign
}
