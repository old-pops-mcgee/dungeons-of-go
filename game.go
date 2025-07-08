package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
}

func initGame() Game {
	return Game{}
}

func (g *Game) unloadGame() {
}

func (g *Game) render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Hello, World!", 20, 20, 20, rl.Blue)
	rl.EndDrawing()
}

func (g *Game) update() {}

func (g *Game) handleInput() {}
