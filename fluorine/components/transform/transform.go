package transform

import (
	"github.com/DrTeePot/game/fluorine/store"
)

const (
	TransformName = "TRANSFORM"

	setPosition = iota
	increasePosition
)

func CreateTransformComponent() store.UniversalComponent_float32 {
	transformReducer := newTransformReducer()

	return store.NewUniversalComponent_float32(
		TransformName,
		transformReducer,
	)
}
