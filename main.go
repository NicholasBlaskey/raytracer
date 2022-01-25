package main

import (
	"fmt"
	"math"

	"math/rand"

	"github.com/nicholasblaskey/raytracer/camera"
	"github.com/nicholasblaskey/raytracer/light"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
	"time"
	//"github.com/nicholasblaskey/raytracer/obj"
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

/*
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
*/

/*
// Draw scene with a plane
func main() {
	n := 100
	checker := material.CheckerPattern(
		tuple.Color(1.0, 1.0, 1.0),
		tuple.Color(0.0, 0.0, 0.0),
	)

	pattern := material.GradientPattern(
		tuple.Color(0.7, 0.3, 0.3),
		tuple.Color(0.3, 0.3, 0.7),
	)

	floor := shape.NewPlane()
	floor.Material.Color = tuple.Color(1.0, 0.9, 0.9)
	floor.Material.Specular = 0.0
	floor.Material.Pattern = checker

	leftWall := shape.NewPlane()
	leftWall.Transform = matrix.Translate(0.0, 0.0, 2.0).Mul4(
		matrix.RotateX(math.Pi / 2))
	leftWall.Material.Pattern = checker

	middle := shape.NewSphere()
	middle.Transform = matrix.Translate(-0.5, 2.0, 0.5)
	middle.Material.Color = tuple.Color(0.1, 1.0, 0.5)
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3
	middle.Material.Pattern = pattern

	right := shape.NewSphere()
	right.Transform = matrix.Translate(1.5, 0.5, -0.5).Mul4(
		matrix.Scale(0.5, 0.5, 0.5))
	right.Material.Color = tuple.Color(0.5, 1.0, 0.01)
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3
	right.Material.Pattern = pattern

	left := shape.NewSphere()
	left.Transform = matrix.Translate(-1.5, 0.33, -0.75).Mul4(
		matrix.Scale(0.33, 0.33, 0.33))
	left.Material.Color = tuple.Color(1.0, 0.8, 0.1)
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3
	left.Material.Pattern = pattern

	w := world.New()
	l := light.NewPointLight(tuple.Point(-10.0, 1.0, -10.0),
		tuple.Color(1.0, 1.0, 1.0))
	w.Light = &l
	w.Objects = append(w.Objects, middle, left, right, floor, leftWall)
	//w.Objects = append(w.Objects, middle)

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

/*
// Draw scene with reflection (cool)
func main() {
	fmt.Println("start")
	//n := 600
	n := 300
	checker := material.CheckerPattern(
		tuple.Color(1.0, 1.0, 1.0),
		tuple.Color(0.0, 0.0, 0.0),
	)

	stripes := material.StripePattern(
		tuple.Color(1.0, 1.0, 1.0),
		tuple.Color(0.0, 0.0, 0.0),
	)

	floor := shape.NewPlane()
	floor.Material.Color = tuple.Color(1.0, 0.9, 0.9)
	floor.Material.Specular = 0.0
	floor.Material.Pattern = checker

	leftWall := shape.NewPlane()
	leftWall.Transform = matrix.Translate(0.0, 0.0, 10.0).Mul4(
		matrix.RotateX(math.Pi / 2))
	leftWall.Material.Pattern = checker

	backWall := shape.NewPlane()
	backWall.Transform = matrix.Translate(0.0, 0.0, -12.0).Mul4(
		matrix.RotateX(math.Pi / 2))
	backWall.Material.Pattern = stripes

	glass := shape.NewGlassSphere()
	glass.Transform = matrix.Translate(-0.5, 1.0, -1.0).Mul4(
		matrix.Scale(1.0, 1.0, 1.0))
	//glass.Material.Color = tuple.Color(1.0, 1.0, 1.0) //tuple.Color(0.1, 1.0, 0.5)
	glass.Material.Diffuse = 0.5
	//glass.Material.Specular = 0.3
	//glass.Material.Reflective = 0.5

	air := shape.NewGlassSphere()
	air.Transform = matrix.Translate(-0.5, 1.0, -1.0).Mul4(
		matrix.Scale(0.3, 0.3, 0.3))
	air.Material.Color = tuple.Color(1.0, 1.0, 1.0)
	air.Material.RefractiveIndex = 1.02
	air.Material.Transparency = 0.8

	w := world.New()
	l := light.NewPointLight(tuple.Point(-10.0, 10.0, -10.0),
		tuple.Color(1.0, 1.0, 1.0))
	w.Light = &l
	w.Objects = append(w.Objects, floor, glass, air, leftWall, backWall)

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

/*
func hexagonCorner() *shape.Sphere {
	s := shape.NewSphere()
	s.SetTransform(matrix.Translate(0.0, 0.0, -1.0).Mul4(
		matrix.Scale(0.25, 0.25, 0.25)))
	return s
}

func hexagonEdge() *shape.Cylinder {
	s := shape.NewCylinder()
	s.Min = 0.0
	s.Max = 1.0
	s.SetTransform(matrix.Translate(0.0, 0.0, -1.0).Mul4(
		matrix.RotateY(-math.Pi / 6.0)).Mul4(
		matrix.RotateZ(-math.Pi / 2.0)).Mul4(
		matrix.Scale(0.25, 1.0, 0.25)))
	return s
}

func hexagonSide() *shape.Group {
	s := shape.NewGroup()
	s.AddChild(hexagonCorner())
	s.AddChild(hexagonEdge())
	return s
}

func hexagon() *shape.Group {
	hex := shape.NewGroup()
	for n := 0; n < 6; n++ {
		side := hexagonSide()
		side.SetTransform(matrix.RotateY(float64(n) * math.Pi / 3.0))
		hex.AddChild(side)
	}
	return hex
}

// Draw scene with cubes.
func main() {
	//n := 600
	n := 300

	hex := hexagon()

	w := world.New()
	l := light.NewPointLight(tuple.Point(-10.0, 10.0, -10.0),
		tuple.Color(1.0, 1.0, 1.0))
	w.Light = &l
	w.Objects = append(w.Objects, hex)

	c := camera.New(n*2, n, math.Pi/3.0)
	c.Transform = matrix.View(
		tuple.Point(-1.0, 3.0, -5.0),
		tuple.Point(0.0, 1.0, 0.0),
		tuple.Vector(0.0, 1.0, 0.0))

	canv := c.Render(w)
	if err := canv.Save("test.ppm"); err != nil {
		panic(err)
	}

}
*/

/*
// Draw teapot.
func main() {
	//n := 600
	n := 600
	checker := material.CheckerPattern(
		tuple.Color(1.0, 1.0, 1.0),
		tuple.Color(0.0, 0.0, 0.0),
	)

	stripes := material.StripePattern(
		tuple.Color(1.0, 1.0, 1.0),
		tuple.Color(0.0, 0.0, 0.0),
	)

	floor := shape.NewPlane()
	floor.Material.Color = tuple.Color(1.0, 0.9, 0.9)
	floor.Material.Specular = 0.0
	floor.Material.Pattern = checker

	leftWall := shape.NewPlane()
	leftWall.Transform = matrix.Translate(0.0, 0.0, 10.0).Mul4(
		matrix.RotateX(math.Pi / 2))
	leftWall.Material.Pattern = checker

	backWall := shape.NewPlane()
	backWall.Transform = matrix.Translate(0.0, 0.0, -12.0).Mul4(
		matrix.RotateX(math.Pi / 2))
	backWall.Material.Pattern = stripes

	fmt.Println("LOAD TEAPOT")
	teapot, err := obj.FileToGroup("models/lowResTeapot.obj")
	//teapot, err := obj.FileToGroup("models/highResTeapot.obj")
	teapot.SetTransform(matrix.Translate(0.0, 0.45, 0.0).Mul4(
		matrix.RotateX(-math.Pi / 2)).Mul4(
		matrix.Scale(0.10, 0.10, 0.10)))
	if err != nil {
		panic(err)
	}
	fmt.Println("END LOAD TEAPOT", teapot.Bounds())
	//shape.DrawBoundingBoxes = true

	//teapot.Material.Transparency = 1.0
	//teapot.Material.RefractiveIndex = 1.52
	//teapot.Material.Reflective = 1.0

	w := world.New()
	l := light.NewPointLight(tuple.Point(-10.0, 10.0, -10.0),
		tuple.Color(1.0, 1.0, 1.0))
	w.Light = &l
	w.Objects = append(w.Objects, floor, leftWall, backWall, teapot)

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

func generateSphere(xRange, zRange, sizeRange interval) *shape.Sphere {
	x := xRange.min + (xRange.max-xRange.min)*rand.Float64()
	z := zRange.min + (zRange.max-zRange.min)*rand.Float64()
	size := sizeRange.min + (sizeRange.max-sizeRange.min)*(rand.ExpFloat64()/2.0)

	/*
		x = 0.0
		z = 0.0
		size = 10.0
	*/

	s := shape.NewGlassSphere()
	s.Transform = matrix.Translate(x, size, z).Mul4(
		matrix.Scale(size, size, size))
	//s.Material.Reflective = 0.0
	//s.Material.RefractiveIndex = 0.0

	/*
		if rand.Int31n(2) == 0 && false {
			// Metalic
			fmt.Println("A")
			//s.Material.Reflective = 0.8
			//s.Material.Shininess = 800
			//s.Material.Specular = 0.6
			//s.Material.Color = tuple.Color(0.6, 0.6, 0.6)
		} else {
			// Glass
			fmt.Println("B")
			//s.Material.Reflective = 0.5
			s.Material.Transparency = 0.8
			s.Material.RefractiveIndex = 1.52
		}
	*/
	return s
}

type node struct {
	topLeft  *node
	topRight *node
	botLeft  *node
	botRight *node
	group    *shape.Group
}

func makeNodes(depth int) *node {
	if depth == 0 {
		return nil
	}

	n := &node{group: shape.NewGroup()}
	n.topLeft = makeNodes(depth - 1)
	n.topRight = makeNodes(depth - 1)
	n.botLeft = makeNodes(depth - 1)
	n.botRight = makeNodes(depth - 1)

	if n.topLeft != nil {
		n.group.AddChild(n.topLeft.group)
		n.group.AddChild(n.topRight.group)
		n.group.AddChild(n.botLeft.group)
		n.group.AddChild(n.botRight.group)
	}

	return n
}

func addToGroups(root *node, s *shape.Sphere, p tuple.Tuple, xRange, zRange interval) {
	if root.topLeft == nil {
		root.group.AddChild(s)
		return
	}

	xLow, xHigh := xRange.halfInterval()
	zLow, zHigh := zRange.halfInterval()

	if xLow.contains(p[0]) {
		if zLow.contains(p[2]) {
			addToGroups(root.topLeft, s, p, xLow, zLow)
		} else {
			addToGroups(root.topRight, s, p, xLow, zHigh)
		}
	} else {
		if zLow.contains(p[2]) {
			addToGroups(root.botLeft, s, p, xHigh, zLow)
		} else {
			addToGroups(root.botRight, s, p, xHigh, zHigh)
		}
	}
}

type interval struct {
	min float64
	max float64
}

func (i *interval) contains(x float64) bool {
	return x >= i.min && x <= i.max
}

func (i *interval) halfInterval() (interval, interval) {
	mid := (i.min + i.max) / 2.0

	l := interval{i.min, mid}
	h := interval{mid, i.max}
	return l, h
}

// Draws a ton of spheres.
func main() {
	start := time.Now()

	n := 900

	floor := shape.NewPlane()
	floor.Material.Color = tuple.Color(1.0, 0.9, 0.9)
	floor.Material.Specular = 0.0
	floor.Material.Pattern = material.CheckerPattern(
		tuple.Color(1.0, 1.0, 1.0),
		tuple.Color(0.0, 0.0, 0.0),
	)
	floor.Material.Pattern.SetTransform(matrix.Scale(10.0, 10.0, 10.0))

	ceil := shape.NewPlane()
	ceil.Transform = matrix.Translate(0.0, 50.0, 0.0)
	ceil.Material.Color = tuple.Color(0.3, 0.3, 0.7)
	ceil.Material.Specular = 0.0

	fmt.Println("FMT")

	w := world.New()
	l := light.NewPointLight(tuple.Point(25.0, 25.0, 0.0), //tuple.Point(-10.0, 30.0, -10.0),
		tuple.Color(1.0, 1.0, 1.0))
	w.Light = &l
	w.Objects = append(w.Objects, floor, ceil)

	// Generate spheres
	xRange := interval{-128, 128}
	zRange := interval{-32, 256}
	sizeRange := interval{0.30, 3.0}

	var spheres []*shape.Sphere
	for i := 0; i < 400; i++ {
		spheres = append(spheres, generateSphere(xRange, zRange, sizeRange))
	}

	// Make the bounding box.
	root := makeNodes(3)
	for _, s := range spheres {
		p := tuple.Point(0.0, 0.0, 0.0) // Sphere center
		p = s.Transform.Mul4x1(p)
		addToGroups(root, s, p, xRange, zRange)
	}

	shape.DrawBoundingBoxes = false
	w.Objects = append(w.Objects, root.group)
	/*
		for _, s := range spheres {
			w.Objects = append(w.Objects, s)
		}
	*/

	c := camera.New(n*2, n, math.Pi/3.0)
	c.Transform = matrix.View(
		tuple.Point(0.0, 25.0, -40.0),
		tuple.Point(0.0, 10.0, 0.0),
		tuple.Vector(0.0, 1.0, 0.0))
	canv := c.Render(w)
	if err := canv.Save("test.ppm"); err != nil {
		panic(err)
	}

	fmt.Println("took approximately", time.Now().Sub(start))
}
