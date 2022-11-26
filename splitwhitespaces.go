package piscine

func SplitWhiteSpaces(s string) []string {
	var tab []string
	fsplite := ""
	for i, strA := range s {
		if strA > 32 {
			fsplite += string(strA)
		} else if fsplite != "" {
			tab = append(tab, fsplite)
			fsplite = ""
		}
		if i == len(s)-1 {
			tab = append(tab, fsplite)
		}
	}
	return tab
}
