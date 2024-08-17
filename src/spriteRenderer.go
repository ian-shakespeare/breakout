package breakout

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type SpriteRenderer struct {
	shader Shader
	Vao    uint32
}

func NewSpriteRenderer(shader Shader) SpriteRenderer {
	vertices := []float32{
		0.5, 0.5, 1.0, 1.0,
		0.5, -0.5, 1.0, 0.0,
		-0.5, -0.5, 0.0, 0.0,
		-0.5, 0.5, 0.0, 1.0,
	}

	indices := []uint32{
		0, 1, 3,
		1, 2, 3,
	}

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)

	var ebo uint32
	gl.GenBuffers(1, &ebo)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	vertAttrib := uint32(gl.GetAttribLocation(shader.Id, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointerWithOffset(vertAttrib, 2, gl.FLOAT, false, 4*4, 0)

	texCoordAttrib := uint32(gl.GetAttribLocation(shader.Id, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointerWithOffset(texCoordAttrib, 2, gl.FLOAT, false, 4*4, 2*4)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return SpriteRenderer{shader: shader, Vao: vao}
}

func (s SpriteRenderer) Draw(texture Texture, position mgl32.Vec2, size mgl32.Vec2, angle float32, color mgl32.Vec3) {
	s.shader.Use()

	translation := mgl32.Translate3D(position.X(), position.Y(), 0)
	rotation := mgl32.HomogRotate3D(angle, mgl32.Vec3{0, 0, 1})
	scale := mgl32.Scale3D(size.X(), size.Y(), 1)

	model := mgl32.Ident4().Mul4(translation).Mul4(rotation).Mul4(scale)
	s.shader.SetMatrix4("model", model)

	s.shader.SetVector3f("spriteColor", color)

	texture.Bind()
	gl.BindVertexArray(s.Vao)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
}

func (s SpriteRenderer) Delete() {
	s.shader.Delete()
}
