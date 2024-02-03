package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var pixelDiagonal = float32(0)

func main() {
	println("soap bubble started")

	config := NewConfig().Load()

	ebiten.SetWindowTitle("soap bubble")
	ebiten.SetFullscreen(config.Fullscreen)
	ebiten.SetWindowSize(config.WindowWidth, config.WindowHeight)

	h := config.PixelHeight
	w := config.PixelWidth
	pixelDiagonal = float32(math.Sqrt(float64(h*h + w*w)))

	g := NewGame(config)
	if err := g.Load(); err != nil {
		g.Reset()
	}

	go g.SavePeriodically()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
