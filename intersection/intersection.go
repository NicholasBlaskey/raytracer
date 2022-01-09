package intersection

import (
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/nicholasblaskey/raytracer/material"
)

const EPSILON = 0.000000001 // TODO expose this and use it system wide.

type Intersectable interface {
	GetTransform() matrix.Mat4
	SetTransform(matrix.Mat4)
	//
	GetMaterial() *material.Material
	SetMaterial(*material.Material)
	//
	NormalAt(tuple.Tuple) tuple.Tuple
	Intersections(ray.Ray) []*Intersection
}

func Intersections(i Intersectable, r ray.Ray,
	intersectFunc func(ray.Ray) []*Intersection) []*Intersection {

	return intersectFunc(r.Transform(i.GetTransform().Inv()))
}

func NormalAt(i Intersectable, worldPoint tuple.Tuple,
	normFunc func(tuple.Tuple) tuple.Tuple) tuple.Tuple {

	invT := i.GetTransform().Inv()
	objectPoint := invT.Mul4x1(worldPoint)
	objectNormal := normFunc(objectPoint)
	worldNormal := invT.T().Mul4x1(objectNormal)
	worldNormal[3] = 0.0

	return worldNormal.Normalize()
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
	Obj        Intersectable
	T          float64
	Point      tuple.Tuple
	OverPoint  tuple.Tuple
	UnderPoint tuple.Tuple
	Eyev       tuple.Tuple
	Normalv    tuple.Tuple
	Reflectv   tuple.Tuple
	N1         float64
	N2         float64
	Inside     bool
}

// TODO add multiple light sources here
func (i *Intersection) PrepareComputations(r ray.Ray, xs []*Intersection) *Computations {
	c := &Computations{T: i.T, Obj: i.Obj}

	c.Point = r.PositionAt(c.T)

	c.Eyev = r.Direction.Mul(-1)
	c.Normalv = c.Obj.NormalAt(c.Point)

	if c.Normalv.Dot(c.Eyev) < 0.0 {
		c.Inside = true
		c.Normalv = c.Normalv.Mul(-1.0)
	}

	c.UnderPoint = c.Point.Add(c.Normalv.Mul(-EPSILON))
	c.OverPoint = c.Point.Add(c.Normalv.Mul(EPSILON))
	c.Reflectv = r.Direction.Reflect(c.Normalv)

	var containers []Intersectable
	for _, inter := range xs {
		if inter == i {
			if len(containers) == 0 {
				c.N1 = 1.0
			} else {
				c.N1 = containers[len(containers)-1].GetMaterial().RefractiveIndex
			}
		}

		// If containers has inter.Obj then add inter.Obj to containers
		// otherwise remove it.
		found := false
		for j, obj := range containers {
			if obj == inter.Obj {
				containers[j] = containers[len(containers)-1]
				containers = containers[:len(containers)-1]
				found = true
			}
		}
		if !found {
			containers = append(containers, inter.Obj)
		}

		if inter == i {
			if len(containers) == 0 {
				c.N2 = 1.0
			} else {
				c.N2 = containers[len(containers)-1].GetMaterial().RefractiveIndex
			}
			break
		}
	}

	return c
}
