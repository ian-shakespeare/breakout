package breakout

import (
	"math"
	"sync"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type GameState uint32

const (
	GAME_ACTIVE GameState = 0
	GAME_MENU   GameState = 1
	GAME_WIN    GameState = 2

	PLAYER_VELOCITY float32 = 1
	BALL_RADIUS     float32 = 12.5
)

var INITIAL_BALL_VELOCITY mgl32.Vec2 = mgl32.Vec2{0.5, -0.350}

type Game struct {
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

func NewGame(width uint32, height uint32) Game {
	shader, err := LoadShader("assets/shaders/sprite.vert", "assets/shaders/sprite.frag", "sprite")
	if err != nil {
		panic(err)
	}
	shader.Use()

	if _, err := LoadTexture("assets/textures/block.png", "block"); err != nil {
		panic(err)
	}
	if _, err := LoadTexture("assets/textures/block_solid.png", "block_solid"); err != nil {
		panic(err)
	}
	if _, err := LoadTexture("assets/textures/background.jpg", "background"); err != nil {
		panic(err)
	}
	paddle, err := LoadTexture("assets/textures/paddle.png", "paddle")
	if err != nil {
		panic(err)
	}
	awesomeface, err := LoadTexture("assets/textures/awesomeface.png", "awesomeface")
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

	playerSize := mgl32.Vec2{0.2 * float32(width), 0.05 * float32(height)}
	player := Entity{
		Position:  mgl32.Vec2{0.5 * float32(width), float32(height) - 0.5*playerSize.Y()},
		Size:      playerSize,
		Velocity:  mgl32.Vec2{0, 0},
		Color:     mgl32.Vec3{1, 1, 1},
		Angle:     0,
		IsSolid:   true,
		Destroyed: false,
		Sprite:    paddle,
	}

	ballPosition := player.Position.Sub(mgl32.Vec2{0, (0.5 * player.Size.Y()) + 2*BALL_RADIUS})
	ball := NewBall(ballPosition, BALL_RADIUS, INITIAL_BALL_VELOCITY, awesomeface)

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

	if g.ball.entity.Position.Y() >= float32(g.height) {
		g.levels[g.currentLevel].Reset()
		g.resetPlayer()
		return
	}

	wg := sync.WaitGroup{}

	go g.handleCollisions(&wg)

	wg.Wait()
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
		mgl32.Vec3{1, 1, 1},
	)
	g.levels[g.currentLevel].Draw(&g.renderer)
	g.player.Draw(&g.renderer)
	g.ball.entity.Draw(&g.renderer)
}

func (g *Game) Delete() {
	g.renderer.Delete()
}

func (g *Game) resetPlayer() {
	player := Entity{
		Position:  mgl32.Vec2{0.5 * float32(g.width), float32(g.height) - 0.5*g.player.Size.Y()},
		Size:      g.player.Size,
		Velocity:  mgl32.Vec2{0, 0},
		Color:     g.player.Color,
		Angle:     0,
		IsSolid:   true,
		Destroyed: false,
		Sprite:    g.player.Sprite,
	}
	g.player = player

	ballPosition := player.Position.Sub(mgl32.Vec2{0, player.Size.Y() + 2*BALL_RADIUS})
	g.ball.Reset(ballPosition, INITIAL_BALL_VELOCITY)
}

func (g *Game) handleCollisions(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for i := 0; i < len(g.levels[g.currentLevel].bricks); i += 1 {
		brick := &g.levels[g.currentLevel].bricks[i]

		if !brick.Destroyed {
			collision := g.ball.Collides(brick)
			if !collision.Collided {
				continue
			}

			if !brick.IsSolid {
				brick.Destroyed = true
			}

			if collision.Direction == LEFT || collision.Direction == RIGHT {
				g.ball.entity.InvertXVelocity()
				penetration := g.ball.radius - float32(math.Abs(float64(collision.Difference.X())))
				if collision.Direction == LEFT {
					g.ball.entity.SetX(g.ball.entity.Position.X() + penetration)
				} else {
					g.ball.entity.SetX(g.ball.entity.Position.X() - penetration)
				}
			} else {
				g.ball.entity.InvertYVelocity()
				penetration := g.ball.radius - float32(math.Abs(float64(collision.Difference.Y())))
				if collision.Direction == UP {
					g.ball.entity.SetY(g.ball.entity.Position.Y() - penetration)
				} else {
					g.ball.entity.SetY(g.ball.entity.Position.Y() + penetration)
				}
			}
		}
	}

	paddleCollision := g.ball.Collides(&g.player)
	if paddleCollision.Collided {
		halfWidth := 0.5 * g.player.Size.X()
		center := g.player.Position.X() + halfWidth
		distance := (g.ball.entity.Position.X() + g.ball.radius) - center
		percentage := distance / halfWidth

		strength := float32(2)
		oldVelocity := g.ball.entity.Velocity
		g.ball.entity.Velocity = mgl32.Vec2{
			INITIAL_BALL_VELOCITY.X() * percentage * strength,
			-1 * float32(math.Abs(float64(g.ball.entity.Velocity.Y()))),
		}
		g.ball.entity.Velocity = g.ball.entity.Velocity.Normalize().Mul(oldVelocity.Len())
	}
}
