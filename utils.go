package main

import "fmt"

func HideTerminalCursor() {
	fmt.Printf("\033[?25l")
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Contains(slice []int, num int) bool {
	for i := range slice {
		if slice[i] == num {
			return true
		}
	}
	return false
}

func ContainsPoint(points []Point, p Point) bool {
	for _, point := range points {
		if point.x == p.x && point.y == p.y {
			return true
		}
	}
	return false
}

/*
\033[0m"

\033[?25l"
*/
