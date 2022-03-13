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
	if len(args) != 4 {
		fmt.Println("go run main.go <width> <height> <delay> <count>")
		return
	}

	var width int
	var height int
	var delay int
	var count int

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
	delay = Max(1, delay)

	count, err = strconv.Atoi(args[3])
	if err != nil {
		panic(err)
	}
	count = Max(1, count)

	players := make([]Player, 0, count)
	addPlayers(&players, count, width, height)

	field := GenerateField(height, width)
	buf := make([]rune, 0, len(field)*len(field[0]))

	screen.Clear()
	HideTerminalCursor()

	for {
		screen.MoveTopLeft()

		DrawField(&field, buf, &players)
		for p := range players {
			players[p].Next()
			// players[p].CollisionPlayer(&players)
			players[p].CollisionRectangle(&field)
			players[p].Move()

		}
		fmt.Println(time.Now().Clock())

		time.Sleep(time.Second / time.Duration(delay))
	}
}
