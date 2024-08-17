package breakout

import (
	"time"

	"github.com/go-gl/mathgl/mgl32"
)

type GameState uint32

const GAME_ACTIVE = 0
const GAME_MENU = 1
const GAME_WIN = 2

type Game struct {
	State    GameState
	Width    uint32
	Height   uint32
	Renderer SpriteRenderer
}

func NewGame(width uint32, height uint32) Game {
	return Game{
		State:    GAME_ACTIVE,
		Width:    width,
		Height:   height,
		Renderer: SpriteRenderer{},
	}
}

func (g Game) Init() {
	shader, err := LoadShader("shaders/sprite.vert", "shaders/sprite.frag", "sprite")
	if err != nil {
		panic(err)
	}

	projection := mgl32.Ortho(0.0, float32(g.Width), float32(g.Height), 0.0, -1.0, 1.0)
	shader.SetInteger("image", 0)
	shader.SetMatrix4("projection", projection)
	g.Renderer = NewSpriteRenderer(shader)
	LoadTexture("textures/awesomeface.png", "awesomeface")
}

func (g Game) ProcessInput(deltaTime time.Duration) {
}

func (g Game) Update(deltaTime time.Duration) {
}

func (g Game) Render() {
	texture := GetTexture("awesomeface")
	position := mgl32.Vec2{200.0, 200.0}
	size := mgl32.Vec2{300.0, 400.0}
	rotation := float32(45.0)
	color := mgl32.Vec3{1.0, 1.0, 1.0}
	g.Renderer.DrawSprite(texture, position, size, rotation, color)
}
