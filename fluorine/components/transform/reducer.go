package transform

import (
	"github.com/DrTeePot/game/fluorine/action"
	"github.com/DrTeePot/game/fluorine/store"
)

func newTransformReducer() store.FloatReducer {
	return store.NewFloatReducer(3, reducerFunction)
}

func reducerFunction(
	s store.State_float32,
	a action.Action_float32,
) store.State_float32 {
	switch a.Instruction() {
	case setPosition:
		x := a.Value()[0]
		y := a.Value()[1]
		z := a.Value()[2]

		// TODO make this accurately account for scale and rotation
		newPosition := []float32{x, y, z, 0, 0, 0}

		return s.Assign(a.Entity(), newPosition)

	case increasePosition:
		x := a.Value()[0]
		y := a.Value()[1]
		z := a.Value()[2]
		data := s.GetEntity(a.Entity())

		newPosition := []float32{
			x + data[0],
			y + data[1],
			z + data[2],
			0,
			0,
			0,
		}

		return s.Assign(a.Entity(), newPosition)

	default:
		return s
	}
}
