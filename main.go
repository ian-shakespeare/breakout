package main

import (
	breakout "breakout/src"
	"runtime"
	"time"

	_ "image/png"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const WINDOW_WIDTH = 800
const WINDOW_HEIGHT = 600

func init() {
	runtime.LockOSThread()
}

func main() {
	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Breakout", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	// Initialize GL
	if err = gl.Init(); err != nil {
		panic(err)
	}
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

  game := breakout.NewGame(WINDOW_WIDTH, WINDOW_HEIGHT)

  lastFrame := time.Now()

	for !window.ShouldClose() {
    now := time.Now()
    deltaTime := now.Sub(lastFrame)
    lastFrame = now

    game.ProcessInput(deltaTime)
    game.Update(deltaTime)

		gl.Clear(gl.COLOR_BUFFER_BIT)
    game.Render()

		window.SwapBuffers()
		glfw.PollEvents()
	}

  game.Delete()
}
