package main

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EntityAction int

const (
	Move EntityAction = iota
	Melee
	Stand
)

type Entity struct {
	game              *Game
	viewRadius        int
	drawableEntity    DrawableEntity
	movementActionSet map[MovementAction]bool
}

func initEntity(g *Game, m rl.Vector2, gl Glyph, t color.RGBA) *Entity {
	return &Entity{
		game:              g,
		viewRadius:        6,
		drawableEntity:    initDrawableEntity(g, m, gl, t),
		movementActionSet: map[MovementAction]bool{},
	}
}

func (e *Entity) render() {
	e.drawableEntity.render()
}

func (e *Entity) update() {
	var movementDelta MovementDelta

	// Process movements from each movement action
	for movement := range e.movementActionSet {
		tempMovementDelta := MOVEMENT_DELTAS[movement]
		movementDelta.dx += tempMovementDelta.dx
		movementDelta.dy += tempMovementDelta.dy

		// Clear the movement from the action set
		delete(e.movementActionSet, movement)
	}

	// Clamp the movement deltas to ensure we don't process to big a step
	movementDelta.dx = int(rl.Clamp(float32(movementDelta.dx), -1, 1))
	movementDelta.dy = int(rl.Clamp(float32(movementDelta.dy), -1, 1))

	// Find the target coordinates
	targetCoords := e.drawableEntity.mapCoords
	targetCoords.X += float32(movementDelta.dx)
	targetCoords.Y += float32(movementDelta.dy)

	entityAction := e.getEntityActionForTarget(targetCoords)

	switch entityAction {
	case Stand:
		// Do nothing
	case Move:
		e.drawableEntity.mapCoords = targetCoords
	case Melee:
		fmt.Println("I'm attacking the entity!")
	}
}

func (e *Entity) getEntityActionForTarget(targetCoords rl.Vector2) EntityAction {
	// Validate the target position is in bounds
	if !e.game.gameMap.InBounds(int(targetCoords.X), int(targetCoords.Y)) {
		return Stand
	}

	// Validate the target position doesn't have an entity in it
	for _, otherEntity := range e.game.gameMap.Entities {
		if rl.Vector2Equals(targetCoords, otherEntity.drawableEntity.mapCoords) {
			return Melee
		}
	}

	// Validate the target position is walkable, assuming it's in bounds
	targetIndex := e.game.gameMap.CoordToIndex(targetCoords)
	if e.game.gameMap.Tiles[targetIndex].Walkable {
		return Move
	}
	return Stand
}
