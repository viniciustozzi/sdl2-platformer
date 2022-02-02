package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	tex  *sdl.Texture
	x, y float64
}

const (
	playerSpeed = 5
	playerSize  = 120
)

func newPlayer(renderer *sdl.Renderer) *element {
	player := newElement()

	player.position = vector{
		x: screenWidth / 2.0,
		y: 0,
	}

	player.addComponent(newSpriteRenderer(player, renderer, "sprites/player.png", 4))
	player.addComponent(newGravity(player))

	return player
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeed
	} else if keys[sdl.SCANCODE_UP] == 1 {
		p.y -= playerSpeed
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		p.y += playerSpeed
	}
}
