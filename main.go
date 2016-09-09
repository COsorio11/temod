package main

import (
	"engo.io/engo"
	scene "github.com/engoengine/temod/Scene"
)

func main() {
	opts := engo.RunOptions{
		Title:          "Hello World",
		Width:          800,
		Height:         800,
		StandardInputs: true,
	}
	engo.Run(opts, &scene.MyScene{})
}
