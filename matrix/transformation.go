package matrix

func Translate(x, y, z float64) Mat4 {
	return Mat4{
		1.0, 0.0, 0.0, 0.0,
		0., 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		x, y, z, 1.0,
	}
}

func Scale(x, y, z float64) Mat4 {
	return Mat4{
		x, 0.0, 0.0, 0.0,
		0.0, y, 0.0, 0.0,
		0.0, 0.0, z, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}
