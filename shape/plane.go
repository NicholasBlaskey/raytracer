package shape

import (
	"math"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Plane struct {
	Transform matrix.Mat4
	Material  *material.Material
}

func NewPlane() *Plane {
	return &Plane{matrix.Ident4(), material.New()}
}

func (s *Plane) localIntersections(r ray.Ray) []*intersection.Intersection {
	// Ray is parallel to the plane? (plane is defined by (0, +1, 0)
	if math.Abs(r.Direction[1]) < intersection.EPSILON {
		return nil
	}

	t := -r.Origin[1] / r.Direction[1]

	return []*intersection.Intersection{
		&intersection.Intersection{Obj: s, T: t},
	}
}

func (s *Plane) Intersections(origR ray.Ray) []*intersection.Intersection {
	return intersection.Intersections(s, origR, s.localIntersections)
}

func (s *Plane) localNormalAt(objectPoint tuple.Tuple) tuple.Tuple {
	return tuple.Vector(0.0, 1.0, 0.0)
}

// TODO See if there is a performance hit doing this.
func (s *Plane) NormalAt(worldPoint tuple.Tuple) tuple.Tuple {
	return intersection.NormalAt(s, worldPoint, s.localNormalAt)
}

func (s *Plane) GetMaterial() *material.Material {
	return s.Material
}

func (s *Plane) SetMaterial(m *material.Material) {
	s.Material = m
}

func (s *Plane) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *Plane) SetTransform(m matrix.Mat4) {
	s.Transform = m
}
