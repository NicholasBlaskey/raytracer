package shape

import (
	"math"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Cylinder struct {
	Transform matrix.Mat4
	Material  *material.Material
}

func NewCylinder() *Cylinder {
	return &Cylinder{matrix.Ident4(), material.New()}
}

func (s *Cylinder) localIntersections(r ray.Ray) []*intersection.Intersection {
	a := r.Direction[0]*r.Direction[0] + r.Direction[2]*r.Direction[2]

	if math.Abs(a) < intersection.EPSILON { // Ray is parallel to y axis
		return nil
	}

	b := 2*r.Origin[0]*r.Direction[0] + 2*r.Origin[2]*r.Direction[2]
	c := r.Origin[0]*r.Origin[0] + r.Origin[2]*r.Origin[2] - 1

	disc := b*b - 4*a*c
	if disc < 0 {
		return nil
	}

	t0 := (-b - math.Sqrt(disc)) / (2.0 * a)
	t1 := (-b + math.Sqrt(disc)) / (2.0 * a)

	return []*intersection.Intersection{
		&intersection.Intersection{Obj: s, T: t0},
		&intersection.Intersection{Obj: s, T: t1},
	}
}

func (s *Cylinder) Intersections(origR ray.Ray) []*intersection.Intersection {
	return intersection.Intersections(s, origR, s.localIntersections)
}

func (s *Cylinder) localNormalAt(p tuple.Tuple) tuple.Tuple {
	return tuple.Vector(p[0], 0.0, p[2])
}

func (s *Cylinder) NormalAt(worldPoint tuple.Tuple) tuple.Tuple {
	return intersection.NormalAt(s, worldPoint, s.localNormalAt)
}

func (s *Cylinder) GetMaterial() *material.Material {
	return s.Material
}

func (s *Cylinder) SetMaterial(m *material.Material) {
	s.Material = m
}

func (s *Cylinder) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *Cylinder) SetTransform(m matrix.Mat4) {
	s.Transform = m
}
