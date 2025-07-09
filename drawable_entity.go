package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const BASE_SPRITE_WIDTH int = 16
const BASE_SPRITE_HEIGHT int = 16

type DrawableEntity struct {
	game         *Game
	x            int
	y            int
	spritesheetX int // X position
	spritesheetY int
	scale        int
	tint         color.RGBA
}

func initDrawableEntity(g *Game, x int, y int, sx int, sy int, s int, t color.RGBA) DrawableEntity {
	return DrawableEntity{
		game:         g,
		x:            x,
		y:            y,
		spritesheetX: sx,
		spritesheetY: sy,
		scale:        s,
		tint:         t,
	}
}

func (de *DrawableEntity) render() {
	rl.DrawTexturePro(
		de.game.spritesheet,
		rl.NewRectangle(
			float32(de.spritesheetX*BASE_SPRITE_WIDTH),
			float32(de.spritesheetY*BASE_SPRITE_HEIGHT),
			float32(BASE_SPRITE_WIDTH),
			float32(BASE_SPRITE_HEIGHT),
		),
		rl.NewRectangle(
			float32(de.x*de.scale*BASE_SPRITE_WIDTH),
			float32(de.y*de.scale*BASE_SPRITE_HEIGHT),
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
	de.x = int(rl.Clamp(float32(de.x), 0, float32(WindowGridWidth-1)))
	de.y = int(rl.Clamp(float32(de.y), 0, float32(WindowGridHeight-1)))
}
