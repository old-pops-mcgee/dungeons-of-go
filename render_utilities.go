package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RenderTileBasedGraphic(g *Game, gl Glyph, m rl.Vector2, tint color.RGBA) {
	rl.DrawTexturePro(
		g.spritesheet,
		rl.NewRectangle(
			float32(gl.GX*BASE_SPRITE_WIDTH),
			float32(gl.GY*BASE_SPRITE_HEIGHT),
			float32(BASE_SPRITE_WIDTH),
			float32(BASE_SPRITE_HEIGHT),
		),
		rl.NewRectangle(
			m.X*float32(BASE_SPRITE_WIDTH),
			m.Y*float32(BASE_SPRITE_HEIGHT),
			float32(BASE_SPRITE_WIDTH),
			float32(BASE_SPRITE_HEIGHT),
		),
		rl.Vector2{X: 0, Y: 0},
		0,
		tint,
	)
}
