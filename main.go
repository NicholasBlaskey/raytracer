package main

import (
	//"fmt"

	"github.com/nicholasblaskey/raytracer/canvas"
	"github.com/nicholasblaskey/raytracer/tuple"
)

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
