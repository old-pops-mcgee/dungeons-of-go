package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	spritesheet rl.Texture2D
	player      Player
}

func initGame() Game {
	game := Game{
		spritesheet: rl.LoadTexture("assets/16x16-RogueYun-AgmEdit.png"),
	}
	game.player = initPlayer(&game, 4, 4, 0, 4, Scale, rl.White)
	return game
}

func (g *Game) unloadGame() {
	rl.UnloadTexture(g.spritesheet)
}

func (g *Game) render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	g.player.render()
	rl.EndDrawing()
}

func (g *Game) update() {
	g.player.update()
}

func (g *Game) handleInput() {
	for key, action := range MOVEMENT_KEYS {
		if rl.IsKeyDown(key) {
			g.player.movementActionSet[action] = true
		}
	}
}
