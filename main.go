package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(int32(WindowGridWidth*Scale*BASE_SPRITE_WIDTH), int32(WindowGridHeight*Scale*BASE_SPRITE_HEIGHT), "Hello World")
	rl.InitAudioDevice()
	rl.SetTargetFPS(60)

	game := initGame()

	game.run()

	game.unloadGame()

	rl.CloseAudioDevice()
	rl.CloseWindow()
}
