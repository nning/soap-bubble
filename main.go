package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const pixelHeight = 540
const pixelWidth = 960

var maxBubbles = 1
var pixelDiagonal = float32(math.Sqrt(pixelHeight*pixelHeight + pixelWidth*pixelWidth))

func main() {
	println("soap bubble started")

	ebiten.SetWindowTitle("soap bubble")
	ebiten.SetFullscreen(false)
	ebiten.SetWindowSize(1920, 1080)

	g := &Game{}
	if err := g.Load(); err != nil {
		g.Reset()
	}

	go g.SavePeriodically()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
