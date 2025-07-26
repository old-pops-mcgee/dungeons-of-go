package main

import rl "github.com/gen2brain/raylib-go/raylib"

func BresenhamLine(pointA rl.Vector2, pointB rl.Vector2) []rl.Vector2 {
	outputs := []rl.Vector2{}

	x0 := pointA.X
	y0 := pointA.Y
	x1 := pointB.X
	y1 := pointB.Y

	dx := x1 - x0
	dy := y1 - y0
	d := 2*dy - dx
	y := y0

	for x := x0; x <= x1; x += 1 {
		outputs = append(outputs, rl.Vector2{X: x, Y: y})
		if d > 0 {
			y++
			d -= 2 * dx
		}
		d += 2 * dy
	}

	return outputs
}
