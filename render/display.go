package render

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	Width  = 1260
	Height = 720
)

// CURENTLY UNUSED. NOT COMPLETE
type DisplayManager struct {
	width  int
	height int
	window *glfw.Window
}

func NewDisplayManager() *DisplayManager {
	return &DisplayManager{
		width:  Width,
		height: Height,
	}
}
