package main

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
}

func InitPlayer(cP Point) Player {
	return Player{
		currentPosition: cP,
		nextPosition:    cP,
		dirx:            1,
		diry:            1,
		speed:           1,
	}
}

func (player *Player) Next() {
	player.nextPosition.x = player.currentPosition.x + player.dirx*player.speed
	player.nextPosition.y = player.currentPosition.y + player.diry*player.speed
}

// func (player *Player) ResetNextPositon() {
// 	player.nextPosition.x = player.currentPosition.x
// 	player.nextPosition.y = player.currentPosition.y
// }

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
