package main

import rl "github.com/gen2brain/raylib-go/raylib"

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
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Hello, World!", 20, 20, 20, rl.Blue)
	rl.EndDrawing()
}

func (g *Game) update() {}

func (g *Game) handleInput() {}
