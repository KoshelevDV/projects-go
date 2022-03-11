package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	// velocity := 1
	const width, height = 20, 10

	field := [height][width]string{}

	dirx := 1
	diry := 1
	currentPosition := [2]int{2, 1}
	nextPosition := [2]int{1, 1}

	generateField(height, width, &field)
	screen.Clear()
	for {
		screen.MoveTopLeft()
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if i == currentPosition[0] && j == currentPosition[1] {
					fmt.Printf("%s", "@")
				} else {
					fmt.Printf("%s", field[i][j])
				}
			}
			fmt.Println()
		}
		fmt.Printf("\nx: %-2d y: %-2d", currentPosition[1], currentPosition[0])
		NextPosition(&currentPosition, &nextPosition, &dirx, &diry)
		time.Sleep(time.Millisecond * 100)
	}
}

func NextPosition(currentPosition, nextPosition *[2]int, dirx, diry *int) {
	nextPosition[0] += *dirx
	nextPosition[1] += *diry
	fmt.Printf("\nx: %-2d y: %-2d", nextPosition[1], nextPosition[0])
	fmt.Printf("\ndirx: %-2d diry: %-2d", *diry, *dirx)

	y := nextPosition[0]
	x := nextPosition[1]
	switch {
	case y == 0 && x == 0:
		*dirx = 1
		*diry = 1
	case y == 0 && x == 20-1:
		*dirx = 1
		*diry = -1
	case y == 10-1 && x == 0:
		*dirx = -1
		*diry = 1
	case y == 0 && x == 20-1:
		*dirx = -1
		*diry = -1
	case x == 0 || x == 20-1:
		*diry *= -1
	case y == 0 || y == 10-1:
		*dirx *= -1
	}
	currentPosition[0] += *dirx
	currentPosition[1] += *diry
	nextPosition[0] = currentPosition[0]
	nextPosition[1] = currentPosition[1]
}

func generateField(height, width int, field *[10][20]string) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			switch {
			case i == 0 && j == 0:
				field[i][j] = "█"
			case i == 0 && j == width-1:
				field[i][j] = "█"
			case i == height-1 && j == 0:
				field[i][j] = "█"
			case i == height-1 && j == width-1:
				field[i][j] = "█"
			case j == 0 || j == width-1:
				field[i][j] = "█"
			case i == 0 || i == height-1:
				field[i][j] = "█"
			default:
				field[i][j] = " "
			}
		}
	}

}
