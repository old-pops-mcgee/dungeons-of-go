package main

import (
	"fmt"
)

type Planner interface {
	planMovementActionSet(e *Entity) map[MovementAction]bool
}

type PlayerPlanner struct {
}

func (p PlayerPlanner) planMovementActionSet(e *Entity) map[MovementAction]bool {
	actionSet := map[MovementAction]bool{}

	for k, v := range e.movementActionSet {
		actionSet[k] = v
		delete(e.movementActionSet, k)
	}

	return actionSet
}

type HostileEnemyPlanner struct {
}

func (h HostileEnemyPlanner) planMovementActionSet(e *Entity) map[MovementAction]bool {
	actionSet := map[MovementAction]bool{}
	actionSet[MovementAction(Stand)] = true
	fmt.Println("Contemplating actions")
	return actionSet
}
