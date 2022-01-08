package shape

import (
	"math"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Sphere struct {
	Transform matrix.Mat4
	Material  *material.Material
}

func NewSphere() *Sphere {
	return &Sphere{matrix.Ident4(), material.New()}
}

func NewGlassSphere() *Sphere {
	m := material.New()
	m.Transparency = 1.0
	m.RefractiveIndex = 1.5
	return &Sphere{matrix.Ident4(), m}
}

func (s *Sphere) localIntersections(r ray.Ray) []*intersection.Intersection {
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

func (s *Sphere) Intersections(origR ray.Ray) []*intersection.Intersection {
	return intersection.Intersections(s, origR, s.localIntersections)
}

func (s *Sphere) localNormalAt(objectPoint tuple.Tuple) tuple.Tuple {
	return objectPoint.Sub(tuple.Point(0.0, 0.0, 0.0))
}

// TODO See if there is a performance hit doing this.
func (s *Sphere) NormalAt(worldPoint tuple.Tuple) tuple.Tuple {
	return intersection.NormalAt(s, worldPoint, s.localNormalAt)
}

func (s *Sphere) GetMaterial() *material.Material {
	return s.Material
}

func (s *Sphere) SetMaterial(m *material.Material) {
	s.Material = m
}

func (s *Sphere) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *Sphere) SetTransform(m matrix.Mat4) {
	s.Transform = m
}
