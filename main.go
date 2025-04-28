package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	StarCount    = 100
	BaseSpeed    = 2.0 // Base speed of stars
	MaxDistance  = 5.0 // Higher values = stars appear further away
	MinDistance  = 1.0 // Minimum distance value
)

// Star represents a single star in the sky
type Star struct {
	X        float64 // X position
	Y        float64 // Y position
	Distance float64 // Distance factor (1.0 = closest, MaxDistance = furthest)
	Speed    float64 // The speed at which this star falls
}

// Game holds the game state
type Game struct {
	stars    []Star
	rng      *rand.Rand
	scrolled float64       // Total amount scrolled
	pixelImg *ebiten.Image // Single pixel image for drawing stars
}

func NewGame() *Game {
	g := &Game{
		stars:    make([]Star, StarCount),
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
		pixelImg: ebiten.NewImage(1, 1), // Create a 1x1 pixel image
	}

	// Initialize stars with random positions and distances
	// Ensure even distribution of distances
	distanceStep := (MaxDistance - MinDistance) / float64(StarCount)

	for i := 0; i < StarCount; i++ {
		distance := MinDistance + float64(i)*distanceStep
		// Add a small random variation to distances while keeping order
		distance += g.rng.Float64()*distanceStep - distanceStep/2
		if distance < MinDistance {
			distance = MinDistance
		} else if distance > MaxDistance {
			distance = MaxDistance
		}

		g.stars[i] = Star{
			X:        g.rng.Float64() * ScreenWidth,
			Y:        g.rng.Float64() * ScreenHeight,
			Distance: distance,
			Speed:    BaseSpeed / distance, // Speed inversely proportional to distance
		}
	}

	return g
}

func (g *Game) Update() error {
	// Move each star downward based on its speed
	for i := range g.stars {
		g.stars[i].Y += g.stars[i].Speed

		// If star falls off the bottom of the screen, create a new one at the top
		if g.stars[i].Y > ScreenHeight {
			g.stars[i].Y = 0
			g.stars[i].X = g.rng.Float64() * ScreenWidth
			// Keep the same distance (and thus speed) to ensure even distribution
		}
	}

	g.scrolled += 1.0 // Track total scrolled amount

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Fill the screen with black
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw each star
	for _, s := range g.stars {
		// Calculate color based on distance
		// Close stars (distance = 1.0) are bright white
		// Far stars (distance = MaxDistance) are dim blue
		brightness := uint8(255 * (MaxDistance - s.Distance) / (MaxDistance - MinDistance))
		blue := uint8(200 * s.Distance / MaxDistance)

		// Set the pixel color
		g.pixelImg.Fill(color.RGBA{brightness, brightness, brightness + blue, 255})

		// Draw the pixel at the star's position
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(s.X, s.Y)
		screen.DrawImage(g.pixelImg, op)
	}

	// Display FPS
	debugMsg := fmt.Sprintf(`
TPS:            %0.2f
FPS:            %0.2f
PRESS ESCAPE TO QUIT`,
		ebiten.ActualTPS(), ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, debugMsg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Infinite Scrolling Star Field")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
