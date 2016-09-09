package main

import (
	"engo.io/engo"
	scene "github.com/engoengine/temod/Scene"
)

func main() {
	opts := engo.RunOptions{
		Title:  "Hello World",
		Width:  400,
		Height: 400,
	}
	engo.Run(opts, &scene.MyScene{})
}
