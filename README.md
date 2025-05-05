# go-starfield

[![Go Reference](https://pkg.go.dev/badge/github.com/jtbonhomme/go-startfield)](https://pkg.go.dev/github.com/jtbonhomme/go-startfield)
[![issues](https://img.shields.io/github/issues/jtbonhomme/go-startfield)](https://github.com/jtbonhomme/go-startfield/issues)
![GitHub Release](https://img.shields.io/github/v/release/jtbonhomme/go-startfield)
[![license](https://img.shields.io/github/license/jtbonhomme/go-startfield)](https://github.com/jtbonhomme/go-startfield/blob/main/LICENSE)

`go-startfield` is a Go library that provides an Entity-Component-System (ECS) framework tailored for use with the [Ebiten](https://ebiten.org/) game library. It simplifies the development of complex 2D games by organizing game logic into entities, components, and systems.

![](go-starfield.gif)

This repository provides a simple package to reproduce an infinite scrolling field of stars.

## Demo

Visit: https://jtbonhomme.github.io/go-startfield/

## Usage

```go
import (
    stars "github.com/jtbonhomme/go-startfield"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	BaseSpeed    = -2.0 // Base speed of stars
	MaxDistance  = 5.0  // Higher values = stars appear further away
	MinDistance  = 1.0  // Minimum distance value
)

type Game struct {
	starField *stars.StarField
}

func NewGame() *Game {
	return &Game{
		starField: stars.New(ScreenWidth, ScreenHeight, 100, BaseSpeed, MaxDistance, MinDistance),
	}
}

func (g *Game) Update() error {
	// Update the star field
	g.starField.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Fill the screen with black
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw the star field
	g.starField.Draw(screen)
}
```

Check the provided [example](example) for a functional test program.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the library.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
