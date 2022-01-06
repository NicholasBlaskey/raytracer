package material

import (
	"math"

	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Pattern interface {
	At(tuple.Tuple) tuple.Tuple
	AtObject(Object, tuple.Tuple) tuple.Tuple
	GetTransform() matrix.Mat4
	SetTransform(matrix.Mat4)
}

type Stripe struct {
	Color1    tuple.Tuple
	Color2    tuple.Tuple
	Transform matrix.Mat4
}

func StripePattern(c1, c2 tuple.Tuple) *Stripe {
	return &Stripe{c1, c2, matrix.Ident4()}
}

func (s *Stripe) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *Stripe) SetTransform(m matrix.Mat4) {
	s.Transform = m
}

func (s *Stripe) At(p tuple.Tuple) tuple.Tuple {
	if int(math.Floor(p[0]))%2 == 0 {
		return s.Color1
	}
	return s.Color2
}

// Hmmm rethink this abstraction.
// Could be intersectable but that would cause an import cycle?
// Could move intersectable to its own package?
type Object interface {
	GetTransform() matrix.Mat4
}

func (s *Stripe) AtObject(obj Object, worldPoint tuple.Tuple) tuple.Tuple {
	objPoint := obj.GetTransform().Inv().Mul4x1(worldPoint)
	patternPoint := s.GetTransform().Inv().Mul4x1(objPoint)

	return s.At(patternPoint)
}
