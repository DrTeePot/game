package transform

import (
	"github.com/DrTeePot/game/fluorine/action"
)

func SetPosition(entity uint32, x, y, z float32) action.Action_float32 {
	return action.Create_float32(
		TransformName,
		setPosition,
		entity,
		[]float32{x, y, z},
	)
}

func IncreasePosition(entity uint32, x, y, z float32) action.Action_float32 {
	return action.Create_float32(
		TransformName,
		increasePosition,
		entity,
		[]float32{x, y, z},
	)
}
