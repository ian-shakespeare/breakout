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
const WINDOW_HEIGHT = 800

func init() {
	runtime.LockOSThread()
}

func main() {
	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
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
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	game := breakout.NewGame(WINDOW_WIDTH, WINDOW_HEIGHT)

	window.SetKeyCallback(
		func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
			if key == glfw.KeyEscape {
				window.SetShouldClose(true)
			}

			if action == glfw.Press {
				game.Keys[key] = true
			} else if action == glfw.Release {
				game.Keys[key] = false
			}
		},
	)

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
