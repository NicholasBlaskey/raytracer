package shape

import (
	"math"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Cone struct {
	Transform matrix.Mat4
	Material  *material.Material
	Min       float64
	Max       float64
	Closed    bool
}

func NewCone() *Cone {
	return &Cone{
		Transform: matrix.Ident4(),
		Material:  material.New(),
		Min:       math.Inf(-1.0),
		Max:       math.Inf(1.0)}
}

func checkConeCap(r ray.Ray, t float64) bool {
	x := r.Origin[0] + t*r.Direction[0]
	z := r.Origin[2] + t*r.Direction[2]

	return (x*x + z*z) <= 1
}

func (s *Cone) intersectCaps(r ray.Ray,
	xs []*intersection.Intersection) []*intersection.Intersection {

	if !s.Closed || math.Abs(r.Direction[1]) < intersection.EPSILON {
		return xs
	}

	t := (s.Min - r.Origin[1]) / r.Direction[1]
	if checkConeCap(r, t) {
		xs = append(xs, &intersection.Intersection{Obj: s, T: t})
	}

	t = (s.Max - r.Origin[1]) / r.Direction[1]
	if checkConeCap(r, t) {
		xs = append(xs, &intersection.Intersection{Obj: s, T: t})
	}
	return xs
}

func (s *Cone) localIntersections(r ray.Ray) []*intersection.Intersection {
	a := r.Direction[0]*r.Direction[0] -
		r.Direction[1]*r.Direction[1] +
		r.Direction[2]*r.Direction[2]

	b := 2*r.Origin[0]*r.Direction[0] -
		2*r.Origin[1]*r.Direction[1] +
		2*r.Origin[2]*r.Direction[2]

	c := r.Origin[0]*r.Origin[0] -
		r.Origin[1]*r.Origin[1] +
		r.Origin[2]*r.Origin[2]

	if math.Abs(a) < intersection.EPSILON { // Ray is parallel to one of cones half
		if b == 0 { // Ray misses
			return nil
		}
		t := -c / (2.0 * b)
		//return s.intersectCaps(r, nil)
		return []*intersection.Intersection{&intersection.Intersection{Obj: s, T: t}}
	}

	disc := b*b - 4*a*c
	if disc < 0 {
		return nil
	}

	t0 := (-b - math.Sqrt(disc)) / (2.0 * a)
	t1 := (-b + math.Sqrt(disc)) / (2.0 * a)

	if t0 > t1 { // Is this needed? Dont think so?
		t0, t1 = t1, t0
	}

	var xs []*intersection.Intersection
	y0 := r.Origin[1] + t0*r.Direction[1]
	if s.Min < y0 && y0 < s.Max {
		xs = append(xs, &intersection.Intersection{Obj: s, T: t0})
	}

	y1 := r.Origin[1] + t1*r.Direction[1]
	if s.Min < y1 && y1 < s.Max {
		xs = append(xs, &intersection.Intersection{Obj: s, T: t1})
	}

	return s.intersectCaps(r, xs)
}

func (s *Cone) Intersections(origR ray.Ray) []*intersection.Intersection {
	return intersection.Intersections(s, origR, s.localIntersections)
}

func (s *Cone) localNormalAt(p tuple.Tuple) tuple.Tuple {
	dist := p[0]*p[0] + p[2]*p[2]
	if s.Closed && dist < 1.0 && p[1] >= s.Max-intersection.EPSILON {
		return tuple.Vector(0, 1.0, 0.0)
	}

	if s.Closed && dist < 1.0 && p[1] <= s.Min+intersection.EPSILON {
		return tuple.Vector(0, -1.0, 0.0)
	}

	return tuple.Vector(p[0], 0.0, p[2])
}

func (s *Cone) NormalAt(worldPoint tuple.Tuple) tuple.Tuple {
	return intersection.NormalAt(s, worldPoint, s.localNormalAt)
}

func (s *Cone) GetMaterial() *material.Material {
	return s.Material
}

func (s *Cone) SetMaterial(m *material.Material) {
	s.Material = m
}

func (s *Cone) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *Cone) SetTransform(m matrix.Mat4) {
	s.Transform = m
}
