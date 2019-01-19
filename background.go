package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	backgroundHeight = 800
	backgroundWidth  = backgroundHeight * 1920 / 1080
)

type background struct {
	texture *sdl.Texture
	x, y    int
}

func createBackground(renderer *sdl.Renderer, x, y int) (b background) {
	b.texture = textureFromBMP(renderer, "./sprites/background.bmp")

	b.x = x
	b.y = y

	return b
}

func (b *background) draw(renderer *sdl.Renderer) {
	renderer.Copy(b.texture,
		&sdl.Rect{X: 0, Y: 0, W: 1920, H: 1080},
		&sdl.Rect{X: int32(b.x), Y: int32(b.y), W: backgroundWidth, H: backgroundHeight},
	)
}
