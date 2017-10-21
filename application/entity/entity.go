package entity

import (
	"github.com/DrTeePot/game/components"

	"github.com/go-gl/mathgl/mgl32"
)

type Entity interface {
	ID() int32            // TODO should this be a uuid
	Init()                // for initilizing any components
	components.Rotateable // TODO remove these. Specific entities only
	components.Positionable
	components.Scaleable
	components.Renderable
}

type entity struct {
	// TODO identity component
	*components.TexturedModel // implements Renderable
	*components.Transform3D
}

func NewEntity(
	mesh string,
	texture string,
	position mgl32.Vec3,
	rotation mgl32.Vec3,
) Entity {
	e := &entity{}

	e.SetShine(10)
	e.SetReflectivity(1)
	e.MoveTo(position)
	e.SetScale(1)
	e.RotateTo(rotation)

	return e
}

func (e entity) ID() int32 { return 1 }
func (e entity) Init()     {}
