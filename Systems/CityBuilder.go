package systems

import (
	"fmt"

	"engo.io/ecs"
	"engo.io/engo"
)

type CityBuildingSystem struct{}

func (*CityBuildingSystem) Remove(ecs.BasicEntity) {}

func (*CityBuildingSystem) Update(dt float32) {
	if engo.Input.Button("AddCity").JustPressed() {
		fmt.Println("The gamer pressed F1")
	}
}

func (*CityBuildingSystem) New(*ecs.World) {
	fmt.Println("CityBuildingSystem was added to the Scene")
}
