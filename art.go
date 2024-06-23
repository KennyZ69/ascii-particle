package particles

import (
	"math"
	"math/rand"
)

// Just saying "Art" because I dont know yet what will I want to work with
// Ideas are for now: kotlikovy; tennis racket with flames on it.

type Art struct {
	ParticleSystem
}

func ascii(row, col int, count [][]int) rune {
	return '}'
}

func reset(p *Particle, params *ParticleParams) {
	p.lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
	p.speed = math.Floor(params.MaxSpeed * rand.Float64())

	maxX := math.Floor(float64(params.X) / 2)
	x := math.Max(-maxX, math.Min(rand.NormFloat64(), maxX))
	p.x = x + maxX
	p.y = 0
}

func nextPos(p *Particle, delta int64) {
	p.lifetime -= delta
	if p.lifetime <= 0 {
		return
	}
	p.y += (float64(delta) / 1000.0) * p.speed
}

func NewArt(width, height int) Art {
	// // ! width of particle system must be odd, for the 0 to be centered, even force it to be like that
	// if width%2 == 0 {
	// 	fmt.Println("the width must be odd")
	// 	return Art{}
	// }
	Assert("width of particle system must be odd", width%2 == 1)
	return Art{
		ParticleSystem: NewParticleSystem(ParticleParams{
			MaxLife:      5,
			MaxSpeed:     1,
			reset:        reset,
			ascii:        ascii,
			nextPosition: nextPos,
			Count:        100,
		}),
	}
}
