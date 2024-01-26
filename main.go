package main

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	pixels []Pixel
}

type Pixel struct {
	x, y  int
	color color.RGBA
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		g.pixels = append(g.pixels, Pixel{
			x:     x,
			y:     y,
			color: color.RGBA{255, 255, 255, 255},
		})

		// fmt.Println(g.pixels)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{128, 128, 255, 0})

	for _, pixel := range g.pixels {
		screen.Set(pixel.x, pixel.y, pixel.color)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 960, 540
}

func main() {
	println("soap bubble started")

	ebiten.SetWindowTitle("soap bubble")
	ebiten.SetFullscreen(false)

	ebiten.SetWindowSize(1920, 1080)

	g := &Game{}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
