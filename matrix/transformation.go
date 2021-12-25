package matrix

func Translate(x, y, z float64) Mat4 {
	return Mat4{
		1.0, 0.0, 0.0, 0.0,
		0., 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		x, y, z, 1.0,
	}
}
