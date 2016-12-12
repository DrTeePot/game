package entity

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Position         mgl32.Vec3
	Pitch, Yaw, Roll float32
}

func (c *Camera) Move() {
	// TODO make a wrapper for all this adding and subtracting
	// may also be better to interact with c.Position[0] = instead
	// of re-assigning it
	// if watcher.Down(keyboard.W) {
	// 	// if we are pressing forward, it means
	// 	// we are moving negative z since positive z
	// 	// is towards us
	// 	c.Position = c.Position.Sub(mgl32.Vec3{0, 0, 0.02})
	// 	fmt.Println("W Down")
	// }
	// if watcher.Down(keyboard.D) {
	// 	c.Position = c.Position.Add(mgl32.Vec3{0.02, 0, 0})
	// }
	// if watcher.Down(keyboard.A) {
	// 	c.Position = c.Position.Sub(mgl32.Vec3{0.02, 0, 0})
	// }

}
