package world

import (
	"sort"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/light"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type World struct {
	Objects []intersection.Intersectable
	Light   *light.Point
}

func New() *World {
	return &World{}
}

func (w *World) Intersect(r ray.Ray) []*intersection.Intersection {
	var intersections []*intersection.Intersection
	for _, obj := range w.Objects {
		intersections = append(intersections, obj.Intersections(r)...)
	}

	sort.Slice(intersections, func(i, j int) bool {
		// Less function
		return intersections[i].T < intersections[j].T
	})

	return intersections
}

func (w *World) ShadeHit(comps *intersection.Computations, remaining int) tuple.Tuple {
	inShadow := w.IsShadowed(comps.OverPoint)
	surface := comps.Obj.GetMaterial().Lighting(
		comps.Obj, *w.Light, comps.Point, comps.Eyev, comps.Normalv, inShadow)

	reflected := w.ReflectedColor(comps, remaining)

	return surface.Add(reflected)
}

func (w *World) ColorAt(r ray.Ray, remaining int) tuple.Tuple {
	xs := w.Intersect(r)
	inter := intersection.Hit(xs)

	if inter == nil {
		return tuple.Color(0.0, 0.0, 0.0)
	}

	comps := inter.PrepareComputations(r, xs)

	return w.ShadeHit(comps, remaining)
}

func (w *World) ReflectedColor(comps *intersection.Computations, remaining int) tuple.Tuple {
	if remaining <= 0 || comps.Obj.GetMaterial().Reflective == 0 {
		return tuple.Color(0.0, 0.0, 0.0)
	}

	reflectRay := ray.New(comps.OverPoint, comps.Reflectv)
	c := w.ColorAt(reflectRay, remaining-1)

	return c.Mul(comps.Obj.GetMaterial().Reflective)
}

func (w *World) IsShadowed(p tuple.Tuple) bool {
	v := w.Light.Position.Sub(p)
	distance := v.Magnitude()
	dir := v.Normalize()

	intersections := w.Intersect(ray.Ray{p, dir})
	h := intersection.Hit(intersections)

	return h != nil && h.T < distance
}
