package main

import "fmt"

func DrawField(field *[][]string, buf []rune, balls *[]Player) {
	// buf := make([]rune, 0, len(*field)*len((*field)[0]))
	buf = buf[:0]
	for y := 0; y < len(*field); y++ {
	loop:
		for x := 0; x < len((*field)[0]); x++ {
			for _, ball := range *balls {
				if ball.currentPosition.x == x && ball.currentPosition.y == y {
					buf = append(buf, '☻')
					continue loop
				}
			}
			buf = append(buf, []rune((*field)[y][x])...)

			// if player.currentPosition.x == x && player.currentPosition.y == y {
			// 	//fmt.Printf("\033[31m%s\033[0m", "☻")
			// 	buf = append(buf, '☻')
			// } else if player1.currentPosition.x == x && player1.currentPosition.y == y {
			// 	buf = append(buf, '☻')
			// } else {
			// 	// fmt.Printf("%s", (*field)[y][x])
			// 	buf = append(buf, []rune((*field)[y][x])...)
			// }
		}
		// fmt.Printf("\n")
		buf = append(buf, '\n')
	}
	fmt.Print(string(buf))
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
