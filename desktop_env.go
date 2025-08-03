//go:build !web

package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Define desktop values
var GridWidth int = 80
var GridHeight int = 50
var WindowWidth int = GridWidth * BASE_SPRITE_WIDTH
var WindowHeight int = GridHeight * BASE_SPRITE_HEIGHT

func (game *Game) run() {
	for !rl.WindowShouldClose() {
		game.handleInput()
		game.update()
		game.render()
	}
}
