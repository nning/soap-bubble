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
	g.Winds = make(Winds, 0)
	g.Winds = append(g.Winds, NewWind(564, 364, -1, -0.8, 1)) // 45° NW

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
