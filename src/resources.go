package breakout

import (
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

var shaders = make(map[string]Shader)
var textures = make(map[string]Texture)

func LoadShader(vertexFilename string, fragmentFilename string, name string) (Shader, error) {
	vertexSource, err := os.ReadFile(vertexFilename)
	if err != nil {
		return Shader{}, err
	}

	fragmentSource, err := os.ReadFile(fragmentFilename)
	if err != nil {
		return Shader{}, err
	}

	shader, err := NewShader(string(vertexSource), string(fragmentSource))
	if err != nil {
		return Shader{}, err
	}

	shaders[name] = shader
	return shader, nil
}

func GetShader(name string) Shader {
	return shaders[name]
}

func LoadTexture(filename string, name string) (Texture, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	texture := NewTexture(img)

	textures[name] = texture
	return texture, nil
}

func GetTexture(name string) Texture {
	return textures[name]
}

func ClearResources() {
	for key := range shaders {
		delete(shaders, key)
	}

	for key := range textures {
		delete(textures, key)
	}
}
