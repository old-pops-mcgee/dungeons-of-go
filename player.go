package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	game              *Game
	viewRadius        int
	drawableEntity    DrawableEntity
	movementActionSet map[MovementAction]bool
}

func initPlayer(g *Game, m rl.Vector2, gl Glyph, t color.RGBA) *Player {
	return &Player{
		game:              g,
		viewRadius:        6,
		drawableEntity:    initDrawableEntity(g, m, gl, t),
		movementActionSet: map[MovementAction]bool{},
	}
}

func (p *Player) render() {
	p.drawableEntity.render()
}

func (p *Player) update() {
	var movementDelta MovementDelta

	// Process movements from each movement action
	for movement := range p.movementActionSet {
		tempMovementDelta := MOVEMENT_DELTAS[movement]
		movementDelta.dx += tempMovementDelta.dx
		movementDelta.dy += tempMovementDelta.dy

		// Clear the movement from the action set
		delete(p.movementActionSet, movement)
	}

	// Clamp the movement deltas to ensure we don't process to big a step
	movementDelta.dx = int(rl.Clamp(float32(movementDelta.dx), -1, 1))
	movementDelta.dy = int(rl.Clamp(float32(movementDelta.dy), -1, 1))

	// Find the target coordinates
	targetCoords := p.drawableEntity.mapCoords
	targetCoords.X += float32(movementDelta.dx)
	targetCoords.Y += float32(movementDelta.dy)

	if p.isValidMovementTarget(targetCoords) {
		p.drawableEntity.mapCoords = targetCoords
	}
}

func (p *Player) isValidMovementTarget(targetCoords rl.Vector2) bool {
	// Validate the target position is in bounds
	if !p.game.gameMap.InBounds(int(targetCoords.X), int(targetCoords.Y)) {
		return false
	}

	// Validate the target position is walkable, assuming it's in bounds
	targetIndex := p.game.gameMap.CoordToIndex(targetCoords)
	return p.game.gameMap.Tiles[targetIndex].Walkable
}
