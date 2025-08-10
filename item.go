package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ItemTemplate struct {
	glyph Glyph
	color color.RGBA
}

type Item struct {
	game           *Game
	drawableEntity *DrawableEntity
}

var (
	Corpse = ItemTemplate{glyph: CorpseGlyph, color: rl.Red}
)

func (i *ItemTemplate) Spawn(g *Game, m rl.Vector2) *Item {
	return initItem(g, m, i.glyph, i.color)
}

func initItem(g *Game, m rl.Vector2, gl Glyph, t color.RGBA) *Item {
	return &Item{
		game:           g,
		drawableEntity: initDrawableEntity(g, m, gl, t),
	}
}

func (i *Item) render() {
	i.drawableEntity.render()
}
