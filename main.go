package main

import (
	"fmt"
	"math"

	"github.com/nicholasblaskey/raytracer/canvas"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/tuple"
)

/*
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