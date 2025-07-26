package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const PLAYER_INPUT_COOLDOWN int = 6

var cameraZoom float32 = 2

type Game struct {
	spritesheet                rl.Texture2D
	player                     Player
	playerInputCooldownCounter int
	gameMap                    GameMap
	camera                     rl.Camera2D
}

func initGame() Game {
	game := Game{
		spritesheet:                rl.LoadTexture("assets/16x16-RogueYun-AgmEdit.png"),
		playerInputCooldownCounter: PLAYER_INPUT_COOLDOWN,
	}
	game.player = initPlayer(&game, MapCoords{X: 25, Y: 20}, PlayerGlyph, rl.White)
	game.gameMap = GenerateDungeon(&game, GridWidth, GridHeight)
	game.camera = rl.Camera2D{
		Target:   game.getCameraTarget(),
		Offset:   rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2},
		Rotation: 0,
		Zoom:     cameraZoom,
	}
	return game
}

func (g *Game) unloadGame() {
	rl.UnloadTexture(g.spritesheet)
}

func (g *Game) render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.BeginMode2D(g.camera)
	g.gameMap.render()
	g.player.render()
	rl.EndMode2D()
	rl.EndDrawing()
}

func (g *Game) update() {
	// Update the player
	g.player.update()

	// Update the camera
	g.camera.Target = g.getCameraTarget()

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

func (g *Game) getCameraTarget() rl.Vector2 {
	return rl.Vector2{X: float32(g.player.drawableEntity.mapCoords.X * BASE_SPRITE_WIDTH), Y: float32(g.player.drawableEntity.mapCoords.Y * BASE_SPRITE_HEIGHT)}
}
