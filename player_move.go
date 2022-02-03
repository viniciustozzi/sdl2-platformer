package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type playerMove struct {
	container *element
	speed     float64
	dir       vector
}

func newPlayerMove(container *element) *playerMove {
	return &playerMove{
		container: container,
		speed:     5,
		dir:       vector{x: 0, y: 0},
	}
}

func (m *playerMove) onUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		m.dir.x = -1
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		m.dir.x = 1
	} else {
		m.dir.x = 0
	}

	x := m.container.pos.x + (m.speed * m.dir.x * delta)
	y := m.container.pos.y + (m.speed * m.dir.y * delta)
	m.container.pos = vector{x: x, y: y}

	return nil
}

func (m *playerMove) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (m *playerMove) onCollision(other *element) error {
	return nil
}
