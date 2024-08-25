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

func NewBall(position mgl32.Vec2, radius float32, velocity mgl32.Vec2, sprite Texture) Ball {
	entity := Entity{
		Position:  position,
		Size:      mgl32.Vec2{2 * radius, 2 * radius},
		Velocity:  velocity,
		Color:     mgl32.Vec3{1, 1, 1},
		Angle:     0,
		IsSolid:   false,
		Destroyed: false,
		Sprite:    sprite,
	}
	stuck := true

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

	b.stuck = true
}

func (b *Ball) Collides(other *Entity) Collision {
	center := b.entity.Position.Add(mgl32.Vec2{b.radius, b.radius})
	halfExtents := other.Size.Mul(0.5)
	otherCenter := other.Position.Add(halfExtents)

	diff := center.Sub(otherCenter)
	clamped := mgl32.Vec2{
		mgl32.Clamp(diff.X(), -halfExtents.X(), halfExtents.X()),
		mgl32.Clamp(diff.Y(), -halfExtents.Y(), halfExtents.Y()),
	}

	closestPoint := otherCenter.Add(clamped)
	diff = closestPoint.Sub(center)

	if diff.Len() <= b.radius {
		return NewCollision(true, DirectionOf(diff), diff)
	} else {
		return NewCollision(false, UP, mgl32.Vec2{0, 0})
	}
}
