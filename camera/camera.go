package camera

import (
	"github.com/nicholasblaskey/raytracer/matrix"
)

type Camera struct {
	HSize       int
	VSize       int
	FieldOfView float64
	Transform   matrix.Mat4
}

func New(hsize, vsize int, fov float64) *Camera {
	return &Camera{hsize, vsize, fov, matrix.Ident4()}
}
