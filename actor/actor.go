package actor

import (
	// "github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Actor struct {
	Position mgl32.Vec3
	Rotation mgl32.Vec3
	Scale    mgl32.Vec3
}

func (a *Actor) SetPosition(newPosition mgl32.Vec3) {
	a.Position = newPosition
}

func (a *Actor) SetRotation(newRotation mgl32.Vec3) {
	a.Rotation = newRotation
}

func (a *Actor) SetScale(newScale mgl32.Vec3) {
	a.Scale = newScale
}

func Update() {}
