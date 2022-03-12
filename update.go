package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/inancgumus/screen"
)

func Update() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("go run main.go <width> <height> <delay>")
		return
	}

	var width int
	var height int
	var delay int

	width, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	width = Max(0, width)

	height, err = strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}
	height = Max(0, height)

	delay, err = strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}
	delay = Max(10, delay)

	player := InitPlayer(Point{x: 1, y: 1})

	field := GenerateField(height, width)
	screen.Clear()
	HideTerminalCursor()
	for {
		screen.MoveTopLeft()

		DrawField(&field, &player)
		player.Next()
		player.CollisionRectangle(&field)
		player.Move()
		player.ResetNextPositon()
		time.Sleep(time.Millisecond * time.Duration(delay))
	}
}
