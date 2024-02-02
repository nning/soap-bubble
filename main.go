package main

import (
	"image/color"
	"log"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const maxBubbles = 3

type Game struct {
	bubbles []*Bubble
}

type Bubble struct {
	x, y  float32
	r     float32
	color color.RGBA
}

func (b *Bubble) Draw(screen *ebiten.Image) {
	vector.StrokeCircle(screen, b.x, b.y, b.r, 0.7, color.Color(b.color), false)
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) || inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if len(g.bubbles) < maxBubbles {
			// rand int between 50 and 100
			r := rand.Intn(50) + 50

			bubble := Bubble{
				x:     float32(x),
				y:     float32(y),
				r:     float32(r),
				color: color.RGBA{255, 255, 255, 255},
			}

			g.bubbles = append(g.bubbles, &bubble)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.bubbles = make([]*Bubble, 0)
	}

	for _, bubble := range g.bubbles {
		bubble.y += 0.5
		// bubble.x = bubble.x + rand.Intn(3) - 1
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{128, 128, 255, 0})

	for _, bubble := range g.bubbles {
		bubble.Draw(screen)
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
