package world

import (
	"sort"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/light"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/shape"
)

type World struct {
	Objects []interface{} // TODO rework this to be a shape type
	Light   *light.Point
}

func New() *World {
	return &World{}
}

func (w *World) Intersect(r ray.Ray) []*intersection.Intersection {
	var intersections []*intersection.Intersection
	for _, obj := range w.Objects {
		// TODO figure out objects interface
		s := obj.(*shape.Sphere)
		intersections = append(intersections, s.Intersections(r)...)
	}

	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].T < intersections[j].T
	})

	return intersections
}
