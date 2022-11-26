package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for j := 1; j <= y; j++ {
			if j == 1 || j == y {
				for i := 1; i <= x; i++ {
					if (j == 1 && i == 1) || ((j == y && i == x) && (y > 1 && x > 1)) {
						z01.PrintRune('/')
					} else if j == 1 && i == x || j == y && i == 1 {
						z01.PrintRune('\\')
					} else {
						z01.PrintRune('*')
					}
				}
			} else {
				for i := 1; i <= x; i++ {
					if i == 1 || i == x {
						z01.PrintRune('*')
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune('\n')
		}
	} else {
		return
	}
}

func main() {
	arguments := os.Args[1:]

	nbrInt, _ := strconv.Atoi(arguments[0])
	nbrInt1, _ := strconv.Atoi(arguments[1])
	QuadB(nbrInt, nbrInt1)

}
