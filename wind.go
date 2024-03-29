package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Winds []*Wind

type Wind struct {
	Point
	Vector

	Speed        float32
	EdgeX, EdgeY float32
}

func NewWind(x, y, vx, vy, speed float32) *Wind {
	edgeX := x + (pixelDiagonal * vx)
	edgeY := y + (pixelDiagonal * vy)

	return &Wind{Point{x, y}, Vector{vx, vy}, speed, edgeX, edgeY}
}

func (w *Wind) Update() {
}

func (w *Wind) Draw(screen *ebiten.Image) {
	vector.StrokeLine(
		screen,
		w.X,
		w.Y,
		w.EdgeX,
		w.EdgeY,
		float32(1),
		color.Color(color.RGBA{143, 143, 255, 255}),
		false,
	)
}
