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
		g_factor:  1.0,
	}
}

func (g *gravity) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (g *gravity) onUpdate() error {
	g.container.position.y += g.g_factor
	return nil
}
