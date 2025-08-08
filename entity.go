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

type EntityTemplate struct {
	viewRadius int
	glyph      Glyph
	color      color.RGBA
	maxHP      int
	defense    int
	power      int
}

var (
	Player = EntityTemplate{viewRadius: 6, glyph: PlayerGlyph, color: rl.White, maxHP: 30, defense: 2, power: 5}
	Troll  = EntityTemplate{viewRadius: 4, glyph: TrollGlyph, color: rl.DarkGreen, maxHP: 16, defense: 1, power: 4}
	Goblin = EntityTemplate{viewRadius: 6, glyph: GoblinGlyph, color: rl.Lime, maxHP: 10, defense: 0, power: 3}
)

func (e *EntityTemplate) Spawn(g *Game, m rl.Vector2) *Entity {
	return initEntity(g, m, e.glyph, e.color, e.viewRadius, e.maxHP, e.defense, e.power)
}

type Entity struct {
	game              *Game
	viewRadius        int
	drawableEntity    DrawableEntity
	movementActionSet map[MovementAction]bool
	maxHP             int
	currentHP         int
	defense           int
	power             int
}

func initEntity(g *Game, m rl.Vector2, gl Glyph, t color.RGBA, vr int, mh int, d int, p int) *Entity {
	return &Entity{
		game:              g,
		viewRadius:        vr,
		drawableEntity:    initDrawableEntity(g, m, gl, t),
		movementActionSet: map[MovementAction]bool{},
		maxHP:             mh,
		currentHP:         mh,
		defense:           d,
		power:             p,
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
