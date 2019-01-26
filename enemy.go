package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	enemySize           = 80
	enemyAnimationDelay = time.Millisecond * 250
)

type enemy struct {
	texture *sdl.Texture
	x, y    float64
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (e enemy) {
	e.texture = textureFromBMP(renderer, "./sprites/alienblaster.bmp")

	e.x = x
	e.y = y

	return
}

func createEnemies(x, y int, renderer *sdl.Renderer) (enms enemies, err error) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			x := (float64(i)/5)*screenWidth + enemySize/2 + 15
			y := float64(j)*enemySize + enemySize/2 + float64(j*10)
			enm := newBasicEnemy(renderer, x, y)
			enms = append(enms, &enm)
		}
	}
	return enms, nil
}

func (e *enemy) draw(renderer *sdl.Renderer) {
	x := e.x - enemySize/2
	y := e.y - enemySize/2

	renderer.Copy(e.texture,
		&sdl.Rect{X: 0, Y: 0, W: 128, H: 128},
		&sdl.Rect{X: int32(x), Y: int32(y), W: enemySize, H: enemySize},
	)
}

func (e *enemy) bulletInCoordinates() (*bullet, bool) {
	for _, b := range bulletPool {
		if b.active &&
			b.x >= e.x-enemySize/2 &&
			b.x <= e.x+enemySize/2 &&
			b.y <= e.y+enemySize/2 &&
			b.y >= e.y-enemySize/2 {
			return b, true
		}
	}

	return nil, false
}

func (e *enemy) update() {
	if b, ok := e.bulletInCoordinates(); ok {
		b.active = false
		// fmt.Println("got")
		// e.destroy()
	}
}

type enemies []*enemy

func (e enemies) drawAndUpdate(renderer *sdl.Renderer) {
	for _, enemy := range e {
		enemy.draw(renderer)
		enemy.update()
	}
}
