package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	fov "github.com/norendren/go-fov/fov"
)

var cameraZoom float32 = 2

var roomMaxSize int = 10
var roomMinSize int = 6
var maxRooms int = 30

type Game struct {
	spritesheet rl.Texture2D
	player      Player
	gameMap     GameMap
	camera      rl.Camera2D
	FOVCalc     *fov.View
}

func initGame() Game {
	game := Game{
		spritesheet: rl.LoadTexture("assets/16x16-RogueYun-AgmEdit.png"),
	}
	game.player = initPlayer(&game, rl.Vector2{X: 25, Y: 20}, PlayerGlyph, rl.White)
	// This function assigns the new dungeon to the game map
	GenerateDungeon(&game, maxRooms, roomMaxSize, roomMinSize, GridWidth, GridHeight)
	game.camera = rl.Camera2D{
		Target:   game.getCameraTarget(),
		Offset:   rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2},
		Rotation: 0,
		Zoom:     cameraZoom,
	}
	game.FOVCalc = fov.New()
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

	// Update the FOV
	g.FOVCalc.Compute(&g.gameMap, int(g.player.drawableEntity.mapCoords.X), int(g.player.drawableEntity.mapCoords.Y), g.player.viewRadius)

	// Update the camera
	g.camera.Target = g.getCameraTarget()
}

func (g *Game) handleInput() {
	for key, action := range MOVEMENT_KEYS {
		if rl.IsKeyDown(key) {
			g.player.movementActionSet[action] = true
		}
	}
}

func (g *Game) getCameraTarget() rl.Vector2 {
	return rl.Vector2{X: g.player.drawableEntity.mapCoords.X * float32(BASE_SPRITE_WIDTH), Y: g.player.drawableEntity.mapCoords.Y * float32(BASE_SPRITE_HEIGHT)}
}
