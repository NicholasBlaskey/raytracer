package shape

import (
	"math"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type SmoothTriangle struct {
	Transform   matrix.Mat4
	Material    *material.Material
	Parent      intersection.Intersectable
	P0          tuple.Tuple
	P1          tuple.Tuple
	P2          tuple.Tuple
	N0          tuple.Tuple
	N1          tuple.Tuple
	N2          tuple.Tuple
	E0          tuple.Tuple
	E1          tuple.Tuple
	Normal      tuple.Tuple
	boundingBox intersection.Bounds
}

func NewSmoothTriangle(p0, p1, p2, n0, n1, n2 tuple.Tuple) *SmoothTriangle {
	e0 := p1.Sub(p0)
	e1 := p2.Sub(p0)
	normal := e1.Cross(e0).Normalize()

	min := tuple.Point(
		math.Min(p0[0], math.Min(p1[0], p2[0])),
		math.Min(p0[1], math.Min(p1[1], p2[1])),
		math.Min(p0[2], math.Min(p1[2], p2[2])),
	)
	max := tuple.Point(
		math.Max(p0[0], math.Max(p1[0], p2[0])),
		math.Max(p0[1], math.Max(p1[1], p2[1])),
		math.Max(p0[2], math.Max(p1[2], p2[2])),
	)

	return &SmoothTriangle{
		Transform: matrix.Ident4(),
		Material:  material.New(),

		P0: p0, P1: p1, P2: p2,
		N0: n0, N1: n1, N2: n2,
		E0:          e0,
		E1:          e1,
		Normal:      normal,
		boundingBox: intersection.Bounds{min, max},
	}
}

func (s *SmoothTriangle) localIntersections(r ray.Ray) []*intersection.Intersection {
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
		&intersection.Intersection{Obj: s, T: t, U: u, V: v},
	}
}

func (s *SmoothTriangle) Intersections(origR ray.Ray) []*intersection.Intersection {
	return intersection.Intersections(s, origR, s.localIntersections)
}

func (s *SmoothTriangle) NormalAt(worldPoint tuple.Tuple, i *intersection.Intersection) tuple.Tuple {
	localNormalAt := func(p tuple.Tuple) tuple.Tuple {
		return s.N1.Mul(i.U).Add(
			s.N2.Mul(i.V)).Add(
			s.N0.Mul(1 - i.U - i.V))
	}

	return intersection.NormalAt(s, worldPoint, localNormalAt)
}

func (s *SmoothTriangle) GetMaterial() *material.Material {
	return s.Material
}

func (s *SmoothTriangle) SetMaterial(m *material.Material) {
	s.Material = m
}

func (s *SmoothTriangle) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *SmoothTriangle) SetTransform(m matrix.Mat4) {
	s.Transform = m
}

func (s *SmoothTriangle) GetParent() intersection.Intersectable {
	return s.Parent
}

func (s *SmoothTriangle) SetParent(p intersection.Intersectable) {
	s.Parent = p
}

func (s *SmoothTriangle) Bounds() intersection.Bounds {
	return s.boundingBox
}
