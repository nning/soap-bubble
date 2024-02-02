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
	bubbles Bubbles
	winds   Winds
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) || inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		fmt.Println(ebiten.CursorPosition())
	}

	// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	// 	x, y := ebiten.CursorPosition()

	// 	if len(g.bubbles) < maxBubbles {
	// 		r := rand.Intn(50) + 50
	// 		g.bubbles = append(g.bubbles, NewBubble(x, y, r))
	// 	}
	// }

	if len(g.bubbles) < maxBubbles {
		r := rand.Intn(50) + 50

		// x := rand.Intn(pixelWidth-r) + r // TODO ensure bubble is not drawn outside of screen
		// y := rand.Intn(pixelHeight / 2)

		x := pixelWidth / 2
		y := pixelHeight / 4

		g.bubbles = append(g.bubbles, NewBubble(x, y, r))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.bubbles = make(Bubbles, 0)
	}

	g.UpdateBurstings()

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

	for _, wind := range g.winds {
		wind.Draw(screen)
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

func (g *Game) UpdateBurstings() {
	for _, bubble := range g.bubbles {
		if bubble.LowerXBounds() >= pixelHeight {
			bubble.bursting = true
		}
	}

	if len(g.bubbles) > 1 {
		if isBubbleCollision(g.bubbles[0], g.bubbles[1]) {
			g.bubbles[0].bursting = true
			g.bubbles[1].bursting = true
		}
	}

	if len(g.bubbles) > 2 {
		if isBubbleCollision(g.bubbles[1], g.bubbles[2]) {
			g.bubbles[1].bursting = true
			g.bubbles[2].bursting = true
		}
	}
}
