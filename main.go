package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(int32(WindowWidth), int32(WindowHeight), "Hello World")
	rl.InitAudioDevice()
	rl.SetTargetFPS(8)

	game := initGame()

	game.run()

	game.unloadGame()

	rl.CloseAudioDevice()
	rl.CloseWindow()
}
