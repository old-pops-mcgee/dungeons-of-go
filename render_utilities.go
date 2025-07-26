package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RenderTileBasedGraphic(g *Game, gl Glyph, m MapCoords, scale int, tint color.RGBA) {
	rl.DrawTexturePro(
		g.spritesheet,
		rl.NewRectangle(
			float32(gl.GX*BASE_SPRITE_WIDTH),
			float32(gl.GY*BASE_SPRITE_HEIGHT),
			float32(BASE_SPRITE_WIDTH),
			float32(BASE_SPRITE_HEIGHT),
		),
		rl.NewRectangle(
			float32(m.X*scale*BASE_SPRITE_WIDTH),
			float32(m.Y*scale*BASE_SPRITE_HEIGHT),
			float32(BASE_SPRITE_WIDTH*scale),
			float32(BASE_SPRITE_HEIGHT*scale),
		),
		rl.Vector2{X: 0, Y: 0},
		0,
		tint,
	)
}
