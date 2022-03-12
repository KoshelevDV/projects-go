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
	velocity        int
}

func InitPlayer(cP Point) Player {
	return Player{
		currentPosition: cP,
		nextPosition:    cP,
		dirx:            1,
		diry:            1,
		velocity:        1,
	}
}

func (player *Player) Next() {
	player.nextPosition.x += player.dirx * player.velocity
	player.nextPosition.y += player.diry * player.velocity
}

func (player *Player) ResetNextPositon() {
	player.nextPosition.x = player.currentPosition.x
	player.nextPosition.y = player.currentPosition.y
}

func (player *Player) Move() {
	player.currentPosition.x += player.velocity * player.dirx
	player.currentPosition.y += player.velocity * player.diry
}

func (player *Player) CollisionRectangle(field *[][]string) {
	switch {
	// top left corner
	case player.nextPosition.x <= 0 && player.nextPosition.y <= 0:
		fallthrough
	// top right corner
	case player.nextPosition.x >= len((*field)[0])-1 && player.nextPosition.y <= 0:
		fallthrough
	// bottom left corner
	case player.nextPosition.x <= 0 && player.nextPosition.y >= len(*field)-1:
		fallthrough
	// bottom right corner
	case player.nextPosition.x >= len((*field)[0])-1 && player.nextPosition.y >= len(*field)-1:
		player.dirx *= -1
		player.diry *= -1
	case player.nextPosition.x <= 0 || player.nextPosition.x >= len((*field)[0])-1:
		player.dirx *= -1
	case player.nextPosition.y <= 0 || player.nextPosition.y >= len(*field)-1:
		player.diry *= -1
	}
}
