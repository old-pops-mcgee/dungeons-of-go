package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	game              *Game
	drawableEntity    DrawableEntity
	movementActionSet map[MovementAction]bool
}

func initPlayer(g *Game, m rl.Vector2, gl Glyph, t color.RGBA) Player {
	return Player{
		game:              g,
		drawableEntity:    initDrawableEntity(g, m, gl, t),
		movementActionSet: map[MovementAction]bool{},
	}
}

func (p *Player) render() {
	p.drawableEntity.render()
}

func (p *Player) update() {
	// Process movements from movement action
	for movement := range p.movementActionSet {
		movementDelta := MOVEMENT_DELTAS[movement]

		// Find the target coordinates
		targetCoords := p.drawableEntity.mapCoords
		targetCoords.X += float32(movementDelta.dx)
		targetCoords.Y += float32(movementDelta.dy)

		if p.isValidMovementTarget(targetCoords) {
			p.drawableEntity.mapCoords = targetCoords
		}

		// Clear the movement from the action set
		delete(p.movementActionSet, movement)
	}
}

func (p *Player) isValidMovementTarget(targetCoords rl.Vector2) bool {
	// Validate the target position is in bounds
	if !p.game.gameMap.IsInBounds(targetCoords) {
		return false
	}

	// Validate the target position is walkable, assuming it's in bounds
	targetIndex := p.game.gameMap.CoordToIndex(targetCoords)
	return p.game.gameMap.Tiles[targetIndex].Walkable
}
