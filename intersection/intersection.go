package intersection

import (
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Intersectable interface {
	NormalAt(tuple.Tuple) tuple.Tuple
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
	Obj     Intersectable
	T       float64
	Point   tuple.Tuple
	Eyev    tuple.Tuple
	Normalv tuple.Tuple
}

func (i *Intersection) PrepareComputations(r ray.Ray) *Computations {
	c := &Computations{T: i.T, Obj: i.Obj}

	c.Point = r.PositionAt(c.T)
	c.Eyev = r.Direction.Mul(-1)
	c.Normalv = c.Obj.NormalAt(c.Point)

	return c
}
