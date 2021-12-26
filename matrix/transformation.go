package matrix

import (
	"math"
)

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

func RotateX(theta float64) Mat4 {
	return Mat4{
		1.0, 0.0, 0.0, 0.0,
		0.0, math.Cos(theta), math.Sin(theta), 0.0,
		0.0, -math.Sin(theta), math.Cos(theta), 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func RotateY(theta float64) Mat4 {
	return Mat4{
		math.Cos(theta), 0.0, -math.Sin(theta), 0.0,
		0.0, 1.0, 0.0, 0.0,
		math.Sin(theta), 0.0, math.Cos(theta), 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func RotateZ(theta float64) Mat4 {
	return Mat4{
		math.Cos(theta), math.Sin(theta), 0.0, 0.0,
		-math.Sin(theta), math.Cos(theta), 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}
