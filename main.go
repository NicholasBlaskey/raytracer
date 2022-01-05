package main

import (
	//"fmt"
	//"math"

	"math"

	"github.com/nicholasblaskey/raytracer/camera"
	"github.com/nicholasblaskey/raytracer/light"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/shape"
	"github.com/nicholasblaskey/raytracer/tuple"
	"github.com/nicholasblaskey/raytracer/world"
)

/*
// Draw a projectile.
func main() {
	cur := tuple.Point(0.0, 1.0, 0.0)
	velocity := tuple.Vector(1.0, 1.8, 0.0).Normalize().Mul(11.25)

	gravity := tuple.Vector(0.0, -0.15, 0.0)
	wind := tuple.Vector(0.0, -0.01, 0.0)
	e := gravity.Add(wind)

	c := canvas.New(900, 500)

	for cur[0] < float64(c.Width) && c.Height-int(cur[1]) < c.Height {
		x, y := int(cur[0]), c.Height-int(cur[1])
		c.WritePixel(tuple.Color(1.0, 0.0, 0.0), x, y)

		cur = cur.Add(velocity)
		velocity = velocity.Add(e)
	}

	if err := c.Save("test.ppm"); err != nil {
		panic(err)
	}
}
*/

/*
// Draw a clock.
func main() {
	w, h := 100.0, 100.0
	c := canvas.New(int(w), int(h))

	rot := matrix.RotateZ(math.Pi / 6)
	cur := tuple.Point(0.0, 0.8, 0.0)
	for i := 0; i < 12; i++ {
		// Transform to screen coords.
		x, y := (cur[0]+1)*(w/2), (cur[1]+1)*(h/2)
		c.WritePixel(tuple.Color(0.0, 0.0, 1.0), int(x), int(y))

		fmt.Println(cur, x, y)
		cur = rot.Mul4x1(cur)
	}

	if err := c.Save("test.ppm"); err != nil {
		panic(err)
	}
}
*/

/*
// Draw sphere (no shading)
func main() {
	w, h := 500, 500
	c := canvas.New(w, h)

	wallZ := 10.0
	rayOrigin := tuple.Point(0.0, 0.0, -5.0)
	wallSize := 7.0
	pixelSize := wallSize / float64(w)
	sphere := shape.NewSphere()
	sphere.Transform = matrix.Translate(0.1, 0.2, 0.3).Mul4(
		matrix.Scale(0.1, 0.8, 0.3))

	for y := 0; y < h; y++ {
		wallY := (wallSize / 2.0) - pixelSize*float64(y)
		for x := 0; x < w; x++ {
			wallX := (-wallSize / 2.0) + pixelSize*float64(x)

			pos := tuple.Point(wallX, wallY, wallZ)
			r := ray.New(rayOrigin, pos.Sub(rayOrigin).Normalize())
			xs := sphere.Intersections(r)

			if xs != nil {
				c.WritePixel(tuple.Color(0.0, 0.5, 0.5), x, y)
			}

		}
	}

	if err := c.Save("test.ppm"); err != nil {
		panic(err)
	}
}
*/

/*
// Draw sphere (shading)
func main() {
	m := material.New()
	m.Color = tuple.Color(1.0, 0.2, 1.0)
	m.Diffuse = 0.1
	m.Specular = 1.0
	m.Shininess = 0.5

	w, h := 500, 500
	c := canvas.New(w, h)

	wallZ := 10.0
	rayOrigin := tuple.Point(0.0, 0.0, -5.0)
	wallSize := 7.0
	pixelSize := wallSize / float64(w)
	sphere := shape.NewSphere()
	sphere.Material = m

	sphere.Transform = matrix.Ident4()
	//sphere.Transform = matrix.Translate(0.1, 0.2, 0.3).Mul4(
	//	matrix.Scale(0.5, 0.3, 0.5))

	lightPos := tuple.Point(-10.0, 10.0, -10.0)
	lightCol := tuple.Color(1.0, 1.0, 1.0)
	light := light.NewPointLight(lightPos, lightCol)

	for y := 0; y < h; y++ {
		wallY := (wallSize / 2.0) - pixelSize*float64(y)
		for x := 0; x < w; x++ {
			wallX := (-wallSize / 2.0) + pixelSize*float64(x)

			pos := tuple.Point(wallX, wallY, wallZ)
			r := ray.New(rayOrigin, pos.Sub(rayOrigin).Normalize())
			xs := sphere.Intersections(r)

			if xs != nil {
				point := r.PositionAt(xs[0].T)
				normal := sphere.NormalAt(point)
				eye := r.Direction.Mul(-1)
				col := sphere.Material.Lighting(light, point, eye, normal)

				c.WritePixel(col, x, y)
			}

		}
	}

	if err := c.Save("test.ppm"); err != nil {
		panic(err)
	}
}
*/

/*
// Draw scene
func main() {
	n := 600

	floor := shape.NewSphere()
	floor.Transform = matrix.Scale(10.0, 0.01, 10.0)
	floor.Material.Color = tuple.Color(1.0, 0.9, 0.9)
	floor.Material.Specular = 0.0

	leftWall := shape.NewSphere()
	leftWall.Transform = matrix.Translate(0.0, 0.0, 5.0).Mul4(
		matrix.RotateY(-math.Pi / 4.0)).Mul4(
		matrix.RotateX(math.Pi / 2.0)).Mul4(
		matrix.Scale(10.0, 0.01, 10.0))
	leftWall.Material = floor.Material

	rightWall := shape.NewSphere()
	rightWall.Transform = matrix.Translate(0.0, 0.0, 5.0).Mul4(
		matrix.RotateY(math.Pi / 4.0)).Mul4(
		matrix.RotateX(math.Pi / 2.0)).Mul4(
		matrix.Scale(10.0, 0.01, 10.0))
	rightWall.Material = floor.Material

	middle := shape.NewSphere()
	middle.Transform = matrix.Translate(-0.5, 1.0, 0.5)
	middle.Material.Color = tuple.Color(0.1, 1.0, 0.5)
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3

	right := shape.NewSphere()
	right.Transform = matrix.Translate(1.5, 0.5, -0.5).Mul4(
		matrix.Scale(0.5, 0.5, 0.5))
	right.Material.Color = tuple.Color(0.5, 1.0, 0.01)
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3

	left := shape.NewSphere()
	left.Transform = matrix.Translate(-1.5, 0.33, -0.75).Mul4(
		matrix.Scale(0.33, 0.33, 0.33))
	left.Material.Color = tuple.Color(1.0, 0.8, 0.1)
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3

	w := world.New()
	l := light.NewPointLight(tuple.Point(-10.0, 10.0, -10.0),
		tuple.Color(1.0, 1.0, 1.0))
	w.Light = &l
	w.Objects = append(w.Objects, floor, leftWall, rightWall, middle, right, left)

	c := camera.New(n*2, n, math.Pi/3.0)
	c.Transform = matrix.View(
		tuple.Point(0.0, 1.5, -5.0),
		tuple.Point(0.0, 1.0, 0.0),
		tuple.Vector(0.0, 1.0, 0.0))

	canv := c.Render(w)
	if err := canv.Save("test.ppm"); err != nil {
		panic(err)
	}

}
*/

// Draw a scene with shadows
func main() {
	n := 600

	floor := shape.NewSphere()
	floor.Transform = matrix.Scale(10.0, 0.01, 10.0)
	floor.Material.Color = tuple.Color(1.0, 0.9, 0.9)
	floor.Material.Specular = 0.0

	leftWall := shape.NewSphere()
	leftWall.Transform = matrix.Translate(0.0, 0.0, 5.0).Mul4(
		matrix.RotateY(-math.Pi / 4.0)).Mul4(
		matrix.RotateX(math.Pi / 2.0)).Mul4(
		matrix.Scale(10.0, 0.01, 10.0))
	leftWall.Material = floor.Material

	rightWall := shape.NewSphere()
	rightWall.Transform = matrix.Translate(0.0, 0.0, 5.0).Mul4(
		matrix.RotateY(math.Pi / 4.0)).Mul4(
		matrix.RotateX(math.Pi / 2.0)).Mul4(
		matrix.Scale(10.0, 0.01, 10.0))
	rightWall.Material = floor.Material

	middle := shape.NewSphere()
	middle.Transform = matrix.Translate(-0.5, 1.0, 0.5)
	middle.Material.Color = tuple.Color(0.3, 0.7, 0.3)
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3

	right := shape.NewSphere()
	right.Transform = matrix.Translate(0.9, 0.5, 0.3).Mul4(
		matrix.Scale(0.5, 0.5, 0.5))
	right.Material.Color = tuple.Color(0.7, 0.3, 0.3)
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3

	left := shape.NewSphere()
	left.Transform = matrix.Translate(-1.0, 2.1, -0.4).Mul4(
		matrix.RotateZ(-math.Pi / 3)).Mul4(
		matrix.Scale(0.33, 0.45, 0.33))
	left.Material.Color = tuple.Color(0.3, 0.3, 0.7)
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3

	w := world.New()
	l := light.NewPointLight(tuple.Point(-10.0, 10.0, -10.0),
		tuple.Color(1.0, 1.0, 1.0))
	w.Light = &l
	w.Objects = append(w.Objects, floor, leftWall, rightWall, middle, right, left)

	c := camera.New(n*2, n, math.Pi/3.0)
	c.Transform = matrix.View(
		tuple.Point(0.0, 1.5, -5.0),
		tuple.Point(0.0, 1.0, 0.0),
		tuple.Vector(0.0, 1.0, 0.0))

	canv := c.Render(w)
	if err := canv.Save("test.ppm"); err != nil {
		panic(err)
	}

}
