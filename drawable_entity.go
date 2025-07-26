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
	tint      color.RGBA
}

func initDrawableEntity(g *Game, m MapCoords, gl Glyph, t color.RGBA) DrawableEntity {
	return DrawableEntity{
		game:      g,
		mapCoords: m,
		glyph:     gl,
		tint:      t,
	}
}

func (de *DrawableEntity) render() {
	RenderTileBasedGraphic(de.game, de.glyph, de.mapCoords, de.tint)
}

func (de *DrawableEntity) update() {
	// Clamp the player position to the screen
	de.mapCoords.X = int(rl.Clamp(float32(de.mapCoords.X), 0, float32(WindowGridWidth-1)))
	de.mapCoords.Y = int(rl.Clamp(float32(de.mapCoords.Y), 0, float32(WindowGridHeight-1)))
}
