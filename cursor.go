package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Cursor struct {
	x, y int
	r    float32
}

func (c *Cursor) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(c.x), float32(c.y), c.r, color.RGBA{127, 127, 127, 127}, false)
}
