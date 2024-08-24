package breakout

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
)

type Level struct {
	width       uint32
	height      uint32
	bricks      []Entity
	isCompleted bool
}

func LoadLevel(filename string, width uint32, height uint32) (Level, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Level{}, err
	}
	defer file.Close()

	tileData := make([][]uint32, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		columns := strings.Split(scanner.Text(), "")
		tiles := make([]uint32, 0)
		for _, tile := range columns {
			val, err := strconv.ParseUint(tile, 10, 32)
			if err != nil {
				return Level{}, err
			}
			tiles = append(tiles, uint32(val))
		}
		tileData = append(tileData, tiles)
	}

	return tileDataToLevel(tileData, width, height), nil
}

func (l *Level) Draw(renderer *SpriteRenderer) {
	for _, brick := range l.bricks {
		if !brick.Destroyed {
			brick.Draw(renderer)
		}
	}
}

func tileDataToLevel(tileData [][]uint32, levelWidth uint32, levelHeight uint32) Level {
	height := len(tileData)
	width := len(tileData[0])

	brickWidth := float32(levelWidth) / float32(width)
	brickHeight := float32(levelHeight) / float32(height)
	brickTexture := GetTexture("block")

	bricks := make([]Entity, 0)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var color mgl32.Vec3

			switch tileData[y][x] {
			case 0:
				continue
			case 1:
				color = mgl32.Vec3{1.0, 0.25, 0.2}
				break
			case 2:
				color = mgl32.Vec3{0.2, 1.0, 0.4}
				break
			case 3:
				color = mgl32.Vec3{0.2, 0.4, 1.0}
				break
			case 4:
				color = mgl32.Vec3{1.0, 0.2, 1.0}
				break
			case 5:
				color = mgl32.Vec3{1.0, 1.0, 0.2}
				break
			}

			halfWidth := 0.5 * brickWidth
			halfHeight := 0.5 * brickHeight

			bricks = append(bricks, Entity{
				Position:  mgl32.Vec2{brickWidth*float32(x) + halfWidth, brickHeight*float32(y) + halfHeight},
				Size:      mgl32.Vec2{brickWidth, brickHeight},
				Velocity:  mgl32.Vec2{0, 0},
				Color:     color,
				Angle:     0,
				IsSolid:   true,
				Destroyed: false,
				Sprite:    brickTexture,
			})
		}
	}

	return Level{width: levelWidth, height: levelHeight, bricks: bricks, isCompleted: false}
}
