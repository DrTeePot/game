package mesh

import (
	"github.com/DrTeePot/game/fluorine/store"
)

const (
	MeshComponent = "MESH"

	setMesh = iota
)

func CreateComponent() store.UniversalComponent_float32 {
	reducer := newMeshReducer()

	return store.NewUniversalComponent_float32(
		MeshComponent,
		reducer,
	)
}
