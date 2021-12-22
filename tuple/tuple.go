package tuple

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

func (t Tuple) Add(x Tuple) Tuple {
	return Tuple{t[0] + x[0], t[1] + x[1], t[2] + x[2], t[3] + x[3]}
}

func (t Tuple) Sub(x Tuple) Tuple {
	return Tuple{t[0] - x[0], t[1] - x[1], t[2] - x[2], t[3] - x[3]}
}

func (t Tuple) Negate() Tuple {
	return Tuple{-t[0], -t[1], -t[2], -t[3]}
}
