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
	Transform   matrix.Mat4
	Material    *material.Material
	Children    []intersection.Intersectable
	Parent      intersection.Intersectable
	boundingBox intersection.Bounds
}

func (s *Group) AddChild(c intersection.Intersectable) {
	c.SetParent(s)
	s.Children = append(s.Children, c)
}

func NewGroup() *Group {
	return &Group{Transform: matrix.Ident4(), Material: material.New()}
}

func (s *Group) checkAxis(min, max, origin, direction float64) (float64, float64) {
	tMinNumerator := (min - origin)
	tMaxNumerator := (max - origin)

	tMin := tMinNumerator / direction
	tMax := tMaxNumerator / direction

	if tMin > tMax {
		return tMax, tMin
	}
	return tMin, tMax
}

func (s *Group) missesBoundingBox(r ray.Ray) bool {
	b := s.Bounds()
	xTMin, xTMax := s.checkAxis(b.Min[0], b.Max[0], r.Origin[0], r.Direction[0])
	yTMin, yTMax := s.checkAxis(b.Min[1], b.Max[1], r.Origin[1], r.Direction[1])
	zTMin, zTMax := s.checkAxis(b.Min[2], b.Max[2], r.Origin[2], r.Direction[2])

	tMin := math.Max(xTMin, math.Max(yTMin, zTMin))
	tMax := math.Min(xTMax, math.Min(yTMax, zTMax))

	return tMin > tMax
}

func (s *Group) localIntersections(r ray.Ray) []*intersection.Intersection {
	if s.missesBoundingBox(r) {
		return nil
	}

	var xs []*intersection.Intersection
	for _, c := range s.Children {
		xs = append(xs, c.Intersections(r)...)
	}

	sort.Slice(xs, func(i, j int) bool {
		return xs[i].T < xs[j].T
	})

	return xs
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

func (s *Group) NormalAt(worldPoint tuple.Tuple, i *intersection.Intersection) tuple.Tuple {
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

func (s *Group) Bounds() intersection.Bounds {
	// TODO implmenet this correctly.
	//min := tuple.Point(-1.0, -1.0, -1.0)
	//max := tuple.Point(+1.0, +1.0, +1.0)

	min := tuple.Point(-10, -10, -10) //math.Inf(-1), math.Inf(-1), math.Inf(-1))
	max := tuple.Point(10, 10, 10)    //math.Inf(+1), math.Inf(+1), math.Inf(+1))

	s.boundingBox = intersection.Bounds{min, max}

	return s.boundingBox
}
