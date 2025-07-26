package main

type RectangularRoom struct {
	gMap        *GameMap
	TopLeft     MapCoords
	BottomRight MapCoords
}

func GetNewRectangularRoom(gm *GameMap, tl MapCoords, width int, height int) RectangularRoom {
	room := RectangularRoom{
		gMap:    gm,
		TopLeft: tl,
	}

	room.BottomRight = MapCoords{X: tl.X + width, Y: tl.Y + height}
	return room
}

func (r *RectangularRoom) GetCenter() MapCoords {
	return MapCoords{
		X: (r.TopLeft.X + r.BottomRight.X) / 2,
		Y: (r.TopLeft.Y + r.BottomRight.Y) / 2,
	}
}

func (r *RectangularRoom) GetInnerIndices() []int {
	indices := []int{}
	for xCoord := r.TopLeft.X + 1; xCoord < r.BottomRight.X; xCoord++ {
		for yCoord := r.TopLeft.Y + 1; yCoord < r.BottomRight.Y; yCoord++ {
			indices = append(indices, r.gMap.CoordToIndex(MapCoords{X: xCoord, Y: yCoord}))
		}
	}
	return indices
}

func GenerateDungeon(g *Game, mapWidth int, mapHeight int) GameMap {
	gMap := NewGameMap(g, mapWidth, mapHeight)

	r1 := GetNewRectangularRoom(&gMap, MapCoords{X: 20, Y: 15}, 10, 15)
	r2 := GetNewRectangularRoom(&gMap, MapCoords{X: 35, Y: 15}, 10, 15)

	for _, tileIndex := range r1.GetInnerIndices() {
		gMap.Tiles[tileIndex] = Floor
	}

	for _, tileIndex := range r2.GetInnerIndices() {
		gMap.Tiles[tileIndex] = Floor
	}

	return gMap
}
