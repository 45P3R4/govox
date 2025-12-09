package window

import (
	"log"
	"vox/input"
	"vox/view"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func CreateWindow(width int, height int, title string) *glfw.Window {
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.DoubleBuffer, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	window.MakeContextCurrent()

	window.SetSizeCallback(sizeCallback)
	window.SetKeyCallback(input.KeyCallback)
	window.SetCursorPosCallback(input.MousePositionCallback)

	return window
}

func sizeCallback(window *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	view.AspectRatio = float32(width) / float32(height)
	view.InitPerspectiveProjetion(0)
}
