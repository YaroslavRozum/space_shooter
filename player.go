package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 0.3
	playerSize  = 80
)

type player struct {
	texture *sdl.Texture
	x, y    float64
}

func newPlayer(renderer *sdl.Renderer) (p player) {
	p.texture = textureFromBMP(renderer, "./sprites/kspaceduel.bmp")

	p.x = screenWidth / 2
	p.y = screenHeight - playerSize

	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	x := p.x - playerSize/2
	y := p.y - playerSize/2
	renderer.CopyEx(p.texture,
		&sdl.Rect{X: 0, Y: 0, W: 128, H: 128},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 80, H: 80},
		180,
		&sdl.Point{X: playerSize / 2, Y: playerSize / 2},
		sdl.FLIP_NONE,
	)
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeed
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeed
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		p.y -= playerSpeed
	}
	if keys[sdl.SCANCODE_DOWN] == 1 {
		p.y += playerSpeed
	}

	if p.x+playerSize/2 > screenWidth {
		p.x = screenWidth - playerSize/2
	}
	if p.x-playerSize/2 < 0 {
		p.x = 0 + playerSize/2
	}
	if p.y+playerSize/2 > screenHeight {
		p.y = screenHeight - playerSize/2
	}
	if p.y-playerSize/2 < screenHeight/2 {
		p.y = screenHeight/2 + playerSize/2
	}
}
