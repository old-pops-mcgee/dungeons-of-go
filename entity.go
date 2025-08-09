package main

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	fov "github.com/norendren/go-fov/fov"
)

type EntityAction interface {
	performAction(e *Entity)
}

type MoveAction struct {
	targetCoords rl.Vector2
}

func NewMoveAction(target rl.Vector2) MoveAction {
	return MoveAction{targetCoords: target}
}

func (m *MoveAction) performAction(e *Entity) {
	if !e.isPlayer {
		entityCell := e.game.pathGrid.Get(int(e.drawableEntity.mapCoords.X), int(e.drawableEntity.mapCoords.Y))
		if entityCell.Cost >= 6 {
			entityCell.Cost -= 5
		}

		targetCell := e.game.pathGrid.Get(int(m.targetCoords.X), int(m.targetCoords.Y))
		targetCell.Cost += 5
	}
	e.drawableEntity.mapCoords = m.targetCoords
}

type StandAction struct{}

func (s *StandAction) performAction(e *Entity) {
	// Do nothing
}

type MeleeAction struct {
	targetCoords rl.Vector2
}

func (m *MeleeAction) performAction(e *Entity) {
	fmt.Printf("I'm attacking the entity at %f, %f\n", m.targetCoords.X, m.targetCoords.Y)
}

type EntityTemplate struct {
	viewRadius int
	glyph      Glyph
	color      color.RGBA
	planner    Planner
	maxHP      int
	defense    int
	power      int
}

var (
	Player = EntityTemplate{viewRadius: 6, glyph: PlayerGlyph, color: rl.White, planner: PlayerPlanner{}, maxHP: 30, defense: 2, power: 5}
	Troll  = EntityTemplate{viewRadius: 4, glyph: TrollGlyph, color: rl.DarkGreen, planner: HostileEnemyPlanner{}, maxHP: 16, defense: 1, power: 4}
	Goblin = EntityTemplate{viewRadius: 6, glyph: GoblinGlyph, color: rl.Lime, planner: HostileEnemyPlanner{}, maxHP: 10, defense: 0, power: 3}
)

func (e *EntityTemplate) Spawn(g *Game, m rl.Vector2) *Entity {
	return initEntity(g, m, e.glyph, e.color, e.planner, e.viewRadius, e.maxHP, e.defense, e.power)
}

type Entity struct {
	game              *Game
	viewRadius        int
	planner           Planner
	drawableEntity    *DrawableEntity
	movementActionSet map[MovementAction]bool
	maxHP             int
	currentHP         int
	defense           int
	power             int
	isPlayer          bool
	FOVCalc           *fov.View
}

func initEntity(g *Game, m rl.Vector2, gl Glyph, t color.RGBA, pl Planner, vr int, mh int, d int, p int) *Entity {
	return &Entity{
		game:              g,
		viewRadius:        vr,
		planner:           pl,
		drawableEntity:    initDrawableEntity(g, m, gl, t),
		movementActionSet: map[MovementAction]bool{},
		maxHP:             mh,
		currentHP:         mh,
		defense:           d,
		power:             p,
		isPlayer:          false,
		FOVCalc:           fov.New(),
	}
}

func (e *Entity) render() {
	e.drawableEntity.render()
}

func (e *Entity) update() {
	entityAction := e.planner.planNextAction(e)
	entityAction.performAction(e)

	// Update the FOV
	e.FOVCalc.Compute(e.game.gameMap, int(e.drawableEntity.mapCoords.X), int(e.drawableEntity.mapCoords.Y), e.viewRadius)
}

func (e *Entity) getEntityActionForTarget(targetCoords rl.Vector2) EntityAction {
	// Validate the target position is in bounds
	if !e.game.gameMap.InBounds(int(targetCoords.X), int(targetCoords.Y)) {
		return &StandAction{}
	}

	// Validate the target position doesn't have an entity in it
	for _, otherEntity := range e.game.gameMap.Entities {
		if rl.Vector2Equals(targetCoords, otherEntity.drawableEntity.mapCoords) {
			return &MeleeAction{targetCoords: targetCoords}
		}
	}

	// Validate the target position isn't the player (relevant for non-player entities)
	if !e.isPlayer && rl.Vector2Equals(targetCoords, e.game.player.drawableEntity.mapCoords) {
		return &MeleeAction{targetCoords: targetCoords}
	}

	// Validate the target position is walkable, assuming it's in bounds
	targetIndex := e.game.gameMap.CoordToIndex(targetCoords)
	if e.game.gameMap.Tiles[targetIndex].Walkable {
		return &MoveAction{targetCoords: targetCoords}
	}
	return &StandAction{}
}
