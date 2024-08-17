package breakout

import (
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type (
	GameState uint32

	Game struct {
		width    uint32
		height   uint32
		renderer SpriteRenderer
	}
)

const (
	GAME_ACTIVE GameState = 0
	GAME_MENU   GameState = 1
	GAME_WIN    GameState = 2
)

func NewGame(width uint32, height uint32) Game {
	shader, err := LoadShader("assets/shaders/sprite.vert", "assets/shaders/sprite.frag", "sprite")
	if err != nil {
		panic(err)
	}
	shader.Use()

	if _, err := LoadTexture("assets/textures/awesomeface.png", "awesomeface"); err != nil {
		panic(err)
	}
	shader.SetInteger("tex", 0)
	gl.ActiveTexture(gl.TEXTURE0)

	projection := mgl32.Ortho(0, float32(width), float32(height), 0, -1, 1)
	shader.SetMatrix4("projection", projection)

	renderer := NewSpriteRenderer(shader)

	return Game{width, height, renderer}
}

func (g Game) ProcessInput(deltaTime time.Duration) {
}

func (g Game) Update(deltaTime time.Duration) {
}

func (g Game) Render() {
	texture := GetTexture("awesomeface")
	g.renderer.Draw(texture, mgl32.Vec2{100, 100}, mgl32.Vec2{200, 200}, 0)
}

func (g Game) Delete() {
	g.renderer.Delete()
}
