package camera

import (
	"math"

	"github.com/schollz/progressbar/v3"

	"github.com/nicholasblaskey/raytracer/canvas"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"
	"github.com/nicholasblaskey/raytracer/world"
)

type Camera struct {
	HSize       int
	VSize       int
	FieldOfView float64
	Transform   matrix.Mat4
	PixelSize   float64
	halfWidth   float64
	halfHeight  float64
}

func New(hSize, vSize int, fov float64) *Camera {

	halfView := math.Tan(fov / 2.0)
	aspect := float64(hSize) / float64(vSize)

	var halfWidth, halfHeight float64
	if aspect >= 1.0 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		halfHeight = halfView
	}
	pixelSize := (halfWidth * 2.0) / float64(hSize)

	return &Camera{
		HSize:       hSize,
		VSize:       vSize,
		FieldOfView: fov,
		Transform:   matrix.Ident4(),
		PixelSize:   pixelSize,
		halfWidth:   halfWidth,
		halfHeight:  halfHeight,
	}
}

func (c *Camera) RayForPixel(x, y int) ray.Ray {
	xOffset := (float64(x) + 0.5) * c.PixelSize
	yOffset := (float64(y) + 0.5) * c.PixelSize

	worldX := c.halfWidth - xOffset
	worldY := c.halfHeight - yOffset

	transformInv := c.Transform.Inv()

	pixel := transformInv.Mul4x1(tuple.Point(worldX, worldY, -1.0))
	origin := transformInv.Mul4x1(tuple.Point(0.0, 0.0, 0.0))
	dir := pixel.Sub(origin).Normalize()

	return ray.Ray{origin, dir}
}

func (c *Camera) Render(w *world.World) *canvas.Canvas {
	canv := canvas.New(c.HSize, c.VSize)
	bar := progressbar.Default(int64(canv.Width * canv.Height))
	for y := 0; y < c.VSize; y++ {
		for x := 0; x < c.HSize; x++ {

			bar.Add(1)

			ray := c.RayForPixel(x, y)
			color := w.ColorAt(ray)
			canv.WritePixel(color, x, y)
		}
	}
	return canv
}
