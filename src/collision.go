package breakout

import "github.com/go-gl/mathgl/mgl32"

type Direction = int32

const (
	UP    Direction = 0
	RIGHT Direction = 1
	DOWN  Direction = 2
	LEFT  Direction = 3
)

func DirectionOf(target mgl32.Vec2) Direction {
	compass := []mgl32.Vec2{
		{0, 1},  // UP
		{1, 0},  // RIGHT
		{0, -1}, // DOWN
		{-1, 0}, // LEFT
	}
	maxVal := float32(0)
	bestMatch := UP
	for i, d := range compass {
		product := target.Normalize().Dot(d)
		if product > maxVal {
			maxVal = product
			bestMatch = Direction(i)
		}
	}
	return bestMatch
}

type Collision struct {
	Collided   bool
	Direction  Direction
	Difference mgl32.Vec2
}

func NewCollision(collided bool, direction Direction, difference mgl32.Vec2) Collision {
	return Collision{
		Collided:   collided,
		Direction:  direction,
		Difference: difference,
	}
}
