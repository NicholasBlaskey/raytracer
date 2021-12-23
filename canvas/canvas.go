package canvas

import (
	"fmt"
	"math"

	"github.com/nicholasblaskey/raytracer/tuple"
)

const (
	maxLineLength = 70
)

type Canvas struct {
	Width  int
	Height int
	Pixels []float64
}

func New(w, h int) *Canvas {
	return &Canvas{w, h, make([]float64, w*h*3)}
}

func (c *Canvas) at(x, y int) int {
	return x*3 + y*3*c.Width
}

func (c *Canvas) WritePixel(pixel tuple.Tuple, x, y int) {
	c.Pixels[c.at(x, y)] = pixel[0]
	c.Pixels[c.at(x, y)+1] = pixel[1]
	c.Pixels[c.at(x, y)+2] = pixel[2]
}

func (c *Canvas) ReadPixel(x, y int) tuple.Tuple {
	return tuple.Tuple{
		c.Pixels[c.at(x, y)],
		c.Pixels[c.at(x, y)+1],
		c.Pixels[c.at(x, y)+2],
	}
}

func (c *Canvas) ToPPM() string {
	ppm := fmt.Sprintf("%s\n%d %d\n%d", "P3", c.Width, c.Height, 255)

	lineLen := 0
	pixelCounter := 0
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			col := c.ReadPixel(x, y)

			for i := 0; i < 3; i++ {
				v := col[i]
				if v > 1.0 {
					v = 1.0
				}
				if v < 0.0 {
					v = 0.0
				}

				toAdd := fmt.Sprintf("%d", int(math.Round(v*255.0)))
				if len(toAdd)+1+lineLen > maxLineLength || pixelCounter%15 == 0 {
					ppm += "\n"
					lineLen = 0
				}
				ppm += toAdd + " "
				pixelCounter += 1
			}
		}
	}

	ppm += "\n"

	return ppm
}

func (c *Canvas) Save(path string) error {
	return nil
}
