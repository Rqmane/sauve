package piscine

func Join(strs []string, sep string) string {
	var Nouvstr string
	count := 0
	for range strs {
		count++
	}
	for i, v := range strs {
		if i == count-1 {
			Nouvstr += v
		} else {
			Nouvstr += v + sep
		}
	}
	return Nouvstr
}
