//go:build web

package main

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets
var ASSETS embed.FS

var WindowGridWidth int = 80
var WindowGridHeight int = 50
var Scale int = 1

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
