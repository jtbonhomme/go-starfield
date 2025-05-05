package stars

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Star represents a single star in the sky
type Star struct {
	X        float64 // X position
	Y        float64 // Y position
	Distance float64 // Distance factor (1.0 = closest, MaxDistance = furthest)
	Speed    float64 // The speed at which this star falls
}

// StarField represents a collection of stars
type StarField struct {
	stars       []Star
	rng         *rand.Rand
	width       int
	height      int
	pixelImg    *ebiten.Image // Single pixel image for drawing stars
	BaseSpeed   float64       // Base speed of stars. A positive value makes stars fall down, negative makes them go up
	MaxDistance float64       // Higher values = stars appear further away
	MinDistance float64       // Minimum distance value
}

// New creates a new StarField with the given screen dimensions, star count, and speed/distance parameters
func New(width, height, starCount int, baseSpeed, maxDistance, minDistance float64) *StarField {
	sf := &StarField{
		stars:       make([]Star, starCount),
		rng:         rand.New(rand.NewSource(time.Now().UnixNano())),
		width:       width,
		height:      height,
		pixelImg:    ebiten.NewImage(1, 1), // Create a 1x1 pixel image
		BaseSpeed:   baseSpeed,
		MaxDistance: maxDistance,
		MinDistance: minDistance,
	}

	// Initialize stars with random positions and distances
	distanceStep := (maxDistance - minDistance) / float64(starCount)

	for i := 0; i < starCount; i++ {
		distance := minDistance + float64(i)*distanceStep
		distance += sf.rng.Float64()*distanceStep - distanceStep/2
		if distance < minDistance {
			distance = minDistance
		} else if distance > maxDistance {
			distance = maxDistance
		}

		sf.stars[i] = Star{
			X:        sf.rng.Float64() * float64(width),
			Y:        sf.rng.Float64() * float64(height),
			Distance: distance,
			Speed:    baseSpeed / distance, // Speed inversely proportional to distance
		}
	}

	return sf
}

// Update updates the positions of the stars
func (sf *StarField) Update() {
	for i := range sf.stars {
		sf.stars[i].Y += sf.stars[i].Speed

		// If star falls off the bottom of the screen, reset it to the top
		if sf.stars[i].Y > float64(sf.height) {
			sf.stars[i].Y = 0
			sf.stars[i].X = sf.rng.Float64() * float64(sf.width)
		}

		// If star reach the top of the screen, reset it to the bottom
		if sf.stars[i].Y < 0 {
			sf.stars[i].Y = float64(sf.height)
			sf.stars[i].X = sf.rng.Float64() * float64(sf.width)
		}
	}
}

// Draw draws the stars onto the given screen
func (sf *StarField) Draw(screen *ebiten.Image) {
	for _, s := range sf.stars {
		// Calculate color based on distance
		brightness := uint8(255 * (sf.MaxDistance - s.Distance) / (sf.MaxDistance - sf.MinDistance))
		blue := uint8(200 * s.Distance / sf.MaxDistance)

		// Set the pixel color
		sf.pixelImg.Fill(color.RGBA{brightness, brightness, brightness + blue, 255})

		// Draw the pixel at the star's position
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(s.X, s.Y)
		screen.DrawImage(sf.pixelImg, op)
	}
}
