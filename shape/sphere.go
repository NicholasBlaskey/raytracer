package shape

import (
	"math"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Sphere struct {
	Transform matrix.Mat4
}

func NewSphere() *Sphere {
	return &Sphere{matrix.Ident4()}
}

func (s *Sphere) Intersections(origR ray.Ray) []*intersection.Intersection {
	r := origR.Transform(s.Transform.Inv())

	sphereToRay := r.Origin.Sub(tuple.Point(0.0, 0.0, 0.0))

	a := r.Direction.Dot(r.Direction)
	b := 2.0 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1.0

	discriminant := b*b - 4*a*c

	if discriminant < 0.0 {
		return nil
	}

	t0 := (-b - math.Sqrt(discriminant)) / (2.0 * a)
	t1 := (-b + math.Sqrt(discriminant)) / (2.0 * a)

	return []*intersection.Intersection{
		&intersection.Intersection{Obj: s, T: t0},
		&intersection.Intersection{Obj: s, T: t1}}
}

func (s *Sphere) NormalAt(worldPoint tuple.Tuple) tuple.Tuple {
	invT := s.Transform.Inv()

	objectPoint := invT.Mul4x1(worldPoint)
	objectNormal := objectPoint.Sub(tuple.Point(0.0, 0.0, 0.0))

	worldNormal := invT.T().Mul4x1(objectNormal)
	worldNormal[3] = 0.0

	return worldNormal.Normalize()
}
