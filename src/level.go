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

func (l *Level) Reset() {
	for i := 0; i < len(l.bricks); i += 1 {
		brick := &l.bricks[i]

		brick.Destroyed = false
	}
}

func tileDataToLevel(tileData [][]uint32, levelWidth uint32, levelHeight uint32) Level {
	height := len(tileData)
	width := len(tileData[0])

	blockWidth := float32(levelWidth) / float32(width)
	blockHeight := float32(levelHeight) / float32(height)
	blockSprite := GetTexture("block")
	solidBlockSprite := GetTexture("block_solid")

	bricks := make([]Entity, 0)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var color mgl32.Vec3
			solid := false
			sprite := blockSprite

			switch tileData[y][x] {
			case 0:
				continue
			case 1:
				color = mgl32.Vec3{0.8, 0.8, 0.7}
				solid = true
				sprite = solidBlockSprite
				break
			case 2:
				color = mgl32.Vec3{0.2, 0.6, 1.0}
				break
			case 3:
				color = mgl32.Vec3{0.0, 0.7, 0.0}
				break
			case 4:
				color = mgl32.Vec3{0.8, 0.8, 0.4}
				break
			case 5:
				color = mgl32.Vec3{1.0, 0.5, 0.0}
				break
			}

			halfWidth := 0.5 * blockWidth
			halfHeight := 0.5 * blockHeight

			bricks = append(bricks, Entity{
				Position:  mgl32.Vec2{blockWidth*float32(x) + halfWidth, blockHeight*float32(y) + halfHeight},
				Size:      mgl32.Vec2{blockWidth, blockHeight},
				Velocity:  mgl32.Vec2{0, 0},
				Color:     color,
				Angle:     0,
				IsSolid:   solid,
				Destroyed: false,
				Sprite:    sprite,
			})
		}
	}

	return Level{width: levelWidth, height: levelHeight, bricks: bricks, isCompleted: false}
}
