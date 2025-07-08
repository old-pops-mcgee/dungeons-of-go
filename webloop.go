//go:build web

package main

import rl "github.com/gen2brain/raylib-go/raylib"

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
