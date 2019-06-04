package main

import (
	"time"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	enemySize           = 80
	enemyAnimationDelay = time.Millisecond * 90
)

type enemy struct {
	texture *sdl.Texture
	x, y    float64
	active 	bool
}

func newBasicEnemy(texture *sdl.Texture, x, y float64) (e enemy) {
	e.texture = texture

	e.x = x
	e.y = y
	e.active = true

	return
}

func createEnemies(x, y int, renderer *sdl.Renderer) (enms enemies, err error) {
	texture := textureFromBMP(renderer, "./sprites/alienblaster.bmp")
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			x := (float64(i)/5)*screenWidth + enemySize/2 + 15
			y := float64(j)*enemySize + enemySize/2 + float64(j*10)
			enm := newBasicEnemy(texture, x, y)
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

func (e *enemy) bulletsInCoordinates() ([]*bullet, bool) {
	bulletsToReturn := []*bullet{}
	for i, b := range bulletPool {
		if b.enemyInCoordinates(e){
			if i != 0 && i  != len(bulletPool) - 1 {
				if bulletPool[i - 1].enemyInCoordinates(e){
					bulletsToReturn = append(bulletsToReturn, bulletPool[i - 1])
				}
				if bulletPool[i + 1].enemyInCoordinates(e){
					bulletsToReturn = append(bulletsToReturn, bulletPool[i + 1])
				}
			} else if  i == len(bulletPool) - 1 {
				if bulletPool[i - 1].enemyInCoordinates(e){
					bulletsToReturn = append(bulletsToReturn, bulletPool[i - 1])
				}
			} else if i == 0 {
				if bulletPool[i + 1].enemyInCoordinates(e){
					bulletsToReturn = append(bulletsToReturn, bulletPool[i + 1])
				}
			}

			return append(bulletsToReturn, b), true
		}
	}

	return nil, false
}

func (e *enemy) update() {
	if b, ok := e.bulletsInCoordinates(); ok {
		for _, bull := range b {
			bull.active = false
		}
		e.active = false
	}
}

type enemies []*enemy

func (e enemies) drawAndUpdate(renderer *sdl.Renderer) {
	newEnem := []*enemy{}
	for _ , enem := range e {
		if enem.active {
			enem.draw(renderer)
			enem.update()
			newEnem = append(newEnem, enem)
		}
	}
	e = newEnem
}
