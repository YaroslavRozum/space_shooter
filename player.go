package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed          = 0.5
	playerSize           = 80
	playerBulletCooldown = time.Millisecond * 250
)

type player struct {
	texture   *sdl.Texture
	x, y      float64
	lastShoot time.Time
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

func (p *player) shoot(x, y float64) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.x = x
		bul.y = y
		bul.angle = 270 * (math.Pi / 180)
		p.lastShoot = time.Now()
	}
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 && p.x-playerSize/2 > 0 {
		p.x -= playerSpeed
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 && p.x+playerSize/2 < screenWidth {
		p.x += playerSpeed
	}
	if keys[sdl.SCANCODE_UP] == 1 && p.y-playerSize/2 > screenHeight/2 {
		p.y -= playerSpeed
	}
	if keys[sdl.SCANCODE_DOWN] == 1 && p.y+playerSize/2 < screenHeight {
		p.y += playerSpeed
	}
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShoot) >= playerBulletCooldown {
			p.shoot(p.x-10, p.y-playerSize/2)
			p.shoot(p.x+10, p.y-playerSize/2)
		}
	}
}
