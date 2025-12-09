package input

import "github.com/go-gl/glfw/v3.3/glfw"

var InputMap map[glfw.Key]bool

func InitInput() {
	InputMap = make(map[glfw.Key]bool)
}

func KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	switch action {
	case glfw.Press:
		InputMap[key] = true
	case glfw.Release:
		InputMap[key] = false
	}
}
