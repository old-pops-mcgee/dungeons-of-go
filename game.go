package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	spritesheet rl.Texture2D
	player      DrawableEntity
}

func initGame() Game {
	game := Game{
		spritesheet: rl.LoadTexture("assets/16x16-RogueYun-AgmEdit.png"),
	}
	game.player = initDrawableEntity(&game, 4, 4, 0, 4, 2, rl.White)
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

func (g *Game) update() {}

func (g *Game) handleInput() {}
