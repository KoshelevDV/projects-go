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

/*
\033[0m"

\033[?25l"
*/
