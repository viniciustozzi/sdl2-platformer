package main

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type circle_collision struct {
	container   *element
	center      vector
	radius      float64
	onCollision func(elem *element)
}

func newCircleCollision(
	container *element,
	radius float64,
	onCollision func(elem *element)) *circle_collision {
	return &circle_collision{
		container:   container,
		center:      vector{x: container.pos.x, y: container.pos.y},
		radius:      radius,
		onCollision: onCollision,
	}
}

func collides(c1, c2 *circle_collision) bool {
	dist := math.Sqrt(math.Pow(c2.center.x-c1.center.x, 2) +
		math.Pow(c2.center.y-c1.center.y, 2))

	return dist <= c1.radius+c2.radius
}

func (c *circle_collision) onUpdate() error {
	for _, elem := range elements {
		//TODO Remove myself from this list, or I will clide with myself
		c_col := elem.getComponent(&circle_collision{}).(*circle_collision)
		if c_col == c {
			return nil
		}

		if collides(c, c_col) {
			c.onCollision(c_col.container)
		}
	}

	return nil
}

func (c *circle_collision) onDraw(renderer *sdl.Renderer) error {
	gfx.CircleRGBA(
		renderer,
		int32(c.container.pos.x),
		int32(c.container.pos.y),
		int32(c.radius), 255, 0, 0, 255)

	return nil
}
