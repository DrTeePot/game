package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type actionFunc func()
type inputActions map[glfw.Key]map[glfw.Action]actionFunc

type keyboardListener struct {
	window  *glfw.Window
	actions inputActions
}

// NewKeyboardListener creates a listener on the window
// Must only be called from the main thread
func NewKeyboardListener(window *glfw.Window) *keyboardListener {
	listener := &keyboardListener{
		window:  window,
		actions: make(inputActions),
	}

	_ = window.SetKeyCallback(listener.glfwCallback)

	return listener
}

func (l *keyboardListener) On(
	key glfw.Key,
	action glfw.Action,
	callback actionFunc) {

	// the map is nil if uninitialized
	if l.actions[key] == nil {
		l.actions[key] = make(map[glfw.Action]actionFunc)
	}
	l.actions[key][action] = callback
}

// OnKeyPress sets a callback for the Press event for the given key
func (l *keyboardListener) OnKeyPress(key glfw.Key, callback actionFunc) {
	// should this be direct access of use the On function?
	l.On(key, glfw.Press, callback)
}

// OnKeyRelease sets a callback for the Release event for the given key
func (l *keyboardListener) OnKeyRelease(key glfw.Key, callback actionFunc) {
	// should this be direct access of use the On function?
	l.On(key, glfw.Release, callback)
}

// OnMovementKey allows setting a movement key with a callback for
//  on press and release
func (l *keyboardListener) OnMovementKey(
	key glfw.Key,
	onPress actionFunc,
	onRelease actionFunc,
) {
	l.OnKeyPress(key, onPress)
	l.OnKeyRelease(key, onRelease)
}

func (l *keyboardListener) glfwCallback(
	w *glfw.Window,
	key glfw.Key,
	scancode int,
	action glfw.Action,
	mods glfw.ModifierKey,
) {
	// make sure we have actions for this key
	actionMap, ok := l.actions[key]
	if !ok {
		return
	}

	// if callback for the action exists, call it
	if call, ok := actionMap[action]; ok {
		call()
	}
}
