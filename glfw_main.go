package main

import (
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func glfwInit() {
	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}
}

func glInit() {
	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}
}
