# Breakout

> The Arcade Classic Implemented with Go and OpenGL[^1]

## Getting Started

### Building From Source

For POSIX systems, use the following make commands to generate a binary.

```sh
# Build a standalone binary
make build

# Build and run
make run

# Clean up binary
make clean
```

For non-POSIX systems, use the builtin go commands to generate an executable.

## Controls

- `SPACE` Start playing.
- `A` / `D` Move left and right, respectively.
- `R` Reset level and player.
- `ESC` Quit the game.

[^1]: This project uses `go-gl`, `mgl32`, and `glfw` as dependencies. Built as a port of the [Learn OpenGL C++ Breakout](https://learnopengl.com/In-Practice/2D-Game/Breakout).
