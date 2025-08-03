package main

import (
	"math"
	"math/rand/v2"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectangularRoom struct {
	gMap        *GameMap
	TopLeft     rl.Vector2
	BottomRight rl.Vector2
}

func GetNewRectangularRoom(gm *GameMap, tl rl.Vector2, width int, height int) RectangularRoom {
	room := RectangularRoom{
		gMap:    gm,
		TopLeft: tl,
	}

	room.BottomRight = rl.Vector2{X: tl.X + float32(width), Y: tl.Y + float32(height)}
	return room
}

func (r *RectangularRoom) GetCenter() rl.Vector2 {
	return rl.Vector2{
		X: float32(math.Floor(float64(r.TopLeft.X+r.BottomRight.X) / 2)),
		Y: float32(math.Floor(float64(r.TopLeft.Y+r.BottomRight.Y) / 2)),
	}
}

func (r *RectangularRoom) GetInnerIndices() []int {
	indices := []int{}
	for xCoord := r.TopLeft.X + 1; xCoord < r.BottomRight.X; xCoord++ {
		for yCoord := r.TopLeft.Y + 1; yCoord < r.BottomRight.Y; yCoord++ {
			indices = append(indices, r.gMap.CoordToIndex(rl.Vector2{X: xCoord, Y: yCoord}))
		}
	}
	return indices
}

func (r *RectangularRoom) Intersects(other RectangularRoom) bool {
	return r.TopLeft.X <= other.BottomRight.X &&
		r.BottomRight.X >= other.TopLeft.X &&
		r.TopLeft.Y <= other.BottomRight.Y &&
		r.BottomRight.Y >= other.TopLeft.Y
}

func GenerateTunnelIndices(g *Game, pointA rl.Vector2, pointB rl.Vector2) []int {
	coords := []rl.Vector2{}
	indices := []int{}
	// Get the random corner so we can create an L-shaped tunnel
	var corner rl.Vector2
	if rand.Float32() < 0.5 {
		// Move horizontally, then vertically
		corner = rl.Vector2{X: pointB.X, Y: pointA.Y}
	} else {
		// Move vertically, then horizontally
		corner = rl.Vector2{X: pointA.X, Y: pointB.Y}
	}

	coords = append(coords, BresenhamLine(pointA, corner)...)
	coords = append(coords, BresenhamLine(corner, pointB)...)

	for _, coord := range coords {
		indices = append(indices, g.gameMap.CoordToIndex(coord))
	}

	return indices
}

func GenerateDungeon(g *Game, maxRooms int, maxMonsters int, roomMaxSize int, roomMinSize int, mapWidth int, mapHeight int) GameMap {
	gMap := NewGameMap(g, mapWidth, mapHeight)
	g.gameMap = gMap

	roomList := []RectangularRoom{}

	for range maxRooms {
		// Generate the properties of the random room
		roomWidth := rand.IntN(roomMaxSize-roomMinSize) + roomMinSize
		roomHeight := rand.IntN(roomMaxSize-roomMinSize) + roomMinSize
		x := rand.IntN(gMap.Width - roomWidth - 1)
		y := rand.IntN(gMap.Height - roomHeight - 1)
		newRoom := GetNewRectangularRoom(&gMap, rl.Vector2{X: float32(x), Y: float32(y)}, roomWidth, roomHeight)

		// Ensure the room doesn't intersect with an existing room
		canAddRoom := true
		if slices.ContainsFunc(roomList, newRoom.Intersects) {
			canAddRoom = false
		}

		// If the room is valid, put it in
		if canAddRoom {
			for _, tileIndex := range newRoom.GetInnerIndices() {
				gMap.Tiles[tileIndex] = Floor
			}

			if len(roomList) == 0 {
				// Put the player in the first room
				g.player.drawableEntity.mapCoords = newRoom.GetCenter()
			} else {
				for _, tileIndex := range GenerateTunnelIndices(g, newRoom.GetCenter(), roomList[len(roomList)-1].GetCenter()) {
					gMap.Tiles[tileIndex] = Floor
				}
			}

			gMap.PlaceEntities(newRoom, maxMonsters)

			roomList = append(roomList, newRoom)
		}

	}

	return gMap
}
