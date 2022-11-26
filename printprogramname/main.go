package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]
	for i, name := range arg {
		if i > 1 {
			z01.PrintRune(rune(name))
		}
	}
	z01.PrintRune('\n')
}
