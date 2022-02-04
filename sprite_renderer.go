package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container     *element
	tex           *sdl.Texture
	width, height float64
	scale         int32
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string, scale int) *spriteRenderer {
	tex := textureFromBMP(renderer, filename)

	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}

	return &spriteRenderer{
		container: container,
		tex:       textureFromBMP(renderer, filename),
		width:     float64(width),
		height:    float64(height),
		scale:     int32(scale),
	}
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	// Converting coordinates to top left of sprite
	//x := sr.container.pos.x - sr.width/2.0
	//y := sr.container.pos.y - sr.height/2.0

	x := sr.container.pos.x
	y := sr.container.pos.y


	renderer.CopyEx(
		sr.tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(sr.width), H: int32(sr.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(sr.width) * sr.scale, H: int32(sr.height) * sr.scale},
		sr.container.rotation,
		&sdl.Point{X: int32(sr.width) / 2, Y: int32(sr.height) / 2},
		sdl.FLIP_NONE)

	return nil
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
}

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	imgRw := sdl.RWFromFile(filename, "rb")
	defer imgRw.Close()

	img, err := img.LoadPNGRW(imgRw)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}

	return tex
}
