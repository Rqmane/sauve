package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1:]
	var value []int
	k := 96
	for _, index := range name {
		if index == "--upper" {
			k = 64
			continue
		}
		num := 0
		for _, j := range index {
			num = num*10 + int(rune(j)-'0')
		}
		value = append(value, num)
	}
	for i := 0; i < len(value); i++ {
		if len(value) == 0 {
			break
		} else if value[i] > 26 {
			z01.PrintRune(rune(' '))
			continue
		}
		z01.PrintRune(rune(value[i]) + rune(k))
	}
	if len(value) > 0 {
		z01.PrintRune('\n')
	}
}
