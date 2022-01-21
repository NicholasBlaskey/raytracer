package shape

import (
	//"math"
	//"sort"
	"fmt"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type BoundingBox struct {
	Transform matrix.Mat4
	Material  *material.Material
	Children  []intersection.Intersectable
	Parent    intersection.Intersectable
	group     *Group
}

func NewBoundingBox(group *Group) *BoundingBox {
	// Calculate min, max, for x, y, z

	return &BoundingBox{Transform: matrix.Ident4(), Material: material.New(), group: group}
}

func (s *BoundingBox) localIntersections(r ray.Ray) []*intersection.Intersection {
	return s.group.localIntersections(r)
}

func (s *BoundingBox) Intersections(origR ray.Ray) []*intersection.Intersection {
	// Are we in the bounding box? If not just return nil.

	// Otherwise check the intersection of everything in group
	return intersection.Intersections(s.group, origR, s.group.localIntersections)
}

func (s *BoundingBox) localNormalAt(p tuple.Tuple) tuple.Tuple {
	panic("SHOULD NOT BE CALLED")
	return tuple.Point(0.0, 0.0, 0.0)
}

func (s *BoundingBox) NormalAt(worldPoint tuple.Tuple, i *intersection.Intersection) tuple.Tuple {
	panic("SHOULD NOT BE CALLED")
	return tuple.Vector(0.0, 0.0, 0.0)
}

func (s *BoundingBox) GetMaterial() *material.Material {
	panic("SHOULD NOT BE CALLED")
	return s.Material
}

func (s *BoundingBox) SetMaterial(m *material.Material) {
	panic("SHOULD NOT BE CALLED")
	s.Material = m
}

func (s *BoundingBox) GetTransform() matrix.Mat4 {
	return s.group.Transform
}

func (s *BoundingBox) SetTransform(m matrix.Mat4) {
	fmt.Println(m)
	s.group.Transform = m
}

func (s *BoundingBox) GetParent() intersection.Intersectable {
	return s.Parent
}

func (s *BoundingBox) SetParent(p intersection.Intersectable) {
	s.Parent = p
}
