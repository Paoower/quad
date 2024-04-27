package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	// print first line
	PrintFirstLine(x, 'o', 'o', '-')
	// print intermediate lines
	for i := 0; i < y-2; i++ {
		PrintIntermediateLine(x, '|')
	}
	// print last line
	if y > 1 {
		PrintLastLine(x, 'o', 'o', '-')
	}
}

func QuadB(x, y int) {
	// print first line
	PrintFirstLine(x, '/', 92, '*')
	// print intermediate lines
	for i := 0; i < y-2; i++ {
		PrintIntermediateLine(x, '*')
	}
	// print last line
	if y > 1 {
		PrintLastLine(x, 92, '/', '*')
	}
}

func QuadC(x, y int) {
	// print first line
	PrintFirstLine(x, 'A', 'A', 'B')
	// print intermediate lines
	for i := 0; i < y-2; i++ {
		PrintIntermediateLine(x, 'B')
	}
	// print last line
	if y > 1 {
		PrintLastLine(x, 'C', 'C', 'B')
	}
}

func QuadD(x, y int) {
	// print first line
	PrintFirstLine(x, 'A', 'C', 'B')
	// print intermediate lines
	for i := 0; i < y-2; i++ {
		PrintIntermediateLine(x, 'B')
	}
	// print last line
	if y > 1 {
		PrintLastLine(x, 'A', 'C', 'B')
	}
}

func QuadE(x, y int) {
	// print first line
	PrintFirstLine(x, 'A', 'C', 'B')
	// print intermediate lines
	for i := 0; i < y-2; i++ {
		PrintIntermediateLine(x, 'B')
	}
	// print last line
	if y > 1 {
		PrintLastLine(x, 'C', 'A', 'B')
	}
}

func PrintFirstLine(x int, firstEdge, lastEdge, inter rune) {
	z01.PrintRune(firstEdge)
	if x > 1 {
		for i := 0; i < x-2; i++ {
			z01.PrintRune(inter)
		}
		z01.PrintRune(lastEdge)
	}
	z01.PrintRune('\n')
}

func PrintLastLine(x int, firstEdge, lastEdge, inter rune) {
	z01.PrintRune(firstEdge)
	if x > 1 {
		for i := 0; i < x-2; i++ {
			z01.PrintRune(inter)
		}
		z01.PrintRune(lastEdge)
	}
	z01.PrintRune('\n')
}

func PrintIntermediateLine(x int, edge rune) {
	for i := 0; i <= x-1; i++ {
		if i == 0 || i == x-1 {
			z01.PrintRune(edge)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
