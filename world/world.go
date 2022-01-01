package world

import (
	"github.com/nicholasblaskey/raytracer/light"
)

type World struct {
	Objects []interface{} // TODO rework this to be a shape type
	Light   *light.Point
}

func New() *World {
	return &World{}
}
