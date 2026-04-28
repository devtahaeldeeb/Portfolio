package main

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 {
				// first row: ABBBC
				if j == 0 {
					z01.PrintRune('A')
				} else if j == x-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			} else if i == y-1 {
				// last row: CBBBA
				if j == 0 {
					z01.PrintRune('C')
				} else if j == x-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			} else {
				// middle rows B   B
				if j == 0 || j == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
