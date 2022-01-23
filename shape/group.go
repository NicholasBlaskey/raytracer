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

var DrawBoundingBoxes = false

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
	s.calculateBoundingBox()
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

func (s *Group) missesBoundingBox(r ray.Ray) (float64, float64, bool) {
	b := s.Bounds()
	xTMin, xTMax := s.checkAxis(b.Min[0], b.Max[0], r.Origin[0], r.Direction[0])
	yTMin, yTMax := s.checkAxis(b.Min[1], b.Max[1], r.Origin[1], r.Direction[1])
	zTMin, zTMax := s.checkAxis(b.Min[2], b.Max[2], r.Origin[2], r.Direction[2])

	tMin := math.Max(xTMin, math.Max(yTMin, zTMin))
	tMax := math.Min(xTMax, math.Min(yTMax, zTMax))

	return tMin, tMax, tMin > tMax
}

func (s *Group) localIntersections(r ray.Ray) []*intersection.Intersection {
	if tMin, tMax, misses := s.missesBoundingBox(r); misses {
		return nil
	} else {
		if DrawBoundingBoxes {
			s.Material.Color = tuple.Tuple{0.3, 0.5, 0.3}
			return []*intersection.Intersection{
				&intersection.Intersection{Obj: s, T: tMin},
				&intersection.Intersection{Obj: s, T: tMax},
			}
		}
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
	// Only called in case of bounding box.
	// Convert from p being in [min, max] to being in [0, 1]
	// Solve for a in p = min*a+max*(1-a) where a \in [0, 1]
	a := p.Sub(s.boundingBox.Min).TupleDiv(s.boundingBox.Max.Sub(s.boundingBox.Min))
	// Convert from [0, 1] to [-1, 1]
	a = a.SubScalar(1.0 / 2.0).Mul(2.0)

	xAbs, yAbs, zAbs := math.Abs(a[0]), math.Abs(a[1]), math.Abs(a[2])

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

func (s *Group) calculateBoundingBox() {
	// Know the only way we call this is in add child,
	// so s.Children will always have len > 1.
	b := s.Children[0].Bounds()
	s.boundingBox = b
	for _, c := range s.Children[1:] {
		bounds := c.Bounds()
		minMax := []tuple.Tuple{bounds.Min, bounds.Max}

		// Get all 8 points of the children's bounding box.
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				for k := 0; k < 2; k++ {
					p := tuple.Point(minMax[i][0], minMax[j][1], minMax[k][2])
					groupP := c.GetTransform().Mul4x1(p)

					s.boundingBox.Min = s.boundingBox.Min.Min(groupP)
					s.boundingBox.Max = s.boundingBox.Max.Max(groupP)
				}
			}
		}
	}

	if s.Parent != nil {
		parent, ok := s.Parent.(*Group)
		if !ok {
			panic("Somehow object ended up with non group parent. Something serious went wrong")
		}
		parent.calculateBoundingBox()
	}
}

func (s *Group) Bounds() intersection.Bounds {
	return s.boundingBox
}
