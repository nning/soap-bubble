package main

import (
	"image/color"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Bubbles Bubbles
	Winds   Winds
	Paused  bool
}

var firstPosition *Point
var secondPosition *Point

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) || inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Paused = !g.Paused
	}

	// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	// 	fmt.Println(ebiten.CursorPosition())
	// }

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if firstPosition == nil {
			x, y := ebiten.CursorPosition()
			firstPosition = &Point{float32(x), float32(y)}
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		secondPosition = &Point{float32(x), float32(y)}

		g.AddWind(firstPosition, secondPosition)

		firstPosition = nil
		secondPosition = nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset()
	}

	if g.Paused {
		return nil
	}

	if len(g.Bubbles) < maxBubbles {
		r := rand.Intn(50) + 50

		// x := rand.Intn(pixelWidth-r) + r // TODO ensure bubble is not drawn outside of screen
		// y := rand.Intn(pixelHeight / 2)

		x := pixelWidth / 2
		y := pixelHeight / 4

		g.Bubbles = append(g.Bubbles, NewBubble(x, y, r))
	}

	g.UpdateBurstings()
	g.UpdateBubbleWinds()

	for _, bubble := range g.Bubbles {
		bubble.Update()

		if bubble.Bursting && bubble.StrokeWidth <= 0 {
			g.RemoveBubble(bubble)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{128, 128, 255, 0})

	for _, bubble := range g.Bubbles {
		bubble.Draw(screen)
	}

	for _, wind := range g.Winds {
		wind.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return pixelWidth, pixelHeight
}

func (g *Game) RemoveBubble(bubble *Bubble) {
	var k int

	for i, b := range g.Bubbles {
		if b == bubble {
			k = i
			break
		}
	}

	g.Bubbles = append(g.Bubbles[:k], g.Bubbles[k+1:]...)
}

// TODO calculate for all bubbles and in parallel
func (g *Game) UpdateBurstings() {
	// TODO calculate intersections with edges
	for _, bubble := range g.Bubbles {
		if bubble.LowerEdge() >= pixelHeight {
			bubble.Bursting = true
		}

		if bubble.UpperEdge() <= 0 {
			bubble.Bursting = true
		}

		if bubble.RightEdge() >= pixelWidth {
			bubble.Bursting = true
		}

		if bubble.LeftEdge() <= 0 {
			bubble.Bursting = true
		}
	}

	if len(g.Bubbles) > 1 {
		if isBubbleCollision(g.Bubbles[0], g.Bubbles[1]) {
			g.Bubbles[0].Bursting = true
			g.Bubbles[1].Bursting = true
		}
	}

	if len(g.Bubbles) > 2 {
		if isBubbleCollision(g.Bubbles[1], g.Bubbles[2]) {
			g.Bubbles[1].Bursting = true
			g.Bubbles[2].Bursting = true
		}
	}
}

func (g *Game) UpdateBubbleWinds() {
	for _, bubble := range g.Bubbles {
		for _, wind := range g.Winds {
			isCollision, collision := isWindCollision(bubble, wind)
			if !isCollision || collision == nil {
				continue
			}

			// g.Paused = true

			// wind strength based on where on wind line bubble collided
			d := 1 - dist(collision.X, collision.Y, wind.X, wind.Y)/pixelDiagonal

			bubble.X += wind.VX * wind.Speed * d
			bubble.Y += wind.VY * wind.Speed * d

			// TODO bubble falls right down after wind affection, we should have
			//      a vector for the bubble's movement and add the wind vector
			// 		to it
		}
	}
}

func (g *Game) Reset() {
	g.Bubbles = make(Bubbles, 0)

	g.Winds = make(Winds, 0)
	// g.Winds = append(g.Winds, NewWind(564, 364, -1, -0.8, 1)) // 45° NW
	// g.Winds = append(g.Winds, NewWind(320, 91, 0, 1, 10))     // 90° S

	g.Paused = false
}

func (g *Game) AddWind(a, b *Point) {
	vx := b.X - a.X
	vy := b.Y - a.Y

	// calculate normal vector
	norm := dist(vx, vy, 0, 0)
	vx /= norm
	vy /= norm

	speed := dist(a.X, a.Y, b.X, b.Y) / 100

	g.Winds = append(g.Winds, NewWind(a.X, a.Y, vx, vy, speed))
}
