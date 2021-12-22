package canvas

import (
	"fmt"

	"github.com/nicholasblaskey/raytracer/tuple"
)

type Canvas struct {
	Width  int
	Height int
	Pixels []float64
}

func New(w, h int) *Canvas {
	return &Canvas{w, h, make([]float64, w*h*3)}
}

func (c *Canvas) WritePixel(pixel tuple.Tuple, x, y int) {
	c.Pixels[y*3+x+0] = pixel[0]
	c.Pixels[y*3+x+1] = pixel[1]
	c.Pixels[y*3+x+2] = pixel[2]
}

func (c *Canvas) ReadPixel(x, y int) tuple.Tuple {
	return tuple.Tuple{
		c.Pixels[y*3+x+0],
		c.Pixels[y*3+x+1],
		c.Pixels[y*3+x+2],
	}
}

func (c *Canvas) ToPPM() string {
	ppm := fmt.Sprintf("%s\n%d %d\n%d", "P3", c.Width, c.Height, 255)

	return ppm
}

func (c *Canvas) Save(path string) error {
	return nil
}
