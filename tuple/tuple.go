package tuple

import (
	"math"
)

type Tuple [4]float64

func New(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

// TODO should color be its own type? Or is tuple fine?
// We will allocate more data than we need with 4 points?!?!?
// This is prob fine for now.
func Color(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func (t Tuple) Add(x Tuple) Tuple {
	return Tuple{t[0] + x[0], t[1] + x[1], t[2] + x[2], t[3] + x[3]}
}

func (t Tuple) Sub(x Tuple) Tuple {
	return Tuple{t[0] - x[0], t[1] - x[1], t[2] - x[2], t[3] - x[3]}
}

func (t Tuple) Negate() Tuple {
	return Tuple{-t[0], -t[1], -t[2], -t[3]}
}

func (t Tuple) Mul(x float64) Tuple {
	return Tuple{t[0] * x, t[1] * x, t[2] * x, t[3] * x}
}

func (t Tuple) ColorMul(x Tuple) Tuple {
	return Tuple{t[0] * x[0], t[1] * x[1], t[2] * x[2], t[3] * x[3]}
}

func (t Tuple) Div(x float64) Tuple {
	return Tuple{t[0] / x, t[1] / x, t[2] / x, t[3] / x}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t[0]*t[0] + t[1]*t[1] + t[2]*t[2] + t[3]*t[3])
}

func (t Tuple) Normalize() Tuple {
	return t.Div(t.Magnitude())
}

func (t Tuple) Dot(x Tuple) float64 {
	return t[0]*x[0] + t[1]*x[1] + t[2]*x[2] + t[3]*x[3]
}

func (t Tuple) Cross(x Tuple) Tuple {
	return Tuple{
		t[1]*x[2] - t[2]*x[1],
		t[2]*x[0] - t[0]*x[2],
		t[0]*x[1] - t[1]*x[0],
		0.0,
	}
}

func (t Tuple) Reflect(n Tuple) Tuple {
	return t.Sub(n.Mul(2 * t.Dot(n)))
}

func (a Tuple) Min(b Tuple) Tuple {
	return Point(
		math.Min(a[0], b[0]),
		math.Min(a[1], b[1]),
		math.Min(a[2], b[2]),
	)
}

func (a Tuple) Max(b Tuple) Tuple {
	return Point(
		math.Max(a[0], b[0]),
		math.Max(a[1], b[1]),
		math.Max(a[2], b[2]),
	)
}
