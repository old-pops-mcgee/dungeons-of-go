package main

type Glyph struct {
	GX int // X position on the spritesheet
	GY int // Y position on the spritesheet
}

var PlayerGlyph Glyph = Glyph{GX: 0, GY: 4} // @
var FloorGlyph Glyph = Glyph{GX: 14, GY: 2} // .
var WallGlyph Glyph = Glyph{GX: 3, GY: 2}   // #
var BoxGlyph Glyph = Glyph{GX: 12, GY: 14}  // Box
