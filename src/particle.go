package breakout

import "github.com/go-gl/mathgl/mgl32"

type Particle struct {
	Position mgl32.Vec2
	Velocity mgl32.Vec2
	Color    mgl32.Vec4
	Life     float32
}

func NewParticle() Particle {
	return Particle{
		Position: mgl32.Vec2{0, 0},
		Velocity: mgl32.Vec2{0, 0},
		Color:    mgl32.Vec4{0, 0, 0, 1},
		Life:     0,
	}
}
