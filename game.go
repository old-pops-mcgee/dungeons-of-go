package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	spritesheet rl.Texture2D
}

func initGame() Game {
	return Game{
		spritesheet: rl.LoadTexture("assets/16x16-RogueYun-AgmEdit.png"),
	}
}

func (g *Game) unloadGame() {
	rl.UnloadTexture(g.spritesheet)
}

func (g *Game) render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.DrawTexturePro(
		g.spritesheet, rl.NewRectangle(16, 0, 16, 16), rl.NewRectangle(80, 80, 16, 16), rl.Vector2{X: 0, Y: 0}, 0, rl.White)
	rl.EndDrawing()
}

func (g *Game) update() {}

func (g *Game) handleInput() {}
