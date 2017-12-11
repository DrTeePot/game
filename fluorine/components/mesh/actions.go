package mesh

import (
	"github.com/DrTeePot/game/fluorine/action"
)

func SetMesh(entity uint32, m float32) action.Action_float32 {
	return action.Create_float32(
		MeshComponent,
		setMesh,
		entity,
		[]float32{m},
	)
}
