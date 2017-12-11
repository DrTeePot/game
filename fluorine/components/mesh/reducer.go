package mesh

import (
	"github.com/DrTeePot/game/fluorine/action"
	"github.com/DrTeePot/game/fluorine/store"
)

func newMeshReducer() store.FloatReducer {
	return store.NewFloatReducer(1, reducerFunction)
}

func reducerFunction(
	s store.State_float32,
	a action.Action_float32,
) store.State_float32 {
	switch a.Instruction() {
	case setMesh:
		newMeshID := []float32{a.Value()[0]}

		return s.Assign(a.Entity(), newMeshID)

	default:
		return s
	}
}
