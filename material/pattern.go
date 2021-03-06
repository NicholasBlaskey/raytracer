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

// Hmmm rethink this abstraction.
// Could be intersectable but that would cause an import cycle?
// Could move intersectable to its own package?
type Object interface {
	GetTransform() matrix.Mat4
}

// TODO this currently does not handle groups!
// Need to add WorldToObject here.
// Reconsider this too.
func WorldToPattern(p Pattern, obj Object, worldPoint tuple.Tuple) tuple.Tuple {
	objPoint := obj.GetTransform().Inv().Mul4x1(worldPoint)
	patternPoint := p.GetTransform().Inv().Mul4x1(objPoint)

	return patternPoint
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

func (s *Stripe) AtObject(obj Object, worldPoint tuple.Tuple) tuple.Tuple {
	return s.At(WorldToPattern(s, obj, worldPoint))
}

type Gradient struct {
	Color1    tuple.Tuple
	Color2    tuple.Tuple
	Transform matrix.Mat4
}

func GradientPattern(c1, c2 tuple.Tuple) *Gradient {
	return &Gradient{c1, c2, matrix.Ident4()}
}

func (s *Gradient) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *Gradient) SetTransform(m matrix.Mat4) {
	s.Transform = m
}

// This doesn't seem right for spheres. Come back to this.
// The pattern only goes half way across the sphere for some reaosn.a
func (s *Gradient) At(p tuple.Tuple) tuple.Tuple {
	dist := s.Color2.Sub(s.Color1)
	//fraction := p[0] - math.Floor(p[0])
	fraction := p[0] - float64(int(p[0]))
	return s.Color1.Add(dist.Mul(fraction))
}

func (s *Gradient) AtObject(obj Object, worldPoint tuple.Tuple) tuple.Tuple {
	// TODO consider gradient pattern more
	p := WorldToPattern(s, obj, worldPoint)

	return s.At(p.Add(tuple.Vector(1.0, 1.0, 1.0)).Mul(1.0 / 2.0))
}

type Ring struct {
	Color1    tuple.Tuple
	Color2    tuple.Tuple
	Transform matrix.Mat4
}

func RingPattern(c1, c2 tuple.Tuple) *Ring {
	return &Ring{c1, c2, matrix.Ident4()}
}

func (s *Ring) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *Ring) SetTransform(m matrix.Mat4) {
	s.Transform = m
}

func (s *Ring) At(p tuple.Tuple) tuple.Tuple {
	if int(math.Floor(math.Sqrt(p[0]*p[0]+p[2]*p[2])))%2 == 0 {
		return s.Color1
	}
	return s.Color2
}

func (s *Ring) AtObject(obj Object, worldPoint tuple.Tuple) tuple.Tuple {
	return s.At(WorldToPattern(s, obj, worldPoint))
}

type Checker struct {
	Color1    tuple.Tuple
	Color2    tuple.Tuple
	Transform matrix.Mat4
}

func CheckerPattern(c1, c2 tuple.Tuple) *Checker {
	return &Checker{c1, c2, matrix.Ident4()}
}

func (s *Checker) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *Checker) SetTransform(m matrix.Mat4) {
	s.Transform = m
}

func floor(x float64) float64 {
	// TODO look into this. There is weird floating point causing
	// some achne on checkerboard pattern.
	const epsilon = 0.00000000001
	if x < epsilon && x > -epsilon {
		x = 0.0
	}

	return math.Floor(x)
}

func (s *Checker) At(p tuple.Tuple) tuple.Tuple {
	if int(floor(p[0])+floor(p[1])+floor(p[2]))%2 == 0 {
		return s.Color1
	}
	return s.Color2
}

func (s *Checker) AtObject(obj Object, worldPoint tuple.Tuple) tuple.Tuple {
	return s.At(WorldToPattern(s, obj, worldPoint))
}
