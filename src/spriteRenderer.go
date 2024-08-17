package breakout

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type SpriteRenderer struct {
	shader Shader
	vao    uint32
}

func NewSpriteRenderer(shader Shader) SpriteRenderer {
	vertices := []float32{
		0.0, 1.0, 0.0, 1.0,
		1.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 0.0,

		0.0, 1.0, 0.0, 1.0,
		1.0, 1.0, 1.0, 1.0,
		1.0, 0.0, 1.0, 0.0,
	}

	var vao uint32
	var vbo uint32

	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 4*4, nil)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return SpriteRenderer{shader, vao}
}

func (s SpriteRenderer) DrawSprite(texture Texture, position mgl32.Vec2, size mgl32.Vec2, rotation float32, color mgl32.Vec3) {
	s.shader.Use()
	translation := mgl32.Translate2D(position.X(), position.Y())

	prerotation := mgl32.Translate2D(0.5*size.X(), 0.5*size.Y())
	rotaton := mgl32.HomogRotate2D(rotation)
	postrotation := mgl32.Translate2D(-0.5*size.X(), -0.5*size.Y())

	scale := mgl32.Scale2D(size.X(), size.Y())
	model := mgl32.Ident3().Mul3(translation).Mul3(prerotation).Mul3(rotaton).Mul3(postrotation).Mul3(scale)

	s.shader.SetMatrix4("model", model.Mat4())
	s.shader.SetVector3f("spriteColor", color)

	gl.BindVertexArray(s.vao)
	gl.ActiveTexture(gl.TEXTURE0)
	texture.Bind()

	gl.DrawArrays(gl.TRIANGLES, 0, 6)

	gl.BindVertexArray(0)
}
