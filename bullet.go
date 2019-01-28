package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bullet struct {
	texture *sdl.Texture
	x, y    float64
	angle   float64

	active bool
}

const (
	bulletWidth  = 40
	bulletHeight = 20
	bulletSpeed  = 1
)

func newBullet(renderer *sdl.Renderer) (b bullet) {
	b.texture = textureFromBMP(renderer, "./sprites/bullet.bmp")

	return
}

func (b *bullet) draw(renderer *sdl.Renderer) {
	if !b.active {
		return
	}

	x := b.x - bulletWidth/2
	y := b.y - bulletHeight/2

	renderer.CopyEx(b.texture,
		&sdl.Rect{X: 0, Y: 0, W: 246, H: 119},
		&sdl.Rect{X: int32(x), Y: int32(y), W: bulletWidth, H: bulletHeight},
		90,
		&sdl.Point{X: bulletWidth / 2, Y: bulletHeight / 2},
		sdl.FLIP_NONE,
	)
}

func (b *bullet) update() {
	if !b.active {
		return
	}
	b.x += bulletSpeed * math.Cos(b.angle)
	b.y += bulletSpeed * math.Sin(b.angle)
	if b.y < 0 {
		b.active = false
	}
}

type bullets []*bullet

var bulletPool bullets

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bull := newBullet(renderer)
		bulletPool = append(bulletPool, &bull)
	}
}

func (b bullets) getBullet() (*bullet, bool) {
	for _, bul := range b {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}

func (b bullets) drawAndUpdateBullets(renderer *sdl.Renderer) {
	for _, bul := range bulletPool {
		bul.draw(renderer)
		bul.update()
	}
}
