package matrix

import (
	"math"

	"fmt"

	"github.com/nicholasblaskey/raytracer/tuple"
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

func Shear(xy, xz, yx, yz, zx, zy float64) Mat4 {
	return Mat4{
		1.0, yx, zx, 0.0,
		xy, 1.0, zy, 0.0,
		xz, yz, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func View(from, to, up tuple.Tuple) Mat4 {
	fmt.Println(from, to, up)

	forward := to.Sub(from).Normalize()
	upn := up.Normalize()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)

	orientation := Mat4{
		left[0], trueUp[0], -forward[0], 0.0,
		left[1], trueUp[1], -forward[1], 0.0,
		left[2], trueUp[2], -forward[2], 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	return orientation.Mul4(Translate(-from[0], -from[1], -from[2]))

	//return Translate(-from[0], -from[1], -from[2]).Mul4(orientation)

}
