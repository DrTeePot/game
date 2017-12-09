package transform

import (
	"github.com/DrTeePot/game/fluorine/store"
)

const (
	transformName = "TRANSFORM"

	setPosition = iota
	increasePosition
)

func CreateTransformComponent() store.UniversalComponent_float32 {
	transformReducer := newTransformReducer()

	return store.NewUniversalComponent_float32(
		transformName,
		transformReducer,
	)
}
