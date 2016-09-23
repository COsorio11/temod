package entities

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type Hero struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}
