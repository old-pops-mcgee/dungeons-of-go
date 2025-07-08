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
	game.player = initDrawableEntity(&game, 4, 4, 0, 4, Scale, rl.White)
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
	switch rl.GetKeyPressed() {
	case rl.KeyLeft, rl.KeyKp4, rl.KeyH:
		g.player.x -= 1
	case rl.KeyRight, rl.KeyKp6, rl.KeyL:
		g.player.x += 1
	case rl.KeyUp, rl.KeyKp8, rl.KeyK:
		g.player.y -= 1
	case rl.KeyDown, rl.KeyKp2, rl.KeyJ:
		g.player.y += 1
	case rl.KeyY, rl.KeyKp7:
		g.player.y -= 1
		g.player.x -= 1
	case rl.KeyU, rl.KeyKp9:
		g.player.y -= 1
		g.player.x += 1
	case rl.KeyB, rl.KeyKp1:
		g.player.y += 1
		g.player.x -= 1
	case rl.KeyN, rl.KeyKp3:
		g.player.y += 1
		g.player.x += 1
	}
}
