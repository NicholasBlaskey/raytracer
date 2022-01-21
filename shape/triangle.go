package shape

import (
	"math"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Triangle struct {
	Transform   matrix.Mat4
	Material    *material.Material
	Parent      intersection.Intersectable
	P0          tuple.Tuple
	P1          tuple.Tuple
	P2          tuple.Tuple
	E0          tuple.Tuple
	E1          tuple.Tuple
	Normal      tuple.Tuple
	boundingBox intersection.Bounds
}

func NewTriangle(p0, p1, p2 tuple.Tuple) *Triangle {
	e0 := p1.Sub(p0)
	e1 := p2.Sub(p0)
	normal := e1.Cross(e0).Normalize()

	min := p0.Min(p1).Min(p2)
	max := p0.Max(p1).Max(p2)
	return &Triangle{
		Transform:   matrix.Ident4(),
		Material:    material.New(),
		P0:          p0,
		P1:          p1,
		P2:          p2,
		E0:          e0,
		E1:          e1,
		Normal:      normal,
		boundingBox: intersection.Bounds{min, max},
	}
}

func (s *Triangle) localIntersections(r ray.Ray) []*intersection.Intersection {
	dirCrossE1 := r.Direction.Cross(s.E1)
	det := s.E0.Dot(dirCrossE1)
	if math.Abs(det) < intersection.EPSILON {
		return nil
	}

	f := 1.0 / det

	p0ToOrigin := r.Origin.Sub(s.P0)
	u := f * p0ToOrigin.Dot(dirCrossE1)

	if u < 0 || u > 1 {
		return nil
	}

	originCrossE0 := p0ToOrigin.Cross(s.E0)
	v := f * r.Direction.Dot(originCrossE0)
	if v < 0 || u+v > 1.0 {
		return nil
	}

	t := f * s.E1.Dot(originCrossE0)
	return []*intersection.Intersection{
		&intersection.Intersection{Obj: s, T: t},
	}
}

func (s *Triangle) Intersections(origR ray.Ray) []*intersection.Intersection {
	return intersection.Intersections(s, origR, s.localIntersections)
}

func (s *Triangle) localNormalAt(p tuple.Tuple) tuple.Tuple {
	return s.Normal
}

func (s *Triangle) NormalAt(worldPoint tuple.Tuple, i *intersection.Intersection) tuple.Tuple {
	return intersection.NormalAt(s, worldPoint, s.localNormalAt)
}

func (s *Triangle) GetMaterial() *material.Material {
	return s.Material
}

func (s *Triangle) SetMaterial(m *material.Material) {
	s.Material = m
}

func (s *Triangle) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *Triangle) SetTransform(m matrix.Mat4) {
	s.Transform = m
}

func (s *Triangle) GetParent() intersection.Intersectable {
	return s.Parent
}

func (s *Triangle) SetParent(p intersection.Intersectable) {
	s.Parent = p
}

func (s *Triangle) Bounds() intersection.Bounds {
	return s.boundingBox
}
