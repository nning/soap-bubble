package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const pixelHeight = 540
const pixelWidth = 960

var maxBubbles = 1

func main() {
	println("soap bubble started")

	ebiten.SetWindowTitle("soap bubble")
	ebiten.SetFullscreen(false)
	ebiten.SetWindowSize(1920, 1080)

	g := &Game{}
	g.winds = make(Winds, 0)
	g.winds = append(g.winds, &Wind{x: 564, y: 364, angle: 130, speed: 50})

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
