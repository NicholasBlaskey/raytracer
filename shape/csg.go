package shape

import (
	"sort"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Operation int

const (
	Union Operation = iota
	Intersection
	Difference
)

func IntersectionAllowed(opt Operation, lHit, inL, inR bool) bool {
	if opt == Union {
		return (lHit && !inR) || (!lHit && !inL)
	} else if opt == Intersection {
		return (lHit && inR) || (!lHit && inL)
	} else { // Difference
		return (lHit && !inR) || (!lHit && inL)
	}
}

type CSG struct {
	Transform matrix.Mat4
	Material  *material.Material
	Operation Operation
	Parent    intersection.Intersectable
	Left      intersection.Intersectable
	Right     intersection.Intersectable
}

func NewCSG(operation Operation, l, r intersection.Intersectable) *CSG {
	s := &CSG{
		Transform: matrix.Ident4(),
		Material:  material.New(),
		Operation: operation,
		Left:      l,
		Right:     r,
	}

	l.SetParent(s)
	r.SetParent(s)
	return s
}

func hitsObject(haystack, needle intersection.Intersectable) bool {
	switch obj := haystack.(type) {
	case *Group:
		for _, c := range obj.Children {
			if c == needle {
				return true
			}
		}
		return false
	case *CSG:
		return obj.Left == needle || obj.Right == needle
	default:
		return haystack == needle
	}
}

func (s *CSG) FilterIntersections(xs []*intersection.Intersection) []*intersection.Intersection {
	inL, inR := false, false
	var res []*intersection.Intersection

	for _, inter := range xs {
		lHit := hitsObject(s.Left, inter.Obj)

		if IntersectionAllowed(s.Operation, lHit, inL, inR) {
			res = append(res, inter)
		}

		if lHit {
			inL = !inL
		} else {
			inR = !inR
		}
	}

	return res
}

func (s *CSG) localIntersections(r ray.Ray) []*intersection.Intersection {
	xsL := s.Left.Intersections(r)
	xsR := s.Right.Intersections(r)
	xs := append(xsL, xsR...)

	sort.Slice(xs, func(i, j int) bool {
		return xs[i].T < xs[i].T
	})

	return xs
}

func (s *CSG) Intersections(origR ray.Ray) []*intersection.Intersection {
	return intersection.Intersections(s, origR, s.localIntersections)
}

func (s *CSG) localNormalAt(p tuple.Tuple) tuple.Tuple {
	panic("Something went wrong! Should never call a csg normal!")
	return tuple.Vector(0.0, 0.0, 0.0)
}

func (s *CSG) NormalAt(worldPoint tuple.Tuple, i *intersection.Intersection) tuple.Tuple {
	return intersection.NormalAt(s, worldPoint, s.localNormalAt)
}

func (s *CSG) GetMaterial() *material.Material {
	return s.Material
}

func (s *CSG) SetMaterial(m *material.Material) {
	s.Material = m
}

func (s *CSG) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *CSG) SetTransform(m matrix.Mat4) {
	s.Transform = m
}

func (s *CSG) GetParent() intersection.Intersectable {
	return s.Parent
}

func (s *CSG) SetParent(p intersection.Intersectable) {
	s.Parent = p
}

func (s *CSG) Bounds() intersection.Bounds {
	min := tuple.Point(-1.0, -1.0, -1.0)
	max := tuple.Point(+1.0, +1.0, +1.0)
	return intersection.Bounds{min, max}
}
