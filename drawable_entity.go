package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const BASE_SPRITE_WIDTH int = 16
const BASE_SPRITE_HEIGHT int = 16

type DrawableEntity struct {
	game      *Game
	mapCoords MapCoords
	glyph     Glyph
	scale     int
	tint      color.RGBA
}

func initDrawableEntity(g *Game, m MapCoords, gl Glyph, s int, t color.RGBA) DrawableEntity {
	return DrawableEntity{
		game:      g,
		mapCoords: m,
		glyph:     gl,
		scale:     s,
		tint:      t,
	}
}

func (de *DrawableEntity) render() {
	rl.DrawTexturePro(
		de.game.spritesheet,
		rl.NewRectangle(
			float32(de.glyph.GX*BASE_SPRITE_WIDTH),
			float32(de.glyph.GY*BASE_SPRITE_HEIGHT),
			float32(BASE_SPRITE_WIDTH),
			float32(BASE_SPRITE_HEIGHT),
		),
		rl.NewRectangle(
			float32(de.mapCoords.X*de.scale*BASE_SPRITE_WIDTH),
			float32(de.mapCoords.Y*de.scale*BASE_SPRITE_HEIGHT),
			float32(BASE_SPRITE_WIDTH*de.scale),
			float32(BASE_SPRITE_HEIGHT*de.scale),
		),
		rl.Vector2{X: 0, Y: 0},
		0,
		de.tint,
	)
}

func (de *DrawableEntity) update() {
	// Clamp the player position to the screen
	de.mapCoords.X = int(rl.Clamp(float32(de.mapCoords.X), 0, float32(WindowGridWidth-1)))
	de.mapCoords.Y = int(rl.Clamp(float32(de.mapCoords.Y), 0, float32(WindowGridHeight-1)))
}
