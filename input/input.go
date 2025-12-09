package input

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

var InputMap map[glfw.Key]bool

var MouseX float64
var MouseY float64

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

func MousePositionCallback(w *glfw.Window, xpos, ypos float64) {
	MouseX = xpos
	MouseY = ypos
}
