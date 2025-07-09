//go:build !web

package main

import rl "github.com/gen2brain/raylib-go/raylib"

func (game *Game) run() {
	for !rl.WindowShouldClose() {
		game.handleInput()
		game.update()
		game.render()
	}
}
