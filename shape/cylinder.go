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

	return []*intersection.Intersection{
		&intersection.Intersection{Obj: s, T: 1.0}, // TEST REMOVE
	}
}

func (s *Cylinder) Intersections(origR ray.Ray) []*intersection.Intersection {
	return intersection.Intersections(s, origR, s.localIntersections)
}

func (s *Cylinder) localNormalAt(p tuple.Tuple) tuple.Tuple {
	xAbs, yAbs, zAbs := math.Abs(p[0]), math.Abs(p[1]), math.Abs(p[2])

	if xAbs >= yAbs && xAbs >= zAbs {
		return tuple.Vector(p[0], 0.0, 0.0)
	} else if yAbs > xAbs && yAbs > zAbs {
		return tuple.Vector(0.0, p[1], 0.0)
	}
	return tuple.Vector(0.0, 0.0, p[2])
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
