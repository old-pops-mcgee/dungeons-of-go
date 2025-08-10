package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Handler interface {
	handleInput() bool
}

type GameHandler struct {
	game *Game
}

func (g *GameHandler) handleInput() bool {
	processedKey := false
	for key, action := range MOVEMENT_KEYS {
		if rl.IsKeyDown(key) {
			g.game.player.movementActionSet[action] = true
			processedKey = true
		}
	}

	return processedKey
}
