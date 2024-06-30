package particles

import (
	"math"
	"slices"
	"strings"
	"time"
)

type Particle struct {
	lifetime int64
	speed    float64

	x float64
	y float64
}

type Reset func(particle *Particle, params *ParticleParams)
type NextPosition func(particle *Particle, delta int64)
type Ascii func(x, y int, count [][]int) string

type ParticleParams struct {
	MaxLife  int64
	MaxSpeed float64

	X int
	Y int

	Count int

	nextPosition NextPosition
	ascii        Ascii
	reset        Reset
}

type ParticleSystem struct {
	ParticleParams

	particles []*Particle
	// expireTime int64
	lastTime int64
}

func NewParticleSystem(params ParticleParams) ParticleSystem {
	particles := make([]*Particle, 0)
	for i := 0; i < params.Count; i++ {
		particles = append(particles, &Particle{})
	}
	return ParticleSystem{
		ParticleParams: params,
		lastTime:       time.Now().UnixMilli(),
		particles:      particles,
	}
}

func (p *ParticleSystem) Start() {
	for _, part := range p.particles {
		p.reset(part, &p.ParticleParams)
	}
}

func (p *ParticleSystem) Update() {
	now := time.Now().UnixMilli()
	delta := now - p.lastTime
	p.lastTime = now

	for _, part := range p.particles {
		p.nextPosition(part, delta)
		if part.y >= float64(p.Y) || part.x >= float64(p.X) {
			p.reset(part, &p.ParticleParams)
		}
	}
}

func (p *ParticleSystem) Display() string {
	count := make([][]int, 0)

	for row := 0; row < p.Y; row++ {
		c := make([]int, 0)
		for col := 0; col < p.X; col++ {
			c = append(c, 0)
		}
		count = append(count, c)
	}
	for _, part := range p.particles {
		row := int(math.Floor(part.y))
		col := int(math.Floor(part.x))

		count[row][col]++
	}
	out := make([][]string, 0)
	for r, row := range count {
		outRow := make([]string, 0)
		for c := range row {
			outRow = append(outRow, p.ascii(r, c, count))
		}
		out = append(out, outRow)
	}

	slices.Reverse(out)
	outStr := make([]string, 0)
	for _, row := range out {
		outStr = append(outStr, strings.Join(row, ""))
	}

	return strings.Join(outStr, "\n")
}
