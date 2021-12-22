package tuple

type Tuple [4]float64

func New(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}
