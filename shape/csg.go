package shape

import (
	"math"

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

func (s *CSG) localIntersections(r ray.Ray) []*intersection.Intersection {
	// TODO optimize this when it is clear we don't need to check axis.
	xTMin, xTMax := s.checkAxis(r.Origin[0], r.Direction[0])
	yTMin, yTMax := s.checkAxis(r.Origin[1], r.Direction[1])
	zTMin, zTMax := s.checkAxis(r.Origin[2], r.Direction[2])

	tMin := math.Max(xTMin, math.Max(yTMin, zTMin))
	tMax := math.Min(xTMax, math.Min(yTMax, zTMax))
	if tMin > tMax {
		return nil
	}

	return []*intersection.Intersection{
		&intersection.Intersection{Obj: s, T: tMin},
		&intersection.Intersection{Obj: s, T: tMax},
	}
}

func (s *CSG) checkAxis(origin, direction float64) (float64, float64) {
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

func (s *CSG) Intersections(origR ray.Ray) []*intersection.Intersection {
	return intersection.Intersections(s, origR, s.localIntersections)
}

func (s *CSG) localNormalAt(p tuple.Tuple) tuple.Tuple {
	xAbs, yAbs, zAbs := math.Abs(p[0]), math.Abs(p[1]), math.Abs(p[2])

	if xAbs >= yAbs && xAbs >= zAbs {
		return tuple.Vector(p[0], 0.0, 0.0)
	} else if yAbs > xAbs && yAbs > zAbs {
		return tuple.Vector(0.0, p[1], 0.0)
	}
	return tuple.Vector(0.0, 0.0, p[2])
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
