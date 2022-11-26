package piscine

func BasicJoin(strs []string) string {
	var Newstr string
	for _, v := range strs {
		Newstr += v
	}
	return Newstr
}
