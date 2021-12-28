package shape

import (
	"math"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Sphere struct {
}

func NewSphere() *Sphere {
	return &Sphere{}
}

func (s *Sphere) Intersections(r ray.Ray) []*intersection.Intersection {
	sphereToRay := r.Origin.Sub(tuple.Point(0.0, 0.0, 0.0))

	a := r.Direction.Dot(r.Direction)
	b := 2.0 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1.0

	discriminant := b*b - 4*a*c

	if discriminant < 0.0 {
		return nil
	}

	t0 := (-b - math.Sqrt(discriminant)) / (2.0 * a)
	t1 := (-b + math.Sqrt(discriminant)) / (2.0 * a)

	return []*intersection.Intersection{
		&intersection.Intersection{Obj: s, T: t0},
		&intersection.Intersection{Obj: s, T: t1}}
}
