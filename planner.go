package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Planner interface {
	planNextAction(e *Entity) EntityAction
}

type PlayerPlanner struct {
}

func (p PlayerPlanner) planNextAction(e *Entity) EntityAction {
	// Generate the movement action set
	movementActionSet := map[MovementAction]bool{}
	var movementDelta MovementDelta

	for k, v := range e.movementActionSet {
		movementActionSet[k] = v
		delete(e.movementActionSet, k)
	}

	// Process movements from each movement action
	for movement := range movementActionSet {
		tempMovementDelta := MOVEMENT_DELTAS[movement]
		movementDelta.dx += tempMovementDelta.dx
		movementDelta.dy += tempMovementDelta.dy
	}

	// Clamp the movement deltas to ensure we don't process to big a step
	movementDelta.dx = int(rl.Clamp(float32(movementDelta.dx), -1, 1))
	movementDelta.dy = int(rl.Clamp(float32(movementDelta.dy), -1, 1))

	// Find the target coordinates
	targetCoords := e.drawableEntity.mapCoords

	targetCoords.X += float32(movementDelta.dx)
	targetCoords.Y += float32(movementDelta.dy)

	return e.getEntityActionForTarget(targetCoords)
}

type HostileEnemyPlanner struct {
}

func (h HostileEnemyPlanner) planNextAction(e *Entity) EntityAction {
	playerCoords := e.game.player.drawableEntity.mapCoords
	playerCell := e.game.pathGrid.Get(int(playerCoords.X), int(playerCoords.Y))
	entityCoords := e.drawableEntity.mapCoords
	entityCell := e.game.pathGrid.Get(int(entityCoords.X), int(entityCoords.Y))

	path := e.game.pathGrid.GetPathFromCells(entityCell, playerCell, true, false)
	nextStep := path.Next()

	nextAction := e.getEntityActionForTarget(rl.Vector2{X: float32(nextStep.X), Y: float32(nextStep.Y)})

	return nextAction
}
