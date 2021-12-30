package material

import (
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Material struct {
	Color     tuple.Tuple
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func New() Material {
	return Material{
		Color:     tuple.Color(1.0, 1.0, 1.0),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}
