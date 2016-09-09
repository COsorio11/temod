package helpers

import (
	"fmt"

	"engo.io/engo/common"
)

func LoadTexture(pathToFile string) *common.Texture {
	texture, err := common.LoadedSprite(pathToFile)
	if err != nil {
		fmt.Println("Default: ", err)
		texture, _ = common.LoadedSprite("textures/city.png")
	}
	return texture
}
