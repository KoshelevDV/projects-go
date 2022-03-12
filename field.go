package main

import "fmt"

func DrawField(field *[][]string, player *Player) {
	for y := 0; y < len(*field); y++ {
		for x := 0; x < len((*field)[0]); x++ {
			if player.currentPosition.x == x && player.currentPosition.y == y {
				fmt.Printf("\033[31m%s\033[0m", "☻")
			} else {
				fmt.Printf("%s", (*field)[y][x])
			}
		}
		fmt.Printf("\n")
	}
}

func GenerateField(height, width int) [][]string {
	field := make([][]string, height)

	for i := 0; i < height; i++ {
		field[i] = make([]string, width)
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
	return field
}
