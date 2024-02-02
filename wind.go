package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Winds []*Wind

type Wind struct {
	x, y  float32
	angle float32
	speed float32
}

func (w *Wind) Update() {
	// w.x += w.speed * cos(w.angle)
	// w.y += w.speed * sin(w.angle)
}

func (w *Wind) Draw(screen *ebiten.Image) {
	vector.StrokeLine(
		screen,
		w.x,
		w.y,
		w.x+w.speed*float32(math.Cos(float64(w.angle))),
		w.y+w.speed*float32(math.Sin(float64(w.angle))),
		float32(1),
		color.Color(color.RGBA{255, 0, 0, 255}),
		false,
	)
}
