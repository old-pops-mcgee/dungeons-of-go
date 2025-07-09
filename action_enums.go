package main

import rl "github.com/gen2brain/raylib-go/raylib"

type MovementAction int

type MovementDelta struct {
	dx int
	dy int
}

const (
	MOVE_N MovementAction = iota
	MOVE_NW
	MOVE_W
	MOVE_SW
	MOVE_S
	MOVE_SE
	MOVE_E
	MOVE_NE
)

// Maps the input keys to the movement actions they should correspond to
var MOVEMENT_KEYS map[int32]MovementAction = map[int32]MovementAction{
	rl.KeyLeft:  MOVE_W,
	rl.KeyKp4:   MOVE_W,
	rl.KeyH:     MOVE_W,
	rl.KeyRight: MOVE_E,
	rl.KeyKp6:   MOVE_E,
	rl.KeyL:     MOVE_E,
	rl.KeyUp:    MOVE_N,
	rl.KeyKp8:   MOVE_N,
	rl.KeyK:     MOVE_N,
	rl.KeyDown:  MOVE_S,
	rl.KeyKp2:   MOVE_S,
	rl.KeyJ:     MOVE_S,
	rl.KeyY:     MOVE_NW,
	rl.KeyKp7:   MOVE_NW,
	rl.KeyU:     MOVE_NE,
	rl.KeyKp9:   MOVE_NE,
	rl.KeyB:     MOVE_SW,
	rl.KeyKp1:   MOVE_SW,
	rl.KeyN:     MOVE_SE,
	rl.KeyKp3:   MOVE_SE,
}

// Maps the movement actions to the movement delta on the game grid
var MOVEMENT_DELTAS map[MovementAction]MovementDelta = map[MovementAction]MovementDelta{
	MOVE_N:  {dx: 0, dy: -1},
	MOVE_NW: {dx: -1, dy: -1},
	MOVE_W:  {dx: -1, dy: 0},
	MOVE_SW: {dx: -1, dy: 1},
	MOVE_S:  {dx: 0, dy: 1},
	MOVE_SE: {dx: 1, dy: 1},
	MOVE_E:  {dx: 1, dy: 0},
	MOVE_NE: {dx: 1, dy: -1},
}
