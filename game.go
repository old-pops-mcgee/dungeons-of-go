package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	paths "github.com/solarlune/paths"
)

const PLAYER_INPUT_COOLDOWN int = 4

var cameraZoom float32 = 2

var roomMaxSize int = 10
var roomMinSize int = 6
var maxRooms int = 30
var maxMonstersPerRoom = 2

type GameState int

const (
	WaitingForInput GameState = iota
	Playing
	WaitingToPlay
)

type Game struct {
	spritesheet                rl.Texture2D
	player                     *Entity
	playerInputCooldownCounter int
	gameMap                    *GameMap
	pathGrid                   *paths.Grid
	camera                     rl.Camera2D
	state                      GameState
}

func initGame() Game {
	game := Game{
		playerInputCooldownCounter: PLAYER_INPUT_COOLDOWN,
		spritesheet:                rl.LoadTexture("assets/16x16-RogueYun-AgmEdit.png"),
		state:                      WaitingForInput,
	}
	game.player = Player.Spawn(&game, rl.Vector2{X: 25, Y: 20})
	game.player.isPlayer = true
	// This function assigns the new dungeon to the game map
	GenerateDungeon(&game, maxRooms, maxMonstersPerRoom, roomMaxSize, roomMinSize, GridWidth, GridHeight)
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
	for _, i := range g.gameMap.Items {
		i.render()
	}
	for _, e := range g.gameMap.Entities {
		if g.player.FOVCalc.IsVisible(int(e.drawableEntity.mapCoords.X), int(e.drawableEntity.mapCoords.Y)) {
			e.render()
		}
	}
	g.player.render()
	rl.EndMode2D()
	rl.EndDrawing()
}

func (g *Game) update() {
	switch g.state {
	case WaitingForInput:
		// Check to see if we should still be playing
		if *g.player.currentHP <= 0 {
			g.state = WaitingToPlay
			fmt.Println("You died")
		} else {
			// Update the player
			g.player.update()

			// Update the camera
			g.camera.Target = g.getCameraTarget()

			// Update the cooldown timer
			g.playerInputCooldownCounter = max(0, g.playerInputCooldownCounter-1)

		}
	case Playing:
		newEntities := []Entity{}
		for _, e := range g.gameMap.Entities {
			e.update()
			if *e.currentHP > 0 {
				newEntities = append(newEntities, e)
			}
		}
		// Update the enemies
		g.gameMap.Entities = newEntities

		// Set the state to WaitingForInput to give player control
		g.state = WaitingForInput
	case WaitingToPlay:
		// Do nothing
	}

}

func (g *Game) handleInput() {
	if g.playerInputCooldownCounter <= 0 {
		processedKey := false
		for key, action := range MOVEMENT_KEYS {
			if rl.IsKeyDown(key) {
				g.player.movementActionSet[action] = true
				processedKey = true
				g.state = Playing
			}
		}

		// If we processed a key, reset the cooldown timer
		if processedKey {
			g.playerInputCooldownCounter = PLAYER_INPUT_COOLDOWN
		}
	}
}

func (g *Game) getCameraTarget() rl.Vector2 {
	return rl.Vector2{X: g.player.drawableEntity.mapCoords.X * float32(BASE_SPRITE_WIDTH), Y: g.player.drawableEntity.mapCoords.Y * float32(BASE_SPRITE_HEIGHT)}
}
