package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func BresenhamLine(pointA rl.Vector2, pointB rl.Vector2) []rl.Vector2 {
	outputs := []rl.Vector2{}

	plotLineLow := func(o *[]rl.Vector2, x0 float32, y0 float32, x1 float32, y1 float32) {
		dx := x1 - x0
		dy := y1 - y0
		yi := float32(1)

		if dy < 0 {
			yi = -1
			dy *= -1
		}

		d := (2 * dy) - dx
		y := y0

		for x := x0; x <= x1; x++ {
			*o = append(*o, rl.Vector2{X: x, Y: y})
			if d > 0 {
				y += yi
				d += (2 * (dy - dx))
			} else {
				d += 2 * dy
			}
		}
	}

	plotLineHigh := func(o *[]rl.Vector2, x0 float32, y0 float32, x1 float32, y1 float32) {
		dx := x1 - x0
		dy := y1 - y0
		xi := float32(1)

		if dx < 0 {
			xi = -1
			dy *= -1
		}

		d := (2 * dx) - dy
		x := x0

		for y := y0; y <= y1; y++ {
			*o = append(*o, rl.Vector2{X: x, Y: y})
			if d > 0 {
				x += xi
				d += (2 * (dx - dy))
			} else {
				d += 2 * dx
			}
		}
	}

	x0 := pointA.X
	y0 := pointA.Y
	x1 := pointB.X
	y1 := pointB.Y

	if math.Abs(float64(y1-y0)) < math.Abs(float64(x1-x0)) {
		if x0 > x1 {
			plotLineLow(&outputs, x1, y1, x0, y0)
		} else {
			plotLineLow(&outputs, x0, y0, x1, y1)
		}
	} else {
		if y0 > y1 {
			plotLineHigh(&outputs, x1, y1, x0, y0)
		} else {
			plotLineHigh(&outputs, x0, y0, x1, y1)
		}
	}

	return outputs
}
