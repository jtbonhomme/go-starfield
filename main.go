package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/jtbonhomme/go-startfield/stars"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	BaseSpeed    = 2.0 // Base speed of stars
	MaxDistance  = 5.0 // Higher values = stars appear further away
	MinDistance  = 1.0 // Minimum distance value
)

type Game struct {
	starField *stars.StarField
}

func NewGame() *Game {
	return &Game{
		starField: stars.NewStarField(ScreenWidth, ScreenHeight, 100, BaseSpeed, MaxDistance, MinDistance),
	}
}

func (g *Game) Update() error {
	// Update the star field
	g.starField.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Fill the screen with black
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw the star field
	g.starField.Draw(screen)

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
