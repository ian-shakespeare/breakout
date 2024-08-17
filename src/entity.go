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

func (e Entity) Draw(renderer *SpriteRenderer) {
	renderer.Draw(e.Sprite, e.Position, e.Size, e.Angle, e.Color)
}
