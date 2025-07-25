//go:build !web

package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Define desktop values
var GridWidth int = 80
var GridHeight int = 50
var WindowWidth int = 2 * GridWidth * BASE_SPRITE_WIDTH
var WindowHeight int = 2 * GridHeight * BASE_SPRITE_HEIGHT

const PLAYER_INPUT_COOLDOWN int = 8

func (game *Game) run() {
	for !rl.WindowShouldClose() {
		game.handleInput()
		game.update()
		game.render()
	}
}
