package main

import "image/color"

type Player struct {
	game              *Game
	drawableEntity    DrawableEntity
	movementActionSet map[MovementAction]bool
}

func initPlayer(g *Game, m MapCoords, gl Glyph, s int, t color.RGBA) Player {
	return Player{
		game:              g,
		drawableEntity:    initDrawableEntity(g, m, gl, s, t),
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
		targetCoords.X += movementDelta.dx
		targetCoords.Y += movementDelta.dy

		if p.isValidMovementTarget(targetCoords) {
			p.drawableEntity.mapCoords = targetCoords
		}

		// Clear the movement from the action set
		delete(p.movementActionSet, movement)
	}
}

func (p *Player) isValidMovementTarget(targetCoords MapCoords) bool {
	// Validate the target position is in bounds
	if !p.game.gameMap.IsInBounds(targetCoords) {
		return false
	}

	// Validate the target position is walkable, assuming it's in bounds
	targetIndex := p.game.gameMap.CoordToIndex(targetCoords)
	return p.game.gameMap.Tiles[targetIndex].Walkable
}
