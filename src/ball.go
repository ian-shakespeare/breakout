package breakout

import (
	"time"

	"github.com/go-gl/mathgl/mgl32"
)

type Ball struct {
	entity Entity

	radius float32
	stuck  bool
}

func NewBall(entity Entity, radius float32, stuck bool) Ball {
	return Ball{
		entity,
		radius,
		stuck,
	}
}

func (b *Ball) Move(deltaTime time.Duration, windowWidth uint32) mgl32.Vec2 {
	if !b.stuck {
		b.entity.Position = b.entity.Position.Add(b.entity.Velocity.Mul(float32(deltaTime.Milliseconds())))

		if b.entity.Position.X() <= 0 {
			b.entity.InvertXVelocity()
			b.entity.SetX(0)
		} else if b.entity.Position.X()+b.entity.Size.X() >= float32(windowWidth) {
			b.entity.InvertXVelocity()
			b.entity.SetX(float32(windowWidth) - b.entity.Size.X())
		}

		if b.entity.Position.Y() <= 0 {
			b.entity.InvertYVelocity()
			b.entity.SetY(0)
		}
	}

	return b.entity.Position
}

func (b *Ball) Reset(position mgl32.Vec2, velocity mgl32.Vec2) {
	b.entity.Position = position
	b.entity.Velocity = velocity
}

func (b *Ball) Collides(other *Entity) bool {
	center := b.entity.Position.Add(mgl32.Vec2{b.radius, b.radius})
	halfExtents := mgl32.Vec2{other.Size.X(), other.Size.Y()}
	otherCenter := other.Position

	diff := center.Sub(otherCenter)
	clamped := mgl32.Vec2{mgl32.Clamp(diff.X(), -halfExtents.X(), halfExtents.X()), mgl32.Clamp(diff.Y(), -halfExtents.Y(), halfExtents.Y())}

	closestPoint := otherCenter.Add(clamped)
	diff = closestPoint.Sub(center)

	return diff.Len() < b.radius
}
