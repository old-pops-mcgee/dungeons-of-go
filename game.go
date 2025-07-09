package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const PLAYER_INPUT_COOLDOWN int = 6

type Game struct {
	spritesheet                rl.Texture2D
	player                     Player
	playerInputCooldownCounter int
}

func initGame() Game {
	game := Game{
		spritesheet:                rl.LoadTexture("assets/16x16-RogueYun-AgmEdit.png"),
		playerInputCooldownCounter: PLAYER_INPUT_COOLDOWN,
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
	// Update the player
	g.player.update()

	// Update the cooldown timer
	g.playerInputCooldownCounter = max(0, g.playerInputCooldownCounter-1)
}

func (g *Game) handleInput() {
	if g.playerInputCooldownCounter <= 0 {
		processedKey := false
		for key, action := range MOVEMENT_KEYS {
			if rl.IsKeyDown(key) {
				g.player.movementActionSet[action] = true
				processedKey = true
			}
		}

		// If we processed a key, reset the cooldown timer
		if processedKey {
			g.playerInputCooldownCounter = PLAYER_INPUT_COOLDOWN
		}
	}

}
