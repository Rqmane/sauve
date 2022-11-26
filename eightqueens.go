package piscine

import (
	"github.com/01-edu/z01"
)

func main() {
	EightQueens()
}

const size = 8

var board [size][size]bool

func goodDirection(g, h, vg, vh int) bool {
	for 0 <= g && g < size &&
		0 <= h && h < size {
		if board[g][h] {
			return false
		}
		g = g + vg
		h = h + vh
	}
	return true
}

func goodSquare(g, h int) bool {
	return goodDirection(g, h, +0, -1) &&
		goodDirection(g, h, +1, -1) &&
		goodDirection(g, h, +1, +0) &&
		goodDirection(g, h, +1, +1) &&
		goodDirection(g, h, +0, +1) &&
		goodDirection(g, h, -1, +1) &&
		goodDirection(g, h, -1, +0) &&
		goodDirection(g, h, -1, -1)
}

func printQueens() {
	g := 0
	for g < size {
		h := 0
		for h < size {
			if board[g][h] {
				gg := rune(h) + 49
				z01.PrintRune(gg)
			}
			h++
		}
		g++
	}
	z01.PrintRune('\n')
}

func trhg(g int) {
	h := 0
	for h < size {
		if goodSquare(g, h) {
			board[g][h] = true

			if g == size-1 {
				printQueens()
			} else {
				trhg(g + 1)
			}

			board[g][h] = false
		}
		h++
	}
}

func EightQueens() {
	trhg(0)
}
