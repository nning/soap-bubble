package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Bubbles []*Bubble

type Bubble struct {
	Point

	R           float32
	Speed       float32
	Color       color.RGBA
	StrokeWidth float32
	Bursting    bool
}

func NewBubble(x, y, r int) *Bubble {
	return &Bubble{
		Point:       Point{float32(x), float32(y)},
		R:           float32(r),
		Speed:       float32(150-r) / 100,
		StrokeWidth: 0.7,
		Color:       color.RGBA{255, 255, 255, 255},
	}
}

func (b *Bubble) Update() {
	if b.Bursting && b.StrokeWidth > 0 {
		b.StrokeWidth -= 0.2
		b.R += 10
	}

	b.Y += b.Speed
}

func (b *Bubble) Draw(screen *ebiten.Image) {
	vector.StrokeCircle(screen, b.X, b.Y, b.R, b.StrokeWidth, color.Color(b.Color), false)
}

func (b *Bubble) LowerEdge() float32 {
	return b.Y + b.R + b.StrokeWidth
}

func (b *Bubble) UpperEdge() float32 {
	return b.Y - b.R - b.StrokeWidth
}

func (b *Bubble) RightEdge() float32 {
	return b.X + b.R + b.StrokeWidth
}

func (b *Bubble) LeftEdge() float32 {
	return b.X - b.R - b.StrokeWidth
}
