package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileGraphic struct {
	TileGlyph Glyph
	FGColor   color.RGBA
	BGColor   color.RGBA
}

type Tile struct {
	Walkable     bool
	Transparent  bool
	DarkGraphic  TileGraphic
	LightGraphic TileGraphic
}

func NewTile(walkable bool, transparent bool, darkGraphic TileGraphic, lightGraphic TileGraphic) Tile {
	return Tile{
		Walkable:     walkable,
		Transparent:  transparent,
		DarkGraphic:  darkGraphic,
		LightGraphic: lightGraphic,
	}
}

var Shroud TileGraphic = TileGraphic{TileGlyph: BoxGlyph, FGColor: rl.Black, BGColor: rl.Black}

/* Tile types */
var Floor Tile = NewTile(
	true,
	true,
	TileGraphic{TileGlyph: FloorGlyph, FGColor: rl.Gray, BGColor: rl.Black},
	TileGraphic{TileGlyph: FloorGlyph, FGColor: rl.RayWhite, BGColor: rl.Black},
)

var Wall Tile = NewTile(
	false,
	false,
	TileGraphic{TileGlyph: WallGlyph, FGColor: rl.DarkBrown, BGColor: rl.Black},
	TileGraphic{TileGlyph: WallGlyph, FGColor: rl.Brown, BGColor: rl.Black},
)
