package light

import (
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Point struct {
	Position  tuple.Tuple
	Intensity tuple.Tuple
}

func NewPointLight(p, i tuple.Tuple) Point {
	return Point{p, i}
}
