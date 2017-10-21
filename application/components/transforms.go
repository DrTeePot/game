package components

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Positionable interface {
	Position() mgl32.Vec3
	MoveTo(mgl32.Vec3)
	MoveBy(mgl32.Vec3)
}

type Rotateable interface {
	Rotation() mgl32.Vec3
	RotateBy(mgl32.Vec3)
	RotateTo(mgl32.Vec3)
}

type Scaleable interface {
	Scale() float32
	SetScale(float32)
}

type Transform3D struct {
	position mgl32.Vec3
	scale    float32
	rotation mgl32.Vec3 // TODO maybe this should be a quat
}

type Transform2D struct {
	// TODO this has no methods currently
	position mgl32.Vec2
	scale    float32
	rotation float32 // radian? degrees?
}

func (p Transform3D) Position() mgl32.Vec3 { return p.position }
func (p *Transform3D) MoveBy(v mgl32.Vec3) { p.position = p.position.Add(v) }
func (p *Transform3D) MoveTo(v mgl32.Vec3) { p.position = v }

func (r Transform3D) Rotation() mgl32.Vec3   { return r.rotation }
func (r *Transform3D) RotateTo(v mgl32.Vec3) { r.rotation = v }
func (r *Transform3D) RotateBy(v mgl32.Vec3) { r.rotation = r.rotation.Add(v) }

func (p Transform3D) Scale() float32      { return p.scale }
func (p *Transform3D) SetScale(s float32) { p.scale = s }
