package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileGraphic struct {
	Glyph   string
	FGColor color.RGBA
	BGColor color.RGBA
}

type Tile struct {
	Walkable    bool
	Transparent bool
	DarkGraphic TileGraphic
}

func NewTile(walkable bool, transparent bool, darkGraphic TileGraphic) Tile {
	return Tile{
		Walkable:    walkable,
		Transparent: transparent,
		DarkGraphic: darkGraphic,
	}
}

/* Tile types */
var Floor Tile = NewTile(true, true, TileGraphic{Glyph: ".", FGColor: rl.Gray, BGColor: rl.Black})

var Wall Tile = NewTile(false, false, TileGraphic{Glyph: "#", FGColor: rl.White, BGColor: rl.Black})
