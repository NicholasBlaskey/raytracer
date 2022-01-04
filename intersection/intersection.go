package intersection

import (
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/nicholasblaskey/raytracer/material"
)

const EPSILON = 0.000000001 // TODO expose this and use it system wide.

type Intersectable interface {
	NormalAt(tuple.Tuple) tuple.Tuple
	GetMaterial() *material.Material
}

type Intersection struct {
	T   float64
	Obj Intersectable
}

func New(t float64, obj Intersectable) *Intersection {
	return &Intersection{t, obj}
}

// See if this function is useful as syntatic sugar?
func Aggregate(intersections ...*Intersection) []*Intersection {
	return intersections
}

func Hit(intersections []*Intersection) *Intersection {
	minT := float64(9999999)
	var minIntersect *Intersection
	for _, i := range intersections {
		if i.T < minT && i.T >= 0.0 {
			minIntersect = i
			minT = i.T
		}
	}

	return minIntersect
}

type Computations struct {
	Obj       Intersectable
	T         float64
	Point     tuple.Tuple
	OverPoint tuple.Tuple
	Eyev      tuple.Tuple
	Normalv   tuple.Tuple
	Inside    bool
}

// TODO add multiple light sources here
func (i *Intersection) PrepareComputations(r ray.Ray) *Computations {
	c := &Computations{T: i.T, Obj: i.Obj}

	c.Point = r.PositionAt(c.T)

	c.Eyev = r.Direction.Mul(-1)
	c.Normalv = c.Obj.NormalAt(c.Point)

	if c.Normalv.Dot(c.Eyev) < 0.0 {
		c.Inside = true
		c.Normalv = c.Normalv.Mul(-1.0)
	}

	c.OverPoint = c.Point.Add(c.Normalv.Mul(EPSILON))

	return c
}
