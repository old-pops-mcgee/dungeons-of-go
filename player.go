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
		p.drawableEntity.mapCoords.X += movementDelta.dx
		p.drawableEntity.mapCoords.Y += movementDelta.dy

		// Clear the movement from the action set
		delete(p.movementActionSet, movement)
	}

	// Clamp the position of the player's drawable entity
	p.drawableEntity.update()
}
