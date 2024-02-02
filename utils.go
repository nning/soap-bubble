package main

import "math"

func isBubbleCollision(a, b *Bubble) bool {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	sr := float64(a.r + b.r)
	d := math.Sqrt((dx * dx) + (dy * dy))

	return d <= sr
}
