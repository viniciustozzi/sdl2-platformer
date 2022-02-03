package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
}

const (
	playerSpeed = 5
	playerSize  = 120
)

func newPlayer(renderer *sdl.Renderer) *element {
	player := newElement("player")

	player.pos = vector{
		x: screenWidth / 2.0,
		y: 0,
	}

	player.addComponent(newSpriteRenderer(player, renderer, "sprites/player.png", 4))
	player.addComponent(newGravity(player))
	player.addComponent(newPlayerMove(player))

	return player
}
