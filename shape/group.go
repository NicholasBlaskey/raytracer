package shape

import (
	"math"
	"sort"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Group struct {
	Transform matrix.Mat4
	Material  *material.Material
	Children  []intersection.Intersectable
	Parent    intersection.Intersectable
}

func (s *Group) AddChild(c intersection.Intersectable) {
	c.SetParent(s)
	s.Children = append(s.Children, c)
}

func NewGroup() *Group {
	return &Group{Transform: matrix.Ident4(), Material: material.New()}
}

func (s *Group) localIntersections(r ray.Ray) []*intersection.Intersection {
	/*
		return []*intersection.Intersection{
			&intersection.Intersection{Obj: s, T: tMin},
			&intersection.Intersection{Obj: s, T: tMax},
		}
	*/
	var xs []*intersection.Intersection
	for _, c := range s.Children {
		xs = append(xs, c.Intersections(r)...)
	}

	sort.Slice(xs, func(i, j int) bool {
		return xs[i].T < xs[j].T
	})

	return xs
}

func (s *Group) checkAxis(origin, direction float64) (float64, float64) {
	tMinNumerator := (-1 - origin)
	tMaxNumerator := (1 - origin)

	tMin := tMinNumerator / direction
	tMax := tMaxNumerator / direction

	/* // Is this needed? Or does go handle infintly floating point division?
	tMin := tMinNumerator * math.Inf(1)
	tMax := tMaxNumerator * math.Inf(1)
	if math.Abs(direction) >= intersection.EPSILON {
		tMin = tMinNumerator / direction
		tMax = tMaxNumerator / direction
	}
	*/

	if tMin > tMax {
		return tMax, tMin
	}
	return tMin, tMax
}

func (s *Group) Intersections(origR ray.Ray) []*intersection.Intersection {
	return intersection.Intersections(s, origR, s.localIntersections)
}

func (s *Group) localNormalAt(p tuple.Tuple) tuple.Tuple {
	xAbs, yAbs, zAbs := math.Abs(p[0]), math.Abs(p[1]), math.Abs(p[2])

	if xAbs >= yAbs && xAbs >= zAbs {
		return tuple.Vector(p[0], 0.0, 0.0)
	} else if yAbs > xAbs && yAbs > zAbs {
		return tuple.Vector(0.0, p[1], 0.0)
	}
	return tuple.Vector(0.0, 0.0, p[2])
}

func (s *Group) NormalAt(worldPoint tuple.Tuple) tuple.Tuple {
	return intersection.NormalAt(s, worldPoint, s.localNormalAt)
}

func (s *Group) GetMaterial() *material.Material {
	return s.Material
}

func (s *Group) SetMaterial(m *material.Material) {
	s.Material = m
}

func (s *Group) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *Group) SetTransform(m matrix.Mat4) {
	s.Transform = m
}

func (s *Group) GetParent() intersection.Intersectable {
	return s.Parent
}

func (s *Group) SetParent(p intersection.Intersectable) {
	s.Parent = p
}
