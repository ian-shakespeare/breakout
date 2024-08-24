package breakout

import (
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	GAME_ACTIVE GameState = 0
	GAME_MENU   GameState = 1
	GAME_WIN    GameState = 2

	PLAYER_VELOCITY float32 = 1
	BALL_RADIUS     float32 = 25
)

var INITIAL_BALL_VELOCITY mgl32.Vec2 = mgl32.Vec2{0.5, -0.350}

type (
	GameState uint32

	Game struct {
		State        GameState
		Keys         []bool
		width        uint32
		height       uint32
		renderer     SpriteRenderer
		levels       []Level
		currentLevel uint32
		player       Entity
		ball         Ball
	}
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
	paddle, err := LoadTexture("assets/textures/paddle.png", "paddle")
	if err != nil {
		panic(err)
	}
	face, err := LoadTexture("assets/textures/awesomeface.png", "awesomeface")
	if err != nil {
		panic(err)
	}

	shader.SetInteger("tex", 0)
	gl.ActiveTexture(gl.TEXTURE0)

	projection := mgl32.Ortho(0, float32(width), float32(height), 0, -1, 1)
	shader.SetMatrix4("projection", projection)

	renderer := NewSpriteRenderer(shader)

	levels := make([]Level, 0)
	standard, err := LoadLevel("assets/levels/standard.lvl", width, height/2)
	if err != nil {
		panic(err)
	}
	fewSmallGaps, err := LoadLevel("assets/levels/few_small_gaps.lvl", width, height/2)
	if err != nil {
		panic(err)
	}
	spaceInvader, err := LoadLevel("assets/levels/space_invader.lvl", width, height/2)
	if err != nil {
		panic(err)
	}
	bounceGalore, err := LoadLevel("assets/levels/bounce_galore.lvl", width, height/2)
	if err != nil {
		panic(err)
	}
	levels = append(levels, standard, fewSmallGaps, spaceInvader, bounceGalore)
	currentLevel := uint32(0)

	State := GAME_ACTIVE
	Keys := make([]bool, 1024)
	for i := range Keys {
		Keys[i] = false
	}

	playerSize := mgl32.Vec2{0.2 * float32(width), 0.1 * float32(height)}
	player := Entity{
		Position:  mgl32.Vec2{0.5 * float32(width), float32(height) - 0.5*playerSize.Y()},
		Size:      playerSize,
		Velocity:  mgl32.Vec2{0, 0},
		Color:     mgl32.Vec3{1.0, 0.6, 0.1},
		Angle:     0,
		IsSolid:   true,
		Destroyed: false,
		Sprite:    paddle,
	}

	ballPosition := player.Position.Add(mgl32.Vec2{0, -BALL_RADIUS * 2})
	ball := NewBall(
		Entity{
			Position:  ballPosition,
			Size:      mgl32.Vec2{BALL_RADIUS, BALL_RADIUS},
			Velocity:  INITIAL_BALL_VELOCITY,
			Color:     mgl32.Vec3{1, 1, 1},
			Angle:     0,
			IsSolid:   true,
			Destroyed: false,
			Sprite:    face,
		},
		BALL_RADIUS,
		true,
	)

	return Game{State, Keys, width, height, renderer, levels, currentLevel, player, ball}
}

func (g *Game) ProcessInput(deltaTime time.Duration) {
	if g.State != GAME_ACTIVE {
		return
	}

	velocity := PLAYER_VELOCITY * float32(deltaTime.Milliseconds())
	playerHalfWidth := g.player.Size.X() / 2
	if g.Keys[glfw.KeyA] {
		if g.player.Position.X()-playerHalfWidth >= 0 {
			g.player.Position = g.player.Position.Sub(mgl32.Vec2{velocity, 0})

			if g.ball.stuck {
				g.ball.entity.SetX(g.ball.entity.Position.X() - velocity)
			}
		}
	}
	if g.Keys[glfw.KeyD] {
		if g.player.Position.X()+playerHalfWidth <= float32(g.width) {
			g.player.Position = g.player.Position.Add(mgl32.Vec2{velocity, 0})

			if g.ball.stuck {
				g.ball.entity.SetX(g.ball.entity.Position.X() + velocity)
			}
		}
	}

	if g.Keys[glfw.KeySpace] {
		g.ball.stuck = false
	}
}

func (g *Game) Update(deltaTime time.Duration) {
	if g.State != GAME_ACTIVE {
		return
	}

	g.ball.Move(deltaTime, g.width)

	for i := 0; i < len(g.levels[g.currentLevel].bricks); i += 1 {
		brick := &g.levels[g.currentLevel].bricks[i]

		if !brick.Destroyed {
			if g.ball.Collides(brick) {
				brick.Destroyed = true
			}
		}
	}
}

func (g *Game) Render() {
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
	g.levels[g.currentLevel].Draw(&g.renderer)
	g.player.Draw(&g.renderer)
	g.ball.entity.Draw(&g.renderer)
}

func (g *Game) Delete() {
	g.renderer.Delete()
}
