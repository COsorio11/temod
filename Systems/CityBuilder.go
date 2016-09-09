package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	entities "github.com/engoengine/temod/Entities"
	helpers "github.com/engoengine/temod/Helpers"
)

type CityBuildingSystem struct {
	world        *ecs.World
	mouseTracker entities.MouseTracker
}

func (*CityBuildingSystem) Remove(ecs.BasicEntity) {}

func (cb *CityBuildingSystem) New(w *ecs.World) {
	cb.world = w
	cb.mouseTracker.BasicEntity = ecs.NewBasic()
	cb.mouseTracker.MouseComponent = common.MouseComponent{Track: true}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&cb.mouseTracker.BasicEntity, &cb.mouseTracker.MouseComponent, nil, nil)
		}
	}

}

func (cb *CityBuildingSystem) Update(dt float32) {
	if engo.Input.Button("AddCity").JustPressed() {
		city := entities.City{BasicEntity: ecs.NewBasic()}
		city.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{cb.mouseTracker.MouseComponent.MouseX, cb.mouseTracker.MouseComponent.MouseY},
			Width:    30,
			Height:   64,
		}

		texture := helpers.LoadTexture("textures/Mushroom2.png")

		city.RenderComponent = common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{0.1, 0.1},
		}

		for _, system := range cb.world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&city.BasicEntity, &city.RenderComponent, &city.SpaceComponent)
			}
		}
	}
}
