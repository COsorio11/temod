package entities

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type City struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}
