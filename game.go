package main

import (
	"fmt"
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

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		fmt.Println(ebiten.CursorPosition())
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.Bubbles = make(Bubbles, 0)
	}

	if g.Paused {
		return nil
	}

	// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	// 	x, y := ebiten.CursorPosition()

	// 	if len(g.Bubbles) < maxBubbles {
	// 		r := rand.Intn(50) + 50
	// 		g.Bubbles = append(g.Bubbles, NewBubble(x, y, r))
	// 	}
	// }

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
	for _, bubble := range g.Bubbles {
		if bubble.LowerXBounds() >= pixelHeight {
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
