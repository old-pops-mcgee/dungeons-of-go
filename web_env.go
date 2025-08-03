//go:build web

package main

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets
var ASSETS embed.FS

var GridWidth int = 80
var GridHeight int = 50
var WindowWidth int = GridWidth * BASE_SPRITE_WIDTH
var WindowHeight int = GridHeight * BASE_SPRITE_HEIGHT

func init() {
	rl.AddFileSystem(ASSETS)
}

func (game *Game) run() {
	var update = func() {
		game.handleInput()
		game.update()
		game.render()
	}

	rl.SetMainLoop(update)

	for !rl.WindowShouldClose() {
		update()
	}
}
