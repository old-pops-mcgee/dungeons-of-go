//go:build !web

package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Define desktop values
var WindowGridWidth int = 80
var WindowGridHeight int = 50
var Scale int = 2

func (game *Game) run() {
	for !rl.WindowShouldClose() {
		game.handleInput()
		game.update()
		game.render()
	}
}
