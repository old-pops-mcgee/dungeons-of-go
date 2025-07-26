package main

import (
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameMap struct {
	game   *Game
	Tiles  []Tile
	Width  int
	Height int
}

func NewGameMap(g *Game, width int, height int) GameMap {
	gameMap := GameMap{
		game:   g,
		Width:  width,
		Height: height,
	}
	gameMap.Tiles = slices.Repeat([]Tile{Wall}, width*height)
	return gameMap
}

func (g *GameMap) CoordToIndex(coords rl.Vector2) int {
	return int(coords.Y*float32(g.Width) + coords.X)
}

func (g *GameMap) IndexToCoord(index int) rl.Vector2 {
	return rl.Vector2{
		X: float32(index % g.Width),
		Y: float32(index / g.Width),
	}
}

func (g *GameMap) IsInBounds(coords rl.Vector2) bool {
	return coords.X >= 0 && coords.X < float32(g.Width) && coords.Y >= 0 && coords.Y < float32(g.Height)
}

func (g *GameMap) render() {
	for index, tile := range g.Tiles {
		RenderTileBasedGraphic(g.game, tile.DarkGraphic.TileGlyph, g.IndexToCoord(index), tile.DarkGraphic.FGColor)
	}
}
