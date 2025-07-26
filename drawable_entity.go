package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const BASE_SPRITE_WIDTH int = 16
const BASE_SPRITE_HEIGHT int = 16

type DrawableEntity struct {
	game      *Game
	mapCoords rl.Vector2
	glyph     Glyph
	tint      color.RGBA
}

func initDrawableEntity(g *Game, m rl.Vector2, gl Glyph, t color.RGBA) DrawableEntity {
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
	de.mapCoords.X = rl.Clamp(de.mapCoords.X, 0, float32(GridWidth-1))
	de.mapCoords.Y = rl.Clamp(de.mapCoords.Y, 0, float32(GridHeight-1))
}
