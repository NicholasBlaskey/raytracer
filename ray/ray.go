package ray

import (
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Ray struct {
	Origin    tuple.Tuple
	Direction tuple.Tuple
}

func New(origin, dir tuple.Tuple) Ray {
	return Ray{origin, dir}
}

func (r Ray) PositionAt(t float64) tuple.Tuple {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (r Ray) Transform(m matrix.Mat4) Ray {
	return Ray{
		m.Mul4x1(r.Origin),
		m.Mul4x1(r.Direction),
	}
}
