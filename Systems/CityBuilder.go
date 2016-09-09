package systems

import (
	"fmt"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	entities "github.com/engoengine/temod/Entities"
)

type CityBuildingSystem struct {
	world        *ecs.World
	mouseTracker entities.MouseTracker
}

func (*CityBuildingSystem) Remove(ecs.BasicEntity) {}

func (cb *CityBuildingSystem) New(w *ecs.World) {
	cb.world = w
	fmt.Println("CityBuildingSystem was added to the Scene")
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
		//fmt.Println("The gamer pressed F1")
		//fmt.Println("MOUSE X", cb.mouseTracker.MouseComponent.MouseX)
		city := entities.City{BasicEntity: ecs.NewBasic()}
		city.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{cb.mouseTracker.MouseComponent.MouseX, cb.mouseTracker.MouseComponent.MouseY},
			Width:    30,
			Height:   64,
		}

		texture, err := common.LoadedSprite("textures/city.png")
		//texture, err := common.PreloadedSpriteSingle("texture/city.png")
		if err != nil {
			panic("Unable to load texture: " + err.Error())
		}

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
