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

	firstPosition  *Point
	secondPosition *Point

	cursor *Cursor
	config *Config
}

func NewGame(config *Config) *Game {
	return &Game{
		Bubbles: make(Bubbles, 0),
		Winds:   make(Winds, 0),
		Paused:  false,
		cursor:  &Cursor{r: 5},
		config:  config,
	}
}

func (g *Game) Update() error {
	if g.config.ShowCursor {
		g.cursor.x, g.cursor.y = ebiten.CursorPosition()
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) || inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Paused = !g.Paused
	}

	if g.config.Debug {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonMiddle) {
			fmt.Println(ebiten.CursorPosition())
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if g.firstPosition == nil {
			x, y := ebiten.CursorPosition()
			g.firstPosition = &Point{float32(x), float32(y)}
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.secondPosition = &Point{float32(x), float32(y)}

		g.AddWind(g.firstPosition, g.secondPosition)

		g.firstPosition = nil
		g.secondPosition = nil
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		x, y := ebiten.CursorPosition()

		for _, bubble := range g.Bubbles {
			if isPointInBubble(bubble, float32(x), float32(y)) {
				bubble.Bursting = true
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset()
	}

	if g.Paused {
		return nil
	}

	g.CreateBubble()

	g.UpdateBurstings()
	g.UpdateBubbleWinds()

	g.UpdateBubbles()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{127, 127, 255, 0})

	for _, wind := range g.Winds {
		wind.Draw(screen)
	}

	for _, bubble := range g.Bubbles {
		bubble.Draw(screen)
	}

	if g.config.ShowCursor {
		g.cursor.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.config.PixelWidth, g.config.PixelHeight
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

func (g *Game) UpdateBurstings() {
	// TODO calculate intersections with edges?
	if g.config.EdgeCollision {
		for _, bubble := range g.Bubbles {
			if bubble.LowerEdge() >= float32(g.config.PixelHeight) {
				bubble.Bursting = true
			}

			if bubble.UpperEdge() <= 0 {
				bubble.Bursting = true
			}

			if bubble.RightEdge() >= float32(g.config.PixelWidth) {
				bubble.Bursting = true
			}

			if bubble.LeftEdge() <= 0 {
				bubble.Bursting = true
			}
		}
	}

	if g.config.BubbleCollision {
		pairs := make([][]*Bubble, 0)
		for i, bubble := range g.Bubbles {
			for j := i + 1; j < len(g.Bubbles); j++ {
				pairs = append(pairs, []*Bubble{bubble, g.Bubbles[j]})
			}
		}

		for _, pair := range pairs {
			if isBubbleCollision(pair[0], pair[1]) {
				pair[0].Bursting = true
				pair[1].Bursting = true
			}
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

			// wind strength based on where on wind line bubble collided
			d := 1 - dist(collision.X, collision.Y, wind.X, wind.Y)/pixelDiagonal

			bubble.VX += wind.VX * wind.Speed * d
			bubble.VY += wind.VY * wind.Speed * d

			// TODO bubble falls right down after wind affection, we should have
			//      a vector for the bubble's movement and add the wind vector
			// 		to it
		}
	}
}

func (g *Game) CreateBubble() {
	if len(g.Bubbles) >= g.config.MaxBubbles {
		return
	}

	r := rand.Intn(50) + 50

	x := rand.Intn(g.config.PixelWidth-r) + r // TODO ensure bubble is not drawn outside of screen
	y := rand.Intn(g.config.PixelHeight / 2)

	// x := pixelWidth / 2
	// y := pixelHeight / 4

	g.Bubbles = append(g.Bubbles, NewBubble(x, y, r))
}

func (g *Game) Reset() {
	g.Bubbles = make(Bubbles, 0)
	g.Winds = make(Winds, 0)

	g.Paused = false
}

func (g *Game) AddWind(a, b *Point) {
	vx := b.X - a.X
	vy := b.Y - a.Y

	// calculate normal vector
	norm := dist(vx, vy, 0, 0)
	vx /= norm
	vy /= norm

	speed := dist(a.X, a.Y, b.X, b.Y) * 0.00075

	g.Winds = append(g.Winds, NewWind(a.X, a.Y, vx, vy, speed))
}

func (g *Game) UpdateBubbles() {
	for _, bubble := range g.Bubbles {
		bubble.Update()

		if bubble.Bursting && bubble.StrokeWidth <= 0 {
			g.RemoveBubble(bubble)
		}

		bubble.X += bubble.VX
		bubble.Y += bubble.VY

		bubble.VX *= 0.9
		bubble.VY *= 0.9
	}
}
