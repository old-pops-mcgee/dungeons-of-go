package main

import (
	"slices"
)

type GameMap struct {
	game   *Game
	Tiles  []Tile
	Width  int
	Height int
}

type MapCoords struct {
	X int
	Y int
}

func NewGameMap(g *Game, width int, height int) GameMap {
	gameMap := GameMap{
		game:   g,
		Width:  width,
		Height: height,
	}
	gameMap.Tiles = slices.Repeat([]Tile{Floor}, width*height)
	gameMap.Tiles[5] = Wall
	gameMap.Tiles[5+width] = Wall
	gameMap.Tiles[5+2*width] = Wall
	return gameMap
}

func (g GameMap) CoordToIndex(coords MapCoords) int {
	return coords.Y*g.Width + coords.X
}

func (g GameMap) IndexToCoord(index int) MapCoords {
	return MapCoords{
		X: index % g.Width,
		Y: index / g.Width,
	}
}

func (g GameMap) IsInBounds(coords MapCoords) bool {
	return g.CoordToIndex(coords) < len(g.Tiles)
}

func (g GameMap) render() {
	for index, tile := range g.Tiles {
		RenderTileBasedGraphic(g.game, tile.DarkGraphic.TileGlyph, g.IndexToCoord(index), Scale, tile.DarkGraphic.FGColor)
	}
}
