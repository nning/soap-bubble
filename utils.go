package main

import "math"

func dist(x1, y1, x2, y2 float32) float32 {
	dx := x2 - x1
	dy := y2 - y1

	return float32(math.Sqrt(float64((dx * dx) + (dy * dy))))
}

func isPointInBubble(b *Bubble, x, y float32) bool {
	return dist(b.x, b.y, x, y) <= b.r
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
	return dist(a.x, a.y, b.x, b.y) <= a.r+b.r
}

// https://www.jeffreythompson.org/collision-detection/line-circle.php
func isWindCollision(b *Bubble, w *Wind) bool {
	if isPointInBubble(b, w.x, w.y) || isPointInBubble(b, w.edgeX, w.edgeY) {
		return true
	}

	len := dist(w.x, w.y, w.edgeX, w.edgeY)
	dot := (((b.x - w.x) * (w.edgeX - w.x)) + ((b.y - w.y) * (w.edgeY - w.y))) / (len * len)

	closestX := w.x + (dot * (w.edgeX - w.x))
	closestY := w.y + (dot * (w.edgeY - w.y))

	if !isPointOnLine(w.x, w.y, w.edgeX, w.edgeY, closestX, closestY) {
		return false
	}

	d := dist(closestX, closestY, b.x, b.y)

	return d <= b.r
}
