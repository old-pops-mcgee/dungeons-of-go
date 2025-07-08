package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 600, "Hello World")
	rl.InitAudioDevice()
	rl.SetTargetFPS(60)

	game := initGame()

	game.run()

	game.unloadGame()

	rl.CloseAudioDevice()
	rl.CloseWindow()
}
