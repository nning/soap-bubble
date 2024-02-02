package main

import (
	"math"
)

type Point struct {
	X, Y float32
}

func dist(x1, y1, x2, y2 float32) float32 {
	dx := x2 - x1
	dy := y2 - y1

	return float32(math.Sqrt(float64((dx * dx) + (dy * dy))))
}

func isPointInBubble(b *Bubble, x, y float32) bool {
	return dist(b.X, b.Y, x, y) <= b.R
}

func isPointOnLine(x1, y1, x2, y2, x, y float32) bool {
	len := dist(x1, y1, x2, y2)

	d1 := dist(x, y, x1, y1)
	d2 := dist(x, y, x2, y2)

	const epsilon = 0.1
	return d1+d2 >= len-epsilon && d1+d2 <= len+epsilon
}

// https://www.jeffreythompson.org/collision-detection/circle-circle.php
func isBubbleCollision(a, b *Bubble) bool {
	return dist(a.X, a.Y, b.X, b.Y) <= a.R+b.R
}

// https://www.jeffreythompson.org/collision-detection/line-circle.php
func isWindCollision(b *Bubble, w *Wind) (bool, *Point) {
	if isPointInBubble(b, w.X, w.Y) || isPointInBubble(b, w.EdgeX, w.EdgeY) {
		return true, nil
	}

	len := dist(w.X, w.Y, w.EdgeX, w.EdgeY)
	dot := (((b.X - w.X) * (w.EdgeX - w.X)) + ((b.Y - w.Y) * (w.EdgeY - w.Y))) / (len * len)

	closestX := w.X + (dot * (w.EdgeX - w.X))
	closestY := w.Y + (dot * (w.EdgeY - w.Y))

	if !isPointOnLine(w.X, w.Y, w.EdgeX, w.EdgeY, closestX, closestY) {
		return false, nil
	}

	d := dist(closestX, closestY, b.X, b.Y)

	if d <= b.R {
		return true, &Point{closestX, closestY}
	}

	return false, nil
}
