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
		level    Level
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

	if _, err := LoadTexture("assets/textures/block.png", "block"); err != nil {
		panic(err)
	}
	if _, err := LoadTexture("assets/textures/gray.png", "background"); err != nil {
		panic(err)
	}

	shader.SetInteger("tex", 0)
	gl.ActiveTexture(gl.TEXTURE0)

	projection := mgl32.Ortho(0, float32(width), float32(height), 0, -1, 1)
	shader.SetMatrix4("projection", projection)

	renderer := NewSpriteRenderer(shader)

	level, err := LoadLevel("assets/levels/default.lvl", width, height/2)
	if err != nil {
		panic(err)
	}

	return Game{width, height, renderer, level}
}

func (g Game) ProcessInput(deltaTime time.Duration) {
}

func (g Game) Update(deltaTime time.Duration) {
}

func (g Game) Render() {
	background := GetTexture("background")
	screenWidth := float32(g.width)
	screenHeight := float32(g.height)
	g.renderer.Draw(
		background,
		mgl32.Vec2{screenWidth / 2, screenHeight / 2},
		mgl32.Vec2{screenWidth, screenHeight},
		0,
		mgl32.Vec3{0.8, 0.1, 0.9},
	)
	g.level.Draw(&g.renderer)
}

func (g Game) Delete() {
	g.renderer.Delete()
}
