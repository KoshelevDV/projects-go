package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Point struct {
	x int
	y int
}

type Player struct {
	currentPosition Point
	nextPosition    Point
	dirx            int
	diry            int
	speed           int
	id              int
}

func Init() func(Point) Player {
	Id := 0
	return func(cP Point) Player {
		Id += 1
		return Player{
			currentPosition: cP,
			nextPosition:    cP,
			dirx:            rand.Intn(1-(-1)+1) + (-1),
			diry:            rand.Intn(1-(-1)+1) + (-1),
			speed:           1,
			id:              Id,
		}
	}
}

func (player *Player) Next() {
	player.nextPosition.x = player.currentPosition.x + player.dirx*player.speed
	player.nextPosition.y = player.currentPosition.y + player.diry*player.speed
}

func (player *Player) Move() {
	player.currentPosition.x += player.speed * player.dirx
	player.currentPosition.y += player.speed * player.diry
}

func (player *Player) CollisionRectangle(field *[][]string) {
	if player.nextPosition.x <= 0 || player.nextPosition.x >= len((*field)[0])-1 {
		player.dirx *= -1
	}
	if player.nextPosition.y <= 0 || player.nextPosition.y >= len(*field)-1 {
		player.diry *= -1
	}
}

func addPlayers(players *[]Player, count, width, height int) {
	rand.Seed(time.Now().Unix())

	used := make([]Point, (width-1)*(height-1))
	for i := 0; i < count; i++ {
		var pp Point
	p:
		for {
			pp = Point{
				x: rand.Intn(width-2) + 1,
				y: rand.Intn(height-2) + 1,
			}

			if !ContainsPoint(used, pp) {
				used = append(used, pp)
				break p
			}
		}

		*players = append(*players, Init()(pp))
		fmt.Println("3")
	}
}

// func (player *Player) CollisionPlayer(players *[]Player) {
// 	for p := range *players {
// 		if player.id != (*players)[p].id {
// 			if player.nextPosition.x == (*players)[p].currentPosition.x {
// 				player.dirx *= -1
// 			}
// 			if player.nextPosition.y == (*players)[p].currentPosition.y {
// 				player.diry *= -1
// 			}
// 		}
// 	}
// }
