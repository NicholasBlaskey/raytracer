package material

import (
	"math"

	"github.com/nicholasblaskey/raytracer/tuple"
)

type Pattern interface {
	At(tuple.Tuple) tuple.Tuple
}

type Stripe struct {
	Color1 tuple.Tuple
	Color2 tuple.Tuple
}

func StripePattern(c1, c2 tuple.Tuple) *Stripe {
	return &Stripe{c1, c2}
}

func (s *Stripe) At(p tuple.Tuple) tuple.Tuple {
	if int(math.Floor(p[0]))%2 == 0 {
		return s.Color1
	}
	return s.Color2
}
