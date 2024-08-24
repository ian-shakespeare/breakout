package breakout

import "github.com/go-gl/mathgl/mgl32"

type Entity struct {
	Position  mgl32.Vec2
	Size      mgl32.Vec2
	Velocity  mgl32.Vec2
	Color     mgl32.Vec3
	Angle     float32
	IsSolid   bool
	Destroyed bool
	Sprite    Texture
}

func (e *Entity) Draw(renderer *SpriteRenderer) {
	renderer.Draw(e.Sprite, e.Position, e.Size, e.Angle, e.Color)
}

func (e *Entity) SetX(x float32) {
	e.Position = mgl32.Vec2{x, e.Position.Y()}
}

func (e *Entity) SetY(y float32) {
	e.Position = mgl32.Vec2{e.Position.X(), y}
}

func (e *Entity) InvertXVelocity() {
	x := -1 * e.Velocity.X()
	e.Velocity = mgl32.Vec2{x, e.Velocity.Y()}
}

func (e *Entity) InvertYVelocity() {
	y := -1 * e.Velocity.Y()
	e.Velocity = mgl32.Vec2{e.Velocity.X(), y}
}

func (e *Entity) Collides(other *Entity) bool {
	collisionX := e.Position.X()+e.Size.X() >= other.Position.X() && other.Position.X()+other.Size.X() >= e.Position.X()
	collisionY := e.Position.Y()+e.Size.Y() >= other.Position.Y() && other.Position.Y()+other.Size.Y() >= e.Position.Y()

	return collisionX && collisionY
}
