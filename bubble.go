package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Bubbles []*Bubble

type Bubble struct {
	x, y        float32
	r           float32
	speed       float32
	color       color.RGBA
	strokeWidth float32
	bursting    bool
}

func NewBubble(x, y, r int) *Bubble {
	return &Bubble{
		x:           float32(x),
		y:           float32(y),
		r:           float32(r),
		speed:       float32(150-r) / 100,
		strokeWidth: 0.7,
		color:       color.RGBA{255, 255, 255, 255},
	}
}

func (b *Bubble) Update() {
	if b.bursting && b.strokeWidth > 0 {
		b.strokeWidth -= 0.2
		b.r += 10
	}

	b.y += b.speed
}

func (b *Bubble) Draw(screen *ebiten.Image) {
	vector.StrokeCircle(screen, b.x, b.y, b.r, b.strokeWidth, color.Color(b.color), false)
}

func (b *Bubble) LowerXBounds() float32 {
	return b.y + b.r + b.strokeWidth
}
