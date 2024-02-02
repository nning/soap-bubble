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
const pixelHeight = 540
const pixelWidth = 960

type Game struct {
	bubbles Bubbles
}

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
	if b.LowerXBounds() >= pixelHeight {
		b.bursting = true
		if b.strokeWidth > 0 {
			b.strokeWidth -= 0.125
		}
	}

	b.y += b.speed
}

func (b *Bubble) Draw(screen *ebiten.Image) {
	vector.StrokeCircle(screen, b.x, b.y, b.r, b.strokeWidth, color.Color(b.color), false)
}

func (b *Bubble) LowerXBounds() float32 {
	return b.y + b.r + b.strokeWidth
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
			r := rand.Intn(50) + 50
			g.bubbles = append(g.bubbles, NewBubble(x, y, r))
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.bubbles = make(Bubbles, 0)
	}

	for _, bubble := range g.bubbles {
		bubble.Update()

		if bubble.bursting && bubble.strokeWidth <= 0 {
			g.RemoveBubble(bubble)
		}
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
	return pixelWidth, pixelHeight
}

func (g *Game) RemoveBubble(bubble *Bubble) {
	var k int

	for i, b := range g.bubbles {
		if b == bubble {
			k = i
			break
		}
	}

	g.bubbles = append(g.bubbles[:k], g.bubbles[k+1:]...)
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
