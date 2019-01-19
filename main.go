package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()
	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}
	return texture
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Init SDL", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Space Shooter Game",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_OPENGL,
	)
	if err != nil {
		fmt.Println("Create Window", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Create Renderer", err)
		return
	}
	defer renderer.Destroy()
	plr := newPlayer(renderer)
	defer plr.texture.Destroy()
	enms, err := createEnemies(5, 3, renderer)
	if err != nil {
		fmt.Printf("Creating enemies failed %v", err)
	}
	initBulletPool(renderer)
	bg := createBackground(renderer, 0, 0)
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.Clear()
		bg.draw(renderer)
		plr.draw(renderer)
		plr.update()
		enms.draw(renderer)

		drawAndUpdateBullets(renderer)

		renderer.Present()
	}
}
