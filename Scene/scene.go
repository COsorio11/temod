package scene

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	systems "github.com/engoengine/temod/Systems"
)

type MyScene struct{}

// Type uniquely defines your game type37G
func (*MyScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*MyScene) Preload() {
	err := engo.Files.Load("textures/kiwi.svg")
	fmt.Println("svg: ", err)
	engo.Files.Load("textures/Mushroom2.png")
	err = engo.Files.Load("textures/bayless.jpeg")
	fmt.Println("jpeg: ", err)
	engo.Files.Load("textures/city.png")
	err = engo.Files.Load("textures/LinkandEpona2.gif")
	fmt.Println("gif: ", err)
	engo.Files.Load("textures/mustacheguy.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (scene *MyScene) Setup(world *ecs.World) {
	kbspeed := float32(400)
	engo.Input.RegisterButton("AddCity", engo.F1)
	common.SetBackground(color.White)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.AnimationSystem{})
	//
	// spriteSheet := common.NewSpritesheetFromFile("textures/mustacheguy.png", 150, 150)
	// animationAction := common.Animation{Name: "run", Frames: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}}
	//
	// hero := entities.Hero{BasicEntity: ecs.NewBasic()}
	//
	// hero.RenderComponent = spriteSheet

	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(&systems.CityBuildingSystem{})
	world.AddSystem(common.NewKeyboardScroller(kbspeed, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis))
	world.AddSystem(&common.EdgeScroller{400, 20})
	world.AddSystem(&common.MouseZoomer{-0.125})

}
