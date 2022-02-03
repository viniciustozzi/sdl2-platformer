package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type gravity struct {
	container *element
	g_factor  float64
}

func newGravity(container *element) *gravity {
	return &gravity{
		container: container,
		g_factor:  3.0,
	}
}

func (g *gravity) onUpdate() error {
	g.container.pos.y += g.g_factor * delta
	return nil
}

func (g *gravity) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (g *gravity) onCollision(other *element) error {
	return nil
}
