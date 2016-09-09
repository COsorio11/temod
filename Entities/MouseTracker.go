package entities

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}
